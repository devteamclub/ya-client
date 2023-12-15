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

// checks if the UpdateOrderStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrderStatusResponse{}

// UpdateOrderStatusResponse Информация об изменении статуса заказа.
type UpdateOrderStatusResponse struct {
	Order *OrderDTO `json:"order,omitempty"`
}

// NewUpdateOrderStatusResponse instantiates a new UpdateOrderStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrderStatusResponse() *UpdateOrderStatusResponse {
	this := UpdateOrderStatusResponse{}
	return &this
}

// NewUpdateOrderStatusResponseWithDefaults instantiates a new UpdateOrderStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrderStatusResponseWithDefaults() *UpdateOrderStatusResponse {
	this := UpdateOrderStatusResponse{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *UpdateOrderStatusResponse) GetOrder() OrderDTO {
	if o == nil || IsNil(o.Order) {
		var ret OrderDTO
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderStatusResponse) GetOrderOk() (*OrderDTO, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *UpdateOrderStatusResponse) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given OrderDTO and assigns it to the Order field.
func (o *UpdateOrderStatusResponse) SetOrder(v OrderDTO) {
	o.Order = &v
}

func (o UpdateOrderStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrderStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullableUpdateOrderStatusResponse struct {
	value *UpdateOrderStatusResponse
	isSet bool
}

func (v NullableUpdateOrderStatusResponse) Get() *UpdateOrderStatusResponse {
	return v.value
}

func (v *NullableUpdateOrderStatusResponse) Set(val *UpdateOrderStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrderStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrderStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrderStatusResponse(val *UpdateOrderStatusResponse) *NullableUpdateOrderStatusResponse {
	return &NullableUpdateOrderStatusResponse{value: val, isSet: true}
}

func (v NullableUpdateOrderStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrderStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

