package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type SuperhostModel struct {
	Ref                     types.String                     `tfsdk:"ref"`
	Uuid                    types.String                     `tfsdk:"uuid"`
	Comment                 types.String                     `tfsdk:"comment"`
	DeleteAssociatedObjects types.Bool                       `tfsdk:"delete_associated_objects"`
	DhcpAssociatedObjects   internaltypes.UnorderedListValue `tfsdk:"dhcp_associated_objects"`
	Disabled                types.Bool                       `tfsdk:"disabled"`
	DnsAssociatedObjects    internaltypes.UnorderedListValue `tfsdk:"dns_associated_objects"`
	ExtAttrs                types.Map                        `tfsdk:"extattrs"`
	ExtAttrsAll             types.Map                        `tfsdk:"extattrs_all"`
	Name                    types.String                     `tfsdk:"name"`
}

var SuperhostAttrTypes = map[string]attr.Type{
	"ref":                       types.StringType,
	"uuid":                      types.StringType,
	"comment":                   types.StringType,
	"delete_associated_objects": types.BoolType,
	"dhcp_associated_objects":   internaltypes.UnorderedListOfStringType,
	"disabled":                  types.BoolType,
	"dns_associated_objects":    internaltypes.UnorderedListOfStringType,
	"extattrs":                  types.MapType{ElemType: types.StringType},
	"extattrs_all":              types.MapType{ElemType: types.StringType},
	"name":                      types.StringType,
}

var SuperhostResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"comment": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The comment for Super Host.",
	},
	"delete_associated_objects": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "True if we have to delete all DNS/DHCP associated objects with Super Host, false by default.",
	},
	"dhcp_associated_objects": schema.ListAttribute{
		CustomType:  internaltypes.UnorderedListOfStringType,
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "A list of DHCP objects refs which are associated with Super Host.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Disable all DNS/DHCP associated objects with Super Host if True, False by default.",
	},
	"dns_associated_objects": schema.ListAttribute{
		CustomType:  internaltypes.UnorderedListOfStringType,
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "A list of object refs of the DNS resource records which are associated with Super Host.",
	},
	"extattrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
		MarkdownDescription: "Extensible attributes associated with the object, including default and internal attributes.",
		ElementType:         types.StringType,
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of the Super Host.",
	},
}

func (m *SuperhostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Superhost {
	if m == nil {
		return nil
	}
	to := &ipam.Superhost{
		Comment:                 flex.ExpandStringPointer(m.Comment),
		DeleteAssociatedObjects: flex.ExpandBoolPointer(m.DeleteAssociatedObjects),
		DhcpAssociatedObjects:   flex.ExpandFrameworkListString(ctx, m.DhcpAssociatedObjects, diags),
		Disabled:                flex.ExpandBoolPointer(m.Disabled),
		DnsAssociatedObjects:    flex.ExpandFrameworkListString(ctx, m.DnsAssociatedObjects, diags),
		ExtAttrs:                ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Name:                    flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenSuperhost(ctx context.Context, from *ipam.Superhost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SuperhostAttrTypes)
	}
	m := SuperhostModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, SuperhostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SuperhostModel) Flatten(ctx context.Context, from *ipam.Superhost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SuperhostModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DhcpAssociatedObjects = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.DhcpAssociatedObjects, diags)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.DnsAssociatedObjects = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.DnsAssociatedObjects, diags)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
}
