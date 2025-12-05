package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type SamlAuthserviceModel struct {
	Ref            types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	Comment        types.String `tfsdk:"comment"`
	Idp            types.Object `tfsdk:"idp"`
	Name           types.String `tfsdk:"name"`
	SessionTimeout types.Int64  `tfsdk:"session_timeout"`
}

var SamlAuthserviceAttrTypes = map[string]attr.Type{
	"ref":             types.StringType,
    "uuid":        types.StringType,
	"comment":         types.StringType,
	"idp":             types.ObjectType{AttrTypes: SamlAuthserviceIdpAttrTypes},
	"name":            types.StringType,
	"session_timeout": types.Int64Type,
}

var SamlAuthserviceResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The descriptive comment for the SAML authentication service.",
	},
	"idp": schema.SingleNestedAttribute{
		Attributes: SamlAuthserviceIdpResourceSchemaAttributes,
		Optional:   true,
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the SAML authentication service.",
	},
	"session_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The session timeout in seconds.",
	},
}

func ExpandSamlAuthservice(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.SamlAuthservice {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SamlAuthserviceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SamlAuthserviceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.SamlAuthservice {
	if m == nil {
		return nil
	}
	to := &security.SamlAuthservice{
		Ref:            flex.ExpandStringPointer(m.Ref),
		Comment:        flex.ExpandStringPointer(m.Comment),
		Idp:            ExpandSamlAuthserviceIdp(ctx, m.Idp, diags),
		Name:           flex.ExpandStringPointer(m.Name),
		SessionTimeout: flex.ExpandInt64Pointer(m.SessionTimeout),
	}
	return to
}

func FlattenSamlAuthservice(ctx context.Context, from *security.SamlAuthservice, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SamlAuthserviceAttrTypes)
	}
	m := SamlAuthserviceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SamlAuthserviceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SamlAuthserviceModel) Flatten(ctx context.Context, from *security.SamlAuthservice, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SamlAuthserviceModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Idp = FlattenSamlAuthserviceIdp(ctx, from.Idp, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.SessionTimeout = flex.FlattenInt64Pointer(from.SessionTimeout)
}
