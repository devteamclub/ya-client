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

// checks if the GetOrderLabelsDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrderLabelsDataResponse{}

// GetOrderLabelsDataResponse Ответ с информацией для печати ярлыков.
type GetOrderLabelsDataResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *OrderLabelDTO `json:"result,omitempty"`
}

// NewGetOrderLabelsDataResponse instantiates a new GetOrderLabelsDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderLabelsDataResponse() *GetOrderLabelsDataResponse {
	this := GetOrderLabelsDataResponse{}
	return &this
}

// NewGetOrderLabelsDataResponseWithDefaults instantiates a new GetOrderLabelsDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderLabelsDataResponseWithDefaults() *GetOrderLabelsDataResponse {
	this := GetOrderLabelsDataResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOrderLabelsDataResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderLabelsDataResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOrderLabelsDataResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *GetOrderLabelsDataResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetOrderLabelsDataResponse) GetResult() OrderLabelDTO {
	if o == nil || IsNil(o.Result) {
		var ret OrderLabelDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderLabelsDataResponse) GetResultOk() (*OrderLabelDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetOrderLabelsDataResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OrderLabelDTO and assigns it to the Result field.
func (o *GetOrderLabelsDataResponse) SetResult(v OrderLabelDTO) {
	o.Result = &v
}

func (o GetOrderLabelsDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderLabelsDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetOrderLabelsDataResponse struct {
	value *GetOrderLabelsDataResponse
	isSet bool
}

func (v NullableGetOrderLabelsDataResponse) Get() *GetOrderLabelsDataResponse {
	return v.value
}

func (v *NullableGetOrderLabelsDataResponse) Set(val *GetOrderLabelsDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderLabelsDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderLabelsDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderLabelsDataResponse(val *GetOrderLabelsDataResponse) *NullableGetOrderLabelsDataResponse {
	return &NullableGetOrderLabelsDataResponse{value: val, isSet: true}
}

func (v NullableGetOrderLabelsDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderLabelsDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


