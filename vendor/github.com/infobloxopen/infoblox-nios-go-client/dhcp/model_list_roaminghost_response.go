/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
	"fmt"
)

// ListRoaminghostResponse - struct for ListRoaminghostResponse
type ListRoaminghostResponse struct {
	ListRoaminghostResponseObject *ListRoaminghostResponseObject
	ArrayOfRoaminghost            *[]Roaminghost
}

// ListRoaminghostResponseObjectAsListRoaminghostResponse is a convenience function that returns ListRoaminghostResponseObject wrapped in ListRoaminghostResponse
func ListRoaminghostResponseObjectAsListRoaminghostResponse(v *ListRoaminghostResponseObject) ListRoaminghostResponse {
	return ListRoaminghostResponse{
		ListRoaminghostResponseObject: v,
	}
}

// []RoaminghostAsListRoaminghostResponse is a convenience function that returns []Roaminghost wrapped in ListRoaminghostResponse
func ArrayOfRoaminghostAsListRoaminghostResponse(v *[]Roaminghost) ListRoaminghostResponse {
	return ListRoaminghostResponse{
		ArrayOfRoaminghost: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListRoaminghostResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListRoaminghostResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListRoaminghostResponseObject)
	if err == nil {
		jsonListRoaminghostResponseObject, _ := json.Marshal(dst.ListRoaminghostResponseObject)
		if string(jsonListRoaminghostResponseObject) == "{}" { // empty struct
			dst.ListRoaminghostResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListRoaminghostResponseObject = nil
	}

	// try to unmarshal data into ArrayOfRoaminghost
	err = newStrictDecoder(data).Decode(&dst.ArrayOfRoaminghost)
	if err == nil {
		jsonArrayOfRoaminghost, _ := json.Marshal(dst.ArrayOfRoaminghost)
		if string(jsonArrayOfRoaminghost) == "{}" { // empty struct
			dst.ArrayOfRoaminghost = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfRoaminghost = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListRoaminghostResponseObject = nil
		dst.ArrayOfRoaminghost = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListRoaminghostResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListRoaminghostResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListRoaminghostResponse) MarshalJSON() ([]byte, error) {
	if src.ListRoaminghostResponseObject != nil {
		return json.Marshal(&src.ListRoaminghostResponseObject)
	}

	if src.ArrayOfRoaminghost != nil {
		return json.Marshal(&src.ArrayOfRoaminghost)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListRoaminghostResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListRoaminghostResponseObject != nil {
		return obj.ListRoaminghostResponseObject
	}

	if obj.ArrayOfRoaminghost != nil {
		return obj.ArrayOfRoaminghost
	}

	// all schemas are nil
	return nil
}

type NullableListRoaminghostResponse struct {
	value *ListRoaminghostResponse
	isSet bool
}

func (v NullableListRoaminghostResponse) Get() *ListRoaminghostResponse {
	return v.value
}

func (v *NullableListRoaminghostResponse) Set(val *ListRoaminghostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRoaminghostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRoaminghostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRoaminghostResponse(val *ListRoaminghostResponse) *NullableListRoaminghostResponse {
	return &NullableListRoaminghostResponse{value: val, isSet: true}
}

func (v NullableListRoaminghostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRoaminghostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
