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

// FeedStatusType Статус прайс-листа. Возможные значения:   * `ERROR` — найдены ошибки.   * `NA` — прайс-лист не загружался более семи дней или на этапе загрузки произошла ошибка.   * `OK` — ошибок не найдено. 
type FeedStatusType string

// List of FeedStatusType
const (
	ERROR FeedStatusType = "ERROR"
	NA FeedStatusType = "NA"
	OK FeedStatusType = "OK"
)

// All allowed values of FeedStatusType enum
var AllowedFeedStatusTypeEnumValues = []FeedStatusType{
	"ERROR",
	"NA",
	"OK",
}

func (v *FeedStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeedStatusType(value)
	for _, existing := range AllowedFeedStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeedStatusType", value)
}

// NewFeedStatusTypeFromValue returns a pointer to a valid FeedStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeedStatusTypeFromValue(v string) (*FeedStatusType, error) {
	ev := FeedStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeedStatusType: valid values are %v", v, AllowedFeedStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeedStatusType) IsValid() bool {
	for _, existing := range AllowedFeedStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeedStatusType value
func (v FeedStatusType) Ptr() *FeedStatusType {
	return &v
}

type NullableFeedStatusType struct {
	value *FeedStatusType
	isSet bool
}

func (v NullableFeedStatusType) Get() *FeedStatusType {
	return v.value
}

func (v *NullableFeedStatusType) Set(val *FeedStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedStatusType(val *FeedStatusType) *NullableFeedStatusType {
	return &NullableFeedStatusType{value: val, isSet: true}
}

func (v NullableFeedStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
