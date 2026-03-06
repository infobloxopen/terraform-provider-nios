package grid

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &UpgradescheduleDataSource{}

func NewUpgradescheduleDataSource() datasource.DataSource {
	return &UpgradescheduleDataSource{}
}

// UpgradescheduleDataSource defines the data source implementation.
type UpgradescheduleDataSource struct {
	client *niosclient.APIClient
}

func (d *UpgradescheduleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "grid_upgradeschedule"
}

type UpgradescheduleModelWithFilter struct {
	Result types.List `tfsdk:"result"`
}

func (m *UpgradescheduleModelWithFilter) FlattenResults(ctx context.Context, from []grid.Upgradeschedule, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Result = flex.FlattenFrameworkListNestedBlock(ctx, from, UpgradescheduleAttrTypes, diags, FlattenUpgradeschedule)
}

func (d *UpgradescheduleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves information about current Upgrade Schedule config.",
		Attributes: map[string]schema.Attribute{
			"result": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: utils.DataSourceAttributeMap(UpgradescheduleResourceSchemaAttributes, &resp.Diagnostics),
				},
				Computed: true,
			},
		},
	}
}

func (d *UpgradescheduleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*niosclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *niosclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *UpgradescheduleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data UpgradescheduleModelWithFilter

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.GridAPI.
		UpgradescheduleAPI.
		List(ctx).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForUpgradeschedule).
		ProxySearch(config.GetProxySearch())

	// Execute the request
	apiRes, _, err := request.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Upgradeschedule, got error: %s", err))
		return
	}

	res := apiRes.ListUpgradescheduleResponseObject.GetResult()

	tflog.Info(ctx, fmt.Sprintf("Retrieved %d results", len(res)))

	// Process the results
	data.FlattenResults(ctx, res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
