package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type RecordNsModel struct {
	Ref              types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	Addresses        types.List   `tfsdk:"addresses"`
	CloudInfo        types.Object `tfsdk:"cloud_info"`
	Creator          types.String `tfsdk:"creator"`
	DnsName          types.String `tfsdk:"dns_name"`
	LastQueried      types.Int64  `tfsdk:"last_queried"`
	MsDelegationName types.String `tfsdk:"ms_delegation_name"`
	Name             types.String `tfsdk:"name"`
	Nameserver       types.String `tfsdk:"nameserver"`
	Policy           types.String `tfsdk:"policy"`
	View             types.String `tfsdk:"view"`
	Zone             types.String `tfsdk:"zone"`
}

var RecordNsAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
    "uuid":        types.StringType,
	"addresses":          types.ListType{ElemType: types.ObjectType{AttrTypes: RecordNsAddressesAttrTypes}},
	"cloud_info":         types.ObjectType{AttrTypes: RecordNsCloudInfoAttrTypes},
	"creator":            types.StringType,
	"dns_name":           types.StringType,
	"last_queried":       types.Int64Type,
	"ms_delegation_name": types.StringType,
	"name":               types.StringType,
	"nameserver":         types.StringType,
	"policy":             types.StringType,
	"view":               types.StringType,
	"zone":               types.StringType,
}

var RecordNsResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"addresses": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordNsAddressesResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of zone name servers.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RecordNsCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The cloud information associated with the record.",
	},
	"creator": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The record creator.",
	},
	"dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the NS record in punycode format.",
	},
	"last_queried": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the last DNS query in Epoch seconds format.",
	},
	"ms_delegation_name": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The MS delegation point name.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "The name of the NS record in FQDN format. This value can be in unicode format.",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"nameserver": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The domain name of an authoritative server for the redirected zone.",
	},
	"policy": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The host name policy for the record.",
	},
	"view": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The name of the DNS view in which the record resides. Example: \"external\".",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the zone in which the record resides. Example: \"zone.com\". If a view is not specified when searching by zone, the default view is used.",
	},
}

func (m *RecordNsModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.RecordNs {
	if m == nil {
		return nil
	}
	to := &dns.RecordNs{
		Addresses:        flex.ExpandFrameworkListNestedBlock(ctx, m.Addresses, diags, ExpandRecordNsAddresses),
		MsDelegationName: flex.ExpandStringPointer(m.MsDelegationName),
		Nameserver:       flex.ExpandStringPointer(m.Nameserver),
	}
	if isCreate {
		to.Name = flex.ExpandStringPointer(m.Name)
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func FlattenRecordNs(ctx context.Context, from *dns.RecordNs, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordNsAttrTypes)
	}
	m := RecordNsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordNsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordNsModel) Flatten(ctx context.Context, from *dns.RecordNs, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordNsModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Addresses = flex.FlattenFrameworkListNestedBlock(ctx, from.Addresses, RecordNsAddressesAttrTypes, diags, FlattenRecordNsAddresses)
	m.CloudInfo = FlattenRecordNsCloudInfo(ctx, from.CloudInfo, diags)
	m.Creator = flex.FlattenStringPointer(from.Creator)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.MsDelegationName = flex.FlattenStringPointer(from.MsDelegationName)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Nameserver = flex.FlattenStringPointer(from.Nameserver)
	m.Policy = flex.FlattenStringPointer(from.Policy)
	m.View = flex.FlattenStringPointer(from.View)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
