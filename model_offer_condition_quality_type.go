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

// OfferConditionQualityType Внешний вид товара:  * `PERFECT` — идеальный. * `EXCELLENT` — отличный. * `GOOD` — хороший. 
type OfferConditionQualityType string

// List of OfferConditionQualityType
const (
	PERFECT OfferConditionQualityType = "PERFECT"
	EXCELLENT OfferConditionQualityType = "EXCELLENT"
	GOOD OfferConditionQualityType = "GOOD"
)

// All allowed values of OfferConditionQualityType enum
var AllowedOfferConditionQualityTypeEnumValues = []OfferConditionQualityType{
	"PERFECT",
	"EXCELLENT",
	"GOOD",
}

func (v *OfferConditionQualityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OfferConditionQualityType(value)
	for _, existing := range AllowedOfferConditionQualityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OfferConditionQualityType", value)
}

// NewOfferConditionQualityTypeFromValue returns a pointer to a valid OfferConditionQualityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOfferConditionQualityTypeFromValue(v string) (*OfferConditionQualityType, error) {
	ev := OfferConditionQualityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OfferConditionQualityType: valid values are %v", v, AllowedOfferConditionQualityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OfferConditionQualityType) IsValid() bool {
	for _, existing := range AllowedOfferConditionQualityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OfferConditionQualityType value
func (v OfferConditionQualityType) Ptr() *OfferConditionQualityType {
	return &v
}

type NullableOfferConditionQualityType struct {
	value *OfferConditionQualityType
	isSet bool
}

func (v NullableOfferConditionQualityType) Get() *OfferConditionQualityType {
	return v.value
}

func (v *NullableOfferConditionQualityType) Set(val *OfferConditionQualityType) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferConditionQualityType) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferConditionQualityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferConditionQualityType(val *OfferConditionQualityType) *NullableOfferConditionQualityType {
	return &NullableOfferConditionQualityType{value: val, isSet: true}
}

func (v NullableOfferConditionQualityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferConditionQualityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
