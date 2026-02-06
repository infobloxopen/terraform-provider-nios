package dhcp

import (
	"context"

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

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type MacfilteraddressModel struct {
	Ref                 types.String                  `tfsdk:"ref"`
	Uuid                types.String                  `tfsdk:"uuid"`
	AuthenticationTime  types.Int64                   `tfsdk:"authentication_time"`
	Comment             types.String                  `tfsdk:"comment"`
	ExpirationTime      types.Int64                   `tfsdk:"expiration_time"`
	ExtAttrs            types.Map                     `tfsdk:"extattrs"`
	Filter              types.String                  `tfsdk:"filter"`
	Fingerprint         types.String                  `tfsdk:"fingerprint"`
	GuestCustomField1   types.String                  `tfsdk:"guest_custom_field1"`
	GuestCustomField2   types.String                  `tfsdk:"guest_custom_field2"`
	GuestCustomField3   types.String                  `tfsdk:"guest_custom_field3"`
	GuestCustomField4   types.String                  `tfsdk:"guest_custom_field4"`
	GuestEmail          types.String                  `tfsdk:"guest_email"`
	GuestFirstName      types.String                  `tfsdk:"guest_first_name"`
	GuestLastName       types.String                  `tfsdk:"guest_last_name"`
	GuestMiddleName     types.String                  `tfsdk:"guest_middle_name"`
	GuestPhone          types.String                  `tfsdk:"guest_phone"`
	IsRegisteredUser    types.Bool                    `tfsdk:"is_registered_user"`
	Mac                 internaltypes.MACAddressValue `tfsdk:"mac"`
	NeverExpires        types.Bool                    `tfsdk:"never_expires"`
	ReservedForInfoblox types.String                  `tfsdk:"reserved_for_infoblox"`
	Username            types.String                  `tfsdk:"username"`
	ExtAttrsAll         types.Map                     `tfsdk:"extattrs_all"`
}

var MacfilteraddressAttrTypes = map[string]attr.Type{
	"ref":                   types.StringType,
	"uuid":                  types.StringType,
	"authentication_time":   types.Int64Type,
	"comment":               types.StringType,
	"expiration_time":       types.Int64Type,
	"extattrs":              types.MapType{ElemType: types.StringType},
	"filter":                types.StringType,
	"fingerprint":           types.StringType,
	"guest_custom_field1":   types.StringType,
	"guest_custom_field2":   types.StringType,
	"guest_custom_field3":   types.StringType,
	"guest_custom_field4":   types.StringType,
	"guest_email":           types.StringType,
	"guest_first_name":      types.StringType,
	"guest_last_name":       types.StringType,
	"guest_middle_name":     types.StringType,
	"guest_phone":           types.StringType,
	"is_registered_user":    types.BoolType,
	"mac":                   internaltypes.MACAddressType{},
	"never_expires":         types.BoolType,
	"reserved_for_infoblox": types.StringType,
	"username":              types.StringType,
	"extattrs_all":          types.MapType{ElemType: types.StringType},
}

var MacfilteraddressResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"authentication_time": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The absolute UNIX time (in seconds) since the address was last authenticated.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the MAC filter address; maximum 256 characters.",
	},
	"expiration_time": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The absolute UNIX time (in seconds) until the address expires.",
	},
	"extattrs": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"filter": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Name of the MAC filter to which this address belongs.",
	},
	"fingerprint": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "DHCP fingerprint for the address.",
	},
	"guest_custom_field1": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Guest custom field 1.",
	},
	"guest_custom_field2": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Guest custom field 2.",
	},
	"guest_custom_field3": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Guest custom field 3.",
	},
	"guest_custom_field4": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Guest custom field 4.",
	},
	"guest_email": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Guest e-mail.",
	},
	"guest_first_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Guest first name.",
	},
	"guest_last_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Guest last name.",
	},
	"guest_middle_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Guest middle name.",
	},
	"guest_phone": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Guest phone number.",
	},
	"is_registered_user": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the user has been authenticated or not.",
	},
	"mac": schema.StringAttribute{
		CustomType: internaltypes.MACAddressType{},
		Required:   true,
		Validators: []validator.String{
			customvalidator.IsValidMacAddress(),
		},
		MarkdownDescription: "MAC Address.",
	},
	"never_expires": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines if MAC address expiration is enabled or disabled.",
	},
	"reserved_for_infoblox": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Reserved for future use.",
	},
	"username": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Username for authenticated DHCP purposes.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object, including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
}

