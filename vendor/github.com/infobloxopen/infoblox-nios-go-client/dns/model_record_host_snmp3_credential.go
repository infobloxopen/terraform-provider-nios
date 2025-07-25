/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the RecordHostSnmp3Credential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordHostSnmp3Credential{}

// RecordHostSnmp3Credential struct for RecordHostSnmp3Credential
type RecordHostSnmp3Credential struct {
	// The SNMPv3 user name.
	User *string `json:"user,omitempty"`
	// Authentication protocol for the SNMPv3 user.
	AuthenticationProtocol *string `json:"authentication_protocol,omitempty"`
	// Authentication password for the SNMPv3 user.
	AuthenticationPassword *string `json:"authentication_password,omitempty"`
	// Privacy protocol for the SNMPv3 user.
	PrivacyProtocol *string `json:"privacy_protocol,omitempty"`
	// Privacy password for the SNMPv3 user.
	PrivacyPassword *string `json:"privacy_password,omitempty"`
	// Comments for the SNMPv3 user.
	Comment *string `json:"comment,omitempty"`
	// Group for the SNMPv3 credential.
	CredentialGroup      *string `json:"credential_group,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecordHostSnmp3Credential RecordHostSnmp3Credential

// NewRecordHostSnmp3Credential instantiates a new RecordHostSnmp3Credential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordHostSnmp3Credential() *RecordHostSnmp3Credential {
	this := RecordHostSnmp3Credential{}
	return &this
}

// NewRecordHostSnmp3CredentialWithDefaults instantiates a new RecordHostSnmp3Credential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordHostSnmp3CredentialWithDefaults() *RecordHostSnmp3Credential {
	this := RecordHostSnmp3Credential{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RecordHostSnmp3Credential) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordHostSnmp3Credential) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RecordHostSnmp3Credential) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *RecordHostSnmp3Credential) SetUser(v string) {
	o.User = &v
}

// GetAuthenticationProtocol returns the AuthenticationProtocol field value if set, zero value otherwise.
func (o *RecordHostSnmp3Credential) GetAuthenticationProtocol() string {
	if o == nil || IsNil(o.AuthenticationProtocol) {
		var ret string
		return ret
	}
	return *o.AuthenticationProtocol
}

// GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordHostSnmp3Credential) GetAuthenticationProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationProtocol) {
		return nil, false
	}
	return o.AuthenticationProtocol, true
}

// HasAuthenticationProtocol returns a boolean if a field has been set.
func (o *RecordHostSnmp3Credential) HasAuthenticationProtocol() bool {
	if o != nil && !IsNil(o.AuthenticationProtocol) {
		return true
	}

	return false
}

// SetAuthenticationProtocol gets a reference to the given string and assigns it to the AuthenticationProtocol field.
func (o *RecordHostSnmp3Credential) SetAuthenticationProtocol(v string) {
	o.AuthenticationProtocol = &v
}

// GetAuthenticationPassword returns the AuthenticationPassword field value if set, zero value otherwise.
func (o *RecordHostSnmp3Credential) GetAuthenticationPassword() string {
	if o == nil || IsNil(o.AuthenticationPassword) {
		var ret string
		return ret
	}
	return *o.AuthenticationPassword
}

// GetAuthenticationPasswordOk returns a tuple with the AuthenticationPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordHostSnmp3Credential) GetAuthenticationPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationPassword) {
		return nil, false
	}
	return o.AuthenticationPassword, true
}

// HasAuthenticationPassword returns a boolean if a field has been set.
func (o *RecordHostSnmp3Credential) HasAuthenticationPassword() bool {
	if o != nil && !IsNil(o.AuthenticationPassword) {
		return true
	}

	return false
}

// SetAuthenticationPassword gets a reference to the given string and assigns it to the AuthenticationPassword field.
func (o *RecordHostSnmp3Credential) SetAuthenticationPassword(v string) {
	o.AuthenticationPassword = &v
}

// GetPrivacyProtocol returns the PrivacyProtocol field value if set, zero value otherwise.
func (o *RecordHostSnmp3Credential) GetPrivacyProtocol() string {
	if o == nil || IsNil(o.PrivacyProtocol) {
		var ret string
		return ret
	}
	return *o.PrivacyProtocol
}

// GetPrivacyProtocolOk returns a tuple with the PrivacyProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordHostSnmp3Credential) GetPrivacyProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyProtocol) {
		return nil, false
	}
	return o.PrivacyProtocol, true
}

// HasPrivacyProtocol returns a boolean if a field has been set.
func (o *RecordHostSnmp3Credential) HasPrivacyProtocol() bool {
	if o != nil && !IsNil(o.PrivacyProtocol) {
		return true
	}

	return false
}

// SetPrivacyProtocol gets a reference to the given string and assigns it to the PrivacyProtocol field.
func (o *RecordHostSnmp3Credential) SetPrivacyProtocol(v string) {
	o.PrivacyProtocol = &v
}

// GetPrivacyPassword returns the PrivacyPassword field value if set, zero value otherwise.
func (o *RecordHostSnmp3Credential) GetPrivacyPassword() string {
	if o == nil || IsNil(o.PrivacyPassword) {
		var ret string
		return ret
	}
	return *o.PrivacyPassword
}

// GetPrivacyPasswordOk returns a tuple with the PrivacyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordHostSnmp3Credential) GetPrivacyPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyPassword) {
		return nil, false
	}
	return o.PrivacyPassword, true
}

// HasPrivacyPassword returns a boolean if a field has been set.
func (o *RecordHostSnmp3Credential) HasPrivacyPassword() bool {
	if o != nil && !IsNil(o.PrivacyPassword) {
		return true
	}

	return false
}

// SetPrivacyPassword gets a reference to the given string and assigns it to the PrivacyPassword field.
func (o *RecordHostSnmp3Credential) SetPrivacyPassword(v string) {
	o.PrivacyPassword = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RecordHostSnmp3Credential) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordHostSnmp3Credential) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *RecordHostSnmp3Credential) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *RecordHostSnmp3Credential) SetComment(v string) {
	o.Comment = &v
}

// GetCredentialGroup returns the CredentialGroup field value if set, zero value otherwise.
func (o *RecordHostSnmp3Credential) GetCredentialGroup() string {
	if o == nil || IsNil(o.CredentialGroup) {
		var ret string
		return ret
	}
	return *o.CredentialGroup
}

// GetCredentialGroupOk returns a tuple with the CredentialGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordHostSnmp3Credential) GetCredentialGroupOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialGroup) {
		return nil, false
	}
	return o.CredentialGroup, true
}

// HasCredentialGroup returns a boolean if a field has been set.
func (o *RecordHostSnmp3Credential) HasCredentialGroup() bool {
	if o != nil && !IsNil(o.CredentialGroup) {
		return true
	}

	return false
}

// SetCredentialGroup gets a reference to the given string and assigns it to the CredentialGroup field.
func (o *RecordHostSnmp3Credential) SetCredentialGroup(v string) {
	o.CredentialGroup = &v
}

func (o RecordHostSnmp3Credential) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordHostSnmp3Credential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.AuthenticationProtocol) {
		toSerialize["authentication_protocol"] = o.AuthenticationProtocol
	}
	if !IsNil(o.AuthenticationPassword) {
		toSerialize["authentication_password"] = o.AuthenticationPassword
	}
	if !IsNil(o.PrivacyProtocol) {
		toSerialize["privacy_protocol"] = o.PrivacyProtocol
	}
	if !IsNil(o.PrivacyPassword) {
		toSerialize["privacy_password"] = o.PrivacyPassword
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CredentialGroup) {
		toSerialize["credential_group"] = o.CredentialGroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecordHostSnmp3Credential) UnmarshalJSON(data []byte) (err error) {
	varRecordHostSnmp3Credential := _RecordHostSnmp3Credential{}

	err = json.Unmarshal(data, &varRecordHostSnmp3Credential)

	if err != nil {
		return err
	}

	*o = RecordHostSnmp3Credential(varRecordHostSnmp3Credential)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		delete(additionalProperties, "authentication_protocol")
		delete(additionalProperties, "authentication_password")
		delete(additionalProperties, "privacy_protocol")
		delete(additionalProperties, "privacy_password")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "credential_group")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecordHostSnmp3Credential struct {
	value *RecordHostSnmp3Credential
	isSet bool
}

func (v NullableRecordHostSnmp3Credential) Get() *RecordHostSnmp3Credential {
	return v.value
}

func (v *NullableRecordHostSnmp3Credential) Set(val *RecordHostSnmp3Credential) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordHostSnmp3Credential) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordHostSnmp3Credential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordHostSnmp3Credential(val *RecordHostSnmp3Credential) *NullableRecordHostSnmp3Credential {
	return &NullableRecordHostSnmp3Credential{value: val, isSet: true}
}

func (v NullableRecordHostSnmp3Credential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordHostSnmp3Credential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
