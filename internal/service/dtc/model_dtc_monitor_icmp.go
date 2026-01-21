package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type DtcMonitorIcmpModel struct {
	Ref         types.String `tfsdk:"ref"`
	Uuid        types.String `tfsdk:"uuid"`
	Comment     types.String `tfsdk:"comment"`
	ExtAttrs    types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll types.Map    `tfsdk:"extattrs_all"`
	Interval    types.Int64  `tfsdk:"interval"`
	Name        types.String `tfsdk:"name"`
	RetryDown   types.Int64  `tfsdk:"retry_down"`
	RetryUp     types.Int64  `tfsdk:"retry_up"`
	Timeout     types.Int64  `tfsdk:"timeout"`
}

var DtcMonitorIcmpAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
	"uuid":         types.StringType,
	"comment":      types.StringType,
	"extattrs":     types.MapType{ElemType: types.StringType},
	"extattrs_all": types.MapType{ElemType: types.StringType},
	"interval":     types.Int64Type,
	"name":         types.StringType,
	"retry_down":   types.Int64Type,
	"retry_up":     types.Int64Type,
	"timeout":      types.Int64Type,
}

var DtcMonitorIcmpResourceSchemaAttributes = map[string]schema.Attribute{
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
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Comment for this DTC monitor; maximum 256 characters.",
	},
	"extattrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(5),
		MarkdownDescription: "The interval for ICMP health check.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The display name for this DTC monitor.",
	},
	"retry_down": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "The value of how many times the server should appear as down to be treated as dead after it was alive.",
	},
	"retry_up": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1),
		MarkdownDescription: "The value of how many times the server should appear as up to be treated as alive after it was dead.",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(15),
		MarkdownDescription: "The timeout for ICMP health check in seconds.",
	},
}

func (m *DtcMonitorIcmpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dtc.DtcMonitorIcmp {
	if m == nil {
		return nil
	}
	to := &dtc.DtcMonitorIcmp{
		Comment:   flex.ExpandStringPointer(m.Comment),
		ExtAttrs:  ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Interval:  flex.ExpandInt64Pointer(m.Interval),
		Name:      flex.ExpandStringPointer(m.Name),
		RetryDown: flex.ExpandInt64Pointer(m.RetryDown),
		RetryUp:   flex.ExpandInt64Pointer(m.RetryUp),
		Timeout:   flex.ExpandInt64Pointer(m.Timeout),
	}
	return to
}

func FlattenDtcMonitorIcmp(ctx context.Context, from *dtc.DtcMonitorIcmp, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DtcMonitorIcmpAttrTypes)
	}
	m := DtcMonitorIcmpModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, DtcMonitorIcmpAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DtcMonitorIcmpModel) Flatten(ctx context.Context, from *dtc.DtcMonitorIcmp, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DtcMonitorIcmpModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Interval = flex.FlattenInt64Pointer(from.Interval)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.RetryDown = flex.FlattenInt64Pointer(from.RetryDown)
	m.RetryUp = flex.FlattenInt64Pointer(from.RetryUp)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
}