func (m *MacfilteraddressModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Macfilteraddress {
	if m == nil {
		return nil
	}
	to := &dhcp.Macfilteraddress{
		AuthenticationTime:  flex.ExpandInt64Pointer(m.AuthenticationTime),
		Comment:             flex.ExpandStringPointer(m.Comment),
		ExpirationTime:      flex.ExpandInt64Pointer(m.ExpirationTime),
		ExtAttrs:            ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Filter:              flex.ExpandStringPointer(m.Filter),
		GuestCustomField1:   flex.ExpandStringPointer(m.GuestCustomField1),
		GuestCustomField2:   flex.ExpandStringPointer(m.GuestCustomField2),
		GuestCustomField3:   flex.ExpandStringPointer(m.GuestCustomField3),
		GuestCustomField4:   flex.ExpandStringPointer(m.GuestCustomField4),
		GuestEmail:          flex.ExpandStringPointer(m.GuestEmail),
		GuestFirstName:      flex.ExpandStringPointer(m.GuestFirstName),
		GuestLastName:       flex.ExpandStringPointer(m.GuestLastName),
		GuestMiddleName:     flex.ExpandStringPointer(m.GuestMiddleName),
		GuestPhone:          flex.ExpandStringPointer(m.GuestPhone),
		Mac:                 flex.ExpandMACAddr(m.Mac),
		NeverExpires:        flex.ExpandBoolPointer(m.NeverExpires),
		ReservedForInfoblox: flex.ExpandStringPointer(m.ReservedForInfoblox),
		Username:            flex.ExpandStringPointer(m.Username),
	}
	return to
}

func FlattenMacfilteraddress(ctx context.Context, from *dhcp.Macfilteraddress, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MacfilteraddressAttrTypes)
	}
	m := MacfilteraddressModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, MacfilteraddressAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MacfilteraddressModel) Flatten(ctx context.Context, from *dhcp.Macfilteraddress, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MacfilteraddressModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AuthenticationTime = flex.FlattenInt64Pointer(from.AuthenticationTime)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExpirationTime = flex.FlattenInt64Pointer(from.ExpirationTime)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.Fingerprint = flex.FlattenStringPointer(from.Fingerprint)
	m.GuestCustomField1 = flex.FlattenStringPointer(from.GuestCustomField1)
	m.GuestCustomField2 = flex.FlattenStringPointer(from.GuestCustomField2)
	m.GuestCustomField3 = flex.FlattenStringPointer(from.GuestCustomField3)
	m.GuestCustomField4 = flex.FlattenStringPointer(from.GuestCustomField4)
	m.GuestEmail = flex.FlattenStringPointer(from.GuestEmail)
	m.GuestFirstName = flex.FlattenStringPointer(from.GuestFirstName)
	m.GuestLastName = flex.FlattenStringPointer(from.GuestLastName)
	m.GuestMiddleName = flex.FlattenStringPointer(from.GuestMiddleName)
	m.GuestPhone = flex.FlattenStringPointer(from.GuestPhone)
	m.IsRegisteredUser = types.BoolPointerValue(from.IsRegisteredUser)
	m.Mac = flex.FlattenMACAddr(from.Mac)
	m.NeverExpires = types.BoolPointerValue(from.NeverExpires)
	m.ReservedForInfoblox = flex.FlattenStringPointer(from.ReservedForInfoblox)
	m.Username = flex.FlattenStringPointer(from.Username)
}
