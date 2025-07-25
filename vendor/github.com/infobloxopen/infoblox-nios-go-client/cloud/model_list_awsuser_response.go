/*
Infoblox CLOUD API

OpenAPI specification for Infoblox NIOS WAPI CLOUD objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloud

import (
	"encoding/json"
	"fmt"
)

// ListAwsuserResponse - struct for ListAwsuserResponse
type ListAwsuserResponse struct {
	ListAwsuserResponseObject *ListAwsuserResponseObject
	ArrayOfAwsuser            *[]Awsuser
}

// ListAwsuserResponseObjectAsListAwsuserResponse is a convenience function that returns ListAwsuserResponseObject wrapped in ListAwsuserResponse
func ListAwsuserResponseObjectAsListAwsuserResponse(v *ListAwsuserResponseObject) ListAwsuserResponse {
	return ListAwsuserResponse{
		ListAwsuserResponseObject: v,
	}
}

// []AwsuserAsListAwsuserResponse is a convenience function that returns []Awsuser wrapped in ListAwsuserResponse
func ArrayOfAwsuserAsListAwsuserResponse(v *[]Awsuser) ListAwsuserResponse {
	return ListAwsuserResponse{
		ArrayOfAwsuser: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListAwsuserResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListAwsuserResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListAwsuserResponseObject)
	if err == nil {
		jsonListAwsuserResponseObject, _ := json.Marshal(dst.ListAwsuserResponseObject)
		if string(jsonListAwsuserResponseObject) == "{}" { // empty struct
			dst.ListAwsuserResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListAwsuserResponseObject = nil
	}

	// try to unmarshal data into ArrayOfAwsuser
	err = newStrictDecoder(data).Decode(&dst.ArrayOfAwsuser)
	if err == nil {
		jsonArrayOfAwsuser, _ := json.Marshal(dst.ArrayOfAwsuser)
		if string(jsonArrayOfAwsuser) == "{}" { // empty struct
			dst.ArrayOfAwsuser = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfAwsuser = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListAwsuserResponseObject = nil
		dst.ArrayOfAwsuser = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListAwsuserResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListAwsuserResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListAwsuserResponse) MarshalJSON() ([]byte, error) {
	if src.ListAwsuserResponseObject != nil {
		return json.Marshal(&src.ListAwsuserResponseObject)
	}

	if src.ArrayOfAwsuser != nil {
		return json.Marshal(&src.ArrayOfAwsuser)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListAwsuserResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListAwsuserResponseObject != nil {
		return obj.ListAwsuserResponseObject
	}

	if obj.ArrayOfAwsuser != nil {
		return obj.ArrayOfAwsuser
	}

	// all schemas are nil
	return nil
}

type NullableListAwsuserResponse struct {
	value *ListAwsuserResponse
	isSet bool
}

func (v NullableListAwsuserResponse) Get() *ListAwsuserResponse {
	return v.value
}

func (v *NullableListAwsuserResponse) Set(val *ListAwsuserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAwsuserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAwsuserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAwsuserResponse(val *ListAwsuserResponse) *NullableListAwsuserResponse {
	return &NullableListAwsuserResponse{value: val, isSet: true}
}

func (v NullableListAwsuserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAwsuserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
