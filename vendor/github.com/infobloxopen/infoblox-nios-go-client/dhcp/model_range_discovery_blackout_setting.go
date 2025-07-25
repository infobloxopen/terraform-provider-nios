/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
)

// checks if the RangeDiscoveryBlackoutSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangeDiscoveryBlackoutSetting{}

// RangeDiscoveryBlackoutSetting struct for RangeDiscoveryBlackoutSetting
type RangeDiscoveryBlackoutSetting struct {
	// Determines whether a blackout is enabled or not.
	EnableBlackout *bool `json:"enable_blackout,omitempty"`
	// The blackout duration in seconds; minimum value is 1 minute.
	BlackoutDuration     *int64                                         `json:"blackout_duration,omitempty"`
	BlackoutSchedule     *RangediscoveryblackoutsettingBlackoutSchedule `json:"blackout_schedule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RangeDiscoveryBlackoutSetting RangeDiscoveryBlackoutSetting

// NewRangeDiscoveryBlackoutSetting instantiates a new RangeDiscoveryBlackoutSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangeDiscoveryBlackoutSetting() *RangeDiscoveryBlackoutSetting {
	this := RangeDiscoveryBlackoutSetting{}
	return &this
}

// NewRangeDiscoveryBlackoutSettingWithDefaults instantiates a new RangeDiscoveryBlackoutSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeDiscoveryBlackoutSettingWithDefaults() *RangeDiscoveryBlackoutSetting {
	this := RangeDiscoveryBlackoutSetting{}
	return &this
}

// GetEnableBlackout returns the EnableBlackout field value if set, zero value otherwise.
func (o *RangeDiscoveryBlackoutSetting) GetEnableBlackout() bool {
	if o == nil || IsNil(o.EnableBlackout) {
		var ret bool
		return ret
	}
	return *o.EnableBlackout
}

// GetEnableBlackoutOk returns a tuple with the EnableBlackout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeDiscoveryBlackoutSetting) GetEnableBlackoutOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableBlackout) {
		return nil, false
	}
	return o.EnableBlackout, true
}

// HasEnableBlackout returns a boolean if a field has been set.
func (o *RangeDiscoveryBlackoutSetting) HasEnableBlackout() bool {
	if o != nil && !IsNil(o.EnableBlackout) {
		return true
	}

	return false
}

// SetEnableBlackout gets a reference to the given bool and assigns it to the EnableBlackout field.
func (o *RangeDiscoveryBlackoutSetting) SetEnableBlackout(v bool) {
	o.EnableBlackout = &v
}

// GetBlackoutDuration returns the BlackoutDuration field value if set, zero value otherwise.
func (o *RangeDiscoveryBlackoutSetting) GetBlackoutDuration() int64 {
	if o == nil || IsNil(o.BlackoutDuration) {
		var ret int64
		return ret
	}
	return *o.BlackoutDuration
}

// GetBlackoutDurationOk returns a tuple with the BlackoutDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeDiscoveryBlackoutSetting) GetBlackoutDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.BlackoutDuration) {
		return nil, false
	}
	return o.BlackoutDuration, true
}

// HasBlackoutDuration returns a boolean if a field has been set.
func (o *RangeDiscoveryBlackoutSetting) HasBlackoutDuration() bool {
	if o != nil && !IsNil(o.BlackoutDuration) {
		return true
	}

	return false
}

// SetBlackoutDuration gets a reference to the given int64 and assigns it to the BlackoutDuration field.
func (o *RangeDiscoveryBlackoutSetting) SetBlackoutDuration(v int64) {
	o.BlackoutDuration = &v
}

// GetBlackoutSchedule returns the BlackoutSchedule field value if set, zero value otherwise.
func (o *RangeDiscoveryBlackoutSetting) GetBlackoutSchedule() RangediscoveryblackoutsettingBlackoutSchedule {
	if o == nil || IsNil(o.BlackoutSchedule) {
		var ret RangediscoveryblackoutsettingBlackoutSchedule
		return ret
	}
	return *o.BlackoutSchedule
}

// GetBlackoutScheduleOk returns a tuple with the BlackoutSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeDiscoveryBlackoutSetting) GetBlackoutScheduleOk() (*RangediscoveryblackoutsettingBlackoutSchedule, bool) {
	if o == nil || IsNil(o.BlackoutSchedule) {
		return nil, false
	}
	return o.BlackoutSchedule, true
}

// HasBlackoutSchedule returns a boolean if a field has been set.
func (o *RangeDiscoveryBlackoutSetting) HasBlackoutSchedule() bool {
	if o != nil && !IsNil(o.BlackoutSchedule) {
		return true
	}

	return false
}

// SetBlackoutSchedule gets a reference to the given RangediscoveryblackoutsettingBlackoutSchedule and assigns it to the BlackoutSchedule field.
func (o *RangeDiscoveryBlackoutSetting) SetBlackoutSchedule(v RangediscoveryblackoutsettingBlackoutSchedule) {
	o.BlackoutSchedule = &v
}

func (o RangeDiscoveryBlackoutSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RangeDiscoveryBlackoutSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableBlackout) {
		toSerialize["enable_blackout"] = o.EnableBlackout
	}
	if !IsNil(o.BlackoutDuration) {
		toSerialize["blackout_duration"] = o.BlackoutDuration
	}
	if !IsNil(o.BlackoutSchedule) {
		toSerialize["blackout_schedule"] = o.BlackoutSchedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RangeDiscoveryBlackoutSetting) UnmarshalJSON(data []byte) (err error) {
	varRangeDiscoveryBlackoutSetting := _RangeDiscoveryBlackoutSetting{}

	err = json.Unmarshal(data, &varRangeDiscoveryBlackoutSetting)

	if err != nil {
		return err
	}

	*o = RangeDiscoveryBlackoutSetting(varRangeDiscoveryBlackoutSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enable_blackout")
		delete(additionalProperties, "blackout_duration")
		delete(additionalProperties, "blackout_schedule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRangeDiscoveryBlackoutSetting struct {
	value *RangeDiscoveryBlackoutSetting
	isSet bool
}

func (v NullableRangeDiscoveryBlackoutSetting) Get() *RangeDiscoveryBlackoutSetting {
	return v.value
}

func (v *NullableRangeDiscoveryBlackoutSetting) Set(val *RangeDiscoveryBlackoutSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeDiscoveryBlackoutSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeDiscoveryBlackoutSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeDiscoveryBlackoutSetting(val *RangeDiscoveryBlackoutSetting) *NullableRangeDiscoveryBlackoutSetting {
	return &NullableRangeDiscoveryBlackoutSetting{value: val, isSet: true}
}

func (v NullableRangeDiscoveryBlackoutSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRangeDiscoveryBlackoutSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
