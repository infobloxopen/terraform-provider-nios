package parentalcontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ParentalcontrolAvpModel struct {
	Ref          types.String `tfsdk:"ref"`
	Comment      types.String `tfsdk:"comment"`
	DomainTypes  types.List   `tfsdk:"domain_types"`
	IsRestricted types.Bool   `tfsdk:"is_restricted"`
	Name         types.String `tfsdk:"name"`
	Type         types.Int64  `tfsdk:"type"`
	UserDefined  types.Bool   `tfsdk:"user_defined"`
	ValueType    types.String `tfsdk:"value_type"`
	VendorId     types.Int64  `tfsdk:"vendor_id"`
	VendorType   types.Int64  `tfsdk:"vendor_type"`
}

var ParentalcontrolAvpAttrTypes = map[string]attr.Type{
	"ref":           types.StringType,
	"comment":       types.StringType,
	"domain_types":  types.ListType{ElemType: types.StringType},
	"is_restricted": types.BoolType,
	"name":          types.StringType,
	"type":          types.Int64Type,
	"user_defined":  types.BoolType,
	"value_type":    types.StringType,
	"vendor_id":     types.Int64Type,
	"vendor_type":   types.Int64Type,
}

var ParentalcontrolAvpResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "The comment for the AVP.",
	},
	"domain_types": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.ValueStringsAre(
				stringvalidator.OneOf(
					"ANCILLARY",
					"IP_SPACE_DIS",
					"NAS_CONTEXT",
					"SUBS_ID")),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The list of domains applicable to AVP.",
	},
	"is_restricted": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if AVP is restricted to domains.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of AVP.",
	},
	"type": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 255),
		},
		MarkdownDescription: "The type of AVP as per RFC 2865/2866.",
	},
	"user_defined": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if AVP was defined by user.",
	},
	"value_type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("BYTE", "DATE", "INTEGER", "INTEGER64", "IPADDR", "IPV6ADDR", "IPV6IFID", "IPV6PREFIX", "OCTETS", "SHORT", "STRING"),
		},
		MarkdownDescription: "The type of value.",
	},
	"vendor_id": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The vendor ID as per RFC 2865/2866.",
	},
	"vendor_type": schema.Int64Attribute{
		Optional: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 255),
		},
		MarkdownDescription: "The vendor type as per RFC 2865/2866.",
	},
}

func ExpandParentalcontrolAvp(ctx context.Context, o types.Object, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolAvp {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ParentalcontrolAvpModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ParentalcontrolAvpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolAvp {
	if m == nil {
		return nil
	}
	to := &parentalcontrol.ParentalcontrolAvp{
		Comment:      flex.ExpandStringPointer(m.Comment),
		DomainTypes:  flex.ExpandFrameworkListString(ctx, m.DomainTypes, diags),
		IsRestricted: flex.ExpandBoolPointer(m.IsRestricted),
		Name:         flex.ExpandStringPointer(m.Name),
		Type:         flex.ExpandInt64Pointer(m.Type),
		ValueType:    flex.ExpandStringPointer(m.ValueType),
		VendorId:     flex.ExpandInt64Pointer(m.VendorId),
		VendorType:   flex.ExpandInt64Pointer(m.VendorType),
	}
	return to
}

func FlattenParentalcontrolAvp(ctx context.Context, from *parentalcontrol.ParentalcontrolAvp, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ParentalcontrolAvpAttrTypes)
	}
	m := ParentalcontrolAvpModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ParentalcontrolAvpAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ParentalcontrolAvpModel) Flatten(ctx context.Context, from *parentalcontrol.ParentalcontrolAvp, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ParentalcontrolAvpModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DomainTypes = flex.FlattenFrameworkListString(ctx, from.DomainTypes, diags)
	m.IsRestricted = types.BoolPointerValue(from.IsRestricted)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Type = flex.FlattenInt64Pointer(from.Type)
	m.UserDefined = types.BoolPointerValue(from.UserDefined)
	m.ValueType = flex.FlattenStringPointer(from.ValueType)
	m.VendorId = flex.FlattenInt64Pointer(from.VendorId)
	m.VendorType = flex.FlattenInt64Pointer(from.VendorType)
}
