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

type AdminuserSshKeysModel struct {
	KeyName  types.String `tfsdk:"key_name"`
	KeyType  types.String `tfsdk:"key_type"`
	KeyValue types.String `tfsdk:"key_value"`
}

var AdminuserSshKeysAttrTypes = map[string]attr.Type{
	"key_name":  types.StringType,
	"key_type":  types.StringType,
	"key_value": types.StringType,
}

var AdminuserSshKeysResourceSchemaAttributes = map[string]schema.Attribute{
	"key_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Unique identifier for the key",
	},
	"key_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "ssh_key_types",
	},
	"key_value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "ssh key text",
	},
}

func ExpandAdminuserSshKeys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdminuserSshKeys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdminuserSshKeysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdminuserSshKeysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdminuserSshKeys {
	if m == nil {
		return nil
	}
	to := &security.AdminuserSshKeys{
		KeyName:  flex.ExpandStringPointer(m.KeyName),
		KeyType:  flex.ExpandStringPointer(m.KeyType),
		KeyValue: flex.ExpandStringPointer(m.KeyValue),
	}
	return to
}

func FlattenAdminuserSshKeys(ctx context.Context, from *security.AdminuserSshKeys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdminuserSshKeysAttrTypes)
	}
	m := AdminuserSshKeysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdminuserSshKeysAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdminuserSshKeysModel) Flatten(ctx context.Context, from *security.AdminuserSshKeys, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdminuserSshKeysModel{}
	}
	m.KeyName = flex.FlattenStringPointer(from.KeyName)
	m.KeyType = flex.FlattenStringPointer(from.KeyType)
	m.KeyValue = flex.FlattenStringPointer(from.KeyValue)
}
