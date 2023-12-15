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

// checks if the EmptyApiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmptyApiResponse{}

// EmptyApiResponse Пустой ответ сервера.
type EmptyApiResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
}

// NewEmptyApiResponse instantiates a new EmptyApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmptyApiResponse() *EmptyApiResponse {
	this := EmptyApiResponse{}
	return &this
}

// NewEmptyApiResponseWithDefaults instantiates a new EmptyApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmptyApiResponseWithDefaults() *EmptyApiResponse {
	this := EmptyApiResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EmptyApiResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmptyApiResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EmptyApiResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *EmptyApiResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

func (o EmptyApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmptyApiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableEmptyApiResponse struct {
	value *EmptyApiResponse
	isSet bool
}

func (v NullableEmptyApiResponse) Get() *EmptyApiResponse {
	return v.value
}

func (v *NullableEmptyApiResponse) Set(val *EmptyApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmptyApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmptyApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmptyApiResponse(val *EmptyApiResponse) *NullableEmptyApiResponse {
	return &NullableEmptyApiResponse{value: val, isSet: true}
}

func (v NullableEmptyApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmptyApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


