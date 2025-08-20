package security

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type AdminroleModel struct {
	Ref         types.String `tfsdk:"ref"`
	Comment     types.String `tfsdk:"comment"`
	Disable     types.Bool   `tfsdk:"disable"`
	ExtAttrs    types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll types.Map    `tfsdk:"extattrs_all"`
	Name        types.String `tfsdk:"name"`
}

var AdminroleAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
	"comment":      types.StringType,
	"disable":      types.BoolType,
	"extattrs":     types.MapType{ElemType: types.StringType},
	"extattrs_all": types.MapType{ElemType: types.StringType},
	"name":         types.StringType,
}

var AdminroleResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The descriptive comment of the Admin Role object.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The disable flag.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object.",
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The name of an admin role.",
	},
}

func (m *AdminroleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.Adminrole {
	if m == nil {
		return nil
	}
	to := &security.Adminrole{
		Comment:  flex.ExpandStringPointer(m.Comment),
		Disable:  flex.ExpandBoolPointer(m.Disable),
		ExtAttrs: ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:     flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenAdminrole(ctx context.Context, from *security.Adminrole, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdminroleAttrTypes)
	}
	m := AdminroleModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, AdminroleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdminroleModel) Flatten(ctx context.Context, from *security.Adminrole, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdminroleModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
}
