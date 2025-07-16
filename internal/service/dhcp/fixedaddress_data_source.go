package dhcp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &FixedaddressDataSource{}

func NewFixedaddressDataSource() datasource.DataSource {
	return &FixedaddressDataSource{}
}

// FixedaddressDataSource defines the data source implementation.
type FixedaddressDataSource struct {
	client *niosclient.APIClient
}

func (d *FixedaddressDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_fixedaddress"
}

type FixedaddressModelWithFilter struct {
	Filters        types.Map  `tfsdk:"filters"`
	ExtAttrFilters types.Map  `tfsdk:"extattrfilters"`
	Result         types.List `tfsdk:"result"`
	Body           types.Map  `tfsdk:"body"`
}

func (m *FixedaddressModelWithFilter) FlattenResults(ctx context.Context, from []dhcp.Fixedaddress, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Result = flex.FlattenFrameworkListNestedBlock(ctx, from, FixedaddressAttrTypes, diags, FlattenFixedaddress)
}

func (d *FixedaddressDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
					Attributes: utils.DataSourceAttributeMap(FixedaddressResourceSchemaAttributes, &resp.Diagnostics),
				},
				Computed: true,
			},
			"body": schema.MapAttribute{
				Description: "The body of the request to be sent to the API. This is used for creating or updating resources.",
				ElementType: types.StringType,
				Optional:    true,
			},
		},
	}
}

func (d *FixedaddressDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *FixedaddressDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data FixedaddressModelWithFilter

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if !data.Body.IsNull() && !data.Body.IsUnknown() {
		// If body is provided, we will use it to create a new fixed address
		apiRes, _, err := d.client.DHCPAPI.
			FixedaddressAPI.
			StructUpdate(ctx).
			StructUpdate(flex.ExpandFrameworkMapString(ctx, data.Filters, &resp.Diagnostics)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForFixedaddress).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Fixedaddress, got error: %s", err))
			return
		}

		res := apiRes.UpdateFixedaddressResponseAsObject.GetResult()
		data.FlattenResults(ctx, res, &resp.Diagnostics)

		// Save updated data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	apiRes, _, err := d.client.DHCPAPI.
		FixedaddressAPI.
		List(ctx).
		Filters(flex.ExpandFrameworkMapString(ctx, data.Filters, &resp.Diagnostics)).
		Extattrfilter(flex.ExpandFrameworkMapString(ctx, data.ExtAttrFilters, &resp.Diagnostics)).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForFixedaddress).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Fixedaddress, got error: %s", err))
		return
	}

	res := apiRes.ListFixedaddressResponseObject.GetResult()
	data.FlattenResults(ctx, res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
