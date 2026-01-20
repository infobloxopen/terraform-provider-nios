package parentalcontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ParentalcontrolSubscriberrecordModel struct {
	Ref                    types.String `tfsdk:"ref"`
	Uuid                   types.String `tfsdk:"uuid"`
	AccountingSessionId    types.String `tfsdk:"accounting_session_id"`
	AltIpAddr              types.String `tfsdk:"alt_ip_addr"`
	Ans0                   types.String `tfsdk:"ans0"`
	Ans1                   types.String `tfsdk:"ans1"`
	Ans2                   types.String `tfsdk:"ans2"`
	Ans3                   types.String `tfsdk:"ans3"`
	Ans4                   types.String `tfsdk:"ans4"`
	BlackList              types.String `tfsdk:"black_list"`
	Bwflag                 types.Bool   `tfsdk:"bwflag"`
	DynamicCategoryPolicy  types.Bool   `tfsdk:"dynamic_category_policy"`
	Flags                  types.String `tfsdk:"flags"`
	IpAddr                 types.String `tfsdk:"ip_addr"`
	Ipsd                   types.String `tfsdk:"ipsd"`
	Localid                types.String `tfsdk:"localid"`
	NasContextual          types.String `tfsdk:"nas_contextual"`
	OpCode                 types.String `tfsdk:"op_code"`
	ParentalControlPolicy  types.String `tfsdk:"parental_control_policy"`
	Prefix                 types.Int64  `tfsdk:"prefix"`
	ProxyAll               types.Bool   `tfsdk:"proxy_all"`
	Site                   types.String `tfsdk:"site"`
	SubscriberId           types.String `tfsdk:"subscriber_id"`
	SubscriberSecurePolicy types.String `tfsdk:"subscriber_secure_policy"`
	UnknownCategoryPolicy  types.Bool   `tfsdk:"unknown_category_policy"`
	WhiteList              types.String `tfsdk:"white_list"`
	WpcCategoryPolicy      types.String `tfsdk:"wpc_category_policy"`
}

var ParentalcontrolSubscriberrecordAttrTypes = map[string]attr.Type{
	"ref":                      types.StringType,
	"uuid":                     types.StringType,
	"accounting_session_id":    types.StringType,
	"alt_ip_addr":              types.StringType,
	"ans0":                     types.StringType,
	"ans1":                     types.StringType,
	"ans2":                     types.StringType,
	"ans3":                     types.StringType,
	"ans4":                     types.StringType,
	"black_list":               types.StringType,
	"bwflag":                   types.BoolType,
	"dynamic_category_policy":  types.BoolType,
	"flags":                    types.StringType,
	"ip_addr":                  types.StringType,
	"ipsd":                     types.StringType,
	"localid":                  types.StringType,
	"nas_contextual":           types.StringType,
	"op_code":                  types.StringType,
	"parental_control_policy":  types.StringType,
	"prefix":                   types.Int64Type,
	"proxy_all":                types.BoolType,
	"site":                     types.StringType,
	"subscriber_id":            types.StringType,
	"subscriber_secure_policy": types.StringType,
	"unknown_category_policy":  types.BoolType,
	"white_list":               types.StringType,
	"wpc_category_policy":      types.StringType,
}

var ParentalcontrolSubscriberrecordResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"accounting_session_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "accounting_session_id",
	},
	"alt_ip_addr": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "alt_ip_addr",
	},
	"ans0": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "ans0",
	},
	"ans1": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "ans1",
	},
	"ans2": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "ans2",
	},
	"ans3": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "ans3",
	},
	"ans4": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "ans4",
	},
	"black_list": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "black_list",
	},
	"bwflag": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "bwflag",
	},
	"dynamic_category_policy": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "dynamic_category_policy",
	},
	"flags": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "flags",
	},
	"ip_addr": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "ip_addr",
	},
	"ipsd": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "ipsd",
	},
	"localid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "localid",
	},
	"nas_contextual": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "nas_contextual",
	},
	"op_code": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "op_code",
	},
	"parental_control_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "parental_control_policy",
	},
	"prefix": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "prefix",
	},
	"proxy_all": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "proxy_all",
	},
	"site": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "site",
	},
	"subscriber_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "subscriber_id",
	},
	"subscriber_secure_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "subscriber_secure_policy",
	},
	"unknown_category_policy": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "unknown_category_policy",
	},
	"white_list": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "white_list",
	},
	"wpc_category_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "wpc_category_policy",
	},
}

