package parentalcontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ParentalcontrolSubscriberModel struct {
	Ref                          types.String `tfsdk:"ref"`
	AltSubscriberId              types.String `tfsdk:"alt_subscriber_id"`
	AltSubscriberIdRegexp        types.String `tfsdk:"alt_subscriber_id_regexp"`
	AltSubscriberIdSubexpression types.Int64  `tfsdk:"alt_subscriber_id_subexpression"`
	Ancillaries                  types.List   `tfsdk:"ancillaries"`
	CatAcctname                  types.String `tfsdk:"cat_acctname"`
	CatPassword                  types.String `tfsdk:"cat_password"`
	CatUpdateFrequency           types.Int64  `tfsdk:"cat_update_frequency"`
	CategoryUrl                  types.String `tfsdk:"category_url"`
	EnableMgmtOnlyNas            types.Bool   `tfsdk:"enable_mgmt_only_nas"`
	EnableParentalControl        types.Bool   `tfsdk:"enable_parental_control"`
	InterimAccountingInterval    types.Int64  `tfsdk:"interim_accounting_interval"`
	IpAnchors                    types.List   `tfsdk:"ip_anchors"`
	IpSpaceDiscRegexp            types.String `tfsdk:"ip_space_disc_regexp"`
	IpSpaceDiscSubexpression     types.Int64  `tfsdk:"ip_space_disc_subexpression"`
	IpSpaceDiscriminator         types.String `tfsdk:"ip_space_discriminator"`
	LocalId                      types.String `tfsdk:"local_id"`
	LocalIdRegexp                types.String `tfsdk:"local_id_regexp"`
	LocalIdSubexpression         types.Int64  `tfsdk:"local_id_subexpression"`
	LogGuestLookups              types.Bool   `tfsdk:"log_guest_lookups"`
	NasContextInfo               types.String `tfsdk:"nas_context_info"`
	PcZoneName                   types.String `tfsdk:"pc_zone_name"`
	ProxyPassword                types.String `tfsdk:"proxy_password"`
	ProxyUrl                     types.String `tfsdk:"proxy_url"`
	ProxyUsername                types.String `tfsdk:"proxy_username"`
	SubscriberId                 types.String `tfsdk:"subscriber_id"`
	SubscriberIdRegexp           types.String `tfsdk:"subscriber_id_regexp"`
	SubscriberIdSubexpression    types.Int64  `tfsdk:"subscriber_id_subexpression"`
	ZveloUpdateFailureInDays     types.Int64  `tfsdk:"zvelo_update_failure_in_days"`
}

var ParentalcontrolSubscriberAttrTypes = map[string]attr.Type{
	"ref":                             types.StringType,
	"alt_subscriber_id":               types.StringType,
	"alt_subscriber_id_regexp":        types.StringType,
	"alt_subscriber_id_subexpression": types.Int64Type,
	"ancillaries":                     types.ListType{ElemType: types.StringType},
	"cat_acctname":                    types.StringType,
	"cat_password":                    types.StringType,
	"cat_update_frequency":            types.Int64Type,
	"category_url":                    types.StringType,
	"enable_mgmt_only_nas":            types.BoolType,
	"enable_parental_control":         types.BoolType,
	"interim_accounting_interval":     types.Int64Type,
	"ip_anchors":                      types.ListType{ElemType: types.StringType},
	"ip_space_disc_regexp":            types.StringType,
	"ip_space_disc_subexpression":     types.Int64Type,
	"ip_space_discriminator":          types.StringType,
	"local_id":                        types.StringType,
	"local_id_regexp":                 types.StringType,
	"local_id_subexpression":          types.Int64Type,
	"log_guest_lookups":               types.BoolType,
	"nas_context_info":                types.StringType,
	"pc_zone_name":                    types.StringType,
	"proxy_password":                  types.StringType,
	"proxy_url":                       types.StringType,
	"proxy_username":                  types.StringType,
	"subscriber_id":                   types.StringType,
	"subscriber_id_regexp":            types.StringType,
	"subscriber_id_subexpression":     types.Int64Type,
	"zvelo_update_failure_in_days":    types.Int64Type,
}

var ParentalcontrolSubscriberResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"alt_subscriber_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of AVP to be used as an alternate subscriber ID for fixed lines.",
	},
	"alt_subscriber_id_regexp": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A character string to control aspects of rewriting of the fields.",
	},
	"alt_subscriber_id_subexpression": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The subexpression indicates which subexpression to extract. If zero, then the text matching the entire regular expression is extracted. If non-zero, then the regex must contain at least that many sub-expression groups. It takes values from 0 to 8.",
	},
	"ancillaries": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of ordered AVP Ancillary Fields.",
	},
	"cat_acctname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Category content account name using the categorization service.",
	},
	"cat_password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Category content account password to access the categorization service.",
	},
	"cat_update_frequency": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Category content updates every number of hours.",
	},
	"category_url": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Category content vendor url to download category data from and upload feedback to, configure for parental control.",
	},
	"enable_mgmt_only_nas": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if NAS RADIUS traffic is accepted over MGMT only.",
	},
	"enable_parental_control": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if parental control is enabled.",
	},
	"interim_accounting_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time for collector to be fully populated. Valid values are from 1 to 65535.",
	},
	"ip_anchors": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The ordered list of IP Anchors AVPs. The list content cannot be changed, but the order of elements.",
	},
	"ip_space_disc_regexp": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A character string to control aspects of rewriting of the fields.",
	},
	"ip_space_disc_subexpression": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The subexpression indicates which subexpression to extract. If zero, then the text matching the entire regular expression is extracted. If non-zero, then the regex must contain at least that many sub-expression groups. It takes values from 0 to 8.",
	},
	"ip_space_discriminator": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of AVP to be used as IP address discriminator.",
	},
	"local_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of AVP to be used as local ID.",
	},
	"local_id_regexp": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A character string to control aspects of rewriting of the fields.",
	},
	"local_id_subexpression": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The subexpression indicates which subexpression to extract. If zero, then the text matching the entire regular expression is extracted. If non-zero, then the regex must contain at least that many sub-expression groups. It takes values from 0 to 8.",
	},
	"log_guest_lookups": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "CEF log all guest lookups, will produce two logs in case of a violation.",
	},
	"nas_context_info": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "NAS contextual information AVP.",
	},
	"pc_zone_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The SOA to store parental control records.",
	},
	"proxy_password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Proxy server password used for authentication.",
	},
	"proxy_url": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Proxy url to download category data from and upload feedback to, configure for parental control. The default value 'None' is no longer valid as it match url regex pattern \"^http|https://\". The new default value does not get saved in database, but rather used for comparision with object created in unit test cases.",
	},
	"proxy_username": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Proxy server username used for authentication.",
	},
	"subscriber_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of AVP to be used as a subscriber.",
	},
	"subscriber_id_regexp": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A character string to control aspects of rewriting of the fields.",
	},
	"subscriber_id_subexpression": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The subexpression indicates which subexpression to extract. If zero, then the text matching the entire regular expression is extracted. If non-zero, then the regex must contain at least that many sub-expression groups. It takes values from 0 to 8.",
	},
	"zvelo_update_failure_in_days": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Number of days since zvelo DB failed to update.",
	},
}

