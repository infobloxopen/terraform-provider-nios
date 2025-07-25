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

// checks if the GridntpsettingntpaclaclistAddressAc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GridntpsettingntpaclaclistAddressAc{}

// GridntpsettingntpaclaclistAddressAc struct for GridntpsettingntpaclaclistAddressAc
type GridntpsettingntpaclaclistAddressAc struct {
	// The address this rule applies to or \"Any\".
	Address *string `json:"address,omitempty"`
	// The permission to use for this address.
	Permission           *string `json:"permission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GridntpsettingntpaclaclistAddressAc GridntpsettingntpaclaclistAddressAc

// NewGridntpsettingntpaclaclistAddressAc instantiates a new GridntpsettingntpaclaclistAddressAc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridntpsettingntpaclaclistAddressAc() *GridntpsettingntpaclaclistAddressAc {
	this := GridntpsettingntpaclaclistAddressAc{}
	return &this
}

// NewGridntpsettingntpaclaclistAddressAcWithDefaults instantiates a new GridntpsettingntpaclaclistAddressAc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridntpsettingntpaclaclistAddressAcWithDefaults() *GridntpsettingntpaclaclistAddressAc {
	this := GridntpsettingntpaclaclistAddressAc{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GridntpsettingntpaclaclistAddressAc) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridntpsettingntpaclaclistAddressAc) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GridntpsettingntpaclaclistAddressAc) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GridntpsettingntpaclaclistAddressAc) SetAddress(v string) {
	o.Address = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *GridntpsettingntpaclaclistAddressAc) GetPermission() string {
	if o == nil || IsNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridntpsettingntpaclaclistAddressAc) GetPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *GridntpsettingntpaclaclistAddressAc) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *GridntpsettingntpaclaclistAddressAc) SetPermission(v string) {
	o.Permission = &v
}

func (o GridntpsettingntpaclaclistAddressAc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GridntpsettingntpaclaclistAddressAc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GridntpsettingntpaclaclistAddressAc) UnmarshalJSON(data []byte) (err error) {
	varGridntpsettingntpaclaclistAddressAc := _GridntpsettingntpaclaclistAddressAc{}

	err = json.Unmarshal(data, &varGridntpsettingntpaclaclistAddressAc)

	if err != nil {
		return err
	}

	*o = GridntpsettingntpaclaclistAddressAc(varGridntpsettingntpaclaclistAddressAc)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "permission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGridntpsettingntpaclaclistAddressAc struct {
	value *GridntpsettingntpaclaclistAddressAc
	isSet bool
}

func (v NullableGridntpsettingntpaclaclistAddressAc) Get() *GridntpsettingntpaclaclistAddressAc {
	return v.value
}

func (v *NullableGridntpsettingntpaclaclistAddressAc) Set(val *GridntpsettingntpaclaclistAddressAc) {
	v.value = val
	v.isSet = true
}

func (v NullableGridntpsettingntpaclaclistAddressAc) IsSet() bool {
	return v.isSet
}

func (v *NullableGridntpsettingntpaclaclistAddressAc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridntpsettingntpaclaclistAddressAc(val *GridntpsettingntpaclaclistAddressAc) *NullableGridntpsettingntpaclaclistAddressAc {
	return &NullableGridntpsettingntpaclaclistAddressAc{value: val, isSet: true}
}

func (v NullableGridntpsettingntpaclaclistAddressAc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridntpsettingntpaclaclistAddressAc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
