/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the TransferOrdersFromShipmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferOrdersFromShipmentRequest{}

// TransferOrdersFromShipmentRequest Запрос переноса заказов из отгрузки.
type TransferOrdersFromShipmentRequest struct {
	// Список заказов, которые вы не успеваете подготовить.
	OrderIds []int64 `json:"orderIds"`
}

type _TransferOrdersFromShipmentRequest TransferOrdersFromShipmentRequest

// NewTransferOrdersFromShipmentRequest instantiates a new TransferOrdersFromShipmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferOrdersFromShipmentRequest(orderIds []int64) *TransferOrdersFromShipmentRequest {
	this := TransferOrdersFromShipmentRequest{}
	this.OrderIds = orderIds
	return &this
}

// NewTransferOrdersFromShipmentRequestWithDefaults instantiates a new TransferOrdersFromShipmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferOrdersFromShipmentRequestWithDefaults() *TransferOrdersFromShipmentRequest {
	this := TransferOrdersFromShipmentRequest{}
	return &this
}

// GetOrderIds returns the OrderIds field value
func (o *TransferOrdersFromShipmentRequest) GetOrderIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.OrderIds
}

// GetOrderIdsOk returns a tuple with the OrderIds field value
// and a boolean to check if the value has been set.
func (o *TransferOrdersFromShipmentRequest) GetOrderIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderIds, true
}

// SetOrderIds sets field value
func (o *TransferOrdersFromShipmentRequest) SetOrderIds(v []int64) {
	o.OrderIds = v
}

func (o TransferOrdersFromShipmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferOrdersFromShipmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderIds"] = o.OrderIds
	return toSerialize, nil
}

func (o *TransferOrdersFromShipmentRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orderIds",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTransferOrdersFromShipmentRequest := _TransferOrdersFromShipmentRequest{}

	err = json.Unmarshal(bytes, &varTransferOrdersFromShipmentRequest)

	if err != nil {
		return err
	}

	*o = TransferOrdersFromShipmentRequest(varTransferOrdersFromShipmentRequest)

	return err
}

type NullableTransferOrdersFromShipmentRequest struct {
	value *TransferOrdersFromShipmentRequest
	isSet bool
}

func (v NullableTransferOrdersFromShipmentRequest) Get() *TransferOrdersFromShipmentRequest {
	return v.value
}

func (v *NullableTransferOrdersFromShipmentRequest) Set(val *TransferOrdersFromShipmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferOrdersFromShipmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferOrdersFromShipmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferOrdersFromShipmentRequest(val *TransferOrdersFromShipmentRequest) *NullableTransferOrdersFromShipmentRequest {
	return &NullableTransferOrdersFromShipmentRequest{value: val, isSet: true}
}

func (v NullableTransferOrdersFromShipmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferOrdersFromShipmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


