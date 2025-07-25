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

// checks if the UpdateOrderedrangesResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrderedrangesResponseAsObject{}

// UpdateOrderedrangesResponseAsObject The response format to update __Orderedranges__ in object format.
type UpdateOrderedrangesResponseAsObject struct {
	Result               *Orderedranges `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateOrderedrangesResponseAsObject UpdateOrderedrangesResponseAsObject

// NewUpdateOrderedrangesResponseAsObject instantiates a new UpdateOrderedrangesResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrderedrangesResponseAsObject() *UpdateOrderedrangesResponseAsObject {
	this := UpdateOrderedrangesResponseAsObject{}
	return &this
}

// NewUpdateOrderedrangesResponseAsObjectWithDefaults instantiates a new UpdateOrderedrangesResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrderedrangesResponseAsObjectWithDefaults() *UpdateOrderedrangesResponseAsObject {
	this := UpdateOrderedrangesResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateOrderedrangesResponseAsObject) GetResult() Orderedranges {
	if o == nil || IsNil(o.Result) {
		var ret Orderedranges
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderedrangesResponseAsObject) GetResultOk() (*Orderedranges, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateOrderedrangesResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Orderedranges and assigns it to the Result field.
func (o *UpdateOrderedrangesResponseAsObject) SetResult(v Orderedranges) {
	o.Result = &v
}

func (o UpdateOrderedrangesResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrderedrangesResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateOrderedrangesResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateOrderedrangesResponseAsObject := _UpdateOrderedrangesResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateOrderedrangesResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateOrderedrangesResponseAsObject(varUpdateOrderedrangesResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateOrderedrangesResponseAsObject struct {
	value *UpdateOrderedrangesResponseAsObject
	isSet bool
}

func (v NullableUpdateOrderedrangesResponseAsObject) Get() *UpdateOrderedrangesResponseAsObject {
	return v.value
}

func (v *NullableUpdateOrderedrangesResponseAsObject) Set(val *UpdateOrderedrangesResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrderedrangesResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrderedrangesResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrderedrangesResponseAsObject(val *UpdateOrderedrangesResponseAsObject) *NullableUpdateOrderedrangesResponseAsObject {
	return &NullableUpdateOrderedrangesResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateOrderedrangesResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrderedrangesResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
