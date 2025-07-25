/*
Infoblox MISC API

OpenAPI specification for Infoblox NIOS WAPI MISC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package misc

import (
	"encoding/json"
)

// checks if the GetCsvimporttaskResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCsvimporttaskResponseObjectAsResult{}

// GetCsvimporttaskResponseObjectAsResult The response format to retrieve __Csvimporttask__ objects.
type GetCsvimporttaskResponseObjectAsResult struct {
	Result *Csvimporttask `json:"result,omitempty"`
}

// NewGetCsvimporttaskResponseObjectAsResult instantiates a new GetCsvimporttaskResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCsvimporttaskResponseObjectAsResult() *GetCsvimporttaskResponseObjectAsResult {
	this := GetCsvimporttaskResponseObjectAsResult{}
	return &this
}

// NewGetCsvimporttaskResponseObjectAsResultWithDefaults instantiates a new GetCsvimporttaskResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCsvimporttaskResponseObjectAsResultWithDefaults() *GetCsvimporttaskResponseObjectAsResult {
	this := GetCsvimporttaskResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetCsvimporttaskResponseObjectAsResult) GetResult() Csvimporttask {
	if o == nil || IsNil(o.Result) {
		var ret Csvimporttask
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCsvimporttaskResponseObjectAsResult) GetResultOk() (*Csvimporttask, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetCsvimporttaskResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Csvimporttask and assigns it to the Result field.
func (o *GetCsvimporttaskResponseObjectAsResult) SetResult(v Csvimporttask) {
	o.Result = &v
}

func (o GetCsvimporttaskResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCsvimporttaskResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetCsvimporttaskResponseObjectAsResult struct {
	value *GetCsvimporttaskResponseObjectAsResult
	isSet bool
}

func (v NullableGetCsvimporttaskResponseObjectAsResult) Get() *GetCsvimporttaskResponseObjectAsResult {
	return v.value
}

func (v *NullableGetCsvimporttaskResponseObjectAsResult) Set(val *GetCsvimporttaskResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCsvimporttaskResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCsvimporttaskResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCsvimporttaskResponseObjectAsResult(val *GetCsvimporttaskResponseObjectAsResult) *NullableGetCsvimporttaskResponseObjectAsResult {
	return &NullableGetCsvimporttaskResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetCsvimporttaskResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCsvimporttaskResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