func ExpandParentalcontrolSubscriber(ctx context.Context, o types.Object, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscriber {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ParentalcontrolSubscriberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ParentalcontrolSubscriberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolSubscriber {
	if m == nil {
		return nil
	}
	to := &parentalcontrol.ParentalcontrolSubscriber{
		Ref:                          flex.ExpandStringPointer(m.Ref),
		AltSubscriberId:              flex.ExpandStringPointer(m.AltSubscriberId),
		AltSubscriberIdRegexp:        flex.ExpandStringPointer(m.AltSubscriberIdRegexp),
		AltSubscriberIdSubexpression: flex.ExpandInt64Pointer(m.AltSubscriberIdSubexpression),
		Ancillaries:                  flex.ExpandFrameworkListString(ctx, m.Ancillaries, diags),
		CatAcctname:                  flex.ExpandStringPointer(m.CatAcctname),
		CatPassword:                  flex.ExpandStringPointer(m.CatPassword),
		CatUpdateFrequency:           flex.ExpandInt64Pointer(m.CatUpdateFrequency),
		CategoryUrl:                  flex.ExpandStringPointer(m.CategoryUrl),
		EnableMgmtOnlyNas:            flex.ExpandBoolPointer(m.EnableMgmtOnlyNas),
		EnableParentalControl:        flex.ExpandBoolPointer(m.EnableParentalControl),
		InterimAccountingInterval:    flex.ExpandInt64Pointer(m.InterimAccountingInterval),
		IpAnchors:                    flex.ExpandFrameworkListString(ctx, m.IpAnchors, diags),
		IpSpaceDiscRegexp:            flex.ExpandStringPointer(m.IpSpaceDiscRegexp),
		IpSpaceDiscSubexpression:     flex.ExpandInt64Pointer(m.IpSpaceDiscSubexpression),
		IpSpaceDiscriminator:         flex.ExpandStringPointer(m.IpSpaceDiscriminator),
		LocalId:                      flex.ExpandStringPointer(m.LocalId),
		LocalIdRegexp:                flex.ExpandStringPointer(m.LocalIdRegexp),
		LocalIdSubexpression:         flex.ExpandInt64Pointer(m.LocalIdSubexpression),
		LogGuestLookups:              flex.ExpandBoolPointer(m.LogGuestLookups),
		NasContextInfo:               flex.ExpandStringPointer(m.NasContextInfo),
		PcZoneName:                   flex.ExpandStringPointer(m.PcZoneName),
		ProxyPassword:                flex.ExpandStringPointer(m.ProxyPassword),
		ProxyUrl:                     flex.ExpandStringPointer(m.ProxyUrl),
		ProxyUsername:                flex.ExpandStringPointer(m.ProxyUsername),
		SubscriberId:                 flex.ExpandStringPointer(m.SubscriberId),
		SubscriberIdRegexp:           flex.ExpandStringPointer(m.SubscriberIdRegexp),
		SubscriberIdSubexpression:    flex.ExpandInt64Pointer(m.SubscriberIdSubexpression),
	}
	return to
}

func FlattenParentalcontrolSubscriber(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscriber, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ParentalcontrolSubscriberAttrTypes)
	}
	m := ParentalcontrolSubscriberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ParentalcontrolSubscriberAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ParentalcontrolSubscriberModel) Flatten(ctx context.Context, from *parentalcontrol.ParentalcontrolSubscriber, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ParentalcontrolSubscriberModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AltSubscriberId = flex.FlattenStringPointer(from.AltSubscriberId)
	m.AltSubscriberIdRegexp = flex.FlattenStringPointer(from.AltSubscriberIdRegexp)
	m.AltSubscriberIdSubexpression = flex.FlattenInt64Pointer(from.AltSubscriberIdSubexpression)
	m.Ancillaries = flex.FlattenFrameworkListString(ctx, from.Ancillaries, diags)
	m.CatAcctname = flex.FlattenStringPointer(from.CatAcctname)
	m.CatPassword = flex.FlattenStringPointer(from.CatPassword)
	m.CatUpdateFrequency = flex.FlattenInt64Pointer(from.CatUpdateFrequency)
	m.CategoryUrl = flex.FlattenStringPointer(from.CategoryUrl)
	m.EnableMgmtOnlyNas = types.BoolPointerValue(from.EnableMgmtOnlyNas)
	m.EnableParentalControl = types.BoolPointerValue(from.EnableParentalControl)
	m.InterimAccountingInterval = flex.FlattenInt64Pointer(from.InterimAccountingInterval)
	m.IpAnchors = flex.FlattenFrameworkListString(ctx, from.IpAnchors, diags)
	m.IpSpaceDiscRegexp = flex.FlattenStringPointer(from.IpSpaceDiscRegexp)
	m.IpSpaceDiscSubexpression = flex.FlattenInt64Pointer(from.IpSpaceDiscSubexpression)
	m.IpSpaceDiscriminator = flex.FlattenStringPointer(from.IpSpaceDiscriminator)
	m.LocalId = flex.FlattenStringPointer(from.LocalId)
	m.LocalIdRegexp = flex.FlattenStringPointer(from.LocalIdRegexp)
	m.LocalIdSubexpression = flex.FlattenInt64Pointer(from.LocalIdSubexpression)
	m.LogGuestLookups = types.BoolPointerValue(from.LogGuestLookups)
	m.NasContextInfo = flex.FlattenStringPointer(from.NasContextInfo)
	m.PcZoneName = flex.FlattenStringPointer(from.PcZoneName)
	m.ProxyPassword = flex.FlattenStringPointer(from.ProxyPassword)
	m.ProxyUrl = flex.FlattenStringPointer(from.ProxyUrl)
	m.ProxyUsername = flex.FlattenStringPointer(from.ProxyUsername)
	m.SubscriberId = flex.FlattenStringPointer(from.SubscriberId)
	m.SubscriberIdRegexp = flex.FlattenStringPointer(from.SubscriberIdRegexp)
	m.SubscriberIdSubexpression = flex.FlattenInt64Pointer(from.SubscriberIdSubexpression)
	m.ZveloUpdateFailureInDays = flex.FlattenInt64Pointer(from.ZveloUpdateFailureInDays)
}
