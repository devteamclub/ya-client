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

// FeedIndexLogsStatusType Статус индексации прайс-листа и проверки на соответствие техническим требованиям.  Возможные значения: * `ERROR` — произошли ошибки. * `OK` — обработан без ошибок. * `WARNING` — наблюдались некритичные проблемы. 
type FeedIndexLogsStatusType string

// List of FeedIndexLogsStatusType
const (
	ERROR FeedIndexLogsStatusType = "ERROR"
	OK FeedIndexLogsStatusType = "OK"
	WARNING FeedIndexLogsStatusType = "WARNING"
)

// All allowed values of FeedIndexLogsStatusType enum
var AllowedFeedIndexLogsStatusTypeEnumValues = []FeedIndexLogsStatusType{
	"ERROR",
	"OK",
	"WARNING",
}

func (v *FeedIndexLogsStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeedIndexLogsStatusType(value)
	for _, existing := range AllowedFeedIndexLogsStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeedIndexLogsStatusType", value)
}

// NewFeedIndexLogsStatusTypeFromValue returns a pointer to a valid FeedIndexLogsStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeedIndexLogsStatusTypeFromValue(v string) (*FeedIndexLogsStatusType, error) {
	ev := FeedIndexLogsStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeedIndexLogsStatusType: valid values are %v", v, AllowedFeedIndexLogsStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeedIndexLogsStatusType) IsValid() bool {
	for _, existing := range AllowedFeedIndexLogsStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeedIndexLogsStatusType value
func (v FeedIndexLogsStatusType) Ptr() *FeedIndexLogsStatusType {
	return &v
}

type NullableFeedIndexLogsStatusType struct {
	value *FeedIndexLogsStatusType
	isSet bool
}

func (v NullableFeedIndexLogsStatusType) Get() *FeedIndexLogsStatusType {
	return v.value
}

func (v *NullableFeedIndexLogsStatusType) Set(val *FeedIndexLogsStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedIndexLogsStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedIndexLogsStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedIndexLogsStatusType(val *FeedIndexLogsStatusType) *NullableFeedIndexLogsStatusType {
	return &NullableFeedIndexLogsStatusType{value: val, isSet: true}
}

func (v NullableFeedIndexLogsStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedIndexLogsStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

