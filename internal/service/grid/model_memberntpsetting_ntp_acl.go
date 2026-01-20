package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberntpsettingNtpAclModel struct {
	AclType  types.String `tfsdk:"acl_type"`
	AcList   types.List   `tfsdk:"ac_list"`
	NamedAcl types.String `tfsdk:"named_acl"`
	Service  types.String `tfsdk:"service"`
}

var MemberntpsettingNtpAclAttrTypes = map[string]attr.Type{
	"acl_type":  types.StringType,
	"ac_list":   types.ListType{ElemType: types.ObjectType{AttrTypes: MemberntpsettingntpaclAcListAttrTypes}},
	"named_acl": types.StringType,
	"service":   types.StringType,
}

var MemberntpsettingNtpAclResourceSchemaAttributes = map[string]schema.Attribute{
	"acl_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The NTP access control list type.",
	},
	"ac_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberntpsettingntpaclAcListResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of NTP access control items.",
	},
	"named_acl": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The NTP access named ACL.",
	},
	"service": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of service with access control for the assigned named ACL.",
	},
}

func ExpandMemberntpsettingNtpAcl(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberntpsettingNtpAcl {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberntpsettingNtpAclModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberntpsettingNtpAclModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberntpsettingNtpAcl {
	if m == nil {
		return nil
	}
	to := &grid.MemberntpsettingNtpAcl{
		AclType:  flex.ExpandStringPointer(m.AclType),
		AcList:   flex.ExpandFrameworkListNestedBlock(ctx, m.AcList, diags, ExpandMemberntpsettingntpaclAcList),
		NamedAcl: flex.ExpandStringPointer(m.NamedAcl),
		Service:  flex.ExpandStringPointer(m.Service),
	}
	return to
}

func FlattenMemberntpsettingNtpAcl(ctx context.Context, from *grid.MemberntpsettingNtpAcl, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberntpsettingNtpAclAttrTypes)
	}
	m := MemberntpsettingNtpAclModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberntpsettingNtpAclAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberntpsettingNtpAclModel) Flatten(ctx context.Context, from *grid.MemberntpsettingNtpAcl, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberntpsettingNtpAclModel{}
	}
	m.AclType = flex.FlattenStringPointer(from.AclType)
	m.AcList = flex.FlattenFrameworkListNestedBlock(ctx, from.AcList, MemberntpsettingntpaclAcListAttrTypes, diags, FlattenMemberntpsettingntpaclAcList)
	m.NamedAcl = flex.FlattenStringPointer(from.NamedAcl)
	m.Service = flex.FlattenStringPointer(from.Service)
}
