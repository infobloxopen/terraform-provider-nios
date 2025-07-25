/*
Infoblox DTC API

OpenAPI specification for Infoblox NIOS WAPI DTC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dtc

import (
	"encoding/json"
)

// checks if the GetDtcMonitorTcpResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDtcMonitorTcpResponseObjectAsResult{}

// GetDtcMonitorTcpResponseObjectAsResult The response format to retrieve __DtcMonitorTcp__ objects.
type GetDtcMonitorTcpResponseObjectAsResult struct {
	Result *DtcMonitorTcp `json:"result,omitempty"`
}

// NewGetDtcMonitorTcpResponseObjectAsResult instantiates a new GetDtcMonitorTcpResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDtcMonitorTcpResponseObjectAsResult() *GetDtcMonitorTcpResponseObjectAsResult {
	this := GetDtcMonitorTcpResponseObjectAsResult{}
	return &this
}

// NewGetDtcMonitorTcpResponseObjectAsResultWithDefaults instantiates a new GetDtcMonitorTcpResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDtcMonitorTcpResponseObjectAsResultWithDefaults() *GetDtcMonitorTcpResponseObjectAsResult {
	this := GetDtcMonitorTcpResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetDtcMonitorTcpResponseObjectAsResult) GetResult() DtcMonitorTcp {
	if o == nil || IsNil(o.Result) {
		var ret DtcMonitorTcp
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDtcMonitorTcpResponseObjectAsResult) GetResultOk() (*DtcMonitorTcp, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetDtcMonitorTcpResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given DtcMonitorTcp and assigns it to the Result field.
func (o *GetDtcMonitorTcpResponseObjectAsResult) SetResult(v DtcMonitorTcp) {
	o.Result = &v
}

func (o GetDtcMonitorTcpResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDtcMonitorTcpResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetDtcMonitorTcpResponseObjectAsResult struct {
	value *GetDtcMonitorTcpResponseObjectAsResult
	isSet bool
}

func (v NullableGetDtcMonitorTcpResponseObjectAsResult) Get() *GetDtcMonitorTcpResponseObjectAsResult {
	return v.value
}

func (v *NullableGetDtcMonitorTcpResponseObjectAsResult) Set(val *GetDtcMonitorTcpResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDtcMonitorTcpResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDtcMonitorTcpResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDtcMonitorTcpResponseObjectAsResult(val *GetDtcMonitorTcpResponseObjectAsResult) *NullableGetDtcMonitorTcpResponseObjectAsResult {
	return &NullableGetDtcMonitorTcpResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetDtcMonitorTcpResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDtcMonitorTcpResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
