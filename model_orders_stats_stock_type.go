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

// OrdersStatsStockType Тип товара:  * `DEFECT` — товар бракованный.  * `FIT` — товар надлежащего качества. 
type OrdersStatsStockType string

// List of OrdersStatsStockType
const (
	FIT OrdersStatsStockType = "FIT"
	FREEZE OrdersStatsStockType = "FREEZE"
	AVAILABLE OrdersStatsStockType = "AVAILABLE"
	QUARANTINE OrdersStatsStockType = "QUARANTINE"
	UTILIZATION OrdersStatsStockType = "UTILIZATION"
	DEFECT OrdersStatsStockType = "DEFECT"
	EXPIRED OrdersStatsStockType = "EXPIRED"
)

// All allowed values of OrdersStatsStockType enum
var AllowedOrdersStatsStockTypeEnumValues = []OrdersStatsStockType{
	"FIT",
	"FREEZE",
	"AVAILABLE",
	"QUARANTINE",
	"UTILIZATION",
	"DEFECT",
	"EXPIRED",
}

func (v *OrdersStatsStockType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrdersStatsStockType(value)
	for _, existing := range AllowedOrdersStatsStockTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrdersStatsStockType", value)
}

// NewOrdersStatsStockTypeFromValue returns a pointer to a valid OrdersStatsStockType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrdersStatsStockTypeFromValue(v string) (*OrdersStatsStockType, error) {
	ev := OrdersStatsStockType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrdersStatsStockType: valid values are %v", v, AllowedOrdersStatsStockTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrdersStatsStockType) IsValid() bool {
	for _, existing := range AllowedOrdersStatsStockTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrdersStatsStockType value
func (v OrdersStatsStockType) Ptr() *OrdersStatsStockType {
	return &v
}

type NullableOrdersStatsStockType struct {
	value *OrdersStatsStockType
	isSet bool
}

func (v NullableOrdersStatsStockType) Get() *OrdersStatsStockType {
	return v.value
}

func (v *NullableOrdersStatsStockType) Set(val *OrdersStatsStockType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersStatsStockType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersStatsStockType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersStatsStockType(val *OrdersStatsStockType) *NullableOrdersStatsStockType {
	return &NullableOrdersStatsStockType{value: val, isSet: true}
}

func (v NullableOrdersStatsStockType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersStatsStockType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

