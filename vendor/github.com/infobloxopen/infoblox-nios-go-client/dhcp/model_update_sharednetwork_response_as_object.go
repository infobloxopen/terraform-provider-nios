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

// checks if the UpdateSharednetworkResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSharednetworkResponseAsObject{}

// UpdateSharednetworkResponseAsObject The response format to update __Sharednetwork__ in object format.
type UpdateSharednetworkResponseAsObject struct {
	Result               *Sharednetwork `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateSharednetworkResponseAsObject UpdateSharednetworkResponseAsObject

// NewUpdateSharednetworkResponseAsObject instantiates a new UpdateSharednetworkResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSharednetworkResponseAsObject() *UpdateSharednetworkResponseAsObject {
	this := UpdateSharednetworkResponseAsObject{}
	return &this
}

// NewUpdateSharednetworkResponseAsObjectWithDefaults instantiates a new UpdateSharednetworkResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSharednetworkResponseAsObjectWithDefaults() *UpdateSharednetworkResponseAsObject {
	this := UpdateSharednetworkResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateSharednetworkResponseAsObject) GetResult() Sharednetwork {
	if o == nil || IsNil(o.Result) {
		var ret Sharednetwork
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSharednetworkResponseAsObject) GetResultOk() (*Sharednetwork, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateSharednetworkResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Sharednetwork and assigns it to the Result field.
func (o *UpdateSharednetworkResponseAsObject) SetResult(v Sharednetwork) {
	o.Result = &v
}

func (o UpdateSharednetworkResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSharednetworkResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateSharednetworkResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateSharednetworkResponseAsObject := _UpdateSharednetworkResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateSharednetworkResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateSharednetworkResponseAsObject(varUpdateSharednetworkResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateSharednetworkResponseAsObject struct {
	value *UpdateSharednetworkResponseAsObject
	isSet bool
}

func (v NullableUpdateSharednetworkResponseAsObject) Get() *UpdateSharednetworkResponseAsObject {
	return v.value
}

func (v *NullableUpdateSharednetworkResponseAsObject) Set(val *UpdateSharednetworkResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSharednetworkResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSharednetworkResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSharednetworkResponseAsObject(val *UpdateSharednetworkResponseAsObject) *NullableUpdateSharednetworkResponseAsObject {
	return &NullableUpdateSharednetworkResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateSharednetworkResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSharednetworkResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
