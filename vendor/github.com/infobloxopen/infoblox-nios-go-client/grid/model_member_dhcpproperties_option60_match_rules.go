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

// checks if the MemberDhcppropertiesOption60MatchRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberDhcppropertiesOption60MatchRules{}

// MemberDhcppropertiesOption60MatchRules struct for MemberDhcppropertiesOption60MatchRules
type MemberDhcppropertiesOption60MatchRules struct {
	// The match value for this DHCP Option 60 match rule.
	MatchValue *string `json:"match_value,omitempty"`
	// The option space for this DHCP Option 60 match rule.
	OptionSpace *string `json:"option_space,omitempty"`
	// Determines if the match value is a substring.
	IsSubstring *bool `json:"is_substring,omitempty"`
	// The offset of match value for this DHCP Option 60 match rule.
	SubstringOffset *int64 `json:"substring_offset,omitempty"`
	// The length of match value for this DHCP Option 60 match rule.
	SubstringLength      *int64 `json:"substring_length,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemberDhcppropertiesOption60MatchRules MemberDhcppropertiesOption60MatchRules

// NewMemberDhcppropertiesOption60MatchRules instantiates a new MemberDhcppropertiesOption60MatchRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberDhcppropertiesOption60MatchRules() *MemberDhcppropertiesOption60MatchRules {
	this := MemberDhcppropertiesOption60MatchRules{}
	return &this
}

// NewMemberDhcppropertiesOption60MatchRulesWithDefaults instantiates a new MemberDhcppropertiesOption60MatchRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberDhcppropertiesOption60MatchRulesWithDefaults() *MemberDhcppropertiesOption60MatchRules {
	this := MemberDhcppropertiesOption60MatchRules{}
	return &this
}

// GetMatchValue returns the MatchValue field value if set, zero value otherwise.
func (o *MemberDhcppropertiesOption60MatchRules) GetMatchValue() string {
	if o == nil || IsNil(o.MatchValue) {
		var ret string
		return ret
	}
	return *o.MatchValue
}

// GetMatchValueOk returns a tuple with the MatchValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDhcppropertiesOption60MatchRules) GetMatchValueOk() (*string, bool) {
	if o == nil || IsNil(o.MatchValue) {
		return nil, false
	}
	return o.MatchValue, true
}

// HasMatchValue returns a boolean if a field has been set.
func (o *MemberDhcppropertiesOption60MatchRules) HasMatchValue() bool {
	if o != nil && !IsNil(o.MatchValue) {
		return true
	}

	return false
}

// SetMatchValue gets a reference to the given string and assigns it to the MatchValue field.
func (o *MemberDhcppropertiesOption60MatchRules) SetMatchValue(v string) {
	o.MatchValue = &v
}

// GetOptionSpace returns the OptionSpace field value if set, zero value otherwise.
func (o *MemberDhcppropertiesOption60MatchRules) GetOptionSpace() string {
	if o == nil || IsNil(o.OptionSpace) {
		var ret string
		return ret
	}
	return *o.OptionSpace
}

// GetOptionSpaceOk returns a tuple with the OptionSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDhcppropertiesOption60MatchRules) GetOptionSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.OptionSpace) {
		return nil, false
	}
	return o.OptionSpace, true
}

// HasOptionSpace returns a boolean if a field has been set.
func (o *MemberDhcppropertiesOption60MatchRules) HasOptionSpace() bool {
	if o != nil && !IsNil(o.OptionSpace) {
		return true
	}

	return false
}

// SetOptionSpace gets a reference to the given string and assigns it to the OptionSpace field.
func (o *MemberDhcppropertiesOption60MatchRules) SetOptionSpace(v string) {
	o.OptionSpace = &v
}

// GetIsSubstring returns the IsSubstring field value if set, zero value otherwise.
func (o *MemberDhcppropertiesOption60MatchRules) GetIsSubstring() bool {
	if o == nil || IsNil(o.IsSubstring) {
		var ret bool
		return ret
	}
	return *o.IsSubstring
}

// GetIsSubstringOk returns a tuple with the IsSubstring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDhcppropertiesOption60MatchRules) GetIsSubstringOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSubstring) {
		return nil, false
	}
	return o.IsSubstring, true
}

// HasIsSubstring returns a boolean if a field has been set.
func (o *MemberDhcppropertiesOption60MatchRules) HasIsSubstring() bool {
	if o != nil && !IsNil(o.IsSubstring) {
		return true
	}

	return false
}

// SetIsSubstring gets a reference to the given bool and assigns it to the IsSubstring field.
func (o *MemberDhcppropertiesOption60MatchRules) SetIsSubstring(v bool) {
	o.IsSubstring = &v
}

// GetSubstringOffset returns the SubstringOffset field value if set, zero value otherwise.
func (o *MemberDhcppropertiesOption60MatchRules) GetSubstringOffset() int64 {
	if o == nil || IsNil(o.SubstringOffset) {
		var ret int64
		return ret
	}
	return *o.SubstringOffset
}

// GetSubstringOffsetOk returns a tuple with the SubstringOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDhcppropertiesOption60MatchRules) GetSubstringOffsetOk() (*int64, bool) {
	if o == nil || IsNil(o.SubstringOffset) {
		return nil, false
	}
	return o.SubstringOffset, true
}

// HasSubstringOffset returns a boolean if a field has been set.
func (o *MemberDhcppropertiesOption60MatchRules) HasSubstringOffset() bool {
	if o != nil && !IsNil(o.SubstringOffset) {
		return true
	}

	return false
}

// SetSubstringOffset gets a reference to the given int64 and assigns it to the SubstringOffset field.
func (o *MemberDhcppropertiesOption60MatchRules) SetSubstringOffset(v int64) {
	o.SubstringOffset = &v
}

// GetSubstringLength returns the SubstringLength field value if set, zero value otherwise.
func (o *MemberDhcppropertiesOption60MatchRules) GetSubstringLength() int64 {
	if o == nil || IsNil(o.SubstringLength) {
		var ret int64
		return ret
	}
	return *o.SubstringLength
}

// GetSubstringLengthOk returns a tuple with the SubstringLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDhcppropertiesOption60MatchRules) GetSubstringLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.SubstringLength) {
		return nil, false
	}
	return o.SubstringLength, true
}

// HasSubstringLength returns a boolean if a field has been set.
func (o *MemberDhcppropertiesOption60MatchRules) HasSubstringLength() bool {
	if o != nil && !IsNil(o.SubstringLength) {
		return true
	}

	return false
}

// SetSubstringLength gets a reference to the given int64 and assigns it to the SubstringLength field.
func (o *MemberDhcppropertiesOption60MatchRules) SetSubstringLength(v int64) {
	o.SubstringLength = &v
}

func (o MemberDhcppropertiesOption60MatchRules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberDhcppropertiesOption60MatchRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchValue) {
		toSerialize["match_value"] = o.MatchValue
	}
	if !IsNil(o.OptionSpace) {
		toSerialize["option_space"] = o.OptionSpace
	}
	if !IsNil(o.IsSubstring) {
		toSerialize["is_substring"] = o.IsSubstring
	}
	if !IsNil(o.SubstringOffset) {
		toSerialize["substring_offset"] = o.SubstringOffset
	}
	if !IsNil(o.SubstringLength) {
		toSerialize["substring_length"] = o.SubstringLength
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemberDhcppropertiesOption60MatchRules) UnmarshalJSON(data []byte) (err error) {
	varMemberDhcppropertiesOption60MatchRules := _MemberDhcppropertiesOption60MatchRules{}

	err = json.Unmarshal(data, &varMemberDhcppropertiesOption60MatchRules)

	if err != nil {
		return err
	}

	*o = MemberDhcppropertiesOption60MatchRules(varMemberDhcppropertiesOption60MatchRules)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "match_value")
		delete(additionalProperties, "option_space")
		delete(additionalProperties, "is_substring")
		delete(additionalProperties, "substring_offset")
		delete(additionalProperties, "substring_length")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemberDhcppropertiesOption60MatchRules struct {
	value *MemberDhcppropertiesOption60MatchRules
	isSet bool
}

func (v NullableMemberDhcppropertiesOption60MatchRules) Get() *MemberDhcppropertiesOption60MatchRules {
	return v.value
}

func (v *NullableMemberDhcppropertiesOption60MatchRules) Set(val *MemberDhcppropertiesOption60MatchRules) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberDhcppropertiesOption60MatchRules) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberDhcppropertiesOption60MatchRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberDhcppropertiesOption60MatchRules(val *MemberDhcppropertiesOption60MatchRules) *NullableMemberDhcppropertiesOption60MatchRules {
	return &NullableMemberDhcppropertiesOption60MatchRules{value: val, isSet: true}
}

func (v NullableMemberDhcppropertiesOption60MatchRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberDhcppropertiesOption60MatchRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
