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

// UpdateDiscoveryGridpropertiesResponse - struct for UpdateDiscoveryGridpropertiesResponse
type UpdateDiscoveryGridpropertiesResponse struct {
	UpdateDiscoveryGridpropertiesResponseAsObject *UpdateDiscoveryGridpropertiesResponseAsObject
	String                                        *string
}

// UpdateDiscoveryGridpropertiesResponseAsObjectAsUpdateDiscoveryGridpropertiesResponse is a convenience function that returns UpdateDiscoveryGridpropertiesResponseAsObject wrapped in UpdateDiscoveryGridpropertiesResponse
func UpdateDiscoveryGridpropertiesResponseAsObjectAsUpdateDiscoveryGridpropertiesResponse(v *UpdateDiscoveryGridpropertiesResponseAsObject) UpdateDiscoveryGridpropertiesResponse {
	return UpdateDiscoveryGridpropertiesResponse{
		UpdateDiscoveryGridpropertiesResponseAsObject: v,
	}
}

// stringAsUpdateDiscoveryGridpropertiesResponse is a convenience function that returns string wrapped in UpdateDiscoveryGridpropertiesResponse
func StringAsUpdateDiscoveryGridpropertiesResponse(v *string) UpdateDiscoveryGridpropertiesResponse {
	return UpdateDiscoveryGridpropertiesResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateDiscoveryGridpropertiesResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateDiscoveryGridpropertiesResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateDiscoveryGridpropertiesResponseAsObject)
	if err == nil {
		jsonUpdateDiscoveryGridpropertiesResponseAsObject, _ := json.Marshal(dst.UpdateDiscoveryGridpropertiesResponseAsObject)
		if string(jsonUpdateDiscoveryGridpropertiesResponseAsObject) == "{}" { // empty struct
			dst.UpdateDiscoveryGridpropertiesResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateDiscoveryGridpropertiesResponseAsObject = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UpdateDiscoveryGridpropertiesResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateDiscoveryGridpropertiesResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateDiscoveryGridpropertiesResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateDiscoveryGridpropertiesResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateDiscoveryGridpropertiesResponseAsObject != nil {
		return json.Marshal(&src.UpdateDiscoveryGridpropertiesResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateDiscoveryGridpropertiesResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateDiscoveryGridpropertiesResponseAsObject != nil {
		return obj.UpdateDiscoveryGridpropertiesResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateDiscoveryGridpropertiesResponse struct {
	value *UpdateDiscoveryGridpropertiesResponse
	isSet bool
}

func (v NullableUpdateDiscoveryGridpropertiesResponse) Get() *UpdateDiscoveryGridpropertiesResponse {
	return v.value
}

func (v *NullableUpdateDiscoveryGridpropertiesResponse) Set(val *UpdateDiscoveryGridpropertiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDiscoveryGridpropertiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDiscoveryGridpropertiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDiscoveryGridpropertiesResponse(val *UpdateDiscoveryGridpropertiesResponse) *NullableUpdateDiscoveryGridpropertiesResponse {
	return &NullableUpdateDiscoveryGridpropertiesResponse{value: val, isSet: true}
}

func (v NullableUpdateDiscoveryGridpropertiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDiscoveryGridpropertiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
