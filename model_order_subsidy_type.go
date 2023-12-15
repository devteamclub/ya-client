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

// OrderSubsidyType Тип субсидии.
type OrderSubsidyType string

// List of OrderSubsidyType
const (
	NOT_SUBSIDY OrderSubsidyType = "NOT_SUBSIDY"
	SBER_SPASIBO OrderSubsidyType = "SBER_SPASIBO"
	YANDEX_CASHBACK OrderSubsidyType = "YANDEX_CASHBACK"
	SUBSIDY OrderSubsidyType = "SUBSIDY"
	DELIVERY OrderSubsidyType = "DELIVERY"
)

// All allowed values of OrderSubsidyType enum
var AllowedOrderSubsidyTypeEnumValues = []OrderSubsidyType{
	"NOT_SUBSIDY",
	"SBER_SPASIBO",
	"YANDEX_CASHBACK",
	"SUBSIDY",
	"DELIVERY",
}

func (v *OrderSubsidyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderSubsidyType(value)
	for _, existing := range AllowedOrderSubsidyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderSubsidyType", value)
}

// NewOrderSubsidyTypeFromValue returns a pointer to a valid OrderSubsidyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderSubsidyTypeFromValue(v string) (*OrderSubsidyType, error) {
	ev := OrderSubsidyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderSubsidyType: valid values are %v", v, AllowedOrderSubsidyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderSubsidyType) IsValid() bool {
	for _, existing := range AllowedOrderSubsidyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderSubsidyType value
func (v OrderSubsidyType) Ptr() *OrderSubsidyType {
	return &v
}

type NullableOrderSubsidyType struct {
	value *OrderSubsidyType
	isSet bool
}

func (v NullableOrderSubsidyType) Get() *OrderSubsidyType {
	return v.value
}

func (v *NullableOrderSubsidyType) Set(val *OrderSubsidyType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubsidyType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubsidyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubsidyType(val *OrderSubsidyType) *NullableOrderSubsidyType {
	return &NullableOrderSubsidyType{value: val, isSet: true}
}

func (v NullableOrderSubsidyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubsidyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

