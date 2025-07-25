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

// checks if the CreateIpv6networkcontainerResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIpv6networkcontainerResponseAsObject{}

// CreateIpv6networkcontainerResponseAsObject The response format to create __Ipv6networkcontainer__ in object format.
type CreateIpv6networkcontainerResponseAsObject struct {
	Result               *Ipv6networkcontainer `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateIpv6networkcontainerResponseAsObject CreateIpv6networkcontainerResponseAsObject

// NewCreateIpv6networkcontainerResponseAsObject instantiates a new CreateIpv6networkcontainerResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIpv6networkcontainerResponseAsObject() *CreateIpv6networkcontainerResponseAsObject {
	this := CreateIpv6networkcontainerResponseAsObject{}
	return &this
}

// NewCreateIpv6networkcontainerResponseAsObjectWithDefaults instantiates a new CreateIpv6networkcontainerResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIpv6networkcontainerResponseAsObjectWithDefaults() *CreateIpv6networkcontainerResponseAsObject {
	this := CreateIpv6networkcontainerResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateIpv6networkcontainerResponseAsObject) GetResult() Ipv6networkcontainer {
	if o == nil || IsNil(o.Result) {
		var ret Ipv6networkcontainer
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIpv6networkcontainerResponseAsObject) GetResultOk() (*Ipv6networkcontainer, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateIpv6networkcontainerResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Ipv6networkcontainer and assigns it to the Result field.
func (o *CreateIpv6networkcontainerResponseAsObject) SetResult(v Ipv6networkcontainer) {
	o.Result = &v
}

func (o CreateIpv6networkcontainerResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIpv6networkcontainerResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateIpv6networkcontainerResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateIpv6networkcontainerResponseAsObject := _CreateIpv6networkcontainerResponseAsObject{}

	err = json.Unmarshal(data, &varCreateIpv6networkcontainerResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateIpv6networkcontainerResponseAsObject(varCreateIpv6networkcontainerResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateIpv6networkcontainerResponseAsObject struct {
	value *CreateIpv6networkcontainerResponseAsObject
	isSet bool
}

func (v NullableCreateIpv6networkcontainerResponseAsObject) Get() *CreateIpv6networkcontainerResponseAsObject {
	return v.value
}

func (v *NullableCreateIpv6networkcontainerResponseAsObject) Set(val *CreateIpv6networkcontainerResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIpv6networkcontainerResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIpv6networkcontainerResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIpv6networkcontainerResponseAsObject(val *CreateIpv6networkcontainerResponseAsObject) *NullableCreateIpv6networkcontainerResponseAsObject {
	return &NullableCreateIpv6networkcontainerResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateIpv6networkcontainerResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIpv6networkcontainerResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
