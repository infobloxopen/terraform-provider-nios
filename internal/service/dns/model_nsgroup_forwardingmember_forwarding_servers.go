package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NsgroupForwardingmemberForwardingServersModel struct {
	Name                  types.String `tfsdk:"name"`
	ForwardersOnly        types.Bool   `tfsdk:"forwarders_only"`
	ForwardTo             types.List   `tfsdk:"forward_to"`
	UseOverrideForwarders types.Bool   `tfsdk:"use_override_forwarders"`
}

var NsgroupForwardingmemberForwardingServersAttrTypes = map[string]attr.Type{
	"name":                    types.StringType,
	"forwarders_only":         types.BoolType,
	"forward_to":              types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupforwardingmemberforwardingserversForwardToAttrTypes}},
	"use_override_forwarders": types.BoolType,
}

var NsgroupForwardingmemberForwardingServersResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of this Grid member in FQDN format.",
	},
	"forwarders_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the appliance sends queries to forwarders only, and not to other internal or Internet root servers.",
	},
	"forward_to": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupforwardingmemberforwardingserversForwardToResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The information for the remote name server to which you want the Infoblox appliance to forward queries for a specified domain name.",
	},
	"use_override_forwarders": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: forward_to",
	},
}

func ExpandNsgroupForwardingmemberForwardingServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.NsgroupForwardingmemberForwardingServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NsgroupForwardingmemberForwardingServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NsgroupForwardingmemberForwardingServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.NsgroupForwardingmemberForwardingServers {
	if m == nil {
		return nil
	}
	to := &dns.NsgroupForwardingmemberForwardingServers{
		Name:                  flex.ExpandStringPointer(m.Name),
		ForwardersOnly:        flex.ExpandBoolPointer(m.ForwardersOnly),
		ForwardTo:             flex.ExpandFrameworkListNestedBlock(ctx, m.ForwardTo, diags, ExpandNsgroupforwardingmemberforwardingserversForwardTo),
		UseOverrideForwarders: flex.ExpandBoolPointer(m.UseOverrideForwarders),
	}
	return to
}

func FlattenNsgroupForwardingmemberForwardingServers(ctx context.Context, from *dns.NsgroupForwardingmemberForwardingServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupForwardingmemberForwardingServersAttrTypes)
	}
	m := NsgroupForwardingmemberForwardingServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NsgroupForwardingmemberForwardingServersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NsgroupForwardingmemberForwardingServersModel) Flatten(ctx context.Context, from *dns.NsgroupForwardingmemberForwardingServers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NsgroupForwardingmemberForwardingServersModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.ForwardersOnly = types.BoolPointerValue(from.ForwardersOnly)
	m.ForwardTo = flex.FlattenFrameworkListNestedBlock(ctx, from.ForwardTo, NsgroupforwardingmemberforwardingserversForwardToAttrTypes, diags, FlattenNsgroupforwardingmemberforwardingserversForwardTo)
	m.UseOverrideForwarders = types.BoolPointerValue(from.UseOverrideForwarders)
}
