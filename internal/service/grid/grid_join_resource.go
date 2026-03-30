package grid

import (
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
		option.WithNIOSHostUrl(data.MemberIP.ValueString()),
		option.WithDebug(true),
	)

	joinReq := gridclient.GridJoin{
		GridName:     data.GridName.ValueStringPointer(),
		Master:       data.Master.ValueStringPointer(),
		SharedSecret: data.SharedSecret.ValueStringPointer(),
	}

	_, httpResp, err := memberClient.GridAPI.
		GridJoinAPI.
		Create(ctx).
		GridJoin(joinReq).
		ReturnAsObject(1).
		Execute()

	// Check for 200 response with HTML body indicating member is already joined to a grid master
	if err != nil && httpResp != nil && httpResp.StatusCode == 200 {
		if httpResp.Body != nil {
			bodyBytes, readErr := io.ReadAll(httpResp.Body)
			defer httpResp.Body.Close()
			if readErr == nil {
				bodyStr := string(bodyBytes)
				// If response is HTML redirect, member is already joined
				if strings.Contains(bodyStr, "<HTML>") && strings.Contains(bodyStr, "REFRESH") {
					// Extract URL from HTML redirect
					masterURL := ""
					if urlStart := strings.Index(bodyStr, "URL="); urlStart != -1 {
						urlStart += 4
						urlEnd := strings.IndexAny(bodyStr[urlStart:], "\"'")
						if urlEnd != -1 {
							masterURL = strings.TrimSpace(bodyStr[urlStart : urlStart+urlEnd])
						}
					}

					var errorMsg string
					if masterURL != "" {
						errorMsg = fmt.Sprintf("The member is already part of grid master: %s", masterURL)
					}

					resp.Diagnostics.AddError("Grid Join Failed", errorMsg)
					resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
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
	tflog.Debug(ctx, "Grid join Initiated", map[string]any{
		"member_ip": data.MemberIP.ValueString(),
		"master":    data.Master.ValueString(),
		"grid_name": data.GridName.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *GridJoinResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GridJoinModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GridJoinResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GridJoinModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GridJoinResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GridJoinModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}
