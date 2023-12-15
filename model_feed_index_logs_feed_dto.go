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

// checks if the FeedIndexLogsFeedDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedIndexLogsFeedDTO{}

// FeedIndexLogsFeedDTO Информация о прайс-листе.
type FeedIndexLogsFeedDTO struct {
	// Идентификатор прайс-листа.
	Id *int64 `json:"id,omitempty"`
}

// NewFeedIndexLogsFeedDTO instantiates a new FeedIndexLogsFeedDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedIndexLogsFeedDTO() *FeedIndexLogsFeedDTO {
	this := FeedIndexLogsFeedDTO{}
	return &this
}

// NewFeedIndexLogsFeedDTOWithDefaults instantiates a new FeedIndexLogsFeedDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedIndexLogsFeedDTOWithDefaults() *FeedIndexLogsFeedDTO {
	this := FeedIndexLogsFeedDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FeedIndexLogsFeedDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndexLogsFeedDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FeedIndexLogsFeedDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *FeedIndexLogsFeedDTO) SetId(v int64) {
	o.Id = &v
}

func (o FeedIndexLogsFeedDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedIndexLogsFeedDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableFeedIndexLogsFeedDTO struct {
	value *FeedIndexLogsFeedDTO
	isSet bool
}

func (v NullableFeedIndexLogsFeedDTO) Get() *FeedIndexLogsFeedDTO {
	return v.value
}

func (v *NullableFeedIndexLogsFeedDTO) Set(val *FeedIndexLogsFeedDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedIndexLogsFeedDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedIndexLogsFeedDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedIndexLogsFeedDTO(val *FeedIndexLogsFeedDTO) *NullableFeedIndexLogsFeedDTO {
	return &NullableFeedIndexLogsFeedDTO{value: val, isSet: true}
}

func (v NullableFeedIndexLogsFeedDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedIndexLogsFeedDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


