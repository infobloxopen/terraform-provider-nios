package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type FixedaddressSnmpCredentialModel struct {
	CommunityString types.String `tfsdk:"community_string"`
	Comment         types.String `tfsdk:"comment"`
	CredentialGroup types.String `tfsdk:"credential_group"`
}

var FixedaddressSnmpCredentialAttrTypes = map[string]attr.Type{
	"community_string": types.StringType,
	"comment":          types.StringType,
	"credential_group": types.StringType,
}

var FixedaddressSnmpCredentialResourceSchemaAttributes = map[string]schema.Attribute{
	"community_string": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The public community string.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comments for the SNMPv1 and SNMPv2 users.",
	},
	"credential_group": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "Group for the SNMPv1 and SNMPv2 credential.",
	},
}

func ExpandFixedaddressSnmpCredential(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.FixedaddressSnmpCredential {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedaddressSnmpCredentialModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FixedaddressSnmpCredentialModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.FixedaddressSnmpCredential {
	if m == nil {
		return nil
	}
	to := &dhcp.FixedaddressSnmpCredential{
		CommunityString: flex.ExpandStringPointer(m.CommunityString),
		Comment:         flex.ExpandStringPointer(m.Comment),
		CredentialGroup: flex.ExpandStringPointer(m.CredentialGroup),
	}
	return to
}

func FlattenFixedaddressSnmpCredential(ctx context.Context, from *dhcp.FixedaddressSnmpCredential, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddressSnmpCredentialAttrTypes)
	}
	m := FixedaddressSnmpCredentialModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FixedaddressSnmpCredentialAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FixedaddressSnmpCredentialModel) Flatten(ctx context.Context, from *dhcp.FixedaddressSnmpCredential, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FixedaddressSnmpCredentialModel{}
	}
	m.CommunityString = flex.FlattenStringPointer(from.CommunityString)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CredentialGroup = flex.FlattenStringPointer(from.CredentialGroup)
}
