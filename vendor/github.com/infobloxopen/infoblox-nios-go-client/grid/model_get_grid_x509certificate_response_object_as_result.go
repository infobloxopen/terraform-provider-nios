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

// checks if the GetGridX509certificateResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGridX509certificateResponseObjectAsResult{}

// GetGridX509certificateResponseObjectAsResult The response format to retrieve __GridX509certificate__ objects.
type GetGridX509certificateResponseObjectAsResult struct {
	Result *GridX509certificate `json:"result,omitempty"`
}

// NewGetGridX509certificateResponseObjectAsResult instantiates a new GetGridX509certificateResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGridX509certificateResponseObjectAsResult() *GetGridX509certificateResponseObjectAsResult {
	this := GetGridX509certificateResponseObjectAsResult{}
	return &this
}

// NewGetGridX509certificateResponseObjectAsResultWithDefaults instantiates a new GetGridX509certificateResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGridX509certificateResponseObjectAsResultWithDefaults() *GetGridX509certificateResponseObjectAsResult {
	this := GetGridX509certificateResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetGridX509certificateResponseObjectAsResult) GetResult() GridX509certificate {
	if o == nil || IsNil(o.Result) {
		var ret GridX509certificate
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGridX509certificateResponseObjectAsResult) GetResultOk() (*GridX509certificate, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetGridX509certificateResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given GridX509certificate and assigns it to the Result field.
func (o *GetGridX509certificateResponseObjectAsResult) SetResult(v GridX509certificate) {
	o.Result = &v
}

func (o GetGridX509certificateResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGridX509certificateResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetGridX509certificateResponseObjectAsResult struct {
	value *GetGridX509certificateResponseObjectAsResult
	isSet bool
}

func (v NullableGetGridX509certificateResponseObjectAsResult) Get() *GetGridX509certificateResponseObjectAsResult {
	return v.value
}

func (v *NullableGetGridX509certificateResponseObjectAsResult) Set(val *GetGridX509certificateResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGridX509certificateResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGridX509certificateResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGridX509certificateResponseObjectAsResult(val *GetGridX509certificateResponseObjectAsResult) *NullableGetGridX509certificateResponseObjectAsResult {
	return &NullableGetGridX509certificateResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetGridX509certificateResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGridX509certificateResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
