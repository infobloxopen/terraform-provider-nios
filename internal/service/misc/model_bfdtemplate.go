package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type BfdtemplateModel struct {
	Ref                 types.String `tfsdk:"ref"`
	Uuid                types.String `tfsdk:"uuid"`
	DetectionMultiplier types.Int64  `tfsdk:"detection_multiplier"`
	MinRxInterval       types.Int64  `tfsdk:"min_rx_interval"`
	MinTxInterval       types.Int64  `tfsdk:"min_tx_interval"`
	Name                types.String `tfsdk:"name"`
}

var BfdtemplateAttrTypes = map[string]attr.Type{
	"ref":                  types.StringType,
	"uuid":                 types.StringType,
	"detection_multiplier": types.Int64Type,
	"min_rx_interval":      types.Int64Type,
	"min_tx_interval":      types.Int64Type,
	"name":                 types.StringType,
}

var BfdtemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Universally Unique ID assigned for this object.",
	},
	"detection_multiplier": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(3),
		Validators: []validator.Int64{
			int64validator.Between(3, 50),
		},
		MarkdownDescription: "The detection time multiplier value for BFD protocol. The negotiated transmit interval, multiplied by this value, provides the detection time for the receiving system in asynchronous BFD mode. Valid values are between 3 and 50.",
	},
	"min_rx_interval": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(100),
		Validators: []validator.Int64{
			int64validator.Between(50, 9999),
		},
		MarkdownDescription: "The minimum receive time (in seconds) for BFD protocol. Valid values are between 50 and 9999.",
	},
	"min_tx_interval": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(100),
		Validators: []validator.Int64{
			int64validator.Between(50, 9999),
		},
		MarkdownDescription: "The minimum transmission time (in seconds) for BFD protocol. Valid values are between 50 and 9999.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the BFD template object.",
	},
}

func (m *BfdtemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Bfdtemplate {
	if m == nil {
		return nil
	}
	to := &misc.Bfdtemplate{
		DetectionMultiplier: flex.ExpandInt64Pointer(m.DetectionMultiplier),
		MinRxInterval:       flex.ExpandInt64Pointer(m.MinRxInterval),
		MinTxInterval:       flex.ExpandInt64Pointer(m.MinTxInterval),
		Name:                flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenBfdtemplate(ctx context.Context, from *misc.Bfdtemplate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(BfdtemplateAttrTypes)
	}
	m := BfdtemplateModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, BfdtemplateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *BfdtemplateModel) Flatten(ctx context.Context, from *misc.Bfdtemplate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = BfdtemplateModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.DetectionMultiplier = flex.FlattenInt64Pointer(from.DetectionMultiplier)
	m.MinRxInterval = flex.FlattenInt64Pointer(from.MinRxInterval)
	m.MinTxInterval = flex.FlattenInt64Pointer(from.MinTxInterval)
	m.Name = flex.FlattenStringPointer(from.Name)
}
