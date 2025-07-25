/*
Infoblox RPZ API

OpenAPI specification for Infoblox NIOS WAPI RPZ objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rpz

import (
	"encoding/json"
	"fmt"
)

// ListRecordRpzAaaaResponse - struct for ListRecordRpzAaaaResponse
type ListRecordRpzAaaaResponse struct {
	ListRecordRpzAaaaResponseObject *ListRecordRpzAaaaResponseObject
	ArrayOfRecordRpzAaaa            *[]RecordRpzAaaa
}

// ListRecordRpzAaaaResponseObjectAsListRecordRpzAaaaResponse is a convenience function that returns ListRecordRpzAaaaResponseObject wrapped in ListRecordRpzAaaaResponse
func ListRecordRpzAaaaResponseObjectAsListRecordRpzAaaaResponse(v *ListRecordRpzAaaaResponseObject) ListRecordRpzAaaaResponse {
	return ListRecordRpzAaaaResponse{
		ListRecordRpzAaaaResponseObject: v,
	}
}

// []RecordRpzAaaaAsListRecordRpzAaaaResponse is a convenience function that returns []RecordRpzAaaa wrapped in ListRecordRpzAaaaResponse
func ArrayOfRecordRpzAaaaAsListRecordRpzAaaaResponse(v *[]RecordRpzAaaa) ListRecordRpzAaaaResponse {
	return ListRecordRpzAaaaResponse{
		ArrayOfRecordRpzAaaa: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListRecordRpzAaaaResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListRecordRpzAaaaResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListRecordRpzAaaaResponseObject)
	if err == nil {
		jsonListRecordRpzAaaaResponseObject, _ := json.Marshal(dst.ListRecordRpzAaaaResponseObject)
		if string(jsonListRecordRpzAaaaResponseObject) == "{}" { // empty struct
			dst.ListRecordRpzAaaaResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListRecordRpzAaaaResponseObject = nil
	}

	// try to unmarshal data into ArrayOfRecordRpzAaaa
	err = newStrictDecoder(data).Decode(&dst.ArrayOfRecordRpzAaaa)
	if err == nil {
		jsonArrayOfRecordRpzAaaa, _ := json.Marshal(dst.ArrayOfRecordRpzAaaa)
		if string(jsonArrayOfRecordRpzAaaa) == "{}" { // empty struct
			dst.ArrayOfRecordRpzAaaa = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfRecordRpzAaaa = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListRecordRpzAaaaResponseObject = nil
		dst.ArrayOfRecordRpzAaaa = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListRecordRpzAaaaResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListRecordRpzAaaaResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListRecordRpzAaaaResponse) MarshalJSON() ([]byte, error) {
	if src.ListRecordRpzAaaaResponseObject != nil {
		return json.Marshal(&src.ListRecordRpzAaaaResponseObject)
	}

	if src.ArrayOfRecordRpzAaaa != nil {
		return json.Marshal(&src.ArrayOfRecordRpzAaaa)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListRecordRpzAaaaResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListRecordRpzAaaaResponseObject != nil {
		return obj.ListRecordRpzAaaaResponseObject
	}

	if obj.ArrayOfRecordRpzAaaa != nil {
		return obj.ArrayOfRecordRpzAaaa
	}

	// all schemas are nil
	return nil
}

type NullableListRecordRpzAaaaResponse struct {
	value *ListRecordRpzAaaaResponse
	isSet bool
}

func (v NullableListRecordRpzAaaaResponse) Get() *ListRecordRpzAaaaResponse {
	return v.value
}

func (v *NullableListRecordRpzAaaaResponse) Set(val *ListRecordRpzAaaaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRecordRpzAaaaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRecordRpzAaaaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRecordRpzAaaaResponse(val *ListRecordRpzAaaaResponse) *NullableListRecordRpzAaaaResponse {
	return &NullableListRecordRpzAaaaResponse{value: val, isSet: true}
}

func (v NullableListRecordRpzAaaaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRecordRpzAaaaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
