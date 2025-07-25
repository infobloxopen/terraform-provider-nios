/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
)

// checks if the GridDnsDnssecTrustedKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GridDnsDnssecTrustedKeys{}

// GridDnsDnssecTrustedKeys struct for GridDnsDnssecTrustedKeys
type GridDnsDnssecTrustedKeys struct {
	// The FQDN of the domain for which the member validates responses to recursive queries.
	Fqdn *string `json:"fqdn,omitempty"`
	// The DNSSEC algorithm used to generate the key.
	Algorithm *string `json:"algorithm,omitempty"`
	// The DNSSEC key.
	Key *string `json:"key,omitempty"`
	// The secure entry point flag, if set it means this is a KSK configuration.
	SecureEntryPoint *bool `json:"secure_entry_point,omitempty"`
	// Responses must be DNSSEC secure for this hierarchy/domain.
	DnssecMustBeSecure   *bool `json:"dnssec_must_be_secure,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GridDnsDnssecTrustedKeys GridDnsDnssecTrustedKeys

// NewGridDnsDnssecTrustedKeys instantiates a new GridDnsDnssecTrustedKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridDnsDnssecTrustedKeys() *GridDnsDnssecTrustedKeys {
	this := GridDnsDnssecTrustedKeys{}
	return &this
}

// NewGridDnsDnssecTrustedKeysWithDefaults instantiates a new GridDnsDnssecTrustedKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridDnsDnssecTrustedKeysWithDefaults() *GridDnsDnssecTrustedKeys {
	this := GridDnsDnssecTrustedKeys{}
	return &this
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *GridDnsDnssecTrustedKeys) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDnssecTrustedKeys) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *GridDnsDnssecTrustedKeys) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *GridDnsDnssecTrustedKeys) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *GridDnsDnssecTrustedKeys) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDnssecTrustedKeys) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *GridDnsDnssecTrustedKeys) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *GridDnsDnssecTrustedKeys) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *GridDnsDnssecTrustedKeys) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDnssecTrustedKeys) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *GridDnsDnssecTrustedKeys) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *GridDnsDnssecTrustedKeys) SetKey(v string) {
	o.Key = &v
}

// GetSecureEntryPoint returns the SecureEntryPoint field value if set, zero value otherwise.
func (o *GridDnsDnssecTrustedKeys) GetSecureEntryPoint() bool {
	if o == nil || IsNil(o.SecureEntryPoint) {
		var ret bool
		return ret
	}
	return *o.SecureEntryPoint
}

// GetSecureEntryPointOk returns a tuple with the SecureEntryPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDnssecTrustedKeys) GetSecureEntryPointOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureEntryPoint) {
		return nil, false
	}
	return o.SecureEntryPoint, true
}

// HasSecureEntryPoint returns a boolean if a field has been set.
func (o *GridDnsDnssecTrustedKeys) HasSecureEntryPoint() bool {
	if o != nil && !IsNil(o.SecureEntryPoint) {
		return true
	}

	return false
}

// SetSecureEntryPoint gets a reference to the given bool and assigns it to the SecureEntryPoint field.
func (o *GridDnsDnssecTrustedKeys) SetSecureEntryPoint(v bool) {
	o.SecureEntryPoint = &v
}

// GetDnssecMustBeSecure returns the DnssecMustBeSecure field value if set, zero value otherwise.
func (o *GridDnsDnssecTrustedKeys) GetDnssecMustBeSecure() bool {
	if o == nil || IsNil(o.DnssecMustBeSecure) {
		var ret bool
		return ret
	}
	return *o.DnssecMustBeSecure
}

// GetDnssecMustBeSecureOk returns a tuple with the DnssecMustBeSecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDnssecTrustedKeys) GetDnssecMustBeSecureOk() (*bool, bool) {
	if o == nil || IsNil(o.DnssecMustBeSecure) {
		return nil, false
	}
	return o.DnssecMustBeSecure, true
}

// HasDnssecMustBeSecure returns a boolean if a field has been set.
func (o *GridDnsDnssecTrustedKeys) HasDnssecMustBeSecure() bool {
	if o != nil && !IsNil(o.DnssecMustBeSecure) {
		return true
	}

	return false
}

// SetDnssecMustBeSecure gets a reference to the given bool and assigns it to the DnssecMustBeSecure field.
func (o *GridDnsDnssecTrustedKeys) SetDnssecMustBeSecure(v bool) {
	o.DnssecMustBeSecure = &v
}

func (o GridDnsDnssecTrustedKeys) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GridDnsDnssecTrustedKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.SecureEntryPoint) {
		toSerialize["secure_entry_point"] = o.SecureEntryPoint
	}
	if !IsNil(o.DnssecMustBeSecure) {
		toSerialize["dnssec_must_be_secure"] = o.DnssecMustBeSecure
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GridDnsDnssecTrustedKeys) UnmarshalJSON(data []byte) (err error) {
	varGridDnsDnssecTrustedKeys := _GridDnsDnssecTrustedKeys{}

	err = json.Unmarshal(data, &varGridDnsDnssecTrustedKeys)

	if err != nil {
		return err
	}

	*o = GridDnsDnssecTrustedKeys(varGridDnsDnssecTrustedKeys)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fqdn")
		delete(additionalProperties, "algorithm")
		delete(additionalProperties, "key")
		delete(additionalProperties, "secure_entry_point")
		delete(additionalProperties, "dnssec_must_be_secure")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGridDnsDnssecTrustedKeys struct {
	value *GridDnsDnssecTrustedKeys
	isSet bool
}

func (v NullableGridDnsDnssecTrustedKeys) Get() *GridDnsDnssecTrustedKeys {
	return v.value
}

func (v *NullableGridDnsDnssecTrustedKeys) Set(val *GridDnsDnssecTrustedKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableGridDnsDnssecTrustedKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableGridDnsDnssecTrustedKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridDnsDnssecTrustedKeys(val *GridDnsDnssecTrustedKeys) *NullableGridDnsDnssecTrustedKeys {
	return &NullableGridDnsDnssecTrustedKeys{value: val, isSet: true}
}

func (v NullableGridDnsDnssecTrustedKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridDnsDnssecTrustedKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