func ExpandParentalcontrolSubscriberrecord(ctx context.Context, o types.Object, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscriberrecord {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ParentalcontrolSubscriberrecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ParentalcontrolSubscriberrecordModel) Expand(ctx context.Context, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscriberrecord {
	if m == nil {
		return nil
	}
	to := &parentalcontrol.ParentalcontrolSubscriberrecord{
		Ref:                    flex.ExpandStringPointer(m.Ref),
		AccountingSessionId:    flex.ExpandStringPointer(m.AccountingSessionId),
		AltIpAddr:              flex.ExpandStringPointer(m.AltIpAddr),
		Ans0:                   flex.ExpandStringPointer(m.Ans0),
		Ans1:                   flex.ExpandStringPointer(m.Ans1),
		Ans2:                   flex.ExpandStringPointer(m.Ans2),
		Ans3:                   flex.ExpandStringPointer(m.Ans3),
		Ans4:                   flex.ExpandStringPointer(m.Ans4),
		BlackList:              flex.ExpandStringPointer(m.BlackList),
		Bwflag:                 flex.ExpandBoolPointer(m.Bwflag),
		DynamicCategoryPolicy:  flex.ExpandBoolPointer(m.DynamicCategoryPolicy),
		Flags:                  flex.ExpandStringPointer(m.Flags),
		IpAddr:                 flex.ExpandStringPointer(m.IpAddr),
		Ipsd:                   flex.ExpandStringPointer(m.Ipsd),
		Localid:                flex.ExpandStringPointer(m.Localid),
		NasContextual:          flex.ExpandStringPointer(m.NasContextual),
		OpCode:                 flex.ExpandStringPointer(m.OpCode),
		ParentalControlPolicy:  flex.ExpandStringPointer(m.ParentalControlPolicy),
		Prefix:                 flex.ExpandInt64Pointer(m.Prefix),
		ProxyAll:               flex.ExpandBoolPointer(m.ProxyAll),
		Site:                   flex.ExpandStringPointer(m.Site),
		SubscriberId:           flex.ExpandStringPointer(m.SubscriberId),
		SubscriberSecurePolicy: flex.ExpandStringPointer(m.SubscriberSecurePolicy),
		UnknownCategoryPolicy:  flex.ExpandBoolPointer(m.UnknownCategoryPolicy),
		WhiteList:              flex.ExpandStringPointer(m.WhiteList),
		WpcCategoryPolicy:      flex.ExpandStringPointer(m.WpcCategoryPolicy),
	}
	return to
}

func FlattenParentalcontrolSubscriberrecord(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscriberrecord, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ParentalcontrolSubscriberrecordAttrTypes)
	}
	m := ParentalcontrolSubscriberrecordModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ParentalcontrolSubscriberrecordAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ParentalcontrolSubscriberrecordModel) Flatten(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscriberrecord, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ParentalcontrolSubscriberrecordModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AccountingSessionId = flex.FlattenStringPointer(from.AccountingSessionId)
	m.AltIpAddr = flex.FlattenStringPointer(from.AltIpAddr)
	m.Ans0 = flex.FlattenStringPointer(from.Ans0)
	m.Ans1 = flex.FlattenStringPointer(from.Ans1)
	m.Ans2 = flex.FlattenStringPointer(from.Ans2)
	m.Ans3 = flex.FlattenStringPointer(from.Ans3)
	m.Ans4 = flex.FlattenStringPointer(from.Ans4)
	m.BlackList = flex.FlattenStringPointer(from.BlackList)
	m.Bwflag = types.BoolPointerValue(from.Bwflag)
	m.DynamicCategoryPolicy = types.BoolPointerValue(from.DynamicCategoryPolicy)
	m.Flags = flex.FlattenStringPointer(from.Flags)
	m.IpAddr = flex.FlattenStringPointer(from.IpAddr)
	m.Ipsd = flex.FlattenStringPointer(from.Ipsd)
	m.Localid = flex.FlattenStringPointer(from.Localid)
	m.NasContextual = flex.FlattenStringPointer(from.NasContextual)
	m.OpCode = flex.FlattenStringPointer(from.OpCode)
	m.ParentalControlPolicy = flex.FlattenStringPointer(from.ParentalControlPolicy)
	m.Prefix = flex.FlattenInt64Pointer(from.Prefix)
	m.ProxyAll = types.BoolPointerValue(from.ProxyAll)
	m.Site = flex.FlattenStringPointer(from.Site)
	m.SubscriberId = flex.FlattenStringPointer(from.SubscriberId)
	m.SubscriberSecurePolicy = flex.FlattenStringPointer(from.SubscriberSecurePolicy)
	m.UnknownCategoryPolicy = types.BoolPointerValue(from.UnknownCategoryPolicy)
	m.WhiteList = flex.FlattenStringPointer(from.WhiteList)
	m.WpcCategoryPolicy = flex.FlattenStringPointer(from.WpcCategoryPolicy)
}
