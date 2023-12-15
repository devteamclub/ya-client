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

// OfferMappingKindType Вид маппинга.
type OfferMappingKindType string

// List of OfferMappingKindType
const (
	ACTIVE OfferMappingKindType = "ACTIVE"
	ALL OfferMappingKindType = "ALL"
)

// All allowed values of OfferMappingKindType enum
var AllowedOfferMappingKindTypeEnumValues = []OfferMappingKindType{
	"ACTIVE",
	"ALL",
}

func (v *OfferMappingKindType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OfferMappingKindType(value)
	for _, existing := range AllowedOfferMappingKindTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OfferMappingKindType", value)
}

// NewOfferMappingKindTypeFromValue returns a pointer to a valid OfferMappingKindType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOfferMappingKindTypeFromValue(v string) (*OfferMappingKindType, error) {
	ev := OfferMappingKindType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OfferMappingKindType: valid values are %v", v, AllowedOfferMappingKindTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OfferMappingKindType) IsValid() bool {
	for _, existing := range AllowedOfferMappingKindTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OfferMappingKindType value
func (v OfferMappingKindType) Ptr() *OfferMappingKindType {
	return &v
}

type NullableOfferMappingKindType struct {
	value *OfferMappingKindType
	isSet bool
}

func (v NullableOfferMappingKindType) Get() *OfferMappingKindType {
	return v.value
}

func (v *NullableOfferMappingKindType) Set(val *OfferMappingKindType) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferMappingKindType) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferMappingKindType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferMappingKindType(val *OfferMappingKindType) *NullableOfferMappingKindType {
	return &NullableOfferMappingKindType{value: val, isSet: true}
}

func (v NullableOfferMappingKindType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferMappingKindType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

