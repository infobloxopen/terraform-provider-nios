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

// checks if the CreateVlanrangeResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVlanrangeResponseAsObject{}

// CreateVlanrangeResponseAsObject The response format to create __Vlanrange__ in object format.
type CreateVlanrangeResponseAsObject struct {
	Result               *Vlanrange `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateVlanrangeResponseAsObject CreateVlanrangeResponseAsObject

// NewCreateVlanrangeResponseAsObject instantiates a new CreateVlanrangeResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVlanrangeResponseAsObject() *CreateVlanrangeResponseAsObject {
	this := CreateVlanrangeResponseAsObject{}
	return &this
}

// NewCreateVlanrangeResponseAsObjectWithDefaults instantiates a new CreateVlanrangeResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVlanrangeResponseAsObjectWithDefaults() *CreateVlanrangeResponseAsObject {
	this := CreateVlanrangeResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateVlanrangeResponseAsObject) GetResult() Vlanrange {
	if o == nil || IsNil(o.Result) {
		var ret Vlanrange
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVlanrangeResponseAsObject) GetResultOk() (*Vlanrange, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateVlanrangeResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Vlanrange and assigns it to the Result field.
func (o *CreateVlanrangeResponseAsObject) SetResult(v Vlanrange) {
	o.Result = &v
}

func (o CreateVlanrangeResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVlanrangeResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateVlanrangeResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateVlanrangeResponseAsObject := _CreateVlanrangeResponseAsObject{}

	err = json.Unmarshal(data, &varCreateVlanrangeResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateVlanrangeResponseAsObject(varCreateVlanrangeResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateVlanrangeResponseAsObject struct {
	value *CreateVlanrangeResponseAsObject
	isSet bool
}

func (v NullableCreateVlanrangeResponseAsObject) Get() *CreateVlanrangeResponseAsObject {
	return v.value
}

func (v *NullableCreateVlanrangeResponseAsObject) Set(val *CreateVlanrangeResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVlanrangeResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVlanrangeResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVlanrangeResponseAsObject(val *CreateVlanrangeResponseAsObject) *NullableCreateVlanrangeResponseAsObject {
	return &NullableCreateVlanrangeResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateVlanrangeResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVlanrangeResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
