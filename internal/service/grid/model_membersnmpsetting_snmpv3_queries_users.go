package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type MembersnmpsettingSnmpv3QueriesUsersModel struct {
	User                   types.String `tfsdk:"user"`
	Comment                types.String `tfsdk:"comment"`
	Ref                    types.String `tfsdk:"ref"`
	AuthenticationProtocol types.String `tfsdk:"authentication_protocol"`
	Disable                types.Bool   `tfsdk:"disable"`
	ExtAttrs               types.Map    `tfsdk:"extattrs"`
	Name                   types.String `tfsdk:"name"`
	PrivacyProtocol        types.String `tfsdk:"privacy_protocol"`
}

var MembersnmpsettingSnmpv3QueriesUsersAttrTypes = map[string]attr.Type{
	"user":                    types.StringType,
	"comment":                 types.StringType,
	"ref":                     types.StringType,
	"authentication_protocol": types.StringType,
	"disable":                 types.BoolType,
	"extattrs":                types.MapType{ElemType: types.StringType},
	"name":                    types.StringType,
	"privacy_protocol":        types.StringType,
}

var MembersnmpsettingSnmpv3QueriesUsersResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the SNMPv3 user object",
	},
	"user": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The SNMPv3 user.",
	},
	"comment": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A descriptive comment for this queries user.",
	},
	"authentication_protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The authentication protocol to be used for this user.",
	},
	"disable": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if SNMPv3 user is disabled or not.",
	},
	"extattrs": schema.MapAttribute{
		ElementType: types.StringType,
		Computed:    true,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the user.",
	},
	"privacy_protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The privacy protocol to be used for this user.",
	},
}

func ExpandMembersnmpsettingSnmpv3QueriesUsers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MembersnmpsettingSnmpv3QueriesUsers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MembersnmpsettingSnmpv3QueriesUsersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MembersnmpsettingSnmpv3QueriesUsersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MembersnmpsettingSnmpv3QueriesUsers {
	if m == nil {
		return nil
	}
	to := &grid.MembersnmpsettingSnmpv3QueriesUsers{
		User:    flex.ExpandStringPointer(m.User),
		Comment: flex.ExpandStringPointer(m.Comment),
	}
	return to
}

func FlattenMembersnmpsettingSnmpv3QueriesUsers(ctx context.Context, from *grid.MembersnmpsettingSnmpv3QueriesUsers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MembersnmpsettingSnmpv3QueriesUsersAttrTypes)
	}
	m := MembersnmpsettingSnmpv3QueriesUsersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MembersnmpsettingSnmpv3QueriesUsersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MembersnmpsettingSnmpv3QueriesUsersModel) Flatten(ctx context.Context, from *grid.MembersnmpsettingSnmpv3QueriesUsers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MembersnmpsettingSnmpv3QueriesUsersModel{}
	}
	//m.User = FlattenGridsnmpsettingSnmpv3QueriesUsersOneOf(ctx ,)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	// m.Ref = flex.FlattenStringPointer(from.Ref)
	// m.AuthenticationProtocol = flex.FlattenStringPointer(from.AuthenticationProtocol)
	// m.Disable = types.BoolPointerValue(from.Disable)
	// m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	// m.Name = flex.FlattenStringPointer(from.Name)
	// m.PrivacyProtocol = flex.FlattenStringPointer(from.PrivacyProtocol)
}

func FlattenGridsnmpsettingSnmpv3QueriesUsersOneOf(ctx context.Context, from *grid.GridsnmpsettingSnmpv3QueriesUsersOneOf, diags *diag.Diagnostics) types.String {
	if from == nil {
		return types.StringNull()
	}
	return flex.FlattenStringPointer(from.Ref)
}