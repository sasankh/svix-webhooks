/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CompletionMessage struct for CompletionMessage
type CompletionMessage struct {
	Content string `json:"content"`
	Role string `json:"role"`
}

// NewCompletionMessage instantiates a new CompletionMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompletionMessage(content string, role string) *CompletionMessage {
	this := CompletionMessage{}
	this.Content = content
	this.Role = role
	return &this
}

// NewCompletionMessageWithDefaults instantiates a new CompletionMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompletionMessageWithDefaults() *CompletionMessage {
	this := CompletionMessage{}
	return &this
}

// GetContent returns the Content field value
func (o *CompletionMessage) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CompletionMessage) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CompletionMessage) SetContent(v string) {
	o.Content = v
}

// GetRole returns the Role field value
func (o *CompletionMessage) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *CompletionMessage) GetRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *CompletionMessage) SetRole(v string) {
	o.Role = v
}

func (o CompletionMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}
	if true {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableCompletionMessage struct {
	value *CompletionMessage
	isSet bool
}

func (v NullableCompletionMessage) Get() *CompletionMessage {
	return v.value
}

func (v *NullableCompletionMessage) Set(val *CompletionMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletionMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletionMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletionMessage(val *CompletionMessage) *NullableCompletionMessage {
	return &NullableCompletionMessage{value: val, isSet: true}
}

func (v NullableCompletionMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletionMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

