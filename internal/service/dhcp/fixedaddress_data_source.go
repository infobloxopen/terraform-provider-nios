package dhcp

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

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
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_fixed_address"
}

type FixedaddressModelWithFilter struct {
	Filters        types.Map    `tfsdk:"filters"`
	ExtAttrFilters types.Map    `tfsdk:"extattrfilters"`
	Result         types.List   `tfsdk:"result"`
	MaxResults     types.Int32  `tfsdk:"max_results"`
	Paging         types.Int32  `tfsdk:"paging"`
	Body           types.Object `tfsdk:"body"`
}

func (m *FixedaddressModelWithFilter) FlattenResults(ctx context.Context, from []dhcp.Fixedaddress, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Result = flex.FlattenFrameworkListNestedBlock(ctx, from, FixedaddressAttrTypes, diags, FlattenFixedaddress)
}

func (d *FixedaddressDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves information about existing Fixed Address.",
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
			"paging": schema.Int32Attribute{
				Optional:    true,
				Description: "Enable (1) or disable (0) paging for the data source query. When enabled, the system retrieves results in pages, allowing efficient handling of large result sets. Paging is enabled by default.",
				Validators: []validator.Int32{
					int32validator.OneOf(0, 1),
				},
			},
			"max_results": schema.Int32Attribute{
				Optional:    true,
				Description: "Maximum number of objects to be returned. Defaults to 1000.",
			},
			"body": schema.SingleNestedAttribute{
				Description: "The body of the request to Search Fixed Address using Microsoft Server.",
				Attributes:  FixedAddressStructResourceSchemaAttributes,
				Optional:    true,
				Computed:    true,
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
	pageCount := 0

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	allResults, err := utils.ReadWithPages(
		func(pageID string, maxResults int32) ([]dhcp.Fixedaddress, string, error) {

			if !data.MaxResults.IsNull() {
				maxResults = data.MaxResults.ValueInt32()
			}
			var paging int32 = 1
			if !data.Paging.IsNull() {
				paging = data.Paging.ValueInt32()
			}

			//Increment the page count
			pageCount++

			if !data.Body.IsNull() && !data.Body.IsUnknown() {
				// If body is provided, Update Call will be used to retrieve the data.
				request := d.client.DHCPAPI.
					FixedaddressAPI.
					MsServerUpdate(ctx).
					FixedAddressStruct(*ExpandFixedAddressStruct(ctx, data.Body, &resp.Diagnostics)).
					Method("GET").
					ReturnAsObject(1).
					ReturnFieldsPlus(readableAttributesForFixedaddress).
					Paging(paging).
					MaxResults(maxResults)

				// Add page ID if provided
				if pageID != "" {
					request = request.PageId(pageID)
				}

				// Execute the request
				apiRes, _, err := request.Execute()
				if err != nil {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Fixedaddress by extattrs, got error: %s", err))
					return nil, "", err
				}

				res := apiRes.ListFixedaddressResponseObject.GetResult()
				tflog.Info(ctx, fmt.Sprintf("Page %d : Retrieved %d results", pageCount, len(res)))

				// Check for next page ID in additional properties
				additionalProperties := apiRes.ListFixedaddressResponseObject.AdditionalProperties
				var nextPageID string
				npId, ok := additionalProperties["next_page_id"]
				if ok {
					if npIdStr, ok := npId.(string); ok {
						nextPageID = npIdStr
					}
				} else {
					tflog.Info(ctx, "No next page ID found. This is the last page.")
				}
				return res, nextPageID, nil

			}

			request := d.client.DHCPAPI.
				FixedaddressAPI.
				List(ctx).
				Filters(flex.ExpandFrameworkMapString(ctx, data.Filters, &resp.Diagnostics)).
				Extattrfilter(flex.ExpandFrameworkMapString(ctx, data.ExtAttrFilters, &resp.Diagnostics)).
				ReturnAsObject(1).
				ReturnFieldsPlus(readableAttributesForFixedaddress).
				Paging(paging).
				MaxResults(maxResults)

			// Add page ID if provided
			if pageID != "" {
				request = request.PageId(pageID)
			}

			// Execute the request
			apiRes, _, err := request.Execute()
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Fixedaddress by extattrs, got error: %s", err))
				return nil, "", err
			}

			res := apiRes.ListFixedaddressResponseObject.GetResult()
			tflog.Info(ctx, fmt.Sprintf("Page %d : Retrieved %d results", pageCount, len(res)))

			// Check for next page ID in additional properties
			additionalProperties := apiRes.ListFixedaddressResponseObject.AdditionalProperties
			var nextPageID string
			npId, ok := additionalProperties["next_page_id"]
			if ok {
				if npIdStr, ok := npId.(string); ok {
					nextPageID = npIdStr
				}
			} else {
				tflog.Info(ctx, "No next page ID found. This is the last page.")
			}
			return res, nextPageID, nil
		},
	)

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Fixedaddress, got error: %s", err))
		return
	}
	tflog.Info(ctx, fmt.Sprintf("Query complete: Total Number of Pages %d : Total results retrieved %d", pageCount, len(allResults)))

	// Process the results
	data.FlattenResults(ctx, allResults, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
