package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DtcRecordAModel struct {
	Ref         types.String `tfsdk:"ref"`
	AutoCreated types.String `tfsdk:"auto_created"`
	Comment     types.String `tfsdk:"comment"`
	Disable     types.Bool   `tfsdk:"disable"`
	DtcServer   types.String `tfsdk:"dtc_server"`
	Ipv4addr    types.String `tfsdk:"ipv4addr"`
	Ttl         types.Int64  `tfsdk:"ttl"`
	UseTtl      types.Bool   `tfsdk:"use_ttl"`
}

var DtcRecordAAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
	"auto_created": types.StringType,
	"comment":      types.StringType,
	"disable":      types.BoolType,
	"dtc_server":   types.StringType,
	"ipv4addr":     types.StringType,
	"ttl":          types.Int64Type,
	"use_ttl":      types.BoolType,
}

var DtcRecordAResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"auto_created": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Flag that indicates whether this record was automatically created by NIOS.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Comment for the record; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the record is disabled or not. False means that the record is enabled.",
	},
	"dtc_server": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the DTC Server object with which the DTC record is associated.",
	},
	"ipv4addr": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 Address of the domain name.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Time to Live (TTL) value.",
	},
	"use_ttl": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ttl",
	},
}

func ExpandDtcRecordA(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcRecordA {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcRecordAModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcRecordAModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcRecordA {
	if m == nil {
		return nil
	}
	to := &dtc.DtcRecordA{
		Comment:   flex.ExpandStringPointer(m.Comment),
		Disable:   flex.ExpandBoolPointer(m.Disable),
		DtcServer: flex.ExpandStringPointer(m.DtcServer),
		Ipv4addr:  flex.ExpandStringPointer(m.Ipv4addr),
		Ttl:       flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:    flex.ExpandBoolPointer(m.UseTtl),
	}
	return to
}

func FlattenDtcRecordA(ctx context.Context, from *dtc.DtcRecordA, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcRecordAAttrTypes)
	}
	m := DtcRecordAModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcRecordAAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcRecordAModel) Flatten(ctx context.Context, from *dtc.DtcRecordA, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcRecordAModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AutoCreated = flex.FlattenStringPointer(from.AutoCreated)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DtcServer = flex.FlattenStringPointer(from.DtcServer)
	m.Ipv4addr = flex.FlattenStringPointer(from.Ipv4addr)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
}
