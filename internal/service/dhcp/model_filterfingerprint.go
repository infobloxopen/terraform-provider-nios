package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
)

type FilterfingerprintModel struct {
	Ref         types.String                     `tfsdk:"ref"`
	Uuid        types.String                     `tfsdk:"uuid"`
	Comment     types.String                     `tfsdk:"comment"`
	ExtAttrs    types.Map                        `tfsdk:"extattrs"`
	ExtAttrsAll types.Map                        `tfsdk:"extattrs_all"`
	Fingerprint internaltypes.UnorderedListValue `tfsdk:"fingerprint"`
	Name        types.String                     `tfsdk:"name"`
}

var FilterfingerprintAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
	"uuid":         types.StringType,
	"comment":      types.StringType,
	"extattrs":     types.MapType{ElemType: types.StringType},
	"extattrs_all": types.MapType{ElemType: types.StringType},
	"fingerprint":  internaltypes.UnorderedListOfStringType,
	"name":         types.StringType,
}

var FilterfingerprintResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"comment": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "The descriptive comment.",
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
		Computed: true,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
		MarkdownDescription: "Extensible attributes associated with the object, including default and internal attributes.",
		ElementType:         types.StringType,
	},
	"fingerprint": schema.ListAttribute{
		CustomType:  internaltypes.UnorderedListOfStringType,
		ElementType: types.StringType,
		Required:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of DHCP Fingerprint objects.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of a DHCP Fingerprint Filter object.",
	},
}

func (m *FilterfingerprintModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Filterfingerprint {
	if m == nil {
		return nil
	}
	to := &dhcp.Filterfingerprint{
		Comment:     flex.ExpandStringPointer(m.Comment),
		ExtAttrs:    ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Fingerprint: flex.ExpandFrameworkListString(ctx, m.Fingerprint, diags),
		Name:        flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenFilterfingerprint(ctx context.Context, from *dhcp.Filterfingerprint, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FilterfingerprintAttrTypes)
	}
	m := FilterfingerprintModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, FilterfingerprintAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FilterfingerprintModel) Flatten(ctx context.Context, from *dhcp.Filterfingerprint, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FilterfingerprintModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Fingerprint = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.Fingerprint, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
}
