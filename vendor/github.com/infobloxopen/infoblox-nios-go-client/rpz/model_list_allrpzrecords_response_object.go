/*
Infoblox RPZ API

OpenAPI specification for Infoblox NIOS WAPI RPZ objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rpz

import (
	"encoding/json"
)

// checks if the ListAllrpzrecordsResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAllrpzrecordsResponseObject{}

// ListAllrpzrecordsResponseObject The response format to retrieve __Allrpzrecords__ objects.
type ListAllrpzrecordsResponseObject struct {
	Result               []Allrpzrecords `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListAllrpzrecordsResponseObject ListAllrpzrecordsResponseObject

// NewListAllrpzrecordsResponseObject instantiates a new ListAllrpzrecordsResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllrpzrecordsResponseObject() *ListAllrpzrecordsResponseObject {
	this := ListAllrpzrecordsResponseObject{}
	return &this
}

// NewListAllrpzrecordsResponseObjectWithDefaults instantiates a new ListAllrpzrecordsResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllrpzrecordsResponseObjectWithDefaults() *ListAllrpzrecordsResponseObject {
	this := ListAllrpzrecordsResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListAllrpzrecordsResponseObject) GetResult() []Allrpzrecords {
	if o == nil || IsNil(o.Result) {
		var ret []Allrpzrecords
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllrpzrecordsResponseObject) GetResultOk() ([]Allrpzrecords, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListAllrpzrecordsResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []Allrpzrecords and assigns it to the Result field.
func (o *ListAllrpzrecordsResponseObject) SetResult(v []Allrpzrecords) {
	o.Result = v
}

func (o ListAllrpzrecordsResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAllrpzrecordsResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListAllrpzrecordsResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListAllrpzrecordsResponseObject := _ListAllrpzrecordsResponseObject{}

	err = json.Unmarshal(data, &varListAllrpzrecordsResponseObject)

	if err != nil {
		return err
	}

	*o = ListAllrpzrecordsResponseObject(varListAllrpzrecordsResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAllrpzrecordsResponseObject struct {
	value *ListAllrpzrecordsResponseObject
	isSet bool
}

func (v NullableListAllrpzrecordsResponseObject) Get() *ListAllrpzrecordsResponseObject {
	return v.value
}

func (v *NullableListAllrpzrecordsResponseObject) Set(val *ListAllrpzrecordsResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllrpzrecordsResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllrpzrecordsResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllrpzrecordsResponseObject(val *ListAllrpzrecordsResponseObject) *NullableListAllrpzrecordsResponseObject {
	return &NullableListAllrpzrecordsResponseObject{value: val, isSet: true}
}

func (v NullableListAllrpzrecordsResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllrpzrecordsResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
