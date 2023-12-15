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

// LicenseType Тип лицензии: * `ALCOHOL` — лицензия на розничную продажу алкогольной продукции. 
type LicenseType string

// List of LicenseType
const (
	ALCOHOL LicenseType = "ALCOHOL"
	UNKNOWN LicenseType = "UNKNOWN"
)

// All allowed values of LicenseType enum
var AllowedLicenseTypeEnumValues = []LicenseType{
	"ALCOHOL",
	"UNKNOWN",
}

func (v *LicenseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LicenseType(value)
	for _, existing := range AllowedLicenseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LicenseType", value)
}

// NewLicenseTypeFromValue returns a pointer to a valid LicenseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLicenseTypeFromValue(v string) (*LicenseType, error) {
	ev := LicenseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LicenseType: valid values are %v", v, AllowedLicenseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LicenseType) IsValid() bool {
	for _, existing := range AllowedLicenseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LicenseType value
func (v LicenseType) Ptr() *LicenseType {
	return &v
}

type NullableLicenseType struct {
	value *LicenseType
	isSet bool
}

func (v NullableLicenseType) Get() *LicenseType {
	return v.value
}

func (v *NullableLicenseType) Set(val *LicenseType) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseType) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseType(val *LicenseType) *NullableLicenseType {
	return &NullableLicenseType{value: val, isSet: true}
}

func (v NullableLicenseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

