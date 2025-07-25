/*
Infoblox DISCOVERY API

OpenAPI specification for Infoblox NIOS WAPI DISCOVERY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discovery

import (
	"encoding/json"
	"fmt"
)

// GetDiscoverySdnnetworkResponse - struct for GetDiscoverySdnnetworkResponse
type GetDiscoverySdnnetworkResponse struct {
	DiscoverySdnnetwork                          *DiscoverySdnnetwork
	GetDiscoverySdnnetworkResponseObjectAsResult *GetDiscoverySdnnetworkResponseObjectAsResult
}

// DiscoverySdnnetworkAsGetDiscoverySdnnetworkResponse is a convenience function that returns DiscoverySdnnetwork wrapped in GetDiscoverySdnnetworkResponse
func DiscoverySdnnetworkAsGetDiscoverySdnnetworkResponse(v *DiscoverySdnnetwork) GetDiscoverySdnnetworkResponse {
	return GetDiscoverySdnnetworkResponse{
		DiscoverySdnnetwork: v,
	}
}

// GetDiscoverySdnnetworkResponseObjectAsResultAsGetDiscoverySdnnetworkResponse is a convenience function that returns GetDiscoverySdnnetworkResponseObjectAsResult wrapped in GetDiscoverySdnnetworkResponse
func GetDiscoverySdnnetworkResponseObjectAsResultAsGetDiscoverySdnnetworkResponse(v *GetDiscoverySdnnetworkResponseObjectAsResult) GetDiscoverySdnnetworkResponse {
	return GetDiscoverySdnnetworkResponse{
		GetDiscoverySdnnetworkResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetDiscoverySdnnetworkResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DiscoverySdnnetwork
	err = newStrictDecoder(data).Decode(&dst.DiscoverySdnnetwork)
	if err == nil {
		jsonDiscoverySdnnetwork, _ := json.Marshal(dst.DiscoverySdnnetwork)
		if string(jsonDiscoverySdnnetwork) == "{}" { // empty struct
			dst.DiscoverySdnnetwork = nil
		} else {
			match++
		}
	} else {
		dst.DiscoverySdnnetwork = nil
	}

	// try to unmarshal data into GetDiscoverySdnnetworkResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetDiscoverySdnnetworkResponseObjectAsResult)
	if err == nil {
		jsonGetDiscoverySdnnetworkResponseObjectAsResult, _ := json.Marshal(dst.GetDiscoverySdnnetworkResponseObjectAsResult)
		if string(jsonGetDiscoverySdnnetworkResponseObjectAsResult) == "{}" { // empty struct
			dst.GetDiscoverySdnnetworkResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetDiscoverySdnnetworkResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DiscoverySdnnetwork = nil
		dst.GetDiscoverySdnnetworkResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetDiscoverySdnnetworkResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetDiscoverySdnnetworkResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetDiscoverySdnnetworkResponse) MarshalJSON() ([]byte, error) {
	if src.DiscoverySdnnetwork != nil {
		return json.Marshal(&src.DiscoverySdnnetwork)
	}

	if src.GetDiscoverySdnnetworkResponseObjectAsResult != nil {
		return json.Marshal(&src.GetDiscoverySdnnetworkResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetDiscoverySdnnetworkResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DiscoverySdnnetwork != nil {
		return obj.DiscoverySdnnetwork
	}

	if obj.GetDiscoverySdnnetworkResponseObjectAsResult != nil {
		return obj.GetDiscoverySdnnetworkResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetDiscoverySdnnetworkResponse struct {
	value *GetDiscoverySdnnetworkResponse
	isSet bool
}

func (v NullableGetDiscoverySdnnetworkResponse) Get() *GetDiscoverySdnnetworkResponse {
	return v.value
}

func (v *NullableGetDiscoverySdnnetworkResponse) Set(val *GetDiscoverySdnnetworkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDiscoverySdnnetworkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDiscoverySdnnetworkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDiscoverySdnnetworkResponse(val *GetDiscoverySdnnetworkResponse) *NullableGetDiscoverySdnnetworkResponse {
	return &NullableGetDiscoverySdnnetworkResponse{value: val, isSet: true}
}

func (v NullableGetDiscoverySdnnetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDiscoverySdnnetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
