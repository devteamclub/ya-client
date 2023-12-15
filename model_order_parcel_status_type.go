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

// OrderParcelStatusType Статус заказа в партнерской службе доставки.
type OrderParcelStatusType string

// List of OrderParcelStatusType
const (
	NEW OrderParcelStatusType = "NEW"
	CREATED OrderParcelStatusType = "CREATED"
	READY_TO_SHIP OrderParcelStatusType = "READY_TO_SHIP"
	ERROR OrderParcelStatusType = "ERROR"
	UNKNOWN OrderParcelStatusType = "UNKNOWN"
)

// All allowed values of OrderParcelStatusType enum
var AllowedOrderParcelStatusTypeEnumValues = []OrderParcelStatusType{
	"NEW",
	"CREATED",
	"READY_TO_SHIP",
	"ERROR",
	"UNKNOWN",
}

func (v *OrderParcelStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderParcelStatusType(value)
	for _, existing := range AllowedOrderParcelStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderParcelStatusType", value)
}

// NewOrderParcelStatusTypeFromValue returns a pointer to a valid OrderParcelStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderParcelStatusTypeFromValue(v string) (*OrderParcelStatusType, error) {
	ev := OrderParcelStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderParcelStatusType: valid values are %v", v, AllowedOrderParcelStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderParcelStatusType) IsValid() bool {
	for _, existing := range AllowedOrderParcelStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderParcelStatusType value
func (v OrderParcelStatusType) Ptr() *OrderParcelStatusType {
	return &v
}

type NullableOrderParcelStatusType struct {
	value *OrderParcelStatusType
	isSet bool
}

func (v NullableOrderParcelStatusType) Get() *OrderParcelStatusType {
	return v.value
}

func (v *NullableOrderParcelStatusType) Set(val *OrderParcelStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderParcelStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderParcelStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderParcelStatusType(val *OrderParcelStatusType) *NullableOrderParcelStatusType {
	return &NullableOrderParcelStatusType{value: val, isSet: true}
}

func (v NullableOrderParcelStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderParcelStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
