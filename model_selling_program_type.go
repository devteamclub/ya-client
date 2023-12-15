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

// SellingProgramType Модель размещения:  * `FBY` — FBY. * `FBS` — FBS. * `DBS` — DBS. * `EXPRESS` — Экспресс. 
type SellingProgramType string

// List of SellingProgramType
const (
	FBY SellingProgramType = "FBY"
	FBS SellingProgramType = "FBS"
	DBS SellingProgramType = "DBS"
	EXPRESS SellingProgramType = "EXPRESS"
)

// All allowed values of SellingProgramType enum
var AllowedSellingProgramTypeEnumValues = []SellingProgramType{
	"FBY",
	"FBS",
	"DBS",
	"EXPRESS",
}

func (v *SellingProgramType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SellingProgramType(value)
	for _, existing := range AllowedSellingProgramTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SellingProgramType", value)
}

// NewSellingProgramTypeFromValue returns a pointer to a valid SellingProgramType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSellingProgramTypeFromValue(v string) (*SellingProgramType, error) {
	ev := SellingProgramType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SellingProgramType: valid values are %v", v, AllowedSellingProgramTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SellingProgramType) IsValid() bool {
	for _, existing := range AllowedSellingProgramTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SellingProgramType value
func (v SellingProgramType) Ptr() *SellingProgramType {
	return &v
}

type NullableSellingProgramType struct {
	value *SellingProgramType
	isSet bool
}

func (v NullableSellingProgramType) Get() *SellingProgramType {
	return v.value
}

func (v *NullableSellingProgramType) Set(val *SellingProgramType) {
	v.value = val
	v.isSet = true
}

func (v NullableSellingProgramType) IsSet() bool {
	return v.isSet
}

func (v *NullableSellingProgramType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSellingProgramType(val *SellingProgramType) *NullableSellingProgramType {
	return &NullableSellingProgramType{value: val, isSet: true}
}

func (v NullableSellingProgramType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSellingProgramType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

