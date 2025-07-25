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

// checks if the MemberBgpAs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberBgpAs{}

// MemberBgpAs struct for MemberBgpAs
type MemberBgpAs struct {
	// The number of this autonomous system.
	As *int64 `json:"as,omitempty"`
	// The AS keepalive timer (in seconds). The valid value is from 1 to 21845.
	Keepalive *int64 `json:"keepalive,omitempty"`
	// The AS holddown timer (in seconds). The valid value is from 3 to 65535.
	Holddown *int64 `json:"holddown,omitempty"`
	// The BGP neighbors for this AS.
	Neighbors []MemberbgpasNeighbors `json:"neighbors,omitempty"`
	// Determines if link detection on the interface is enabled or not.
	LinkDetect           *bool `json:"link_detect,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemberBgpAs MemberBgpAs

// NewMemberBgpAs instantiates a new MemberBgpAs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberBgpAs() *MemberBgpAs {
	this := MemberBgpAs{}
	return &this
}

// NewMemberBgpAsWithDefaults instantiates a new MemberBgpAs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberBgpAsWithDefaults() *MemberBgpAs {
	this := MemberBgpAs{}
	return &this
}

// GetAs returns the As field value if set, zero value otherwise.
func (o *MemberBgpAs) GetAs() int64 {
	if o == nil || IsNil(o.As) {
		var ret int64
		return ret
	}
	return *o.As
}

// GetAsOk returns a tuple with the As field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberBgpAs) GetAsOk() (*int64, bool) {
	if o == nil || IsNil(o.As) {
		return nil, false
	}
	return o.As, true
}

// HasAs returns a boolean if a field has been set.
func (o *MemberBgpAs) HasAs() bool {
	if o != nil && !IsNil(o.As) {
		return true
	}

	return false
}

// SetAs gets a reference to the given int64 and assigns it to the As field.
func (o *MemberBgpAs) SetAs(v int64) {
	o.As = &v
}

// GetKeepalive returns the Keepalive field value if set, zero value otherwise.
func (o *MemberBgpAs) GetKeepalive() int64 {
	if o == nil || IsNil(o.Keepalive) {
		var ret int64
		return ret
	}
	return *o.Keepalive
}

// GetKeepaliveOk returns a tuple with the Keepalive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberBgpAs) GetKeepaliveOk() (*int64, bool) {
	if o == nil || IsNil(o.Keepalive) {
		return nil, false
	}
	return o.Keepalive, true
}

// HasKeepalive returns a boolean if a field has been set.
func (o *MemberBgpAs) HasKeepalive() bool {
	if o != nil && !IsNil(o.Keepalive) {
		return true
	}

	return false
}

// SetKeepalive gets a reference to the given int64 and assigns it to the Keepalive field.
func (o *MemberBgpAs) SetKeepalive(v int64) {
	o.Keepalive = &v
}

// GetHolddown returns the Holddown field value if set, zero value otherwise.
func (o *MemberBgpAs) GetHolddown() int64 {
	if o == nil || IsNil(o.Holddown) {
		var ret int64
		return ret
	}
	return *o.Holddown
}

// GetHolddownOk returns a tuple with the Holddown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberBgpAs) GetHolddownOk() (*int64, bool) {
	if o == nil || IsNil(o.Holddown) {
		return nil, false
	}
	return o.Holddown, true
}

// HasHolddown returns a boolean if a field has been set.
func (o *MemberBgpAs) HasHolddown() bool {
	if o != nil && !IsNil(o.Holddown) {
		return true
	}

	return false
}

// SetHolddown gets a reference to the given int64 and assigns it to the Holddown field.
func (o *MemberBgpAs) SetHolddown(v int64) {
	o.Holddown = &v
}

// GetNeighbors returns the Neighbors field value if set, zero value otherwise.
func (o *MemberBgpAs) GetNeighbors() []MemberbgpasNeighbors {
	if o == nil || IsNil(o.Neighbors) {
		var ret []MemberbgpasNeighbors
		return ret
	}
	return o.Neighbors
}

// GetNeighborsOk returns a tuple with the Neighbors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberBgpAs) GetNeighborsOk() ([]MemberbgpasNeighbors, bool) {
	if o == nil || IsNil(o.Neighbors) {
		return nil, false
	}
	return o.Neighbors, true
}

// HasNeighbors returns a boolean if a field has been set.
func (o *MemberBgpAs) HasNeighbors() bool {
	if o != nil && !IsNil(o.Neighbors) {
		return true
	}

	return false
}

// SetNeighbors gets a reference to the given []MemberbgpasNeighbors and assigns it to the Neighbors field.
func (o *MemberBgpAs) SetNeighbors(v []MemberbgpasNeighbors) {
	o.Neighbors = v
}

// GetLinkDetect returns the LinkDetect field value if set, zero value otherwise.
func (o *MemberBgpAs) GetLinkDetect() bool {
	if o == nil || IsNil(o.LinkDetect) {
		var ret bool
		return ret
	}
	return *o.LinkDetect
}

// GetLinkDetectOk returns a tuple with the LinkDetect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberBgpAs) GetLinkDetectOk() (*bool, bool) {
	if o == nil || IsNil(o.LinkDetect) {
		return nil, false
	}
	return o.LinkDetect, true
}

// HasLinkDetect returns a boolean if a field has been set.
func (o *MemberBgpAs) HasLinkDetect() bool {
	if o != nil && !IsNil(o.LinkDetect) {
		return true
	}

	return false
}

// SetLinkDetect gets a reference to the given bool and assigns it to the LinkDetect field.
func (o *MemberBgpAs) SetLinkDetect(v bool) {
	o.LinkDetect = &v
}

func (o MemberBgpAs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberBgpAs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.As) {
		toSerialize["as"] = o.As
	}
	if !IsNil(o.Keepalive) {
		toSerialize["keepalive"] = o.Keepalive
	}
	if !IsNil(o.Holddown) {
		toSerialize["holddown"] = o.Holddown
	}
	if !IsNil(o.Neighbors) {
		toSerialize["neighbors"] = o.Neighbors
	}
	if !IsNil(o.LinkDetect) {
		toSerialize["link_detect"] = o.LinkDetect
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemberBgpAs) UnmarshalJSON(data []byte) (err error) {
	varMemberBgpAs := _MemberBgpAs{}

	err = json.Unmarshal(data, &varMemberBgpAs)

	if err != nil {
		return err
	}

	*o = MemberBgpAs(varMemberBgpAs)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "as")
		delete(additionalProperties, "keepalive")
		delete(additionalProperties, "holddown")
		delete(additionalProperties, "neighbors")
		delete(additionalProperties, "link_detect")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemberBgpAs struct {
	value *MemberBgpAs
	isSet bool
}

func (v NullableMemberBgpAs) Get() *MemberBgpAs {
	return v.value
}

func (v *NullableMemberBgpAs) Set(val *MemberBgpAs) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberBgpAs) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberBgpAs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberBgpAs(val *MemberBgpAs) *NullableMemberBgpAs {
	return &NullableMemberBgpAs{value: val, isSet: true}
}

func (v NullableMemberBgpAs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberBgpAs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
