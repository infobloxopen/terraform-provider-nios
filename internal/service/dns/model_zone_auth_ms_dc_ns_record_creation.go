package dns

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

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ZoneAuthMsDcNsRecordCreationModel struct {
	Address types.String `tfsdk:"address"`
	Comment types.String `tfsdk:"comment"`
}

var ZoneAuthMsDcNsRecordCreationAttrTypes = map[string]attr.Type{
	"address": types.StringType,
	"comment": types.StringType,
}

var ZoneAuthMsDcNsRecordCreationResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The IPv4 address of the domain controller that is allowed to create NS records.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "Optional user comment.",
	},
}

func ExpandZoneAuthMsDcNsRecordCreation(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneAuthMsDcNsRecordCreation {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthMsDcNsRecordCreationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneAuthMsDcNsRecordCreationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneAuthMsDcNsRecordCreation {
	if m == nil {
		return nil
	}
	to := &dns.ZoneAuthMsDcNsRecordCreation{
		Address: flex.ExpandStringPointer(m.Address),
		Comment: flex.ExpandStringPointer(m.Comment),
	}
	return to
}

func FlattenZoneAuthMsDcNsRecordCreation(ctx context.Context, from *dns.ZoneAuthMsDcNsRecordCreation, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthMsDcNsRecordCreationAttrTypes)
	}
	m := ZoneAuthMsDcNsRecordCreationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthMsDcNsRecordCreationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthMsDcNsRecordCreationModel) Flatten(ctx context.Context, from *dns.ZoneAuthMsDcNsRecordCreation, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthMsDcNsRecordCreationModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Comment = flex.FlattenStringPointer(from.Comment)
}
