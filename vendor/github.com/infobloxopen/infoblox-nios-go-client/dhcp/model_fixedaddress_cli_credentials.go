/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
)

// checks if the FixedaddressCliCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FixedaddressCliCredentials{}

// FixedaddressCliCredentials struct for FixedaddressCliCredentials
type FixedaddressCliCredentials struct {
	// The CLI user name.
	User *string `json:"user,omitempty"`
	// The CLI password.
	Password *string `json:"password,omitempty"`
	// The type of the credential.
	CredentialType *string `json:"credential_type,omitempty"`
	// The commment for the credential.
	Comment *string `json:"comment,omitempty"`
	// The Credentials ID.
	Id *int64 `json:"id,omitempty"`
	// Group for the CLI credential.
	CredentialGroup      *string `json:"credential_group,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FixedaddressCliCredentials FixedaddressCliCredentials

// NewFixedaddressCliCredentials instantiates a new FixedaddressCliCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedaddressCliCredentials() *FixedaddressCliCredentials {
	this := FixedaddressCliCredentials{}
	return &this
}

// NewFixedaddressCliCredentialsWithDefaults instantiates a new FixedaddressCliCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedaddressCliCredentialsWithDefaults() *FixedaddressCliCredentials {
	this := FixedaddressCliCredentials{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *FixedaddressCliCredentials) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedaddressCliCredentials) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *FixedaddressCliCredentials) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *FixedaddressCliCredentials) SetUser(v string) {
	o.User = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *FixedaddressCliCredentials) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedaddressCliCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *FixedaddressCliCredentials) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *FixedaddressCliCredentials) SetPassword(v string) {
	o.Password = &v
}

// GetCredentialType returns the CredentialType field value if set, zero value otherwise.
func (o *FixedaddressCliCredentials) GetCredentialType() string {
	if o == nil || IsNil(o.CredentialType) {
		var ret string
		return ret
	}
	return *o.CredentialType
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedaddressCliCredentials) GetCredentialTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialType) {
		return nil, false
	}
	return o.CredentialType, true
}

// HasCredentialType returns a boolean if a field has been set.
func (o *FixedaddressCliCredentials) HasCredentialType() bool {
	if o != nil && !IsNil(o.CredentialType) {
		return true
	}

	return false
}

// SetCredentialType gets a reference to the given string and assigns it to the CredentialType field.
func (o *FixedaddressCliCredentials) SetCredentialType(v string) {
	o.CredentialType = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *FixedaddressCliCredentials) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedaddressCliCredentials) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *FixedaddressCliCredentials) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *FixedaddressCliCredentials) SetComment(v string) {
	o.Comment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FixedaddressCliCredentials) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedaddressCliCredentials) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FixedaddressCliCredentials) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *FixedaddressCliCredentials) SetId(v int64) {
	o.Id = &v
}

// GetCredentialGroup returns the CredentialGroup field value if set, zero value otherwise.
func (o *FixedaddressCliCredentials) GetCredentialGroup() string {
	if o == nil || IsNil(o.CredentialGroup) {
		var ret string
		return ret
	}
	return *o.CredentialGroup
}

// GetCredentialGroupOk returns a tuple with the CredentialGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedaddressCliCredentials) GetCredentialGroupOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialGroup) {
		return nil, false
	}
	return o.CredentialGroup, true
}

// HasCredentialGroup returns a boolean if a field has been set.
func (o *FixedaddressCliCredentials) HasCredentialGroup() bool {
	if o != nil && !IsNil(o.CredentialGroup) {
		return true
	}

	return false
}

// SetCredentialGroup gets a reference to the given string and assigns it to the CredentialGroup field.
func (o *FixedaddressCliCredentials) SetCredentialGroup(v string) {
	o.CredentialGroup = &v
}

func (o FixedaddressCliCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FixedaddressCliCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.CredentialType) {
		toSerialize["credential_type"] = o.CredentialType
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CredentialGroup) {
		toSerialize["credential_group"] = o.CredentialGroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FixedaddressCliCredentials) UnmarshalJSON(data []byte) (err error) {
	varFixedaddressCliCredentials := _FixedaddressCliCredentials{}

	err = json.Unmarshal(data, &varFixedaddressCliCredentials)

	if err != nil {
		return err
	}

	*o = FixedaddressCliCredentials(varFixedaddressCliCredentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		delete(additionalProperties, "password")
		delete(additionalProperties, "credential_type")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "id")
		delete(additionalProperties, "credential_group")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFixedaddressCliCredentials struct {
	value *FixedaddressCliCredentials
	isSet bool
}

func (v NullableFixedaddressCliCredentials) Get() *FixedaddressCliCredentials {
	return v.value
}

func (v *NullableFixedaddressCliCredentials) Set(val *FixedaddressCliCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedaddressCliCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedaddressCliCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedaddressCliCredentials(val *FixedaddressCliCredentials) *NullableFixedaddressCliCredentials {
	return &NullableFixedaddressCliCredentials{value: val, isSet: true}
}

func (v NullableFixedaddressCliCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedaddressCliCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
