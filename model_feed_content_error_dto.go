/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ya-client

import (
	"encoding/json"
)

// checks if the FeedContentErrorDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedContentErrorDTO{}

// FeedContentErrorDTO Информация об ошибке в содержимом прайс-листа. Выводится, если параметр `content status=ERROR`. 
type FeedContentErrorDTO struct {
	Type *FeedContentErrorType `json:"type,omitempty"`
}

// NewFeedContentErrorDTO instantiates a new FeedContentErrorDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedContentErrorDTO() *FeedContentErrorDTO {
	this := FeedContentErrorDTO{}
	return &this
}

// NewFeedContentErrorDTOWithDefaults instantiates a new FeedContentErrorDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedContentErrorDTOWithDefaults() *FeedContentErrorDTO {
	this := FeedContentErrorDTO{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FeedContentErrorDTO) GetType() FeedContentErrorType {
	if o == nil || IsNil(o.Type) {
		var ret FeedContentErrorType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedContentErrorDTO) GetTypeOk() (*FeedContentErrorType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FeedContentErrorDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FeedContentErrorType and assigns it to the Type field.
func (o *FeedContentErrorDTO) SetType(v FeedContentErrorType) {
	o.Type = &v
}

func (o FeedContentErrorDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedContentErrorDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableFeedContentErrorDTO struct {
	value *FeedContentErrorDTO
	isSet bool
}

func (v NullableFeedContentErrorDTO) Get() *FeedContentErrorDTO {
	return v.value
}

func (v *NullableFeedContentErrorDTO) Set(val *FeedContentErrorDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedContentErrorDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedContentErrorDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedContentErrorDTO(val *FeedContentErrorDTO) *NullableFeedContentErrorDTO {
	return &NullableFeedContentErrorDTO{value: val, isSet: true}
}

func (v NullableFeedContentErrorDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedContentErrorDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

