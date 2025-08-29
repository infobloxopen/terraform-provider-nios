package dns

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NsgroupForwardingmemberModel struct {
	Ref               types.String `tfsdk:"ref"`
	Comment           types.String `tfsdk:"comment"`
	ExtAttrs          types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll       types.Map    `tfsdk:"extattrs_all"`
	ForwardingServers types.List   `tfsdk:"forwarding_servers"`
	Name              types.String `tfsdk:"name"`
}

var NsgroupForwardingmemberAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"comment":            types.StringType,
	"extattrs":           types.MapType{ElemType: types.StringType},
	"extattrs_all":       types.MapType{ElemType: types.StringType},
	"forwarding_servers": types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupForwardingmemberForwardingServersAttrTypes}},
	"name":               types.StringType,
}

var NsgroupForwardingmemberResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "Comment for the Forwarding Member Name Server Group; maximum 256 characters.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object.",
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"forwarding_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupForwardingmemberForwardingServersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Required:            true,
		MarkdownDescription: "The list of forwarding member servers.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The name of the Forwarding Member Name Server Group.",
	},
}

func ExpandNsgroupForwardingmember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.NsgroupForwardingmember {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NsgroupForwardingmemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NsgroupForwardingmemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.NsgroupForwardingmember {
	if m == nil {
		return nil
	}
	to := &dns.NsgroupForwardingmember{
		Comment:           flex.ExpandStringPointer(m.Comment),
		ExtAttrs:          ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ForwardingServers: flex.ExpandFrameworkListNestedBlock(ctx, m.ForwardingServers, diags, ExpandNsgroupForwardingmemberForwardingServers),
		Name:              flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenNsgroupForwardingmember(ctx context.Context, from *dns.NsgroupForwardingmember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupForwardingmemberAttrTypes)
	}
	m := NsgroupForwardingmemberModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, NsgroupForwardingmemberAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NsgroupForwardingmemberModel) Flatten(ctx context.Context, from *dns.NsgroupForwardingmember, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NsgroupForwardingmemberModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.ForwardingServers = flex.FlattenFrameworkListNestedBlock(ctx, from.ForwardingServers, NsgroupForwardingmemberForwardingServersAttrTypes, diags, FlattenNsgroupForwardingmemberForwardingServers)
	m.Name = flex.FlattenStringPointer(from.Name)
}
