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

// checks if the FeedbackListDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedbackListDTO{}

// FeedbackListDTO Отзывы пользователей Яндекс.Маркета об указанном магазине.
type FeedbackListDTO struct {
	Paging *ScrollingPagerDTO `json:"paging,omitempty"`
	// Список отзывов.  Содержит не более 20 отзывов. 
	FeedbackList []FeedbackDTO `json:"feedbackList,omitempty"`
}

// NewFeedbackListDTO instantiates a new FeedbackListDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackListDTO() *FeedbackListDTO {
	this := FeedbackListDTO{}
	return &this
}

// NewFeedbackListDTOWithDefaults instantiates a new FeedbackListDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackListDTOWithDefaults() *FeedbackListDTO {
	this := FeedbackListDTO{}
	return &this
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *FeedbackListDTO) GetPaging() ScrollingPagerDTO {
	if o == nil || IsNil(o.Paging) {
		var ret ScrollingPagerDTO
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackListDTO) GetPagingOk() (*ScrollingPagerDTO, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *FeedbackListDTO) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ScrollingPagerDTO and assigns it to the Paging field.
func (o *FeedbackListDTO) SetPaging(v ScrollingPagerDTO) {
	o.Paging = &v
}

// GetFeedbackList returns the FeedbackList field value if set, zero value otherwise.
func (o *FeedbackListDTO) GetFeedbackList() []FeedbackDTO {
	if o == nil || IsNil(o.FeedbackList) {
		var ret []FeedbackDTO
		return ret
	}
	return o.FeedbackList
}

// GetFeedbackListOk returns a tuple with the FeedbackList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackListDTO) GetFeedbackListOk() ([]FeedbackDTO, bool) {
	if o == nil || IsNil(o.FeedbackList) {
		return nil, false
	}
	return o.FeedbackList, true
}

// HasFeedbackList returns a boolean if a field has been set.
func (o *FeedbackListDTO) HasFeedbackList() bool {
	if o != nil && !IsNil(o.FeedbackList) {
		return true
	}

	return false
}

// SetFeedbackList gets a reference to the given []FeedbackDTO and assigns it to the FeedbackList field.
func (o *FeedbackListDTO) SetFeedbackList(v []FeedbackDTO) {
	o.FeedbackList = v
}

func (o FeedbackListDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedbackListDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	if !IsNil(o.FeedbackList) {
		toSerialize["feedbackList"] = o.FeedbackList
	}
	return toSerialize, nil
}

type NullableFeedbackListDTO struct {
	value *FeedbackListDTO
	isSet bool
}

func (v NullableFeedbackListDTO) Get() *FeedbackListDTO {
	return v.value
}

func (v *NullableFeedbackListDTO) Set(val *FeedbackListDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackListDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackListDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackListDTO(val *FeedbackListDTO) *NullableFeedbackListDTO {
	return &NullableFeedbackListDTO{value: val, isSet: true}
}

func (v NullableFeedbackListDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackListDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

