package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RecordNsAddressesModel struct {
	Address       types.String `tfsdk:"address"`
	AutoCreatePtr types.Bool   `tfsdk:"auto_create_ptr"`
}

var RecordNsAddressesAttrTypes = map[string]attr.Type{
	"address":         types.StringType,
	"auto_create_ptr": types.BoolType,
}

var RecordNsAddressesResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The address of the Zone Name Server.",
	},
	"auto_create_ptr": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Flag to indicate if ptr records need to be auto created.",
	},
}

func ExpandRecordNsAddresses(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordNsAddresses {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordNsAddressesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordNsAddressesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordNsAddresses {
	if m == nil {
		return nil
	}
	to := &dns.RecordNsAddresses{
		Address:       flex.ExpandStringPointer(m.Address),
		AutoCreatePtr: flex.ExpandBoolPointer(m.AutoCreatePtr),
	}
	return to
}

func FlattenRecordNsAddresses(ctx context.Context, from *dns.RecordNsAddresses, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordNsAddressesAttrTypes)
	}
	m := RecordNsAddressesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordNsAddressesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordNsAddressesModel) Flatten(ctx context.Context, from *dns.RecordNsAddresses, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordNsAddressesModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AutoCreatePtr = types.BoolPointerValue(from.AutoCreatePtr)
}
