package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberDnsIpv6GlueRecordAddressesModel struct {
	AttachEmptyRecursiveView types.Bool   `tfsdk:"attach_empty_recursive_view"`
	GlueRecordAddress        types.String `tfsdk:"glue_record_address"`
	View                     types.String `tfsdk:"view"`
	GlueAddressChoice        types.String `tfsdk:"glue_address_choice"`
}

var MemberDnsIpv6GlueRecordAddressesAttrTypes = map[string]attr.Type{
	"attach_empty_recursive_view": types.BoolType,
	"glue_record_address":         types.StringType,
	"view":                        types.StringType,
	"glue_address_choice":         types.StringType,
}

var MemberDnsIpv6GlueRecordAddressesResourceSchemaAttributes = map[string]schema.Attribute{
	"attach_empty_recursive_view": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if empty view with recursion enabled will be written into the conf file.",
	},
	"glue_record_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The address the appliance uses to generate the glue record.",
	},
	"view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the DNS View in which the record resides. Example: \"external\".",
	},
	"glue_address_choice": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The address choice for auto-created glue records for this view.",
	},
}

func ExpandMemberDnsIpv6GlueRecordAddresses(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDnsIpv6GlueRecordAddresses {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDnsIpv6GlueRecordAddressesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDnsIpv6GlueRecordAddressesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDnsIpv6GlueRecordAddresses {
	if m == nil {
		return nil
	}
	to := &grid.MemberDnsIpv6GlueRecordAddresses{
		AttachEmptyRecursiveView: flex.ExpandBoolPointer(m.AttachEmptyRecursiveView),
		GlueRecordAddress:        flex.ExpandStringPointer(m.GlueRecordAddress),
		View:                     flex.ExpandStringPointer(m.View),
		GlueAddressChoice:        flex.ExpandStringPointer(m.GlueAddressChoice),
	}
	return to
}

func FlattenMemberDnsIpv6GlueRecordAddresses(ctx context.Context, from *grid.MemberDnsIpv6GlueRecordAddresses, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDnsIpv6GlueRecordAddressesAttrTypes)
	}
	m := MemberDnsIpv6GlueRecordAddressesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDnsIpv6GlueRecordAddressesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDnsIpv6GlueRecordAddressesModel) Flatten(ctx context.Context, from *grid.MemberDnsIpv6GlueRecordAddresses, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDnsIpv6GlueRecordAddressesModel{}
	}
	m.AttachEmptyRecursiveView = types.BoolPointerValue(from.AttachEmptyRecursiveView)
	m.GlueRecordAddress = flex.FlattenStringPointer(from.GlueRecordAddress)
	m.View = flex.FlattenStringPointer(from.View)
	m.GlueAddressChoice = flex.FlattenStringPointer(from.GlueAddressChoice)
}
