/*
Infoblox THREATINSIGHT API

OpenAPI specification for Infoblox NIOS WAPI THREATINSIGHT objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package threatinsight

import (
	"encoding/json"
)

// checks if the UpdateThreatinsightCloudclientResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateThreatinsightCloudclientResponseAsObject{}

// UpdateThreatinsightCloudclientResponseAsObject The response format to update __ThreatinsightCloudclient__ in object format.
type UpdateThreatinsightCloudclientResponseAsObject struct {
	Result               *ThreatinsightCloudclient `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateThreatinsightCloudclientResponseAsObject UpdateThreatinsightCloudclientResponseAsObject

// NewUpdateThreatinsightCloudclientResponseAsObject instantiates a new UpdateThreatinsightCloudclientResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateThreatinsightCloudclientResponseAsObject() *UpdateThreatinsightCloudclientResponseAsObject {
	this := UpdateThreatinsightCloudclientResponseAsObject{}
	return &this
}

// NewUpdateThreatinsightCloudclientResponseAsObjectWithDefaults instantiates a new UpdateThreatinsightCloudclientResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateThreatinsightCloudclientResponseAsObjectWithDefaults() *UpdateThreatinsightCloudclientResponseAsObject {
	this := UpdateThreatinsightCloudclientResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateThreatinsightCloudclientResponseAsObject) GetResult() ThreatinsightCloudclient {
	if o == nil || IsNil(o.Result) {
		var ret ThreatinsightCloudclient
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreatinsightCloudclientResponseAsObject) GetResultOk() (*ThreatinsightCloudclient, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateThreatinsightCloudclientResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ThreatinsightCloudclient and assigns it to the Result field.
func (o *UpdateThreatinsightCloudclientResponseAsObject) SetResult(v ThreatinsightCloudclient) {
	o.Result = &v
}

func (o UpdateThreatinsightCloudclientResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateThreatinsightCloudclientResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateThreatinsightCloudclientResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateThreatinsightCloudclientResponseAsObject := _UpdateThreatinsightCloudclientResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateThreatinsightCloudclientResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateThreatinsightCloudclientResponseAsObject(varUpdateThreatinsightCloudclientResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateThreatinsightCloudclientResponseAsObject struct {
	value *UpdateThreatinsightCloudclientResponseAsObject
	isSet bool
}

func (v NullableUpdateThreatinsightCloudclientResponseAsObject) Get() *UpdateThreatinsightCloudclientResponseAsObject {
	return v.value
}

func (v *NullableUpdateThreatinsightCloudclientResponseAsObject) Set(val *UpdateThreatinsightCloudclientResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateThreatinsightCloudclientResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateThreatinsightCloudclientResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateThreatinsightCloudclientResponseAsObject(val *UpdateThreatinsightCloudclientResponseAsObject) *NullableUpdateThreatinsightCloudclientResponseAsObject {
	return &NullableUpdateThreatinsightCloudclientResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateThreatinsightCloudclientResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateThreatinsightCloudclientResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
