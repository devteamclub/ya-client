/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ya-client

import (
	"encoding/json"
	"fmt"
)

// OrderPaymentType Тип оплаты заказа:  * `PREPAID` — оплата при оформлении заказа.  * `POSTPAID` — оплата при получении заказа.  Если параметр отсутствует, заказ будет оплачен при получении. 
type OrderPaymentType string

// List of OrderPaymentType
const (
	PREPAID OrderPaymentType = "PREPAID"
	POSTPAID OrderPaymentType = "POSTPAID"
	UNKNOWN OrderPaymentType = "UNKNOWN"
)

// All allowed values of OrderPaymentType enum
var AllowedOrderPaymentTypeEnumValues = []OrderPaymentType{
	"PREPAID",
	"POSTPAID",
	"UNKNOWN",
}

func (v *OrderPaymentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderPaymentType(value)
	for _, existing := range AllowedOrderPaymentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderPaymentType", value)
}

// NewOrderPaymentTypeFromValue returns a pointer to a valid OrderPaymentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderPaymentTypeFromValue(v string) (*OrderPaymentType, error) {
	ev := OrderPaymentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderPaymentType: valid values are %v", v, AllowedOrderPaymentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderPaymentType) IsValid() bool {
	for _, existing := range AllowedOrderPaymentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderPaymentType value
func (v OrderPaymentType) Ptr() *OrderPaymentType {
	return &v
}

type NullableOrderPaymentType struct {
	value *OrderPaymentType
	isSet bool
}

func (v NullableOrderPaymentType) Get() *OrderPaymentType {
	return v.value
}

func (v *NullableOrderPaymentType) Set(val *OrderPaymentType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderPaymentType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderPaymentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderPaymentType(val *OrderPaymentType) *NullableOrderPaymentType {
	return &NullableOrderPaymentType{value: val, isSet: true}
}

func (v NullableOrderPaymentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderPaymentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
