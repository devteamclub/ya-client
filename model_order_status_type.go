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

// OrderStatusType Статус заказа:  * `CANCELLED` — заказ отменен.  * `DELIVERED` — заказ получен покупателем.  * `DELIVERY` — заказ передан в службу доставки.  * `PICKUP` — заказ доставлен в пункт самовывоза.  * `PROCESSING` — заказ находится в обработке.  * `UNPAID` — заказ оформлен, но еще не оплачен (если выбрана оплата при оформлении).  Также могут возвращаться другие значения. Обрабатывать их не требуется. 
type OrderStatusType string

// List of OrderStatusType
const (
	PLACING OrderStatusType = "PLACING"
	RESERVED OrderStatusType = "RESERVED"
	UNPAID OrderStatusType = "UNPAID"
	PROCESSING OrderStatusType = "PROCESSING"
	DELIVERY OrderStatusType = "DELIVERY"
	PICKUP OrderStatusType = "PICKUP"
	DELIVERED OrderStatusType = "DELIVERED"
	CANCELLED OrderStatusType = "CANCELLED"
	PENDING OrderStatusType = "PENDING"
	REJECTED OrderStatusType = "REJECTED"
	PARTIALLY_RETURNED OrderStatusType = "PARTIALLY_RETURNED"
	RETURNED OrderStatusType = "RETURNED"
	CANCELLED_WITHOUT_REFUND OrderStatusType = "CANCELLED_WITHOUT_REFUND"
	UNKNOWN OrderStatusType = "UNKNOWN"
)

// All allowed values of OrderStatusType enum
var AllowedOrderStatusTypeEnumValues = []OrderStatusType{
	"PLACING",
	"RESERVED",
	"UNPAID",
	"PROCESSING",
	"DELIVERY",
	"PICKUP",
	"DELIVERED",
	"CANCELLED",
	"PENDING",
	"REJECTED",
	"PARTIALLY_RETURNED",
	"RETURNED",
	"CANCELLED_WITHOUT_REFUND",
	"UNKNOWN",
}

func (v *OrderStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderStatusType(value)
	for _, existing := range AllowedOrderStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderStatusType", value)
}

// NewOrderStatusTypeFromValue returns a pointer to a valid OrderStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderStatusTypeFromValue(v string) (*OrderStatusType, error) {
	ev := OrderStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderStatusType: valid values are %v", v, AllowedOrderStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderStatusType) IsValid() bool {
	for _, existing := range AllowedOrderStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderStatusType value
func (v OrderStatusType) Ptr() *OrderStatusType {
	return &v
}

type NullableOrderStatusType struct {
	value *OrderStatusType
	isSet bool
}

func (v NullableOrderStatusType) Get() *OrderStatusType {
	return v.value
}

func (v *NullableOrderStatusType) Set(val *OrderStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatusType(val *OrderStatusType) *NullableOrderStatusType {
	return &NullableOrderStatusType{value: val, isSet: true}
}

func (v NullableOrderStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

