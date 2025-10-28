package dhcp

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type FixedaddressMsServerModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
}

var FixedaddressMsServerAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
}

var FixedaddressMsServerResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])$|^[0-9a-z]+(\.[a-z]+)*$`),
				"Must be valid IPv4 address or FQDN with lowercase letters, no trailing dot or spaces",
			),
		},
		MarkdownDescription: "The IPv4 Address or FQDN of the Microsoft server.",
	},
}

func ExpandFixedaddressMsServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dhcp.FixedaddressMsServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedaddressMsServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FixedaddressMsServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.FixedaddressMsServer {
	if m == nil {
		return nil
	}
	to := &dhcp.FixedaddressMsServer{
		Ipv4addr: flex.ExpandStringPointer(m.Ipv4addr),
	}
	return to
}

func FlattenFixedaddressMsServer(ctx context.Context, from *dhcp.FixedaddressMsServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddressMsServerAttrTypes)
	}
	m := FixedaddressMsServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FixedaddressMsServerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FixedaddressMsServerModel) Flatten(ctx context.Context, from *dhcp.FixedaddressMsServer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FixedaddressMsServerModel{}
	}
	m.Ipv4addr = flex.FlattenStringPointer(from.Ipv4addr)
}
