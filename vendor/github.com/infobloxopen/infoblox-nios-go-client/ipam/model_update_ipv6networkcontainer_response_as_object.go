/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the UpdateIpv6networkcontainerResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIpv6networkcontainerResponseAsObject{}

// UpdateIpv6networkcontainerResponseAsObject The response format to update __Ipv6networkcontainer__ in object format.
type UpdateIpv6networkcontainerResponseAsObject struct {
	Result               *Ipv6networkcontainer `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateIpv6networkcontainerResponseAsObject UpdateIpv6networkcontainerResponseAsObject

// NewUpdateIpv6networkcontainerResponseAsObject instantiates a new UpdateIpv6networkcontainerResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIpv6networkcontainerResponseAsObject() *UpdateIpv6networkcontainerResponseAsObject {
	this := UpdateIpv6networkcontainerResponseAsObject{}
	return &this
}

// NewUpdateIpv6networkcontainerResponseAsObjectWithDefaults instantiates a new UpdateIpv6networkcontainerResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIpv6networkcontainerResponseAsObjectWithDefaults() *UpdateIpv6networkcontainerResponseAsObject {
	this := UpdateIpv6networkcontainerResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateIpv6networkcontainerResponseAsObject) GetResult() Ipv6networkcontainer {
	if o == nil || IsNil(o.Result) {
		var ret Ipv6networkcontainer
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpv6networkcontainerResponseAsObject) GetResultOk() (*Ipv6networkcontainer, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateIpv6networkcontainerResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Ipv6networkcontainer and assigns it to the Result field.
func (o *UpdateIpv6networkcontainerResponseAsObject) SetResult(v Ipv6networkcontainer) {
	o.Result = &v
}

func (o UpdateIpv6networkcontainerResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIpv6networkcontainerResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateIpv6networkcontainerResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateIpv6networkcontainerResponseAsObject := _UpdateIpv6networkcontainerResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateIpv6networkcontainerResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateIpv6networkcontainerResponseAsObject(varUpdateIpv6networkcontainerResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateIpv6networkcontainerResponseAsObject struct {
	value *UpdateIpv6networkcontainerResponseAsObject
	isSet bool
}

func (v NullableUpdateIpv6networkcontainerResponseAsObject) Get() *UpdateIpv6networkcontainerResponseAsObject {
	return v.value
}

func (v *NullableUpdateIpv6networkcontainerResponseAsObject) Set(val *UpdateIpv6networkcontainerResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIpv6networkcontainerResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIpv6networkcontainerResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIpv6networkcontainerResponseAsObject(val *UpdateIpv6networkcontainerResponseAsObject) *NullableUpdateIpv6networkcontainerResponseAsObject {
	return &NullableUpdateIpv6networkcontainerResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateIpv6networkcontainerResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIpv6networkcontainerResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
