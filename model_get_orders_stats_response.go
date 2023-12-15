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

// checks if the GetOrdersStatsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrdersStatsResponse{}

// GetOrdersStatsResponse Ответ на запрос отчета по заказам.
type GetOrdersStatsResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *OrdersStatsDTO `json:"result,omitempty"`
}

// NewGetOrdersStatsResponse instantiates a new GetOrdersStatsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrdersStatsResponse() *GetOrdersStatsResponse {
	this := GetOrdersStatsResponse{}
	return &this
}

// NewGetOrdersStatsResponseWithDefaults instantiates a new GetOrdersStatsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrdersStatsResponseWithDefaults() *GetOrdersStatsResponse {
	this := GetOrdersStatsResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOrdersStatsResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersStatsResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOrdersStatsResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *GetOrdersStatsResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetOrdersStatsResponse) GetResult() OrdersStatsDTO {
	if o == nil || IsNil(o.Result) {
		var ret OrdersStatsDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersStatsResponse) GetResultOk() (*OrdersStatsDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetOrdersStatsResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OrdersStatsDTO and assigns it to the Result field.
func (o *GetOrdersStatsResponse) SetResult(v OrdersStatsDTO) {
	o.Result = &v
}

func (o GetOrdersStatsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrdersStatsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetOrdersStatsResponse struct {
	value *GetOrdersStatsResponse
	isSet bool
}

func (v NullableGetOrdersStatsResponse) Get() *GetOrdersStatsResponse {
	return v.value
}

func (v *NullableGetOrdersStatsResponse) Set(val *GetOrdersStatsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrdersStatsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrdersStatsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrdersStatsResponse(val *GetOrdersStatsResponse) *NullableGetOrdersStatsResponse {
	return &NullableGetOrdersStatsResponse{value: val, isSet: true}
}

func (v NullableGetOrdersStatsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrdersStatsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


