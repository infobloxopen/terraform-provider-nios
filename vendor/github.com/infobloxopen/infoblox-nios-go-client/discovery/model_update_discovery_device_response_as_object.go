/*
Infoblox DISCOVERY API

OpenAPI specification for Infoblox NIOS WAPI DISCOVERY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discovery

import (
	"encoding/json"
)

// checks if the UpdateDiscoveryDeviceResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDiscoveryDeviceResponseAsObject{}

// UpdateDiscoveryDeviceResponseAsObject The response format to update __DiscoveryDevice__ in object format.
type UpdateDiscoveryDeviceResponseAsObject struct {
	Result               *DiscoveryDevice `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateDiscoveryDeviceResponseAsObject UpdateDiscoveryDeviceResponseAsObject

// NewUpdateDiscoveryDeviceResponseAsObject instantiates a new UpdateDiscoveryDeviceResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDiscoveryDeviceResponseAsObject() *UpdateDiscoveryDeviceResponseAsObject {
	this := UpdateDiscoveryDeviceResponseAsObject{}
	return &this
}

// NewUpdateDiscoveryDeviceResponseAsObjectWithDefaults instantiates a new UpdateDiscoveryDeviceResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDiscoveryDeviceResponseAsObjectWithDefaults() *UpdateDiscoveryDeviceResponseAsObject {
	this := UpdateDiscoveryDeviceResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateDiscoveryDeviceResponseAsObject) GetResult() DiscoveryDevice {
	if o == nil || IsNil(o.Result) {
		var ret DiscoveryDevice
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDiscoveryDeviceResponseAsObject) GetResultOk() (*DiscoveryDevice, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateDiscoveryDeviceResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given DiscoveryDevice and assigns it to the Result field.
func (o *UpdateDiscoveryDeviceResponseAsObject) SetResult(v DiscoveryDevice) {
	o.Result = &v
}

func (o UpdateDiscoveryDeviceResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDiscoveryDeviceResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateDiscoveryDeviceResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateDiscoveryDeviceResponseAsObject := _UpdateDiscoveryDeviceResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateDiscoveryDeviceResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateDiscoveryDeviceResponseAsObject(varUpdateDiscoveryDeviceResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateDiscoveryDeviceResponseAsObject struct {
	value *UpdateDiscoveryDeviceResponseAsObject
	isSet bool
}

func (v NullableUpdateDiscoveryDeviceResponseAsObject) Get() *UpdateDiscoveryDeviceResponseAsObject {
	return v.value
}

func (v *NullableUpdateDiscoveryDeviceResponseAsObject) Set(val *UpdateDiscoveryDeviceResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDiscoveryDeviceResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDiscoveryDeviceResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDiscoveryDeviceResponseAsObject(val *UpdateDiscoveryDeviceResponseAsObject) *NullableUpdateDiscoveryDeviceResponseAsObject {
	return &NullableUpdateDiscoveryDeviceResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateDiscoveryDeviceResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDiscoveryDeviceResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
