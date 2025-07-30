package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type NetworkviewRemoteForwardZonesModel struct {
	Fqdn                types.String      `tfsdk:"fqdn"`
	ServerAddress       iptypes.IPAddress `tfsdk:"server_address"`
	GssTsigDnsPrincipal types.String      `tfsdk:"gss_tsig_dns_principal"`
	GssTsigDomain       types.String      `tfsdk:"gss_tsig_domain"`
	TsigKey             types.String      `tfsdk:"tsig_key"`
	TsigKeyAlg          types.String      `tfsdk:"tsig_key_alg"`
	TsigKeyName         types.String      `tfsdk:"tsig_key_name"`
	KeyType             types.String      `tfsdk:"key_type"`
}

var NetworkviewRemoteForwardZonesAttrTypes = map[string]attr.Type{
	"fqdn":                   types.StringType,
	"server_address":         iptypes.IPAddressType{},
	"gss_tsig_dns_principal": types.StringType,
	"gss_tsig_domain":        types.StringType,
	"tsig_key":               types.StringType,
	"tsig_key_alg":           types.StringType,
	"tsig_key_name":          types.StringType,
	"key_type":               types.StringType,
}

var NetworkviewRemoteForwardZonesResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
		MarkdownDescription: "The FQDN of the remote server.",
	},
	"server_address": schema.StringAttribute{
		CustomType:          iptypes.IPAddressType{},
		Required:            true,
		MarkdownDescription: "The remote server IP address.",
	},
	"gss_tsig_dns_principal": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The principal name in which GSS-TSIG for dynamic updates is enabled.",
	},
	"gss_tsig_domain": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The domain in which GSS-TSIG for dynamic updates is enabled.",
	},
	"tsig_key": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The TSIG key value.",
	},
	"tsig_key_alg": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("HMAC-MD5", "HMAC-SHA256"),
		},
		MarkdownDescription: "The TSIG key alorithm name.",
	},
	"tsig_key_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the TSIG key. The key name entered here must match the TSIG key name on the external name server.",
	},
	"key_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("GSS-TSIG", "NONE", "TSIG"),
		},
		Default:             stringdefault.StaticString("NONE"),
		MarkdownDescription: "The key type to be used.",
	},
}

func ExpandNetworkviewRemoteForwardZones(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkviewRemoteForwardZones {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkviewRemoteForwardZonesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkviewRemoteForwardZonesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkviewRemoteForwardZones {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkviewRemoteForwardZones{
		Fqdn:                flex.ExpandStringPointer(m.Fqdn),
		ServerAddress:       flex.ExpandIPAddress(m.ServerAddress),
		GssTsigDnsPrincipal: flex.ExpandStringPointer(m.GssTsigDnsPrincipal),
		GssTsigDomain:       flex.ExpandStringPointer(m.GssTsigDomain),
		TsigKey:             flex.ExpandStringPointer(m.TsigKey),
		TsigKeyAlg:          flex.ExpandStringPointer(m.TsigKeyAlg),
		TsigKeyName:         flex.ExpandStringPointer(m.TsigKeyName),
		KeyType:             flex.ExpandStringPointer(m.KeyType),
	}
	return to
}

func FlattenNetworkviewRemoteForwardZones(ctx context.Context, from *ipam.NetworkviewRemoteForwardZones, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkviewRemoteForwardZonesAttrTypes)
	}
	m := NetworkviewRemoteForwardZonesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkviewRemoteForwardZonesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkviewRemoteForwardZonesModel) Flatten(ctx context.Context, from *ipam.NetworkviewRemoteForwardZones, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkviewRemoteForwardZonesModel{}
	}
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.ServerAddress = flex.FlattenIPAddress(from.ServerAddress)
	m.GssTsigDnsPrincipal = flex.FlattenStringPointer(from.GssTsigDnsPrincipal)
	m.GssTsigDomain = flex.FlattenStringPointer(from.GssTsigDomain)
	m.TsigKey = flex.FlattenStringPointer(from.TsigKey)
	m.TsigKeyAlg = flex.FlattenStringPointer(from.TsigKeyAlg)
	m.TsigKeyName = flex.FlattenStringPointer(from.TsigKeyName)
	m.KeyType = flex.FlattenStringPointer(from.KeyType)
}
