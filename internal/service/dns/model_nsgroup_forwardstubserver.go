package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type NsgroupForwardstubserverModel struct {
	Ref             types.String `tfsdk:"ref"`
	Comment         types.String `tfsdk:"comment"`
	ExtAttrs        types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll     types.Map    `tfsdk:"extattrs_all"`
	ExternalServers types.List   `tfsdk:"external_servers"`
	Name            types.String `tfsdk:"name"`
}

var NsgroupForwardstubserverAttrTypes = map[string]attr.Type{
	"ref":              types.StringType,
	"comment":          types.StringType,
	"extattrs":         types.MapType{ElemType: types.StringType},
	"extattrs_all":     types.MapType{ElemType: types.StringType},
	"external_servers": types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupForwardstubserverExternalServersAttrTypes}},
	"name":             types.StringType,
}

var NsgroupForwardstubserverResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the Forward Stub Server Name Server Group; maximum 256 characters.",
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
		PlanModifiers: []planmodifier.Map{
			importmod.MarkUnknownIfPrivateSet(),
		},
	},
	"external_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupForwardstubserverExternalServersResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of external servers.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of this Forward Stub Server Name Server Group.",
	},
}

func (m *NsgroupForwardstubserverModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.NsgroupForwardstubserver {
	if m == nil {
		return nil
	}
	to := &dns.NsgroupForwardstubserver{
		Comment:         flex.ExpandStringPointer(m.Comment),
		ExtAttrs:        ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ExternalServers: flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalServers, diags, ExpandNsgroupForwardstubserverExternalServers),
		Name:            flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenNsgroupForwardstubserver(ctx context.Context, from *dns.NsgroupForwardstubserver, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupForwardstubserverAttrTypes)
	}
	m := NsgroupForwardstubserverModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, NsgroupForwardstubserverAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NsgroupForwardstubserverModel) Flatten(ctx context.Context, from *dns.NsgroupForwardstubserver, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NsgroupForwardstubserverModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.ExternalServers = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalServers, NsgroupForwardstubserverExternalServersAttrTypes, diags, FlattenNsgroupForwardstubserverExternalServers)
	m.Name = flex.FlattenStringPointer(from.Name)
}
