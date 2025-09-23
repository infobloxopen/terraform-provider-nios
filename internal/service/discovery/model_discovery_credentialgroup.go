package discovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/discovery"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type DiscoveryCredentialgroupModel struct {
	Ref  types.String `tfsdk:"ref"`
	Name types.String `tfsdk:"name"`
}

var DiscoveryCredentialgroupAttrTypes = map[string]attr.Type{
	"ref":  types.StringType,
	"name": types.StringType,
}

var DiscoveryCredentialgroupResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the Credential group.",
	},
}

func (m *DiscoveryCredentialgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discovery.DiscoveryCredentialgroup {
	if m == nil {
		return nil
	}
	to := &discovery.DiscoveryCredentialgroup{
		Name: flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenDiscoveryCredentialgroup(ctx context.Context, from *discovery.DiscoveryCredentialgroup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DiscoveryCredentialgroupAttrTypes)
	}
	m := DiscoveryCredentialgroupModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DiscoveryCredentialgroupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DiscoveryCredentialgroupModel) Flatten(ctx context.Context, from *discovery.DiscoveryCredentialgroup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DiscoveryCredentialgroupModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Name = flex.FlattenStringPointer(from.Name)
}
