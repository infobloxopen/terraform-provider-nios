package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type DtcRecordAaaaModel struct {
	Ref         types.String        `tfsdk:"ref"`
	AutoCreated types.Bool          `tfsdk:"auto_created"`
	Comment     types.String        `tfsdk:"comment"`
	Disable     types.Bool          `tfsdk:"disable"`
	DtcServer   types.String        `tfsdk:"dtc_server"`
	Ipv6addr    iptypes.IPv6Address `tfsdk:"ipv6addr"`
	Ttl         types.Int64         `tfsdk:"ttl"`
	UseTtl      types.Bool          `tfsdk:"use_ttl"`
}

var DtcRecordAaaaAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
	"auto_created": types.BoolType,
	"comment":      types.StringType,
	"disable":      types.BoolType,
	"dtc_server":   types.StringType,
	"ipv6addr":     iptypes.IPv6AddressType{},
	"ttl":          types.Int64Type,
	"use_ttl":      types.BoolType,
}

var DtcRecordAaaaResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"auto_created": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag that indicates whether this record was automatically created by NIOS.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for the record; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the record is disabled or not. False means that the record is enabled.",
	},
	"dtc_server": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		MarkdownDescription: "The name of the DTC Server object with which the DTC record is associated.",
	},
	"ipv6addr": schema.StringAttribute{
		CustomType:          iptypes.IPv6AddressType{},
		Required:            true,
		MarkdownDescription: "The IPv6 Address of the domain name.",
	},
	"ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ttl")),
		},
		MarkdownDescription: "The Time to Live (TTL) value.",
	},
	"use_ttl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ttl",
	},
}

func (m *DtcRecordAaaaModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dtc.DtcRecordAaaa {
	if m == nil {
		return nil
	}
	to := &dtc.DtcRecordAaaa{
		Comment:  flex.ExpandStringPointer(m.Comment),
		Disable:  flex.ExpandBoolPointer(m.Disable),
		Ipv6addr: flex.ExpandIPv6Address(m.Ipv6addr),
		Ttl:      flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:   flex.ExpandBoolPointer(m.UseTtl),
	}
	if isCreate {
		to.DtcServer = flex.ExpandStringPointer(m.DtcServer)
	}
	return to
}

func FlattenDtcRecordAaaa(ctx context.Context, from *dtc.DtcRecordAaaa, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcRecordAaaaAttrTypes)
	}
	m := DtcRecordAaaaModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcRecordAaaaAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcRecordAaaaModel) Flatten(ctx context.Context, from *dtc.DtcRecordAaaa, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcRecordAaaaModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AutoCreated = types.BoolPointerValue(from.AutoCreated)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DtcServer = flex.FlattenStringPointer(from.DtcServer)
	m.Ipv6addr = flex.FlattenIPv6Address(from.Ipv6addr)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
}
