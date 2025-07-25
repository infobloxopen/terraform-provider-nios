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

// ListFiltermacResponse - struct for ListFiltermacResponse
type ListFiltermacResponse struct {
	ListFiltermacResponseObject *ListFiltermacResponseObject
	ArrayOfFiltermac            *[]Filtermac
}

// ListFiltermacResponseObjectAsListFiltermacResponse is a convenience function that returns ListFiltermacResponseObject wrapped in ListFiltermacResponse
func ListFiltermacResponseObjectAsListFiltermacResponse(v *ListFiltermacResponseObject) ListFiltermacResponse {
	return ListFiltermacResponse{
		ListFiltermacResponseObject: v,
	}
}

// []FiltermacAsListFiltermacResponse is a convenience function that returns []Filtermac wrapped in ListFiltermacResponse
func ArrayOfFiltermacAsListFiltermacResponse(v *[]Filtermac) ListFiltermacResponse {
	return ListFiltermacResponse{
		ArrayOfFiltermac: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListFiltermacResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListFiltermacResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListFiltermacResponseObject)
	if err == nil {
		jsonListFiltermacResponseObject, _ := json.Marshal(dst.ListFiltermacResponseObject)
		if string(jsonListFiltermacResponseObject) == "{}" { // empty struct
			dst.ListFiltermacResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListFiltermacResponseObject = nil
	}

	// try to unmarshal data into ArrayOfFiltermac
	err = newStrictDecoder(data).Decode(&dst.ArrayOfFiltermac)
	if err == nil {
		jsonArrayOfFiltermac, _ := json.Marshal(dst.ArrayOfFiltermac)
		if string(jsonArrayOfFiltermac) == "{}" { // empty struct
			dst.ArrayOfFiltermac = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfFiltermac = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListFiltermacResponseObject = nil
		dst.ArrayOfFiltermac = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListFiltermacResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListFiltermacResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListFiltermacResponse) MarshalJSON() ([]byte, error) {
	if src.ListFiltermacResponseObject != nil {
		return json.Marshal(&src.ListFiltermacResponseObject)
	}

	if src.ArrayOfFiltermac != nil {
		return json.Marshal(&src.ArrayOfFiltermac)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListFiltermacResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListFiltermacResponseObject != nil {
		return obj.ListFiltermacResponseObject
	}

	if obj.ArrayOfFiltermac != nil {
		return obj.ArrayOfFiltermac
	}

	// all schemas are nil
	return nil
}

type NullableListFiltermacResponse struct {
	value *ListFiltermacResponse
	isSet bool
}

func (v NullableListFiltermacResponse) Get() *ListFiltermacResponse {
	return v.value
}

func (v *NullableListFiltermacResponse) Set(val *ListFiltermacResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFiltermacResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFiltermacResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFiltermacResponse(val *ListFiltermacResponse) *NullableListFiltermacResponse {
	return &NullableListFiltermacResponse{value: val, isSet: true}
}

func (v NullableListFiltermacResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFiltermacResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
