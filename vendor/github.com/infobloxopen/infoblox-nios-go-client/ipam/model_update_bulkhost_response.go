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

// UpdateBulkhostResponse - struct for UpdateBulkhostResponse
type UpdateBulkhostResponse struct {
	UpdateBulkhostResponseAsObject *UpdateBulkhostResponseAsObject
	String                         *string
}

// UpdateBulkhostResponseAsObjectAsUpdateBulkhostResponse is a convenience function that returns UpdateBulkhostResponseAsObject wrapped in UpdateBulkhostResponse
func UpdateBulkhostResponseAsObjectAsUpdateBulkhostResponse(v *UpdateBulkhostResponseAsObject) UpdateBulkhostResponse {
	return UpdateBulkhostResponse{
		UpdateBulkhostResponseAsObject: v,
	}
}

// stringAsUpdateBulkhostResponse is a convenience function that returns string wrapped in UpdateBulkhostResponse
func StringAsUpdateBulkhostResponse(v *string) UpdateBulkhostResponse {
	return UpdateBulkhostResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateBulkhostResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateBulkhostResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateBulkhostResponseAsObject)
	if err == nil {
		jsonUpdateBulkhostResponseAsObject, _ := json.Marshal(dst.UpdateBulkhostResponseAsObject)
		if string(jsonUpdateBulkhostResponseAsObject) == "{}" { // empty struct
			dst.UpdateBulkhostResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateBulkhostResponseAsObject = nil
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
		dst.UpdateBulkhostResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateBulkhostResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateBulkhostResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateBulkhostResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateBulkhostResponseAsObject != nil {
		return json.Marshal(&src.UpdateBulkhostResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateBulkhostResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateBulkhostResponseAsObject != nil {
		return obj.UpdateBulkhostResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateBulkhostResponse struct {
	value *UpdateBulkhostResponse
	isSet bool
}

func (v NullableUpdateBulkhostResponse) Get() *UpdateBulkhostResponse {
	return v.value
}

func (v *NullableUpdateBulkhostResponse) Set(val *UpdateBulkhostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBulkhostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBulkhostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBulkhostResponse(val *UpdateBulkhostResponse) *NullableUpdateBulkhostResponse {
	return &NullableUpdateBulkhostResponse{value: val, isSet: true}
}

func (v NullableUpdateBulkhostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBulkhostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
