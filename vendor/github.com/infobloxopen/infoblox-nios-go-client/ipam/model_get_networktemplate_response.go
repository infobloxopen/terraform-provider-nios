/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"fmt"
)

// GetNetworktemplateResponse - struct for GetNetworktemplateResponse
type GetNetworktemplateResponse struct {
	GetNetworktemplateResponseObjectAsResult *GetNetworktemplateResponseObjectAsResult
	Networktemplate                          *Networktemplate
}

// GetNetworktemplateResponseObjectAsResultAsGetNetworktemplateResponse is a convenience function that returns GetNetworktemplateResponseObjectAsResult wrapped in GetNetworktemplateResponse
func GetNetworktemplateResponseObjectAsResultAsGetNetworktemplateResponse(v *GetNetworktemplateResponseObjectAsResult) GetNetworktemplateResponse {
	return GetNetworktemplateResponse{
		GetNetworktemplateResponseObjectAsResult: v,
	}
}

// NetworktemplateAsGetNetworktemplateResponse is a convenience function that returns Networktemplate wrapped in GetNetworktemplateResponse
func NetworktemplateAsGetNetworktemplateResponse(v *Networktemplate) GetNetworktemplateResponse {
	return GetNetworktemplateResponse{
		Networktemplate: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetNetworktemplateResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetNetworktemplateResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetNetworktemplateResponseObjectAsResult)
	if err == nil {
		jsonGetNetworktemplateResponseObjectAsResult, _ := json.Marshal(dst.GetNetworktemplateResponseObjectAsResult)
		if string(jsonGetNetworktemplateResponseObjectAsResult) == "{}" { // empty struct
			dst.GetNetworktemplateResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetNetworktemplateResponseObjectAsResult = nil
	}

	// try to unmarshal data into Networktemplate
	err = newStrictDecoder(data).Decode(&dst.Networktemplate)
	if err == nil {
		jsonNetworktemplate, _ := json.Marshal(dst.Networktemplate)
		if string(jsonNetworktemplate) == "{}" { // empty struct
			dst.Networktemplate = nil
		} else {
			match++
		}
	} else {
		dst.Networktemplate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetNetworktemplateResponseObjectAsResult = nil
		dst.Networktemplate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetNetworktemplateResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetNetworktemplateResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetNetworktemplateResponse) MarshalJSON() ([]byte, error) {
	if src.GetNetworktemplateResponseObjectAsResult != nil {
		return json.Marshal(&src.GetNetworktemplateResponseObjectAsResult)
	}

	if src.Networktemplate != nil {
		return json.Marshal(&src.Networktemplate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetNetworktemplateResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetNetworktemplateResponseObjectAsResult != nil {
		return obj.GetNetworktemplateResponseObjectAsResult
	}

	if obj.Networktemplate != nil {
		return obj.Networktemplate
	}

	// all schemas are nil
	return nil
}

type NullableGetNetworktemplateResponse struct {
	value *GetNetworktemplateResponse
	isSet bool
}

func (v NullableGetNetworktemplateResponse) Get() *GetNetworktemplateResponse {
	return v.value
}

func (v *NullableGetNetworktemplateResponse) Set(val *GetNetworktemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworktemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworktemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworktemplateResponse(val *GetNetworktemplateResponse) *NullableGetNetworktemplateResponse {
	return &NullableGetNetworktemplateResponse{value: val, isSet: true}
}

func (v NullableGetNetworktemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworktemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
