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

// CreateFilterrelayagentResponse - struct for CreateFilterrelayagentResponse
type CreateFilterrelayagentResponse struct {
	CreateFilterrelayagentResponseAsObject *CreateFilterrelayagentResponseAsObject
	String                                 *string
}

// CreateFilterrelayagentResponseAsObjectAsCreateFilterrelayagentResponse is a convenience function that returns CreateFilterrelayagentResponseAsObject wrapped in CreateFilterrelayagentResponse
func CreateFilterrelayagentResponseAsObjectAsCreateFilterrelayagentResponse(v *CreateFilterrelayagentResponseAsObject) CreateFilterrelayagentResponse {
	return CreateFilterrelayagentResponse{
		CreateFilterrelayagentResponseAsObject: v,
	}
}

// stringAsCreateFilterrelayagentResponse is a convenience function that returns string wrapped in CreateFilterrelayagentResponse
func StringAsCreateFilterrelayagentResponse(v *string) CreateFilterrelayagentResponse {
	return CreateFilterrelayagentResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateFilterrelayagentResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateFilterrelayagentResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateFilterrelayagentResponseAsObject)
	if err == nil {
		jsonCreateFilterrelayagentResponseAsObject, _ := json.Marshal(dst.CreateFilterrelayagentResponseAsObject)
		if string(jsonCreateFilterrelayagentResponseAsObject) == "{}" { // empty struct
			dst.CreateFilterrelayagentResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateFilterrelayagentResponseAsObject = nil
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
		dst.CreateFilterrelayagentResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateFilterrelayagentResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateFilterrelayagentResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateFilterrelayagentResponse) MarshalJSON() ([]byte, error) {
	if src.CreateFilterrelayagentResponseAsObject != nil {
		return json.Marshal(&src.CreateFilterrelayagentResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateFilterrelayagentResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateFilterrelayagentResponseAsObject != nil {
		return obj.CreateFilterrelayagentResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateFilterrelayagentResponse struct {
	value *CreateFilterrelayagentResponse
	isSet bool
}

func (v NullableCreateFilterrelayagentResponse) Get() *CreateFilterrelayagentResponse {
	return v.value
}

func (v *NullableCreateFilterrelayagentResponse) Set(val *CreateFilterrelayagentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFilterrelayagentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFilterrelayagentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFilterrelayagentResponse(val *CreateFilterrelayagentResponse) *NullableCreateFilterrelayagentResponse {
	return &NullableCreateFilterrelayagentResponse{value: val, isSet: true}
}

func (v NullableCreateFilterrelayagentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFilterrelayagentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
