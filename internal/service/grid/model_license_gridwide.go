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

type LicenseGridwideModel struct {
	Ref              types.String `tfsdk:"ref"`
	ExpirationStatus types.String `tfsdk:"expiration_status"`
	ExpiryDate       types.Int64  `tfsdk:"expiry_date"`
	Key              types.String `tfsdk:"key"`
	Limit            types.String `tfsdk:"limit"`
	LimitContext     types.String `tfsdk:"limit_context"`
	Type             types.String `tfsdk:"type"`
}

var LicenseGridwideAttrTypes = map[string]attr.Type{
	"ref":               types.StringType,
	"expiration_status": types.StringType,
	"expiry_date":       types.Int64Type,
	"key":               types.StringType,
	"limit":             types.StringType,
	"limit_context":     types.StringType,
	"type":              types.StringType,
}

var LicenseGridwideResourceSchemaAttributes = map[string]schema.Attribute{
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
	"key": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The license string.",
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

func ExpandLicenseGridwide(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.LicenseGridwide {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m LicenseGridwideModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *LicenseGridwideModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.LicenseGridwide {
	if m == nil {
		return nil
	}
	to := &grid.LicenseGridwide{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenLicenseGridwide(ctx context.Context, from *grid.LicenseGridwide, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(LicenseGridwideAttrTypes)
	}
	m := LicenseGridwideModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, LicenseGridwideAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *LicenseGridwideModel) Flatten(ctx context.Context, from *grid.LicenseGridwide, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = LicenseGridwideModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.ExpirationStatus = flex.FlattenStringPointer(from.ExpirationStatus)
	m.ExpiryDate = flex.FlattenInt64Pointer(from.ExpiryDate)
	m.Key = flex.FlattenStringPointer(from.Key)
	m.Limit = flex.FlattenStringPointer(from.Limit)
	m.LimitContext = flex.FlattenStringPointer(from.LimitContext)
	m.Type = flex.FlattenStringPointer(from.Type)
}
