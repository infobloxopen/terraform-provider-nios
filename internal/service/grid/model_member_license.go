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

type MemberLicenseModel struct {
	Ref              types.String `tfsdk:"ref"`
	ExpirationStatus types.String `tfsdk:"expiration_status"`
	ExpiryDate       types.Int64  `tfsdk:"expiry_date"`
	Hwid             types.String `tfsdk:"hwid"`
	Key              types.String `tfsdk:"key"`
	Kind             types.String `tfsdk:"kind"`
	Limit            types.String `tfsdk:"limit"`
	LimitContext     types.String `tfsdk:"limit_context"`
	Type             types.String `tfsdk:"type"`
}

var MemberLicenseAttrTypes = map[string]attr.Type{
	"ref":               types.StringType,
	"expiration_status": types.StringType,
	"expiry_date":       types.Int64Type,
	"hwid":              types.StringType,
	"key":               types.StringType,
	"kind":              types.StringType,
	"limit":             types.StringType,
	"limit_context":     types.StringType,
	"type":              types.StringType,
}

var MemberLicenseResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"expiration_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The license expiration status.",
	},
	"expiry_date": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The expiration timestamp of the license.",
	},
	"hwid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The hardware ID of the physical node on which the license is installed.",
	},
	"key": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "License string.",
	},
	"kind": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The overall type of license: static or dynamic.",
	},
	"limit": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The license limit value.",
	},
	"limit_context": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The license limit context.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The license type.",
	},
}

func ExpandMemberLicense(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberLicense {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberLicenseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberLicenseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberLicense {
	if m == nil {
		return nil
	}
	to := &grid.MemberLicense{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenMemberLicense(ctx context.Context, from *grid.MemberLicense, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberLicenseAttrTypes)
	}
	m := MemberLicenseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberLicenseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberLicenseModel) Flatten(ctx context.Context, from *grid.MemberLicense, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberLicenseModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.ExpirationStatus = flex.FlattenStringPointer(from.ExpirationStatus)
	m.ExpiryDate = flex.FlattenInt64Pointer(from.ExpiryDate)
	m.Hwid = flex.FlattenStringPointer(from.Hwid)
	m.Key = flex.FlattenStringPointer(from.Key)
	m.Kind = flex.FlattenStringPointer(from.Kind)
	m.Limit = flex.FlattenStringPointer(from.Limit)
	m.LimitContext = flex.FlattenStringPointer(from.LimitContext)
	m.Type = flex.FlattenStringPointer(from.Type)
}
