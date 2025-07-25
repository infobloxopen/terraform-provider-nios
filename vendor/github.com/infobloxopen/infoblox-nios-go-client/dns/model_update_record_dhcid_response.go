/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
)

// UpdateRecordDhcidResponse - struct for UpdateRecordDhcidResponse
type UpdateRecordDhcidResponse struct {
	UpdateRecordDhcidResponseAsObject *UpdateRecordDhcidResponseAsObject
	String                            *string
}

// UpdateRecordDhcidResponseAsObjectAsUpdateRecordDhcidResponse is a convenience function that returns UpdateRecordDhcidResponseAsObject wrapped in UpdateRecordDhcidResponse
func UpdateRecordDhcidResponseAsObjectAsUpdateRecordDhcidResponse(v *UpdateRecordDhcidResponseAsObject) UpdateRecordDhcidResponse {
	return UpdateRecordDhcidResponse{
		UpdateRecordDhcidResponseAsObject: v,
	}
}

// stringAsUpdateRecordDhcidResponse is a convenience function that returns string wrapped in UpdateRecordDhcidResponse
func StringAsUpdateRecordDhcidResponse(v *string) UpdateRecordDhcidResponse {
	return UpdateRecordDhcidResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateRecordDhcidResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateRecordDhcidResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateRecordDhcidResponseAsObject)
	if err == nil {
		jsonUpdateRecordDhcidResponseAsObject, _ := json.Marshal(dst.UpdateRecordDhcidResponseAsObject)
		if string(jsonUpdateRecordDhcidResponseAsObject) == "{}" { // empty struct
			dst.UpdateRecordDhcidResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateRecordDhcidResponseAsObject = nil
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
		dst.UpdateRecordDhcidResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateRecordDhcidResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateRecordDhcidResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateRecordDhcidResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateRecordDhcidResponseAsObject != nil {
		return json.Marshal(&src.UpdateRecordDhcidResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateRecordDhcidResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateRecordDhcidResponseAsObject != nil {
		return obj.UpdateRecordDhcidResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateRecordDhcidResponse struct {
	value *UpdateRecordDhcidResponse
	isSet bool
}

func (v NullableUpdateRecordDhcidResponse) Get() *UpdateRecordDhcidResponse {
	return v.value
}

func (v *NullableUpdateRecordDhcidResponse) Set(val *UpdateRecordDhcidResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRecordDhcidResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRecordDhcidResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRecordDhcidResponse(val *UpdateRecordDhcidResponse) *NullableUpdateRecordDhcidResponse {
	return &NullableUpdateRecordDhcidResponse{value: val, isSet: true}
}

func (v NullableUpdateRecordDhcidResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRecordDhcidResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
