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

// checks if the UpdateZoneForwardResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateZoneForwardResponseAsObject{}

// UpdateZoneForwardResponseAsObject The response format to update __ZoneForward__ in object format.
type UpdateZoneForwardResponseAsObject struct {
	Result               *ZoneForward `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateZoneForwardResponseAsObject UpdateZoneForwardResponseAsObject

// NewUpdateZoneForwardResponseAsObject instantiates a new UpdateZoneForwardResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateZoneForwardResponseAsObject() *UpdateZoneForwardResponseAsObject {
	this := UpdateZoneForwardResponseAsObject{}
	return &this
}

// NewUpdateZoneForwardResponseAsObjectWithDefaults instantiates a new UpdateZoneForwardResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateZoneForwardResponseAsObjectWithDefaults() *UpdateZoneForwardResponseAsObject {
	this := UpdateZoneForwardResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateZoneForwardResponseAsObject) GetResult() ZoneForward {
	if o == nil || IsNil(o.Result) {
		var ret ZoneForward
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateZoneForwardResponseAsObject) GetResultOk() (*ZoneForward, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateZoneForwardResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ZoneForward and assigns it to the Result field.
func (o *UpdateZoneForwardResponseAsObject) SetResult(v ZoneForward) {
	o.Result = &v
}

func (o UpdateZoneForwardResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateZoneForwardResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateZoneForwardResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateZoneForwardResponseAsObject := _UpdateZoneForwardResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateZoneForwardResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateZoneForwardResponseAsObject(varUpdateZoneForwardResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateZoneForwardResponseAsObject struct {
	value *UpdateZoneForwardResponseAsObject
	isSet bool
}

func (v NullableUpdateZoneForwardResponseAsObject) Get() *UpdateZoneForwardResponseAsObject {
	return v.value
}

func (v *NullableUpdateZoneForwardResponseAsObject) Set(val *UpdateZoneForwardResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateZoneForwardResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateZoneForwardResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateZoneForwardResponseAsObject(val *UpdateZoneForwardResponseAsObject) *NullableUpdateZoneForwardResponseAsObject {
	return &NullableUpdateZoneForwardResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateZoneForwardResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateZoneForwardResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
