package misc

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type BfdtemplateModel struct {
	Ref                 types.String `tfsdk:"ref"`
	AuthenticationKey   types.String `tfsdk:"authentication_key"`
	AuthenticationKeyId types.Int64  `tfsdk:"authentication_key_id"`
	AuthenticationType  types.String `tfsdk:"authentication_type"`
	DetectionMultiplier types.Int64  `tfsdk:"detection_multiplier"`
	MinRxInterval       types.Int64  `tfsdk:"min_rx_interval"`
	MinTxInterval       types.Int64  `tfsdk:"min_tx_interval"`
	Name                types.String `tfsdk:"name"`
}

var BfdtemplateAttrTypes = map[string]attr.Type{
	"ref":                   types.StringType,
	"authentication_key":    types.StringType,
	"authentication_key_id": types.Int64Type,
	"authentication_type":   types.StringType,
	"detection_multiplier":  types.Int64Type,
	"min_rx_interval":       types.Int64Type,
	"min_tx_interval":       types.Int64Type,
	"name":                  types.StringType,
}

var BfdtemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"authentication_key": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Sensitive:           true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The authentication key for BFD protocol message-digest authentication.",
	},
	"authentication_key_id": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(1),
		Validators: []validator.Int64{
			int64validator.Between(1, 255),
		},
		MarkdownDescription: "The authentication key identifier for BFD protocol authentication. Valid values are between 1 and 255.",
	},
	"authentication_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("MD5", "METICULOUS-MD5", "METICULOUS-SHA1", "NONE", "SHA1"),
		},
		MarkdownDescription: "The authentication type for BFD protocol.",
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
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^\S.*\S$`),
				"should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The name of the BFD template object.",
	},
}

func (m *BfdtemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Bfdtemplate {
	if m == nil {
		return nil
	}
	to := &misc.Bfdtemplate{
		AuthenticationKey:   flex.ExpandStringPointer(m.AuthenticationKey),
		AuthenticationKeyId: flex.ExpandInt64Pointer(m.AuthenticationKeyId),
		AuthenticationType:  flex.ExpandStringPointer(m.AuthenticationType),
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
	m.AuthenticationKeyId = flex.FlattenInt64Pointer(from.AuthenticationKeyId)
	m.AuthenticationType = flex.FlattenStringPointer(from.AuthenticationType)
	m.DetectionMultiplier = flex.FlattenInt64Pointer(from.DetectionMultiplier)
	m.MinRxInterval = flex.FlattenInt64Pointer(from.MinRxInterval)
	m.MinTxInterval = flex.FlattenInt64Pointer(from.MinTxInterval)
	m.Name = flex.FlattenStringPointer(from.Name)
}
