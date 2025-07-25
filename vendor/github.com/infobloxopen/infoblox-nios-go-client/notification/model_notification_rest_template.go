/*
Infoblox NOTIFICATION API

OpenAPI specification for Infoblox NIOS WAPI NOTIFICATION objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"
)

// checks if the NotificationRestTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationRestTemplate{}

// NotificationRestTemplate struct for NotificationRestTemplate
type NotificationRestTemplate struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The action name.
	ActionName *string `json:"action_name,omitempty"`
	// The time stamp when a template was added.
	AddedOn *int64 `json:"added_on,omitempty"`
	// The comment for this REST API template.
	Comment *string `json:"comment,omitempty"`
	// The JSON formatted content of a template. The data passed by content creates parameters for a template.
	Content *string `json:"content,omitempty"`
	// The event type.
	EventType []string `json:"event_type,omitempty"`
	// The name of a notification REST template.
	Name *string `json:"name,omitempty"`
	// The outbound type for the template.
	OutboundType *string `json:"outbound_type,omitempty"`
	// The notification REST template parameters.
	Parameters []NotificationRestTemplateParameters `json:"parameters,omitempty"`
	// The template type.
	TemplateType *string `json:"template_type,omitempty"`
	// The vendor identifier.
	VendorIdentifier *string `json:"vendor_identifier,omitempty"`
}

// NewNotificationRestTemplate instantiates a new NotificationRestTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationRestTemplate() *NotificationRestTemplate {
	this := NotificationRestTemplate{}
	return &this
}

// NewNotificationRestTemplateWithDefaults instantiates a new NotificationRestTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationRestTemplateWithDefaults() *NotificationRestTemplate {
	this := NotificationRestTemplate{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *NotificationRestTemplate) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRestTemplate) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *NotificationRestTemplate) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *NotificationRestTemplate) SetRef(v string) {
	o.Ref = &v
}

// GetActionName returns the ActionName field value if set, zero value otherwise.
func (o *NotificationRestTemplate) GetActionName() string {
	if o == nil || IsNil(o.ActionName) {
		var ret string
		return ret
	}
	return *o.ActionName
}

// GetActionNameOk returns a tuple with the ActionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRestTemplate) GetActionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ActionName) {
		return nil, false
	}
	return o.ActionName, true
}

// HasActionName returns a boolean if a field has been set.
func (o *NotificationRestTemplate) HasActionName() bool {
	if o != nil && !IsNil(o.ActionName) {
		return true
	}

	return false
}

// SetActionName gets a reference to the given string and assigns it to the ActionName field.
func (o *NotificationRestTemplate) SetActionName(v string) {
	o.ActionName = &v
}

// GetAddedOn returns the AddedOn field value if set, zero value otherwise.
func (o *NotificationRestTemplate) GetAddedOn() int64 {
	if o == nil || IsNil(o.AddedOn) {
		var ret int64
		return ret
	}
	return *o.AddedOn
}

// GetAddedOnOk returns a tuple with the AddedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRestTemplate) GetAddedOnOk() (*int64, bool) {
	if o == nil || IsNil(o.AddedOn) {
		return nil, false
	}
	return o.AddedOn, true
}

// HasAddedOn returns a boolean if a field has been set.
func (o *NotificationRestTemplate) HasAddedOn() bool {
	if o != nil && !IsNil(o.AddedOn) {
		return true
	}

	return false
}

// SetAddedOn gets a reference to the given int64 and assigns it to the AddedOn field.
func (o *NotificationRestTemplate) SetAddedOn(v int64) {
	o.AddedOn = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NotificationRestTemplate) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRestTemplate) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NotificationRestTemplate) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NotificationRestTemplate) SetComment(v string) {
	o.Comment = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *NotificationRestTemplate) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRestTemplate) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *NotificationRestTemplate) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *NotificationRestTemplate) SetContent(v string) {
	o.Content = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *NotificationRestTemplate) GetEventType() []string {
	if o == nil || IsNil(o.EventType) {
		var ret []string
		return ret
	}
	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRestTemplate) GetEventTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *NotificationRestTemplate) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given []string and assigns it to the EventType field.
func (o *NotificationRestTemplate) SetEventType(v []string) {
	o.EventType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NotificationRestTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRestTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NotificationRestTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NotificationRestTemplate) SetName(v string) {
	o.Name = &v
}

// GetOutboundType returns the OutboundType field value if set, zero value otherwise.
func (o *NotificationRestTemplate) GetOutboundType() string {
	if o == nil || IsNil(o.OutboundType) {
		var ret string
		return ret
	}
	return *o.OutboundType
}

// GetOutboundTypeOk returns a tuple with the OutboundType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRestTemplate) GetOutboundTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OutboundType) {
		return nil, false
	}
	return o.OutboundType, true
}

// HasOutboundType returns a boolean if a field has been set.
func (o *NotificationRestTemplate) HasOutboundType() bool {
	if o != nil && !IsNil(o.OutboundType) {
		return true
	}

	return false
}

// SetOutboundType gets a reference to the given string and assigns it to the OutboundType field.
func (o *NotificationRestTemplate) SetOutboundType(v string) {
	o.OutboundType = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *NotificationRestTemplate) GetParameters() []NotificationRestTemplateParameters {
	if o == nil || IsNil(o.Parameters) {
		var ret []NotificationRestTemplateParameters
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRestTemplate) GetParametersOk() ([]NotificationRestTemplateParameters, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *NotificationRestTemplate) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []NotificationRestTemplateParameters and assigns it to the Parameters field.
func (o *NotificationRestTemplate) SetParameters(v []NotificationRestTemplateParameters) {
	o.Parameters = v
}

// GetTemplateType returns the TemplateType field value if set, zero value otherwise.
func (o *NotificationRestTemplate) GetTemplateType() string {
	if o == nil || IsNil(o.TemplateType) {
		var ret string
		return ret
	}
	return *o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRestTemplate) GetTemplateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateType) {
		return nil, false
	}
	return o.TemplateType, true
}

// HasTemplateType returns a boolean if a field has been set.
func (o *NotificationRestTemplate) HasTemplateType() bool {
	if o != nil && !IsNil(o.TemplateType) {
		return true
	}

	return false
}

// SetTemplateType gets a reference to the given string and assigns it to the TemplateType field.
func (o *NotificationRestTemplate) SetTemplateType(v string) {
	o.TemplateType = &v
}

// GetVendorIdentifier returns the VendorIdentifier field value if set, zero value otherwise.
func (o *NotificationRestTemplate) GetVendorIdentifier() string {
	if o == nil || IsNil(o.VendorIdentifier) {
		var ret string
		return ret
	}
	return *o.VendorIdentifier
}

// GetVendorIdentifierOk returns a tuple with the VendorIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRestTemplate) GetVendorIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.VendorIdentifier) {
		return nil, false
	}
	return o.VendorIdentifier, true
}

// HasVendorIdentifier returns a boolean if a field has been set.
func (o *NotificationRestTemplate) HasVendorIdentifier() bool {
	if o != nil && !IsNil(o.VendorIdentifier) {
		return true
	}

	return false
}

// SetVendorIdentifier gets a reference to the given string and assigns it to the VendorIdentifier field.
func (o *NotificationRestTemplate) SetVendorIdentifier(v string) {
	o.VendorIdentifier = &v
}

func (o NotificationRestTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationRestTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.ActionName) {
		toSerialize["action_name"] = o.ActionName
	}
	if !IsNil(o.AddedOn) {
		toSerialize["added_on"] = o.AddedOn
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OutboundType) {
		toSerialize["outbound_type"] = o.OutboundType
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.TemplateType) {
		toSerialize["template_type"] = o.TemplateType
	}
	if !IsNil(o.VendorIdentifier) {
		toSerialize["vendor_identifier"] = o.VendorIdentifier
	}
	return toSerialize, nil
}

type NullableNotificationRestTemplate struct {
	value *NotificationRestTemplate
	isSet bool
}

func (v NullableNotificationRestTemplate) Get() *NotificationRestTemplate {
	return v.value
}

func (v *NullableNotificationRestTemplate) Set(val *NotificationRestTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationRestTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationRestTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationRestTemplate(val *NotificationRestTemplate) *NullableNotificationRestTemplate {
	return &NullableNotificationRestTemplate{value: val, isSet: true}
}

func (v NullableNotificationRestTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationRestTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
