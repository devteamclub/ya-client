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

// LogisticPointType Тип логистической точки:   * WAREHOUSE — склад;   * PICKUP_POINT — обычная точка выдачи заказов (ПВЗ);   * PICKUP_TERMINAL — постамат;   * PICKUP_POST_OFFICE — отделение почтовой связи (ОПС);   * PICKUP_MIXED — торговый зал и пункт выдачи заказов;   * PICKUP_RETAIL — торговый зал. 
type LogisticPointType string

// List of LogisticPointType
const (
	WAREHOUSE LogisticPointType = "WAREHOUSE"
	PICKUP_POINT LogisticPointType = "PICKUP_POINT"
	PICKUP_TERMINAL LogisticPointType = "PICKUP_TERMINAL"
	PICKUP_POST_OFFICE LogisticPointType = "PICKUP_POST_OFFICE"
	PICKUP_MIXED LogisticPointType = "PICKUP_MIXED"
	PICKUP_RETAIL LogisticPointType = "PICKUP_RETAIL"
	UNKNOWN LogisticPointType = "UNKNOWN"
)

// All allowed values of LogisticPointType enum
var AllowedLogisticPointTypeEnumValues = []LogisticPointType{
	"WAREHOUSE",
	"PICKUP_POINT",
	"PICKUP_TERMINAL",
	"PICKUP_POST_OFFICE",
	"PICKUP_MIXED",
	"PICKUP_RETAIL",
	"UNKNOWN",
}

func (v *LogisticPointType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogisticPointType(value)
	for _, existing := range AllowedLogisticPointTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogisticPointType", value)
}

// NewLogisticPointTypeFromValue returns a pointer to a valid LogisticPointType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogisticPointTypeFromValue(v string) (*LogisticPointType, error) {
	ev := LogisticPointType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogisticPointType: valid values are %v", v, AllowedLogisticPointTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogisticPointType) IsValid() bool {
	for _, existing := range AllowedLogisticPointTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogisticPointType value
func (v LogisticPointType) Ptr() *LogisticPointType {
	return &v
}

type NullableLogisticPointType struct {
	value *LogisticPointType
	isSet bool
}

func (v NullableLogisticPointType) Get() *LogisticPointType {
	return v.value
}

func (v *NullableLogisticPointType) Set(val *LogisticPointType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogisticPointType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogisticPointType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogisticPointType(val *LogisticPointType) *NullableLogisticPointType {
	return &NullableLogisticPointType{value: val, isSet: true}
}

func (v NullableLogisticPointType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogisticPointType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
