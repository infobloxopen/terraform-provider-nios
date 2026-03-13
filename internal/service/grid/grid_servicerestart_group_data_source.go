package grid

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
	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GridServicerestartGroupDataSource{}

func NewGridServicerestartGroupDataSource() datasource.DataSource {
	return &GridServicerestartGroupDataSource{}
}

// GridServicerestartGroupDataSource defines the data source implementation.
type GridServicerestartGroupDataSource struct {
	client *niosclient.APIClient
}

func (d *GridServicerestartGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "grid_servicerestart_group"
}

type GridServicerestartGroupModelWithFilter struct {
	Filters        types.Map   `tfsdk:"filters"`
	ExtAttrFilters types.Map   `tfsdk:"extattrfilters"`
	Result         types.List  `tfsdk:"result"`
	MaxResults     types.Int32 `tfsdk:"max_results"`
	Paging         types.Int32 `tfsdk:"paging"`
}

func (m *GridServicerestartGroupModelWithFilter) FlattenResults(ctx context.Context, from []grid.GridServicerestartGroup, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Result = flex.FlattenFrameworkListNestedBlock(ctx, from, GridServicerestartGroupAttrTypes, diags, FlattenGridServicerestartGroup)
}

func (d *GridServicerestartGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves information about existing GRID Service Restart Groups definitions.",
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
					Attributes: utils.DataSourceAttributeMap(GridServicerestartGroupResourceSchemaAttributes, &resp.Diagnostics),
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
		},
	}
}

func (d *GridServicerestartGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *GridServicerestartGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GridServicerestartGroupModelWithFilter
	pageCount := 0

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	allResults, err := utils.ReadWithPages(
		func(pageID string, maxResults int32) ([]grid.GridServicerestartGroup, string, error) {

			if !data.MaxResults.IsNull() {
				maxResults = data.MaxResults.ValueInt32()
			}
			var paging int32 = 1
			if !data.Paging.IsNull() {
				paging = data.Paging.ValueInt32()
			}

			//Increment the page count
			pageCount++

			request := d.client.GridAPI.
				GridServicerestartGroupAPI.
				List(ctx).
				Filters(flex.ExpandFrameworkMapString(ctx, data.Filters, &resp.Diagnostics)).
				Extattrfilter(flex.ExpandFrameworkMapString(ctx, data.ExtAttrFilters, &resp.Diagnostics)).
				ReturnAsObject(1).
				ReturnFieldsPlus(readableAttributesForGridServicerestartGroup).
				Paging(paging).
				MaxResults(maxResults).
				ProxySearch(config.GetProxySearch())

			// Add page ID if provided
			if pageID != "" {
				request = request.PageId(pageID)
			}

			// Execute the request
			apiRes, _, err := request.Execute()
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read GridServicerestartGroup by extattrs, got error: %s", err))
				return nil, "", err
			}

			res := apiRes.ListGridServicerestartGroupResponseObject.GetResult()
			tflog.Info(ctx, fmt.Sprintf("Page %d : Retrieved %d results", pageCount, len(res)))

			// Check for next page ID in additional properties
			additionalProperties := apiRes.ListGridServicerestartGroupResponseObject.AdditionalProperties
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read GridServicerestartGroup, got error: %s", err))
		return
	}
	tflog.Info(ctx, fmt.Sprintf("Query complete: Total Number of Pages %d : Total results retrieved %d", pageCount, len(allResults)))

	// Process the results
	data.FlattenResults(ctx, allResults, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
