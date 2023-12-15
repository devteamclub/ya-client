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

// PriceCompetitivenessType Привлекательность цены:  * `OPTIMAL` — привлекательная; * `AVERAGE` — умеренная; * `LOW` — непривлекательная. 
type PriceCompetitivenessType string

// List of PriceCompetitivenessType
const (
	OPTIMAL PriceCompetitivenessType = "OPTIMAL"
	AVERAGE PriceCompetitivenessType = "AVERAGE"
	LOW PriceCompetitivenessType = "LOW"
)

// All allowed values of PriceCompetitivenessType enum
var AllowedPriceCompetitivenessTypeEnumValues = []PriceCompetitivenessType{
	"OPTIMAL",
	"AVERAGE",
	"LOW",
}

func (v *PriceCompetitivenessType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PriceCompetitivenessType(value)
	for _, existing := range AllowedPriceCompetitivenessTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PriceCompetitivenessType", value)
}

// NewPriceCompetitivenessTypeFromValue returns a pointer to a valid PriceCompetitivenessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPriceCompetitivenessTypeFromValue(v string) (*PriceCompetitivenessType, error) {
	ev := PriceCompetitivenessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PriceCompetitivenessType: valid values are %v", v, AllowedPriceCompetitivenessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PriceCompetitivenessType) IsValid() bool {
	for _, existing := range AllowedPriceCompetitivenessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PriceCompetitivenessType value
func (v PriceCompetitivenessType) Ptr() *PriceCompetitivenessType {
	return &v
}

type NullablePriceCompetitivenessType struct {
	value *PriceCompetitivenessType
	isSet bool
}

func (v NullablePriceCompetitivenessType) Get() *PriceCompetitivenessType {
	return v.value
}

func (v *NullablePriceCompetitivenessType) Set(val *PriceCompetitivenessType) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceCompetitivenessType) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceCompetitivenessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceCompetitivenessType(val *PriceCompetitivenessType) *NullablePriceCompetitivenessType {
	return &NullablePriceCompetitivenessType{value: val, isSet: true}
}

func (v NullablePriceCompetitivenessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceCompetitivenessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
