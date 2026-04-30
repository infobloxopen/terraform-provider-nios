package grid

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	gridclient "github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/infoblox-nios-go-client/option"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GridJoinResource{}

func NewGridJoinResource() resource.Resource {
	return &GridJoinResource{}
}

// GridJoinResource defines the resource implementation.
type GridJoinResource struct {
	client *niosclient.APIClient
}

func (r *GridJoinResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "grid_join"
}

func (r *GridJoinResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages joining a member to an Infoblox Grid.",
		Attributes:          GridJoinResourceSchemaAttributes,
	}
}

func (r *GridJoinResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*niosclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *niosclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *GridJoinResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GridJoinModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	memberClient := niosclient.NewAPIClient(
		option.WithNIOSUsername(data.MemberUsername.ValueString()),
		option.WithNIOSPassword(data.MemberPassword.ValueString()),
		option.WithNIOSHostUrl(data.MemberURL.ValueString()),
		option.WithDebug(true),
	)

	joinReq := gridclient.GridJoin{
		GridName:     data.GridName.ValueStringPointer(),
		Master:       data.Master.ValueStringPointer(),
		SharedSecret: data.SharedSecret.ValueStringPointer(),
	}

	_, httpResp, err := memberClient.GridAPI.
		GridAPI.
		Create(ctx).
		GridJoin(joinReq).
		ReturnAsObject(1).
		Function("join").
		Execute()

	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}

	// Check for 200 response with HTML body indicating member is already joined to a grid master
	if err != nil && httpResp != nil && httpResp.StatusCode == 200 {
		if httpResp.Body != nil {
			bodyBytes, readErr := io.ReadAll(httpResp.Body)
			if readErr == nil {
				isHTMLResponse := strings.Contains(httpResp.Header.Get("Content-Type"), "text/html")
				isHTML := bytes.Contains(bodyBytes, []byte("<HTML"))
				hasMetaRefresh := bytes.Contains(bodyBytes, []byte("HTTP-EQUIV=\"REFRESH\""))
				// If response is HTML redirect, member is already joined
				if isHTMLResponse && isHTML && hasMetaRefresh {
					// Extract URL from HTML redirect
					masterURL := extractRedirectURL(string(bodyBytes))
					errorMsg := "Member is already part of another grid."
					if masterURL != "" {
						errorMsg += fmt.Sprintf(" Master URL: %s", masterURL)
					}
					resp.Diagnostics.AddError("Grid Join Failed", errorMsg)
					return
				}
			}
		}
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Grid Join Error",
			fmt.Sprintf("Failed to join grid: %s", err),
		)
		return
	}

	// Normal 200 response - grid join initiated
	tflog.Debug(ctx, "Grid join initiated successfully. Please verify the grid status manually through the NIOS GUI to confirm the member has joined. If the join operation fails, manual intervention may be required to troubleshoot connectivity issues, credentials, or grid configuration.", map[string]any{
		"member_url": data.MemberURL.ValueString(),
		"master":     data.Master.ValueString(),
		"grid_name":  data.GridName.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *GridJoinResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Read does not perform any verification of the grid join status.
	// The grid join operation is a one-time action.
}

func (r *GridJoinResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is not supported as all grid join attributes are immutable.
}

func (r *GridJoinResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Delete only removes the resource from Terraform state without making any API call.
	// To actually unjoin a member from the grid, you must delete the corresponding nios_grid_member resource,
	// which will remove the member from the grid and make it a standalone Grid.
	// This resource does not directly support unjoining a member through deletion.
}

func extractRedirectURL(body string) string {
	idx := strings.Index(body, "URL=")
	if idx == -1 {
		return ""
	}

	start := idx + 4
	end := strings.IndexAny(body[start:], "\"'")
	if end == -1 {
		return ""
	}

	return strings.TrimSpace(body[start : start+end])
}
