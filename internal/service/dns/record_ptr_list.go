package dns

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
	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ list.ListResource = &RecordPtrList{}
var _ list.ListResourceWithConfigure = &RecordPtrList{}

func NewRecordPtrList() list.ListResource {
	return &RecordPtrList{}
}

// RecordPtrList defines the List implementation.
type RecordPtrList struct {
	client *niosclient.APIClient
}

func (l *RecordPtrList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dns_record_ptr"
}

func (l *RecordPtrList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type RecordPtrListModel struct {
	Filters        types.Map `tfsdk:"filters"`
	ExtAttrFilters types.Map `tfsdk:"extattrfilters"`
}

func (l *RecordPtrList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Query existing DNS PTR Records.",
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

func (l *RecordPtrList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var data RecordPtrListModel
	limit := int32(req.Limit)

	diags := req.Config.Get(ctx, &data)
	if diags.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(diags)
		return
	}

	baseFilters := flex.ExpandFrameworkMapString(ctx, data.Filters, &diags)
	if baseFilters == nil {
		baseFilters = make(map[string]interface{})
	}

	// To exclude SYSTEM records we query STATIC and DYNAMIC separately.
	_, userSetCreator := baseFilters["creator"]

	creatorsToFetch := []string{"STATIC", "DYNAMIC"}
	if userSetCreator {
		creatorsToFetch = []string{""} // empty string = use baseFilters as-is
	}

	var allResults []dns.RecordPtr
	remaining := limit

	for _, creator := range creatorsToFetch {
		if remaining <= 0 {
			break
		}

		// Build a per-creator copy of the filters map.
		iterFilters := make(map[string]interface{}, len(baseFilters))
		for k, v := range baseFilters {
			iterFilters[k] = v
		}
		if creator != "" {
			iterFilters["creator"] = creator
		}

		pageCount := 0
		var totalFetched int32
		iterLimit := remaining // each creator call gets its share of the remaining limit

		results, err := utils.ReadWithPages(
			func(pageID string, maxResultsPerPage int32) ([]dns.RecordPtr, string, error) {
				var paging int32 = 1

				// Adjust page size to not fetch more than the remaining needed results.
				if r := iterLimit - totalFetched; r < maxResultsPerPage {
					maxResultsPerPage = r
				}

				pageCount++

				request := l.client.DNSAPI.
					RecordPtrAPI.
					List(ctx).
					Filters(iterFilters).
					Extattrfilter(flex.ExpandFrameworkMapString(ctx, data.ExtAttrFilters, &diags)).
					ReturnAsObject(1).
					ReturnFieldsPlus(readableAttributesForRecordPtr).
					Paging(paging).
					MaxResults(maxResultsPerPage)

				if pageID != "" {
					request = request.PageId(pageID)
				}

				apiRes, _, err := request.Execute()
				if err != nil {
					return nil, "", err
				}

				res := apiRes.ListRecordPtrResponseObject.GetResult()
				tflog.Info(ctx, fmt.Sprintf("Creator=%q Page %d: retrieved %d results", creator, pageCount, len(res)))

				totalFetched += int32(len(res))

				// If the limit for this creator is reached, stop pagination.
				if totalFetched >= iterLimit {
					tflog.Info(ctx, "Limit reached, stopped fetching more pages.")
					return res, "", nil
				}

				additionalProperties := apiRes.ListRecordPtrResponseObject.AdditionalProperties
				var nextPageID string
				if npId, ok := additionalProperties["next_page_id"]; ok {
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
			diags.AddError("Client Error", fmt.Sprintf("Unable to list RecordPtr, got error: %s", err))
			stream.Results = list.ListResultsStreamDiagnostics(diags)
			return
		}

		allResults = append(allResults, results...)
		remaining -= int32(len(results))
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
				if item.ExtAttrs != nil {
					delete(*item.ExtAttrs, terraformInternalIDEA)
				}
				result1 := FlattenRecordPtr(ctx, &item, &result.Diagnostics)
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
