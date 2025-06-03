package dtc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/Infoblox-CTO/infoblox-nios-go-client/client"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &DtcpoolDataSource{}

func NewDtcpoolDataSource() datasource.DataSource {
	return &DtcpoolDataSource{}
}

// DtcpoolDataSource defines the data source implementation.
type DtcpoolDataSource struct {
	client *niosclient.APIClient
}

func (d *DtcpoolDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "datasource_nios_DtcPool"
}

type DtcPoolModelWithFilter struct {
	Filters        types.Map  `tfsdk:"filters"`
	ExtAttrFilters types.Map  `tfsdk:"extattrfilters"`
	Result         types.List `tfsdk:"result"`
}

func (m *DtcPoolModelWithFilter) FlattenResults(ctx context.Context, from []dtc.DtcPool, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Result = flex.FlattenFrameworkListNestedBlock(ctx, from, DtcPoolAttrTypes, diags, FlattenDtcPool)
}

func (d *DtcpoolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"filters": schema.MapAttribute{
				Description: "Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"extattrfilters": schema.MapAttribute{
				Description: "External Attribute Filters are used to return a more specific list of results by filtering on external attributes. If you specify multiple filters, the results returned will have only resources that match all the specified filters.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: utils.DataSourceAttributeMap(DtcPoolResourceSchemaAttributes, &resp.Diagnostics),
				},
				Computed: true,
			},
		},
	}
}

func (d *DtcpoolDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *DtcpoolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DtcPoolModelWithFilter

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := d.client.DTCAPI.
		DtcpoolAPI.
		Get(ctx).
		Filters(flex.ExpandFrameworkMapString(ctx, data.Filters, &resp.Diagnostics)).
		Extattrfilter(flex.ExpandFrameworkMapString(ctx, data.ExtAttrFilters, &resp.Diagnostics)).
		ReturnAsObject(1).
		ReturnFields2(readableAttributesForDtcPool).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Dtcpool, got error: %s", err))
		return
	}

	res := apiRes.ListDtcPoolResponseObject.GetResult()
	data.FlattenResults(ctx, res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
