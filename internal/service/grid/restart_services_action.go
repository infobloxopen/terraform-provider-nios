package grid

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
)

const (
	restartOptionIfNeeded = "RESTART_IF_NEEDED"
	restartServiceAll     = "ALL"
)

var (
	_ action.Action              = &RestartServicesAction{}
	_ action.ActionWithConfigure = &RestartServicesAction{}
)

type restartServicesActionModel struct {
	Groups        types.List   `tfsdk:"groups"`
	Members       types.List   `tfsdk:"members"`
	Mode          types.String `tfsdk:"mode"`
	RestartOption types.String `tfsdk:"restart_option"`
	Services      types.List   `tfsdk:"services"`
	UserName      types.String `tfsdk:"user_name"`
}

type restartServicesRequest struct {
	Groups        []string `json:"groups,omitempty"`
	Members       []string `json:"members,omitempty"`
	Mode          string   `json:"mode,omitempty"`
	RestartOption string   `json:"restart_option"`
	Services      []string `json:"services"`
	UserName      string   `json:"user_name,omitempty"`
}

// RestartServicesAction invokes the NIOS restartservices Grid function.
type RestartServicesAction struct {
	client *niosclient.APIClient
}

func NewRestartServicesAction() action.Action {
	return &RestartServicesAction{}
}

func (a *RestartServicesAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_restart_services"
}

func (a *RestartServicesAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	serviceValidator := stringvalidator.OneOf("ALL", "DNS", "DHCP", "DHCPV4", "DHCPV6")

	resp.Schema = schema.Schema{
		MarkdownDescription: "Requests a restart of NIOS Grid services.",
		Attributes: map[string]schema.Attribute{
			"groups": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Service Restart Group names to restart. `groups` conflicts with `members`. If both are omitted, NIOS uses the default restart group.",
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
					listvalidator.UniqueValues(),
					listvalidator.ConflictsWith(path.MatchRoot("members")),
				},
			},
			"members": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Grid Member names to restart. `members` conflicts with `groups`. If both are omitted, NIOS uses the default restart group.",
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
					listvalidator.UniqueValues(),
					listvalidator.ConflictsWith(path.MatchRoot("groups")),
				},
			},
			"mode": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Grid restart method. Valid values are `GROUPED`, `SEQUENTIAL`, and `SIMULTANEOUS`.",
				Validators: []validator.String{
					stringvalidator.OneOf("GROUPED", "SEQUENTIAL", "SIMULTANEOUS"),
				},
			},
			"restart_option": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Controls whether NIOS restarts services only when needed or unconditionally. Valid values are `RESTART_IF_NEEDED` and `FORCE_RESTART`. Defaults to `RESTART_IF_NEEDED`.",
				Validators: []validator.String{
					stringvalidator.OneOf("RESTART_IF_NEEDED", "FORCE_RESTART"),
				},
			},
			"services": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Services to restart. Valid values are `ALL`, `DNS`, `DHCP`, `DHCPV4`, and `DHCPV6`. Defaults to `ALL`.",
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(serviceValidator),
				},
			},
			"user_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name of the user requesting the restart, recorded by NIOS.",
			},
		},
	}
}

func (a *RestartServicesAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*niosclient.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Action Configure Type",
			fmt.Sprintf("Expected *niosclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	a.client = client
}

func (a *RestartServicesAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data restartServicesActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := restartServicesRequest{
		RestartOption: restartOptionIfNeeded,
		Services:      []string{restartServiceAll},
	}

	if !data.Groups.IsNull() && !data.Groups.IsUnknown() {
		resp.Diagnostics.Append(data.Groups.ElementsAs(ctx, &payload.Groups, false)...)
	}
	if !data.Members.IsNull() && !data.Members.IsUnknown() {
		resp.Diagnostics.Append(data.Members.ElementsAs(ctx, &payload.Members, false)...)
	}
	if !data.Mode.IsNull() && !data.Mode.IsUnknown() {
		payload.Mode = data.Mode.ValueString()
	}
	if !data.RestartOption.IsNull() && !data.RestartOption.IsUnknown() {
		payload.RestartOption = data.RestartOption.ValueString()
	}
	if !data.Services.IsNull() && !data.Services.IsUnknown() {
		resp.Diagnostics.Append(data.Services.ElementsAs(ctx, &payload.Services, false)...)
	}
	if !data.UserName.IsNull() && !data.UserName.IsUnknown() {
		payload.UserName = data.UserName.ValueString()
	}
	if resp.Diagnostics.HasError() {
		return
	}

	if a.client == nil {
		resp.Diagnostics.AddError("Client Error", "The provider has not configured the NIOS API client.")
		return
	}

	if err := invokeRestartServices(ctx, a.client, payload); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to restart NIOS services, got error: %s", err))
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: "NIOS accepted the service restart request."})
	}
}

func invokeRestartServices(ctx context.Context, client *niosclient.APIClient, payload restartServicesRequest) error {
	gridResponse, _, err := client.GridAPI.GridAPI.
		List(ctx).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		return fmt.Errorf("get Grid reference: %w", err)
	}
	if gridResponse == nil || gridResponse.ListGridResponseObject == nil {
		return fmt.Errorf("get Grid reference: NIOS returned an empty response")
	}

	grids := gridResponse.ListGridResponseObject.GetResult()
	if len(grids) == 0 || grids[0].GetRef() == "" {
		return fmt.Errorf("get Grid reference: NIOS returned no Grid objects")
	}

	baseURL, err := client.GridAPI.Cfg.ServerURLWithContext(ctx, "GridAPIService.Create")
	if err != nil {
		return fmt.Errorf("build restart services URL: %w", err)
	}
	restartURL := strings.TrimRight(baseURL, "/") + "/" + strings.TrimLeft(grids[0].GetRef(), "/")
	query := url.Values{"_function": []string{"restartservices"}}
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	httpRequest, err := client.GridAPI.PrepareRequest(ctx, restartURL, http.MethodPost, payload, headers, query, nil, nil)
	if err != nil {
		return fmt.Errorf("prepare restart services request: %w", err)
	}

	httpResponse, err := client.GridAPI.CallAPI(httpRequest)
	if err != nil {
		return fmt.Errorf("call restart services: %w", err)
	}
	if httpResponse == nil {
		return fmt.Errorf("call restart services: NIOS returned no HTTP response")
	}

	body, readErr := io.ReadAll(io.LimitReader(httpResponse.Body, 64*1024))
	closeErr := httpResponse.Body.Close()
	if readErr != nil {
		return fmt.Errorf("restart services returned HTTP %d; read response: %w", httpResponse.StatusCode, readErr)
	}
	if closeErr != nil {
		return fmt.Errorf("restart services returned HTTP %d; close response: %w", httpResponse.StatusCode, closeErr)
	}
	if httpResponse.StatusCode < http.StatusOK || httpResponse.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("restart services returned HTTP %d: %s", httpResponse.StatusCode, strings.TrimSpace(string(body)))
	}

	return nil
}
