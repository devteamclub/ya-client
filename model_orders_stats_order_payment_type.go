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

// OrdersStatsOrderPaymentType Тип оплаты заказа: - `CREDIT` — заказ оформлен в кредит; - `POSTPAID` — заказ оплачен после того, как был получен; - `PREPAID` — заказ оплачен до того, как был получен. 
type OrdersStatsOrderPaymentType string

// List of OrdersStatsOrderPaymentType
const (
	CREDIT OrdersStatsOrderPaymentType = "CREDIT"
	POSTPAID OrdersStatsOrderPaymentType = "POSTPAID"
	PREPAID OrdersStatsOrderPaymentType = "PREPAID"
	TINKOFF_CREDIT OrdersStatsOrderPaymentType = "TINKOFF_CREDIT"
)

// All allowed values of OrdersStatsOrderPaymentType enum
var AllowedOrdersStatsOrderPaymentTypeEnumValues = []OrdersStatsOrderPaymentType{
	"CREDIT",
	"POSTPAID",
	"PREPAID",
	"TINKOFF_CREDIT",
}

func (v *OrdersStatsOrderPaymentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrdersStatsOrderPaymentType(value)
	for _, existing := range AllowedOrdersStatsOrderPaymentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrdersStatsOrderPaymentType", value)
}

// NewOrdersStatsOrderPaymentTypeFromValue returns a pointer to a valid OrdersStatsOrderPaymentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrdersStatsOrderPaymentTypeFromValue(v string) (*OrdersStatsOrderPaymentType, error) {
	ev := OrdersStatsOrderPaymentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrdersStatsOrderPaymentType: valid values are %v", v, AllowedOrdersStatsOrderPaymentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrdersStatsOrderPaymentType) IsValid() bool {
	for _, existing := range AllowedOrdersStatsOrderPaymentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrdersStatsOrderPaymentType value
func (v OrdersStatsOrderPaymentType) Ptr() *OrdersStatsOrderPaymentType {
	return &v
}

type NullableOrdersStatsOrderPaymentType struct {
	value *OrdersStatsOrderPaymentType
	isSet bool
}

func (v NullableOrdersStatsOrderPaymentType) Get() *OrdersStatsOrderPaymentType {
	return v.value
}

func (v *NullableOrdersStatsOrderPaymentType) Set(val *OrdersStatsOrderPaymentType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersStatsOrderPaymentType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersStatsOrderPaymentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersStatsOrderPaymentType(val *OrdersStatsOrderPaymentType) *NullableOrdersStatsOrderPaymentType {
	return &NullableOrdersStatsOrderPaymentType{value: val, isSet: true}
}

func (v NullableOrdersStatsOrderPaymentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersStatsOrderPaymentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
