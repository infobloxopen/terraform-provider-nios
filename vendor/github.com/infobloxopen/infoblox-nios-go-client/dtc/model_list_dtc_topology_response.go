/*
Infoblox DTC API

OpenAPI specification for Infoblox NIOS WAPI DTC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dtc

import (
	"encoding/json"
	"fmt"
)

// ListDtcTopologyResponse - struct for ListDtcTopologyResponse
type ListDtcTopologyResponse struct {
	ListDtcTopologyResponseObject *ListDtcTopologyResponseObject
	ArrayOfDtcTopology            *[]DtcTopology
}

// ListDtcTopologyResponseObjectAsListDtcTopologyResponse is a convenience function that returns ListDtcTopologyResponseObject wrapped in ListDtcTopologyResponse
func ListDtcTopologyResponseObjectAsListDtcTopologyResponse(v *ListDtcTopologyResponseObject) ListDtcTopologyResponse {
	return ListDtcTopologyResponse{
		ListDtcTopologyResponseObject: v,
	}
}

// []DtcTopologyAsListDtcTopologyResponse is a convenience function that returns []DtcTopology wrapped in ListDtcTopologyResponse
func ArrayOfDtcTopologyAsListDtcTopologyResponse(v *[]DtcTopology) ListDtcTopologyResponse {
	return ListDtcTopologyResponse{
		ArrayOfDtcTopology: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListDtcTopologyResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListDtcTopologyResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListDtcTopologyResponseObject)
	if err == nil {
		jsonListDtcTopologyResponseObject, _ := json.Marshal(dst.ListDtcTopologyResponseObject)
		if string(jsonListDtcTopologyResponseObject) == "{}" { // empty struct
			dst.ListDtcTopologyResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListDtcTopologyResponseObject = nil
	}

	// try to unmarshal data into ArrayOfDtcTopology
	err = newStrictDecoder(data).Decode(&dst.ArrayOfDtcTopology)
	if err == nil {
		jsonArrayOfDtcTopology, _ := json.Marshal(dst.ArrayOfDtcTopology)
		if string(jsonArrayOfDtcTopology) == "{}" { // empty struct
			dst.ArrayOfDtcTopology = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfDtcTopology = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListDtcTopologyResponseObject = nil
		dst.ArrayOfDtcTopology = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListDtcTopologyResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListDtcTopologyResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListDtcTopologyResponse) MarshalJSON() ([]byte, error) {
	if src.ListDtcTopologyResponseObject != nil {
		return json.Marshal(&src.ListDtcTopologyResponseObject)
	}

	if src.ArrayOfDtcTopology != nil {
		return json.Marshal(&src.ArrayOfDtcTopology)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListDtcTopologyResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListDtcTopologyResponseObject != nil {
		return obj.ListDtcTopologyResponseObject
	}

	if obj.ArrayOfDtcTopology != nil {
		return obj.ArrayOfDtcTopology
	}

	// all schemas are nil
	return nil
}

type NullableListDtcTopologyResponse struct {
	value *ListDtcTopologyResponse
	isSet bool
}

func (v NullableListDtcTopologyResponse) Get() *ListDtcTopologyResponse {
	return v.value
}

func (v *NullableListDtcTopologyResponse) Set(val *ListDtcTopologyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDtcTopologyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDtcTopologyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDtcTopologyResponse(val *ListDtcTopologyResponse) *NullableListDtcTopologyResponse {
	return &NullableListDtcTopologyResponse{value: val, isSet: true}
}

func (v NullableListDtcTopologyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDtcTopologyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
