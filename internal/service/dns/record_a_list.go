package dns

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ list.ListResource = &RecordAList{}
var _ list.ListResourceWithConfigure = &RecordAList{}

func NewRecordAList() list.ListResource {
	return &RecordAList{}
}

// RecordAList defines the data source implementation.
type RecordAList struct {
	client *niosclient.APIClient
}

func (l *RecordAList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dns_record_a"
}

func (l *RecordAList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type RecordAListModel struct {
	Filters        types.Map `tfsdk:"filters"`
	ExtAttrFilters types.Map `tfsdk:"extattrfilters"`
}

func (l *RecordAList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves information about existing DNS A Records.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.MapAttribute{
				MarkdownDescription: "Filter parameters for querying DNS A records.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"extattrfilters": schema.MapAttribute{
				MarkdownDescription: "Extensible attribute filters for querying DNS A records.",
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (l *RecordAList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var data RecordAListModel

	diags := req.Config.Get(ctx, &data)
	if diags.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(diags)
		return
	}

	apiRes, _, err := l.client.DNSAPI.RecordAAPI.
		List(ctx).
		Filters(flex.ExpandFrameworkMapString(ctx, data.Filters, &diags)).
		Extattrfilter(flex.ExpandFrameworkMapString(ctx, data.ExtAttrFilters, &diags)).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForRecordA).
		MaxResults(int32(req.Limit)).
		Execute()

	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to list RecordA, got error: %s", err))
		stream.Results = list.ListResultsStreamDiagnostics(diags)
		return
	}

	if apiRes == nil || apiRes.ListRecordAResponseObject == nil {
		// No results found, return empty stream
		stream.Results = func(push func(list.ListResult) bool) {}
		return
	}

	res := apiRes.ListRecordAResponseObject.GetResult()
	if res == nil {
		// No results found, return empty stream
		stream.Results = func(push func(list.ListResult) bool) {}
		return
	}

	stream.Results = func(push func(list.ListResult) bool) {
		for _, item := range res {
			result := req.NewListResult(ctx)

			//result.Diagnostics.Append(result.Identity.Set(ctx, identityData)...)
			result.Diagnostics.Append(result.Identity.SetAttribute(ctx, path.Root("ref"), &item.Ref)...)
			if result.Diagnostics.HasError() {
				if !push(result) {
					return
				}
				continue
			}

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
				result1 := FlattenRecordA(ctx, &item, &result.Diagnostics)
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
