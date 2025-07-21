package validator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ConvertSchemaAttributesToAttrTypes converts schema attributes to attribute types.
func ConvertSchemaAttributesToAttrTypes(schemaAttributes map[string]schema.Attribute) map[string]attr.Type {
	attrTypes := make(map[string]attr.Type)
	for key, attribute := range schemaAttributes {
		attrTypes[key] = attribute.GetType() // Ensure GetType() is used correctly
	}
	return attrTypes
}

type DefaultOptionsModifier struct {
	AttrTypes map[string]attr.Type
}

func (m DefaultOptionsModifier) Description(_ context.Context) string {
	return "Injects default DHCP lease-time option if not provided"
}

func (m DefaultOptionsModifier) MarkdownDescription(_ context.Context) string {
	return m.Description(context.Background())
}

func (m DefaultOptionsModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	// If config is null or unknown, inject default
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		resp.PlanValue = types.ListValueMust(
			types.ObjectType{AttrTypes: m.AttrTypes},
			[]attr.Value{types.ObjectValueMust(m.AttrTypes, defaultOption())},
		)
		return
	}

	// Check if dhcp-lease-time is already present
	elements := req.ConfigValue.Elements()
	found := false
	for _, elem := range elements {
		obj := elem.(types.Object)
		nameAttr := obj.Attributes()["name"].(types.String)
		if nameAttr.ValueString() == "dhcp-lease-time" {
			found = true
			break
		}
	}

	// If found, use the config as-is
	if found {
		resp.PlanValue = req.ConfigValue
		return
	}

	// If not found, append the default
	newList := append(elements, types.ObjectValueMust(m.AttrTypes, defaultOption()))
	resp.PlanValue = types.ListValueMust(types.ObjectType{AttrTypes: m.AttrTypes}, newList)
}

func defaultOption() map[string]attr.Value {
	return map[string]attr.Value{
		"name":         types.StringValue("dhcp-lease-time"),
		"num":          types.Int64Value(51),
		"use_option":   types.BoolValue(false),
		"value":        types.StringValue("43200"),
		"vendor_class": types.StringValue("DHCP"),
	}
}
