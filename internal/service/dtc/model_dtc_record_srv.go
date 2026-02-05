package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type DtcRecordSrvModel struct {
	Ref       types.String `tfsdk:"ref"`
	Uuid      types.String `tfsdk:"uuid"`
	Comment   types.String `tfsdk:"comment"`
	Disable   types.Bool   `tfsdk:"disable"`
	DtcServer types.String `tfsdk:"dtc_server"`
	Name      types.String `tfsdk:"name"`
	Port      types.Int64  `tfsdk:"port"`
	Priority  types.Int64  `tfsdk:"priority"`
	Target    types.String `tfsdk:"target"`
	Ttl       types.Int64  `tfsdk:"ttl"`
	UseTtl    types.Bool   `tfsdk:"use_ttl"`
	Weight    types.Int64  `tfsdk:"weight"`
}

var DtcRecordSrvAttrTypes = map[string]attr.Type{
	"ref":        types.StringType,
	"uuid":       types.StringType,
	"comment":    types.StringType,
	"disable":    types.BoolType,
	"dtc_server": types.StringType,
	"name":       types.StringType,
	"port":       types.Int64Type,
	"priority":   types.Int64Type,
	"target":     types.StringType,
	"ttl":        types.Int64Type,
	"use_ttl":    types.BoolType,
	"weight":     types.Int64Type,
}

var DtcRecordSrvResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
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
	"name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name for an SRV record in unicode format.",
	},
	"port": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The port of the SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"priority": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The priority of the SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
	"target": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "The target of the SRV record in FQDN format. This value can be in unicode format.",
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
	"weight": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 65535),
		},
		MarkdownDescription: "The weight of the SRV record. Valid values are from 0 to 65535 (inclusive), in 32-bit unsigned integer format.",
	},
}

func (m *DtcRecordSrvModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dtc.DtcRecordSrv {
	if m == nil {
		return nil
	}
	to := &dtc.DtcRecordSrv{
		Comment:  flex.ExpandStringPointer(m.Comment),
		Disable:  flex.ExpandBoolPointer(m.Disable),
		Name:     flex.ExpandStringPointer(m.Name),
		Port:     flex.ExpandInt64Pointer(m.Port),
		Priority: flex.ExpandInt64Pointer(m.Priority),
		Target:   flex.ExpandStringPointer(m.Target),
		Ttl:      flex.ExpandInt64Pointer(m.Ttl),
		UseTtl:   flex.ExpandBoolPointer(m.UseTtl),
		Weight:   flex.ExpandInt64Pointer(m.Weight),
	}
	if isCreate {
		to.DtcServer = flex.ExpandStringPointer(m.DtcServer)
	}
	return to
}

func FlattenDtcRecordSrv(ctx context.Context, from *dtc.DtcRecordSrv, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcRecordSrvAttrTypes)
	}
	m := DtcRecordSrvModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DtcRecordSrvAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcRecordSrvModel) Flatten(ctx context.Context, from *dtc.DtcRecordSrv, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcRecordSrvModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DtcServer = flex.FlattenStringPointer(from.DtcServer)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Priority = flex.FlattenInt64Pointer(from.Priority)
	m.Target = flex.FlattenStringPointer(from.Target)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.Weight = flex.FlattenInt64Pointer(from.Weight)
}
