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

// CreateDhcpoptionspaceResponse - struct for CreateDhcpoptionspaceResponse
type CreateDhcpoptionspaceResponse struct {
	CreateDhcpoptionspaceResponseAsObject *CreateDhcpoptionspaceResponseAsObject
	String                                *string
}

// CreateDhcpoptionspaceResponseAsObjectAsCreateDhcpoptionspaceResponse is a convenience function that returns CreateDhcpoptionspaceResponseAsObject wrapped in CreateDhcpoptionspaceResponse
func CreateDhcpoptionspaceResponseAsObjectAsCreateDhcpoptionspaceResponse(v *CreateDhcpoptionspaceResponseAsObject) CreateDhcpoptionspaceResponse {
	return CreateDhcpoptionspaceResponse{
		CreateDhcpoptionspaceResponseAsObject: v,
	}
}

// stringAsCreateDhcpoptionspaceResponse is a convenience function that returns string wrapped in CreateDhcpoptionspaceResponse
func StringAsCreateDhcpoptionspaceResponse(v *string) CreateDhcpoptionspaceResponse {
	return CreateDhcpoptionspaceResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateDhcpoptionspaceResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateDhcpoptionspaceResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateDhcpoptionspaceResponseAsObject)
	if err == nil {
		jsonCreateDhcpoptionspaceResponseAsObject, _ := json.Marshal(dst.CreateDhcpoptionspaceResponseAsObject)
		if string(jsonCreateDhcpoptionspaceResponseAsObject) == "{}" { // empty struct
			dst.CreateDhcpoptionspaceResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateDhcpoptionspaceResponseAsObject = nil
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
		dst.CreateDhcpoptionspaceResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateDhcpoptionspaceResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateDhcpoptionspaceResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateDhcpoptionspaceResponse) MarshalJSON() ([]byte, error) {
	if src.CreateDhcpoptionspaceResponseAsObject != nil {
		return json.Marshal(&src.CreateDhcpoptionspaceResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateDhcpoptionspaceResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateDhcpoptionspaceResponseAsObject != nil {
		return obj.CreateDhcpoptionspaceResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateDhcpoptionspaceResponse struct {
	value *CreateDhcpoptionspaceResponse
	isSet bool
}

func (v NullableCreateDhcpoptionspaceResponse) Get() *CreateDhcpoptionspaceResponse {
	return v.value
}

func (v *NullableCreateDhcpoptionspaceResponse) Set(val *CreateDhcpoptionspaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDhcpoptionspaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDhcpoptionspaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDhcpoptionspaceResponse(val *CreateDhcpoptionspaceResponse) *NullableCreateDhcpoptionspaceResponse {
	return &NullableCreateDhcpoptionspaceResponse{value: val, isSet: true}
}

func (v NullableCreateDhcpoptionspaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDhcpoptionspaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
