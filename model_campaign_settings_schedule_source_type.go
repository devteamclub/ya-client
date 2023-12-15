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

// CampaignSettingsScheduleSourceType Источник информации о расписании работы службы доставки. Возможные значения: * `WEB` — информация получена из настроек личного кабинета магазина на Яндекс Маркете. * `YML` — информация получена из прайс-листа магазина. 
type CampaignSettingsScheduleSourceType string

// List of CampaignSettingsScheduleSourceType
const (
	WEB CampaignSettingsScheduleSourceType = "WEB"
	YML CampaignSettingsScheduleSourceType = "YML"
)

// All allowed values of CampaignSettingsScheduleSourceType enum
var AllowedCampaignSettingsScheduleSourceTypeEnumValues = []CampaignSettingsScheduleSourceType{
	"WEB",
	"YML",
}

func (v *CampaignSettingsScheduleSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CampaignSettingsScheduleSourceType(value)
	for _, existing := range AllowedCampaignSettingsScheduleSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CampaignSettingsScheduleSourceType", value)
}

// NewCampaignSettingsScheduleSourceTypeFromValue returns a pointer to a valid CampaignSettingsScheduleSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignSettingsScheduleSourceTypeFromValue(v string) (*CampaignSettingsScheduleSourceType, error) {
	ev := CampaignSettingsScheduleSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignSettingsScheduleSourceType: valid values are %v", v, AllowedCampaignSettingsScheduleSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignSettingsScheduleSourceType) IsValid() bool {
	for _, existing := range AllowedCampaignSettingsScheduleSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CampaignSettingsScheduleSourceType value
func (v CampaignSettingsScheduleSourceType) Ptr() *CampaignSettingsScheduleSourceType {
	return &v
}

type NullableCampaignSettingsScheduleSourceType struct {
	value *CampaignSettingsScheduleSourceType
	isSet bool
}

func (v NullableCampaignSettingsScheduleSourceType) Get() *CampaignSettingsScheduleSourceType {
	return v.value
}

func (v *NullableCampaignSettingsScheduleSourceType) Set(val *CampaignSettingsScheduleSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignSettingsScheduleSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignSettingsScheduleSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignSettingsScheduleSourceType(val *CampaignSettingsScheduleSourceType) *NullableCampaignSettingsScheduleSourceType {
	return &NullableCampaignSettingsScheduleSourceType{value: val, isSet: true}
}

func (v NullableCampaignSettingsScheduleSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignSettingsScheduleSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
