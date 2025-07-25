/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
	"fmt"
)

// ListGridThreatprotectionResponse - struct for ListGridThreatprotectionResponse
type ListGridThreatprotectionResponse struct {
	ListGridThreatprotectionResponseObject *ListGridThreatprotectionResponseObject
	ArrayOfGridThreatprotection            *[]GridThreatprotection
}

// ListGridThreatprotectionResponseObjectAsListGridThreatprotectionResponse is a convenience function that returns ListGridThreatprotectionResponseObject wrapped in ListGridThreatprotectionResponse
func ListGridThreatprotectionResponseObjectAsListGridThreatprotectionResponse(v *ListGridThreatprotectionResponseObject) ListGridThreatprotectionResponse {
	return ListGridThreatprotectionResponse{
		ListGridThreatprotectionResponseObject: v,
	}
}

// []GridThreatprotectionAsListGridThreatprotectionResponse is a convenience function that returns []GridThreatprotection wrapped in ListGridThreatprotectionResponse
func ArrayOfGridThreatprotectionAsListGridThreatprotectionResponse(v *[]GridThreatprotection) ListGridThreatprotectionResponse {
	return ListGridThreatprotectionResponse{
		ArrayOfGridThreatprotection: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListGridThreatprotectionResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListGridThreatprotectionResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListGridThreatprotectionResponseObject)
	if err == nil {
		jsonListGridThreatprotectionResponseObject, _ := json.Marshal(dst.ListGridThreatprotectionResponseObject)
		if string(jsonListGridThreatprotectionResponseObject) == "{}" { // empty struct
			dst.ListGridThreatprotectionResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListGridThreatprotectionResponseObject = nil
	}

	// try to unmarshal data into ArrayOfGridThreatprotection
	err = newStrictDecoder(data).Decode(&dst.ArrayOfGridThreatprotection)
	if err == nil {
		jsonArrayOfGridThreatprotection, _ := json.Marshal(dst.ArrayOfGridThreatprotection)
		if string(jsonArrayOfGridThreatprotection) == "{}" { // empty struct
			dst.ArrayOfGridThreatprotection = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfGridThreatprotection = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListGridThreatprotectionResponseObject = nil
		dst.ArrayOfGridThreatprotection = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListGridThreatprotectionResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListGridThreatprotectionResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListGridThreatprotectionResponse) MarshalJSON() ([]byte, error) {
	if src.ListGridThreatprotectionResponseObject != nil {
		return json.Marshal(&src.ListGridThreatprotectionResponseObject)
	}

	if src.ArrayOfGridThreatprotection != nil {
		return json.Marshal(&src.ArrayOfGridThreatprotection)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListGridThreatprotectionResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListGridThreatprotectionResponseObject != nil {
		return obj.ListGridThreatprotectionResponseObject
	}

	if obj.ArrayOfGridThreatprotection != nil {
		return obj.ArrayOfGridThreatprotection
	}

	// all schemas are nil
	return nil
}

type NullableListGridThreatprotectionResponse struct {
	value *ListGridThreatprotectionResponse
	isSet bool
}

func (v NullableListGridThreatprotectionResponse) Get() *ListGridThreatprotectionResponse {
	return v.value
}

func (v *NullableListGridThreatprotectionResponse) Set(val *ListGridThreatprotectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListGridThreatprotectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListGridThreatprotectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGridThreatprotectionResponse(val *ListGridThreatprotectionResponse) *NullableListGridThreatprotectionResponse {
	return &NullableListGridThreatprotectionResponse{value: val, isSet: true}
}

func (v NullableListGridThreatprotectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGridThreatprotectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
