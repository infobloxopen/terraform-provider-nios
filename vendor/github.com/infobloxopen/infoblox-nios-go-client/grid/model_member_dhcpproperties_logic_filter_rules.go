/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
)

// checks if the MemberDhcppropertiesLogicFilterRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberDhcppropertiesLogicFilterRules{}

// MemberDhcppropertiesLogicFilterRules struct for MemberDhcppropertiesLogicFilterRules
type MemberDhcppropertiesLogicFilterRules struct {
	// The filter name.
	Filter *string `json:"filter,omitempty"`
	// The filter type. Valid values are: * MAC * NAC * Option
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemberDhcppropertiesLogicFilterRules MemberDhcppropertiesLogicFilterRules

// NewMemberDhcppropertiesLogicFilterRules instantiates a new MemberDhcppropertiesLogicFilterRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberDhcppropertiesLogicFilterRules() *MemberDhcppropertiesLogicFilterRules {
	this := MemberDhcppropertiesLogicFilterRules{}
	return &this
}

// NewMemberDhcppropertiesLogicFilterRulesWithDefaults instantiates a new MemberDhcppropertiesLogicFilterRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberDhcppropertiesLogicFilterRulesWithDefaults() *MemberDhcppropertiesLogicFilterRules {
	this := MemberDhcppropertiesLogicFilterRules{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MemberDhcppropertiesLogicFilterRules) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDhcppropertiesLogicFilterRules) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MemberDhcppropertiesLogicFilterRules) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *MemberDhcppropertiesLogicFilterRules) SetFilter(v string) {
	o.Filter = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MemberDhcppropertiesLogicFilterRules) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDhcppropertiesLogicFilterRules) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MemberDhcppropertiesLogicFilterRules) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MemberDhcppropertiesLogicFilterRules) SetType(v string) {
	o.Type = &v
}

func (o MemberDhcppropertiesLogicFilterRules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberDhcppropertiesLogicFilterRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemberDhcppropertiesLogicFilterRules) UnmarshalJSON(data []byte) (err error) {
	varMemberDhcppropertiesLogicFilterRules := _MemberDhcppropertiesLogicFilterRules{}

	err = json.Unmarshal(data, &varMemberDhcppropertiesLogicFilterRules)

	if err != nil {
		return err
	}

	*o = MemberDhcppropertiesLogicFilterRules(varMemberDhcppropertiesLogicFilterRules)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filter")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemberDhcppropertiesLogicFilterRules struct {
	value *MemberDhcppropertiesLogicFilterRules
	isSet bool
}

func (v NullableMemberDhcppropertiesLogicFilterRules) Get() *MemberDhcppropertiesLogicFilterRules {
	return v.value
}

func (v *NullableMemberDhcppropertiesLogicFilterRules) Set(val *MemberDhcppropertiesLogicFilterRules) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberDhcppropertiesLogicFilterRules) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberDhcppropertiesLogicFilterRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberDhcppropertiesLogicFilterRules(val *MemberDhcppropertiesLogicFilterRules) *NullableMemberDhcppropertiesLogicFilterRules {
	return &NullableMemberDhcppropertiesLogicFilterRules{value: val, isSet: true}
}

func (v NullableMemberDhcppropertiesLogicFilterRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberDhcppropertiesLogicFilterRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
