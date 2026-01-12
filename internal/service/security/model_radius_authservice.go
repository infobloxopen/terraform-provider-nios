package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RadiusAuthserviceModel struct {
	Ref              types.String `tfsdk:"ref"`
	AcctRetries      types.Int64  `tfsdk:"acct_retries"`
	AcctTimeout      types.Int64  `tfsdk:"acct_timeout"`
	AuthRetries      types.Int64  `tfsdk:"auth_retries"`
	AuthTimeout      types.Int64  `tfsdk:"auth_timeout"`
	CacheTtl         types.Int64  `tfsdk:"cache_ttl"`
	Comment          types.String `tfsdk:"comment"`
	Disable          types.Bool   `tfsdk:"disable"`
	EnableCache      types.Bool   `tfsdk:"enable_cache"`
	Mode             types.String `tfsdk:"mode"`
	Name             types.String `tfsdk:"name"`
	RecoveryInterval types.Int64  `tfsdk:"recovery_interval"`
	Servers          types.List   `tfsdk:"servers"`
}

var RadiusAuthserviceAttrTypes = map[string]attr.Type{
	"ref":               types.StringType,
	"acct_retries":      types.Int64Type,
	"acct_timeout":      types.Int64Type,
	"auth_retries":      types.Int64Type,
	"auth_timeout":      types.Int64Type,
	"cache_ttl":         types.Int64Type,
	"comment":           types.StringType,
	"disable":           types.BoolType,
	"enable_cache":      types.BoolType,
	"mode":              types.StringType,
	"name":              types.StringType,
	"recovery_interval": types.Int64Type,
	"servers":           types.ListType{ElemType: types.ObjectType{AttrTypes: RadiusAuthserviceServersAttrTypes}},
}

var RadiusAuthserviceResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"acct_retries": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of times to attempt to contact an accounting RADIUS server.",
	},
	"acct_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of seconds to wait for a response from the RADIUS server.",
	},
	"auth_retries": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of times to attempt to contact an authentication RADIUS server.",
	},
	"auth_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of seconds to wait for a response from the RADIUS server.",
	},
	"cache_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The TTL of cached authentication data in seconds.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The RADIUS descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the RADIUS authentication service is disabled.",
	},
	"enable_cache": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the authentication cache is enabled.",
	},
	"mode": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The way to contact the RADIUS server.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The RADIUS authentication service name.",
	},
	"recovery_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time period to wait before retrying a server that has been marked as down.",
	},
	"servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RadiusAuthserviceServersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The ordered list of RADIUS authentication servers.",
	},
}

func ExpandRadiusAuthservice(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.RadiusAuthservice {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RadiusAuthserviceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RadiusAuthserviceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.RadiusAuthservice {
	if m == nil {
		return nil
	}
	to := &security.RadiusAuthservice{
		Ref:              flex.ExpandStringPointer(m.Ref),
		AcctRetries:      flex.ExpandInt64Pointer(m.AcctRetries),
		AcctTimeout:      flex.ExpandInt64Pointer(m.AcctTimeout),
		AuthRetries:      flex.ExpandInt64Pointer(m.AuthRetries),
		AuthTimeout:      flex.ExpandInt64Pointer(m.AuthTimeout),
		CacheTtl:         flex.ExpandInt64Pointer(m.CacheTtl),
		Comment:          flex.ExpandStringPointer(m.Comment),
		Disable:          flex.ExpandBoolPointer(m.Disable),
		EnableCache:      flex.ExpandBoolPointer(m.EnableCache),
		Mode:             flex.ExpandStringPointer(m.Mode),
		Name:             flex.ExpandStringPointer(m.Name),
		RecoveryInterval: flex.ExpandInt64Pointer(m.RecoveryInterval),
		Servers:          flex.ExpandFrameworkListNestedBlock(ctx, m.Servers, diags, ExpandRadiusAuthserviceServers),
	}
	return to
}

func FlattenRadiusAuthservice(ctx context.Context, from *security.RadiusAuthservice, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RadiusAuthserviceAttrTypes)
	}
	m := RadiusAuthserviceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RadiusAuthserviceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RadiusAuthserviceModel) Flatten(ctx context.Context, from *security.RadiusAuthservice, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RadiusAuthserviceModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AcctRetries = flex.FlattenInt64Pointer(from.AcctRetries)
	m.AcctTimeout = flex.FlattenInt64Pointer(from.AcctTimeout)
	m.AuthRetries = flex.FlattenInt64Pointer(from.AuthRetries)
	m.AuthTimeout = flex.FlattenInt64Pointer(from.AuthTimeout)
	m.CacheTtl = flex.FlattenInt64Pointer(from.CacheTtl)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.EnableCache = types.BoolPointerValue(from.EnableCache)
	m.Mode = flex.FlattenStringPointer(from.Mode)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.RecoveryInterval = flex.FlattenInt64Pointer(from.RecoveryInterval)
	m.Servers = flex.FlattenFrameworkListNestedBlock(ctx, from.Servers, RadiusAuthserviceServersAttrTypes, diags, FlattenRadiusAuthserviceServers)
}
