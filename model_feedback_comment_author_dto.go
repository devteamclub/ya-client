/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the FeedbackCommentAuthorDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedbackCommentAuthorDTO{}

// FeedbackCommentAuthorDTO Информация об авторе ответа.
type FeedbackCommentAuthorDTO struct {
	Type *FeedbackCommentAuthorType `json:"type,omitempty"`
	// Имя автора отзыва или название магазина.
	Name *string `json:"name,omitempty"`
}

// NewFeedbackCommentAuthorDTO instantiates a new FeedbackCommentAuthorDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackCommentAuthorDTO() *FeedbackCommentAuthorDTO {
	this := FeedbackCommentAuthorDTO{}
	return &this
}

// NewFeedbackCommentAuthorDTOWithDefaults instantiates a new FeedbackCommentAuthorDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackCommentAuthorDTOWithDefaults() *FeedbackCommentAuthorDTO {
	this := FeedbackCommentAuthorDTO{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FeedbackCommentAuthorDTO) GetType() FeedbackCommentAuthorType {
	if o == nil || IsNil(o.Type) {
		var ret FeedbackCommentAuthorType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackCommentAuthorDTO) GetTypeOk() (*FeedbackCommentAuthorType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FeedbackCommentAuthorDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FeedbackCommentAuthorType and assigns it to the Type field.
func (o *FeedbackCommentAuthorDTO) SetType(v FeedbackCommentAuthorType) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FeedbackCommentAuthorDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackCommentAuthorDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FeedbackCommentAuthorDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FeedbackCommentAuthorDTO) SetName(v string) {
	o.Name = &v
}

func (o FeedbackCommentAuthorDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedbackCommentAuthorDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableFeedbackCommentAuthorDTO struct {
	value *FeedbackCommentAuthorDTO
	isSet bool
}

func (v NullableFeedbackCommentAuthorDTO) Get() *FeedbackCommentAuthorDTO {
	return v.value
}

func (v *NullableFeedbackCommentAuthorDTO) Set(val *FeedbackCommentAuthorDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackCommentAuthorDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackCommentAuthorDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackCommentAuthorDTO(val *FeedbackCommentAuthorDTO) *NullableFeedbackCommentAuthorDTO {
	return &NullableFeedbackCommentAuthorDTO{value: val, isSet: true}
}

func (v NullableFeedbackCommentAuthorDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackCommentAuthorDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


