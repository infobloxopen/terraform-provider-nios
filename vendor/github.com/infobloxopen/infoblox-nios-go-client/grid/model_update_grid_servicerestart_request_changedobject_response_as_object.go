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

// checks if the UpdateGridServicerestartRequestChangedobjectResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGridServicerestartRequestChangedobjectResponseAsObject{}

// UpdateGridServicerestartRequestChangedobjectResponseAsObject The response format to update __GridServicerestartRequestChangedobject__ in object format.
type UpdateGridServicerestartRequestChangedobjectResponseAsObject struct {
	Result               *GridServicerestartRequestChangedobject `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateGridServicerestartRequestChangedobjectResponseAsObject UpdateGridServicerestartRequestChangedobjectResponseAsObject

// NewUpdateGridServicerestartRequestChangedobjectResponseAsObject instantiates a new UpdateGridServicerestartRequestChangedobjectResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGridServicerestartRequestChangedobjectResponseAsObject() *UpdateGridServicerestartRequestChangedobjectResponseAsObject {
	this := UpdateGridServicerestartRequestChangedobjectResponseAsObject{}
	return &this
}

// NewUpdateGridServicerestartRequestChangedobjectResponseAsObjectWithDefaults instantiates a new UpdateGridServicerestartRequestChangedobjectResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGridServicerestartRequestChangedobjectResponseAsObjectWithDefaults() *UpdateGridServicerestartRequestChangedobjectResponseAsObject {
	this := UpdateGridServicerestartRequestChangedobjectResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateGridServicerestartRequestChangedobjectResponseAsObject) GetResult() GridServicerestartRequestChangedobject {
	if o == nil || IsNil(o.Result) {
		var ret GridServicerestartRequestChangedobject
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGridServicerestartRequestChangedobjectResponseAsObject) GetResultOk() (*GridServicerestartRequestChangedobject, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateGridServicerestartRequestChangedobjectResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given GridServicerestartRequestChangedobject and assigns it to the Result field.
func (o *UpdateGridServicerestartRequestChangedobjectResponseAsObject) SetResult(v GridServicerestartRequestChangedobject) {
	o.Result = &v
}

func (o UpdateGridServicerestartRequestChangedobjectResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGridServicerestartRequestChangedobjectResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateGridServicerestartRequestChangedobjectResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateGridServicerestartRequestChangedobjectResponseAsObject := _UpdateGridServicerestartRequestChangedobjectResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateGridServicerestartRequestChangedobjectResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateGridServicerestartRequestChangedobjectResponseAsObject(varUpdateGridServicerestartRequestChangedobjectResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateGridServicerestartRequestChangedobjectResponseAsObject struct {
	value *UpdateGridServicerestartRequestChangedobjectResponseAsObject
	isSet bool
}

func (v NullableUpdateGridServicerestartRequestChangedobjectResponseAsObject) Get() *UpdateGridServicerestartRequestChangedobjectResponseAsObject {
	return v.value
}

func (v *NullableUpdateGridServicerestartRequestChangedobjectResponseAsObject) Set(val *UpdateGridServicerestartRequestChangedobjectResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGridServicerestartRequestChangedobjectResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGridServicerestartRequestChangedobjectResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGridServicerestartRequestChangedobjectResponseAsObject(val *UpdateGridServicerestartRequestChangedobjectResponseAsObject) *NullableUpdateGridServicerestartRequestChangedobjectResponseAsObject {
	return &NullableUpdateGridServicerestartRequestChangedobjectResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateGridServicerestartRequestChangedobjectResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGridServicerestartRequestChangedobjectResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
