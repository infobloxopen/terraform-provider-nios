/*
Infoblox SECURITY API

OpenAPI specification for Infoblox NIOS WAPI SECURITY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package security

import (
	"encoding/json"
)

// checks if the Snmpuser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Snmpuser{}

// Snmpuser struct for Snmpuser
type Snmpuser struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// Determines an authentication password for the user. This is a write-only attribute.
	AuthenticationPassword *string `json:"authentication_password,omitempty"`
	// The authentication protocol to be used for this user.
	AuthenticationProtocol *string `json:"authentication_protocol,omitempty"`
	// A descriptive comment for the SNMPv3 User.
	Comment *string `json:"comment,omitempty"`
	// Determines if SNMPv3 user is disabled or not.
	Disable *bool `json:"disable,omitempty"`
	// Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.
	ExtAttrs *map[string]ExtAttrs `json:"extattrs,omitempty"`
	// The name of the user.
	Name *string `json:"name,omitempty"`
	// Determines a password for the privacy protocol.
	PrivacyPassword *string `json:"privacy_password,omitempty"`
	// The privacy protocol to be used for this user.
	PrivacyProtocol *string `json:"privacy_protocol,omitempty"`
}

// NewSnmpuser instantiates a new Snmpuser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpuser() *Snmpuser {
	this := Snmpuser{}
	return &this
}

// NewSnmpuserWithDefaults instantiates a new Snmpuser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpuserWithDefaults() *Snmpuser {
	this := Snmpuser{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Snmpuser) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snmpuser) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Snmpuser) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Snmpuser) SetRef(v string) {
	o.Ref = &v
}

// GetAuthenticationPassword returns the AuthenticationPassword field value if set, zero value otherwise.
func (o *Snmpuser) GetAuthenticationPassword() string {
	if o == nil || IsNil(o.AuthenticationPassword) {
		var ret string
		return ret
	}
	return *o.AuthenticationPassword
}

// GetAuthenticationPasswordOk returns a tuple with the AuthenticationPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snmpuser) GetAuthenticationPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationPassword) {
		return nil, false
	}
	return o.AuthenticationPassword, true
}

// HasAuthenticationPassword returns a boolean if a field has been set.
func (o *Snmpuser) HasAuthenticationPassword() bool {
	if o != nil && !IsNil(o.AuthenticationPassword) {
		return true
	}

	return false
}

// SetAuthenticationPassword gets a reference to the given string and assigns it to the AuthenticationPassword field.
func (o *Snmpuser) SetAuthenticationPassword(v string) {
	o.AuthenticationPassword = &v
}

// GetAuthenticationProtocol returns the AuthenticationProtocol field value if set, zero value otherwise.
func (o *Snmpuser) GetAuthenticationProtocol() string {
	if o == nil || IsNil(o.AuthenticationProtocol) {
		var ret string
		return ret
	}
	return *o.AuthenticationProtocol
}

// GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snmpuser) GetAuthenticationProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationProtocol) {
		return nil, false
	}
	return o.AuthenticationProtocol, true
}

// HasAuthenticationProtocol returns a boolean if a field has been set.
func (o *Snmpuser) HasAuthenticationProtocol() bool {
	if o != nil && !IsNil(o.AuthenticationProtocol) {
		return true
	}

	return false
}

// SetAuthenticationProtocol gets a reference to the given string and assigns it to the AuthenticationProtocol field.
func (o *Snmpuser) SetAuthenticationProtocol(v string) {
	o.AuthenticationProtocol = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Snmpuser) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snmpuser) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Snmpuser) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Snmpuser) SetComment(v string) {
	o.Comment = &v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *Snmpuser) GetDisable() bool {
	if o == nil || IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snmpuser) GetDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *Snmpuser) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *Snmpuser) SetDisable(v bool) {
	o.Disable = &v
}

// GetExtAttrs returns the ExtAttrs field value if set, zero value otherwise.
func (o *Snmpuser) GetExtAttrs() map[string]ExtAttrs {
	if o == nil || IsNil(o.ExtAttrs) {
		var ret map[string]ExtAttrs
		return ret
	}
	return *o.ExtAttrs
}

// GetExtAttrsOk returns a tuple with the ExtAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snmpuser) GetExtAttrsOk() (*map[string]ExtAttrs, bool) {
	if o == nil || IsNil(o.ExtAttrs) {
		return nil, false
	}
	return o.ExtAttrs, true
}

// HasExtAttrs returns a boolean if a field has been set.
func (o *Snmpuser) HasExtAttrs() bool {
	if o != nil && !IsNil(o.ExtAttrs) {
		return true
	}

	return false
}

// SetExtAttrs gets a reference to the given map[string]ExtAttrs and assigns it to the ExtAttrs field.
func (o *Snmpuser) SetExtAttrs(v map[string]ExtAttrs) {
	o.ExtAttrs = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Snmpuser) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snmpuser) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Snmpuser) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Snmpuser) SetName(v string) {
	o.Name = &v
}

// GetPrivacyPassword returns the PrivacyPassword field value if set, zero value otherwise.
func (o *Snmpuser) GetPrivacyPassword() string {
	if o == nil || IsNil(o.PrivacyPassword) {
		var ret string
		return ret
	}
	return *o.PrivacyPassword
}

// GetPrivacyPasswordOk returns a tuple with the PrivacyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snmpuser) GetPrivacyPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyPassword) {
		return nil, false
	}
	return o.PrivacyPassword, true
}

// HasPrivacyPassword returns a boolean if a field has been set.
func (o *Snmpuser) HasPrivacyPassword() bool {
	if o != nil && !IsNil(o.PrivacyPassword) {
		return true
	}

	return false
}

// SetPrivacyPassword gets a reference to the given string and assigns it to the PrivacyPassword field.
func (o *Snmpuser) SetPrivacyPassword(v string) {
	o.PrivacyPassword = &v
}

// GetPrivacyProtocol returns the PrivacyProtocol field value if set, zero value otherwise.
func (o *Snmpuser) GetPrivacyProtocol() string {
	if o == nil || IsNil(o.PrivacyProtocol) {
		var ret string
		return ret
	}
	return *o.PrivacyProtocol
}

// GetPrivacyProtocolOk returns a tuple with the PrivacyProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snmpuser) GetPrivacyProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyProtocol) {
		return nil, false
	}
	return o.PrivacyProtocol, true
}

// HasPrivacyProtocol returns a boolean if a field has been set.
func (o *Snmpuser) HasPrivacyProtocol() bool {
	if o != nil && !IsNil(o.PrivacyProtocol) {
		return true
	}

	return false
}

// SetPrivacyProtocol gets a reference to the given string and assigns it to the PrivacyProtocol field.
func (o *Snmpuser) SetPrivacyProtocol(v string) {
	o.PrivacyProtocol = &v
}

func (o Snmpuser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Snmpuser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.AuthenticationPassword) {
		toSerialize["authentication_password"] = o.AuthenticationPassword
	}
	if !IsNil(o.AuthenticationProtocol) {
		toSerialize["authentication_protocol"] = o.AuthenticationProtocol
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	if !IsNil(o.ExtAttrs) {
		toSerialize["extattrs"] = o.ExtAttrs
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PrivacyPassword) {
		toSerialize["privacy_password"] = o.PrivacyPassword
	}
	if !IsNil(o.PrivacyProtocol) {
		toSerialize["privacy_protocol"] = o.PrivacyProtocol
	}
	return toSerialize, nil
}

type NullableSnmpuser struct {
	value *Snmpuser
	isSet bool
}

func (v NullableSnmpuser) Get() *Snmpuser {
	return v.value
}

func (v *NullableSnmpuser) Set(val *Snmpuser) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpuser) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpuser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpuser(val *Snmpuser) *NullableSnmpuser {
	return &NullableSnmpuser{value: val, isSet: true}
}

func (v NullableSnmpuser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpuser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
