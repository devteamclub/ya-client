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

// OrderDeliveryEacType Тип кода подтверждения ЭАПП: - MERCHANT_TO_COURIER — продавец передает код курьеру. - COURIER_TO_MERCHANT — курьер передает код продавцу. - CHECKING_BY_MERCHANT - продавец проверяет код на своей стороне. 
type OrderDeliveryEacType string

// List of OrderDeliveryEacType
const (
	MERCHANT_TO_COURIER OrderDeliveryEacType = "MERCHANT_TO_COURIER"
	COURIER_TO_MERCHANT OrderDeliveryEacType = "COURIER_TO_MERCHANT"
	CHECKING_BY_MERCHANT OrderDeliveryEacType = "CHECKING_BY_MERCHANT"
)

// All allowed values of OrderDeliveryEacType enum
var AllowedOrderDeliveryEacTypeEnumValues = []OrderDeliveryEacType{
	"MERCHANT_TO_COURIER",
	"COURIER_TO_MERCHANT",
	"CHECKING_BY_MERCHANT",
}

func (v *OrderDeliveryEacType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderDeliveryEacType(value)
	for _, existing := range AllowedOrderDeliveryEacTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderDeliveryEacType", value)
}

// NewOrderDeliveryEacTypeFromValue returns a pointer to a valid OrderDeliveryEacType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderDeliveryEacTypeFromValue(v string) (*OrderDeliveryEacType, error) {
	ev := OrderDeliveryEacType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderDeliveryEacType: valid values are %v", v, AllowedOrderDeliveryEacTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderDeliveryEacType) IsValid() bool {
	for _, existing := range AllowedOrderDeliveryEacTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderDeliveryEacType value
func (v OrderDeliveryEacType) Ptr() *OrderDeliveryEacType {
	return &v
}

type NullableOrderDeliveryEacType struct {
	value *OrderDeliveryEacType
	isSet bool
}

func (v NullableOrderDeliveryEacType) Get() *OrderDeliveryEacType {
	return v.value
}

func (v *NullableOrderDeliveryEacType) Set(val *OrderDeliveryEacType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDeliveryEacType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDeliveryEacType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDeliveryEacType(val *OrderDeliveryEacType) *NullableOrderDeliveryEacType {
	return &NullableOrderDeliveryEacType{value: val, isSet: true}
}

func (v NullableOrderDeliveryEacType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDeliveryEacType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

