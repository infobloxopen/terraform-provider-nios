package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type MemberLomUsersModel struct {
	Name     types.String `tfsdk:"name"`
	Password types.String `tfsdk:"password"`
	Role     types.String `tfsdk:"role"`
	Disable  types.Bool   `tfsdk:"disable"`
	Comment  types.String `tfsdk:"comment"`
}

var MemberLomUsersAttrTypes = map[string]attr.Type{
	"name":     types.StringType,
	"password": types.StringType,
	"role":     types.StringType,
	"disable":  types.BoolType,
	"comment":  types.StringType,
}

var MemberLomUsersResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The LOM user name.",
	},
	"password": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The LOM user password.",
	},
	"role": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("USER"),
		Validators: []validator.String{
			stringvalidator.OneOf("OPERATOR", "USER"),
		},
		MarkdownDescription: "The LOM user role which specifies the list of actions that are allowed for the user.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the LOM user is disabled.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The descriptive comment for the LOM user.",
	},
}

func ExpandMemberLomUsers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberLomUsers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberLomUsersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberLomUsersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberLomUsers {
	if m == nil {
		return nil
	}
	to := &grid.MemberLomUsers{
		Name:    flex.ExpandStringPointer(m.Name),
		Role:    flex.ExpandStringPointer(m.Role),
		Disable: flex.ExpandBoolPointer(m.Disable),
		Comment: flex.ExpandStringPointer(m.Comment),
	}
	// TODO: Password is non-updatable, Please create an IsCreate Block for the same
	return to
}

func FlattenMemberLomUsers(ctx context.Context, from *grid.MemberLomUsers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberLomUsersAttrTypes)
	}
	m := MemberLomUsersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberLomUsersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberLomUsersModel) Flatten(ctx context.Context, from *grid.MemberLomUsers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberLomUsersModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Password = flex.FlattenStringPointer(from.Password)
	m.Role = flex.FlattenStringPointer(from.Role)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.Comment = flex.FlattenStringPointer(from.Comment)
}
