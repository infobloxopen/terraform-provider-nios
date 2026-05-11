package dhcp

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ list.ListResource = &FixedaddressList{}
var _ list.ListResourceWithConfigure = &FixedaddressList{}

func NewFixedaddressList() list.ListResource {
	return &FixedaddressList{}
}

// FixedaddressList defines the List implementation.
type FixedaddressList struct {
	client *niosclient.APIClient
}

func (l *FixedaddressList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_fixed_address"
}

func (l *FixedaddressList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*niosclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected List Resource Configure Type",
			fmt.Sprintf("Expected *niosclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	l.client = client
}

type FixedaddressListModel struct {
	Filters        types.Map `tfsdk:"filters"`
	ExtAttrFilters types.Map `tfsdk:"extattrfilters"`
}

func (l *FixedaddressList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Query existing DHCP Fixed Addresses.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.MapAttribute{
				MarkdownDescription: "Filters are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"extattrfilters": schema.MapAttribute{
				MarkdownDescription: "External Attribute Filters are used to return a more specific list of results by filtering on external attributes. If you specify multiple filters, the results returned will have only resources that match all the specified filters.",
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (l *FixedaddressList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var data FixedaddressListModel
	pageCount := 0
	limit := int32(req.Limit)

	diags := req.Config.Get(ctx, &data)
	if diags.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(diags)
		return
	}

	allResults, err := utils.ReadWithPages(
		func(pageID string, maxResultsPerPage int32) ([]dhcp.Fixedaddress, string, error) {

			var paging int32 = 1

			// If total limit is set by user and is less than maxResultsPerPage, use it as maxResultsPerPage for API call to optimize the number of results.
			// If limit > maxResultsPerPage, terraform automatically breaks connection to the provider after limit is reached.
			if limit < maxResultsPerPage {
				maxResultsPerPage = limit
			}

			//Increment the page count
			pageCount++

			request := l.client.DHCPAPI.
				FixedaddressAPI.
				List(ctx).
				Filters(flex.ExpandFrameworkMapString(ctx, data.Filters, &diags)).
				Extattrfilter(flex.ExpandFrameworkMapString(ctx, data.ExtAttrFilters, &diags)).
				ReturnAsObject(1).
				ReturnFieldsPlus(readableAttributesForFixedaddress).
				Paging(paging).
				MaxResults(maxResultsPerPage)

			// Add page ID if provided
			if pageID != "" {
				request = request.PageId(pageID)
			}

			// Execute the request
			apiRes, _, err := request.Execute()
			if err != nil {
				return nil, "", err
			}

			res := apiRes.ListFixedaddressResponseObject.GetResult()
			tflog.Info(ctx, fmt.Sprintf("Page %d : Retrieved %d results", pageCount, len(res)))

			// Check for next page ID in additional properties
			additionalProperties := apiRes.ListFixedaddressResponseObject.AdditionalProperties
			var nextPageID string

			// If limit is reached , we do not need to continue to make API calls, we can return the results and empty nextPageID to stop pagination.
			if len(res) >= int(limit) {
				nextPageID = ""
				tflog.Info(ctx, "Limit reached, stopped fetching more pages.")
				return res, nextPageID, nil
			}

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
		diags.AddError("Client Error", fmt.Sprintf("Unable to list Fixedaddress, got error: %s", err))
		stream.Results = list.ListResultsStreamDiagnostics(diags)
		return
	}

	stream.Results = func(push func(list.ListResult) bool) {
		for _, item := range allResults {
			result := req.NewListResult(ctx)

			// Set the Identity for each result
			result.Diagnostics.Append(result.Identity.SetAttribute(ctx, path.Root("ref"), &item.Ref)...)
			if result.Diagnostics.HasError() {
				if !push(result) {
					return
				}
				continue
			}

			// By default, list only returns the identity.
			// If IncludeResource is true, it gets the full resource and sets it in the result.Resource
			if req.IncludeResource {
				var extAttrsAll types.Map
				item.ExtAttrs, extAttrsAll, diags = RemoveInheritedExtAttrs(ctx, extAttrsAll, *item.ExtAttrs)
				result.Diagnostics.Append(result.Resource.SetAttribute(ctx, path.Root("extattrs_all"), extAttrsAll)...)
				if result.Diagnostics.HasError() {
					if !push(result) {
						return
					}
					continue
				}
				result1 := FlattenFixedaddress(ctx, &item, &result.Diagnostics)
				result.Diagnostics.Append(result.Resource.Set(ctx, &result1)...)
				if result.Diagnostics.HasError() {
					if !push(result) {
						return
					}
					continue
				}

			}

			// Push the result to the stream
			if !push(result) {
				return
			}
		}
	}

}
