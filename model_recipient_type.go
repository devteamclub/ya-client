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

// RecipientType Способ возврата товара покупателем:  * `SHOP` — в точку возврата магазина.  * `DELIVERY_SERVICE` — отправить курьером.  * `POST` — почта. 
type RecipientType string

// List of RecipientType
const (
	SHOP RecipientType = "SHOP"
	DELIVERY_SERVICE RecipientType = "DELIVERY_SERVICE"
	POST RecipientType = "POST"
)

// All allowed values of RecipientType enum
var AllowedRecipientTypeEnumValues = []RecipientType{
	"SHOP",
	"DELIVERY_SERVICE",
	"POST",
}

func (v *RecipientType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecipientType(value)
	for _, existing := range AllowedRecipientTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecipientType", value)
}

// NewRecipientTypeFromValue returns a pointer to a valid RecipientType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecipientTypeFromValue(v string) (*RecipientType, error) {
	ev := RecipientType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecipientType: valid values are %v", v, AllowedRecipientTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecipientType) IsValid() bool {
	for _, existing := range AllowedRecipientTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecipientType value
func (v RecipientType) Ptr() *RecipientType {
	return &v
}

type NullableRecipientType struct {
	value *RecipientType
	isSet bool
}

func (v NullableRecipientType) Get() *RecipientType {
	return v.value
}

func (v *NullableRecipientType) Set(val *RecipientType) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientType) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientType(val *RecipientType) *NullableRecipientType {
	return &NullableRecipientType{value: val, isSet: true}
}

func (v NullableRecipientType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
