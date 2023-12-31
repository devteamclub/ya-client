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

// checks if the GetModelsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetModelsRequest{}

// GetModelsRequest Запрос информации о моделях.
type GetModelsRequest struct {
	// Список моделей.
	Models []int64 `json:"models,omitempty"`
}

// NewGetModelsRequest instantiates a new GetModelsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetModelsRequest() *GetModelsRequest {
	this := GetModelsRequest{}
	return &this
}

// NewGetModelsRequestWithDefaults instantiates a new GetModelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetModelsRequestWithDefaults() *GetModelsRequest {
	this := GetModelsRequest{}
	return &this
}

// GetModels returns the Models field value if set, zero value otherwise.
func (o *GetModelsRequest) GetModels() []int64 {
	if o == nil || IsNil(o.Models) {
		var ret []int64
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetModelsRequest) GetModelsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Models) {
		return nil, false
	}
	return o.Models, true
}

// HasModels returns a boolean if a field has been set.
func (o *GetModelsRequest) HasModels() bool {
	if o != nil && !IsNil(o.Models) {
		return true
	}

	return false
}

// SetModels gets a reference to the given []int64 and assigns it to the Models field.
func (o *GetModelsRequest) SetModels(v []int64) {
	o.Models = v
}

func (o GetModelsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetModelsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Models) {
		toSerialize["models"] = o.Models
	}
	return toSerialize, nil
}

type NullableGetModelsRequest struct {
	value *GetModelsRequest
	isSet bool
}

func (v NullableGetModelsRequest) Get() *GetModelsRequest {
	return v.value
}

func (v *NullableGetModelsRequest) Set(val *GetModelsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetModelsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetModelsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetModelsRequest(val *GetModelsRequest) *NullableGetModelsRequest {
	return &NullableGetModelsRequest{value: val, isSet: true}
}

func (v NullableGetModelsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetModelsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


