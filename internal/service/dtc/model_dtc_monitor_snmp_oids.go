package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DtcMonitorSnmpOidsModel struct {
	Oid       types.String `tfsdk:"oid"`
	Comment   types.String `tfsdk:"comment"`
	Type      types.String `tfsdk:"type"`
	Condition types.String `tfsdk:"condition"`
	First     types.String `tfsdk:"first"`
	Last      types.String `tfsdk:"last"`
}

var DtcMonitorSnmpOidsAttrTypes = map[string]attr.Type{
	"oid":       types.StringType,
	"comment":   types.StringType,
	"type":      types.StringType,
	"condition": types.StringType,
	"first":     types.StringType,
	"last":      types.StringType,
}

var DtcMonitorSnmpOidsResourceSchemaAttributes = map[string]schema.Attribute{
	"oid": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The SNMP OID value for DTC SNMP Monitor health checks.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The comment for a DTC SNMP Health Monitor OID object.",
	},
	"type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("STRING"),
		Validators: []validator.String{
			stringvalidator.OneOf("STRING", "INTEGER"),
		},
		MarkdownDescription: "The value of the condition type for DTC SNMP Monitor health check results.",
	},
	"condition": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("ANY"),
		Validators: []validator.String{
			stringvalidator.OneOf("ANY", "EXACT", "LEQ", "GEQ", "RANGE"),
		},
		MarkdownDescription: "The condition of the validation result for an SNMP health check. The following conditions can be applied to the health check results: 'ANY' accepts any response; 'EXACT' accepts result equal to 'first'; 'LEQ' accepts result which is less than 'first'; 'GEQ' accepts result which is greater than 'first'; 'RANGE' accepts result value of which is between 'first' and 'last'.",
	},
	"first": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The condition's first term to match against the SNMP health check result.",
	},
	"last": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The condition's second term to match against the SNMP health check result with 'RANGE' condition.",
	},
}

func ExpandDtcMonitorSnmpOids(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dtc.DtcMonitorSnmpOids {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DtcMonitorSnmpOidsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DtcMonitorSnmpOidsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcMonitorSnmpOids {
	if m == nil {
		return nil
	}
	to := &dtc.DtcMonitorSnmpOids{
		Oid:       flex.ExpandStringPointer(m.Oid),
		Comment:   flex.ExpandStringPointer(m.Comment),
		Type:      flex.ExpandStringPointer(m.Type),
		Condition: flex.ExpandStringPointer(m.Condition),
		First:     flex.ExpandStringPointer(m.First),
		Last:      flex.ExpandStringPointer(m.Last),
	}
	return to
}

func FlattenDtcMonitorSnmpOids(ctx context.Context, from *dtc.DtcMonitorSnmpOids, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcMonitorSnmpOidsAttrTypes)
	}
	m := DtcMonitorSnmpOidsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcMonitorSnmpOidsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcMonitorSnmpOidsModel) Flatten(ctx context.Context, from *dtc.DtcMonitorSnmpOids, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcMonitorSnmpOidsModel{}
	}
	m.Oid = flex.FlattenStringPointer(from.Oid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.Condition = flex.FlattenStringPointer(from.Condition)
	m.First = flex.FlattenStringPointer(from.First)
	m.Last = flex.FlattenStringPointer(from.Last)
}
