package rir

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/rir"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
)

type RirOrganizationModel struct {
	Ref         types.String `tfsdk:"ref"`
	Uuid        types.String `tfsdk:"uuid"`
	ExtAttrs    types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll types.Map    `tfsdk:"extattrs_all"`
	Id          types.String `tfsdk:"id"`
	Maintainer  types.String `tfsdk:"maintainer"`
	Name        types.String `tfsdk:"name"`
	Password    types.String `tfsdk:"password"`
	Rir         types.String `tfsdk:"rir"`
	SenderEmail types.String `tfsdk:"sender_email"`
}

var RirOrganizationAttrTypes = map[string]attr.Type{
	"ref":          types.StringType,
	"uuid":         types.StringType,
	"extattrs":     types.MapType{ElemType: types.StringType},
	"extattrs_all": types.MapType{ElemType: types.StringType},
	"id":           types.StringType,
	"maintainer":   types.StringType,
	"name":         types.StringType,
	"password":     types.StringType,
	"rir":          types.StringType,
	"sender_email": types.StringType,
}

var RirOrganizationResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The RIR organization identifier.",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"maintainer": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The RIR organization maintainer.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The RIR organization name.",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The password for the maintainer of RIR organization.",
	},
	"rir": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The RIR associated with RIR organization.",
	},
	"sender_email": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The sender e-mail address for RIR organization.",
	},
}

func ExpandRirOrganization(ctx context.Context, o types.Object, diags *diag.Diagnostics) *rir.RirOrganization {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RirOrganizationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RirOrganizationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *rir.RirOrganization {
	if m == nil {
		return nil
	}
	to := &rir.RirOrganization{
		Ref:         flex.ExpandStringPointer(m.Ref),
		ExtAttrs:    ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Id:          flex.ExpandStringPointer(m.Id),
		Maintainer:  flex.ExpandStringPointer(m.Maintainer),
		Name:        flex.ExpandStringPointer(m.Name),
		Password:    flex.ExpandStringPointer(m.Password),
		Rir:         flex.ExpandStringPointer(m.Rir),
		SenderEmail: flex.ExpandStringPointer(m.SenderEmail),
	}
	return to
}

func FlattenRirOrganization(ctx context.Context, from *rir.RirOrganization, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RirOrganizationAttrTypes)
	}
	m := RirOrganizationModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, RirOrganizationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RirOrganizationModel) Flatten(ctx context.Context, from *rir.RirOrganization, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RirOrganizationModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Maintainer = flex.FlattenStringPointer(from.Maintainer)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Password = flex.FlattenStringPointer(from.Password)
	m.Rir = flex.FlattenStringPointer(from.Rir)
	m.SenderEmail = flex.FlattenStringPointer(from.SenderEmail)
}
