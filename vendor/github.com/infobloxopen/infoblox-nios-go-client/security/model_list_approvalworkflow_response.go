/*
Infoblox SECURITY API

OpenAPI specification for Infoblox NIOS WAPI SECURITY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package security

import (
	"encoding/json"
	"fmt"
)

// ListApprovalworkflowResponse - struct for ListApprovalworkflowResponse
type ListApprovalworkflowResponse struct {
	ListApprovalworkflowResponseObject *ListApprovalworkflowResponseObject
	ArrayOfApprovalworkflow            *[]Approvalworkflow
}

// ListApprovalworkflowResponseObjectAsListApprovalworkflowResponse is a convenience function that returns ListApprovalworkflowResponseObject wrapped in ListApprovalworkflowResponse
func ListApprovalworkflowResponseObjectAsListApprovalworkflowResponse(v *ListApprovalworkflowResponseObject) ListApprovalworkflowResponse {
	return ListApprovalworkflowResponse{
		ListApprovalworkflowResponseObject: v,
	}
}

// []ApprovalworkflowAsListApprovalworkflowResponse is a convenience function that returns []Approvalworkflow wrapped in ListApprovalworkflowResponse
func ArrayOfApprovalworkflowAsListApprovalworkflowResponse(v *[]Approvalworkflow) ListApprovalworkflowResponse {
	return ListApprovalworkflowResponse{
		ArrayOfApprovalworkflow: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListApprovalworkflowResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListApprovalworkflowResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListApprovalworkflowResponseObject)
	if err == nil {
		jsonListApprovalworkflowResponseObject, _ := json.Marshal(dst.ListApprovalworkflowResponseObject)
		if string(jsonListApprovalworkflowResponseObject) == "{}" { // empty struct
			dst.ListApprovalworkflowResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListApprovalworkflowResponseObject = nil
	}

	// try to unmarshal data into ArrayOfApprovalworkflow
	err = newStrictDecoder(data).Decode(&dst.ArrayOfApprovalworkflow)
	if err == nil {
		jsonArrayOfApprovalworkflow, _ := json.Marshal(dst.ArrayOfApprovalworkflow)
		if string(jsonArrayOfApprovalworkflow) == "{}" { // empty struct
			dst.ArrayOfApprovalworkflow = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfApprovalworkflow = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListApprovalworkflowResponseObject = nil
		dst.ArrayOfApprovalworkflow = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListApprovalworkflowResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListApprovalworkflowResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListApprovalworkflowResponse) MarshalJSON() ([]byte, error) {
	if src.ListApprovalworkflowResponseObject != nil {
		return json.Marshal(&src.ListApprovalworkflowResponseObject)
	}

	if src.ArrayOfApprovalworkflow != nil {
		return json.Marshal(&src.ArrayOfApprovalworkflow)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListApprovalworkflowResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListApprovalworkflowResponseObject != nil {
		return obj.ListApprovalworkflowResponseObject
	}

	if obj.ArrayOfApprovalworkflow != nil {
		return obj.ArrayOfApprovalworkflow
	}

	// all schemas are nil
	return nil
}

type NullableListApprovalworkflowResponse struct {
	value *ListApprovalworkflowResponse
	isSet bool
}

func (v NullableListApprovalworkflowResponse) Get() *ListApprovalworkflowResponse {
	return v.value
}

func (v *NullableListApprovalworkflowResponse) Set(val *ListApprovalworkflowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListApprovalworkflowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListApprovalworkflowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApprovalworkflowResponse(val *ListApprovalworkflowResponse) *NullableListApprovalworkflowResponse {
	return &NullableListApprovalworkflowResponse{value: val, isSet: true}
}

func (v NullableListApprovalworkflowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApprovalworkflowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
