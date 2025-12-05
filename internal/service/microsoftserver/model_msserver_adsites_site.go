package microsoftserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MsserverAdsitesSiteModel struct {
	Ref      types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	Domain   types.String `tfsdk:"domain"`
	Name     types.String `tfsdk:"name"`
	Networks types.List   `tfsdk:"networks"`
}

var MsserverAdsitesSiteAttrTypes = map[string]attr.Type{
	"ref":      types.StringType,
    "uuid":        types.StringType,
	"domain":   types.StringType,
	"name":     types.StringType,
	"networks": types.ListType{ElemType: types.StringType},
}

var MsserverAdsitesSiteResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"domain": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the Active Directory Domain to which the site belongs.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the site properties object for the Active Directory Sites.",
	},
	"networks": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of networks to which the device interfaces belong.",
	},
}

func ExpandMsserverAdsitesSite(ctx context.Context, o types.Object, diags *diag.Diagnostics) *microsoftserver.MsserverAdsitesSite {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MsserverAdsitesSiteModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MsserverAdsitesSiteModel) Expand(ctx context.Context, diags *diag.Diagnostics) *microsoftserver.MsserverAdsitesSite {
	if m == nil {
		return nil
	}
	to := &microsoftserver.MsserverAdsitesSite{
		Ref:      flex.ExpandStringPointer(m.Ref),
		Domain:   flex.ExpandStringPointer(m.Domain),
		Name:     flex.ExpandStringPointer(m.Name),
		Networks: flex.ExpandFrameworkListString(ctx, m.Networks, diags),
	}
	return to
}

func FlattenMsserverAdsitesSite(ctx context.Context, from *microsoftserver.MsserverAdsitesSite, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MsserverAdsitesSiteAttrTypes)
	}
	m := MsserverAdsitesSiteModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MsserverAdsitesSiteAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MsserverAdsitesSiteModel) Flatten(ctx context.Context, from *microsoftserver.MsserverAdsitesSite, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MsserverAdsitesSiteModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Domain = flex.FlattenStringPointer(from.Domain)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Networks = flex.FlattenFrameworkListString(ctx, from.Networks, diags)
}
