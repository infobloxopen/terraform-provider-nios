package grid

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var (
	_ action.Action              = &RestartServicesAction{}
	_ action.ActionWithConfigure = &RestartServicesAction{}
)

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
	resp.Schema = schema.Schema{
		MarkdownDescription: "Requests a restart of NIOS Grid services.",
		Attributes:          RestartServicesActionSchemaAttributes,
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
	var data RestartServicesModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	gridResponse, _, err := a.client.GridAPI.GridAPI.
		List(ctx).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get Grid Reference, got error: %s", err))
		return
	}

	grids := gridResponse.ListGridResponseObject.GetResult()
	if len(grids) == 0 || grids[0].GetRef() == "" {
		resp.Diagnostics.AddError("Client Error", "Unable to get Grid Reference, got empty result")
		return
	}

	gridRef := grids[0].GetRef()

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := a.client.GridAPI.GridAPI.
		CreateFunctionCall(ctx, utils.ExtractResourceRef(gridRef)).
		GridFunctionCall(*payload).
		ReturnAsObject(1).
		Function("restartservices").
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to restart NIOS services, got error: %s", err))
		return
	}

	if httpResp.StatusCode < http.StatusOK || httpResp.StatusCode >= http.StatusMultipleChoices {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to restart NIOS services, got status code: %d", httpResp.StatusCode))
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: "NIOS accepted the service restart request."})
	}
}
