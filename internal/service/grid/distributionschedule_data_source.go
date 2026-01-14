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
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &DistributionscheduleDataSource{}

func NewDistributionscheduleDataSource() datasource.DataSource {
	return &DistributionscheduleDataSource{}
}

// DistributionscheduleDataSource defines the data source implementation.
type DistributionscheduleDataSource struct {
	client *niosclient.APIClient
}

func (d *DistributionscheduleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "grid_distributionschedule"
}

type DistributionscheduleModelWithFilter struct {
	Result types.List `tfsdk:"result"`
}

func (m *DistributionscheduleModelWithFilter) FlattenResults(ctx context.Context, from []grid.Distributionschedule, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Result = flex.FlattenFrameworkListNestedBlock(ctx, from, DistributionscheduleAttrTypes, diags, FlattenDistributionschedule)
}

func (d *DistributionscheduleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves information about current Distribution Schedule config from the NIOS Grid.",
		Attributes: map[string]schema.Attribute{
			"result": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: utils.DataSourceAttributeMap(DistributionscheduleResourceSchemaAttributes, &resp.Diagnostics),
				},
				Computed: true,
			},
		},
	}
}

func (d *DistributionscheduleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *DistributionscheduleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DistributionscheduleModelWithFilter

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.GridAPI.
		DistributionscheduleAPI.
		List(ctx).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForDistributionschedule)

	// Execute the request
	apiRes, _, err := request.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Distributionschedule, got error: %s", err))
		return
	}

	results := apiRes.ListDistributionscheduleResponseObject.GetResult()

	tflog.Info(ctx, fmt.Sprintf("Retrieved %d results", len(results)))

	// Process the results
	data.FlattenResults(ctx, results, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
