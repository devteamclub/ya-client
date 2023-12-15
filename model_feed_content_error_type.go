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

// FeedContentErrorType Тип ошибки в содержимом прайс-листа. Возможные значения: * `PARSE_ERROR` — ошибка при проверке прайс-листа, не связанная с форматом YML. Например, прайс-лист пустой или его не удалось разархивировать. * `PARSE_XML_ERROR` — несоответствие техническим требованиям формата YML. Например, элементы и их значения описаны некорректно. * `TOO_MANY_REJECTED_OFFERS` — более чем в половине предложений из прайс-листа найдены ошибки. Все предложения из прайс-листа не будут опубликованы на Маркете. 
type FeedContentErrorType string

// List of FeedContentErrorType
const (
	PARSE_ERROR FeedContentErrorType = "PARSE_ERROR"
	PARSE_XML_ERROR FeedContentErrorType = "PARSE_XML_ERROR"
	TOO_MANY_REJECTED_OFFERS FeedContentErrorType = "TOO_MANY_REJECTED_OFFERS"
)

// All allowed values of FeedContentErrorType enum
var AllowedFeedContentErrorTypeEnumValues = []FeedContentErrorType{
	"PARSE_ERROR",
	"PARSE_XML_ERROR",
	"TOO_MANY_REJECTED_OFFERS",
}

func (v *FeedContentErrorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeedContentErrorType(value)
	for _, existing := range AllowedFeedContentErrorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeedContentErrorType", value)
}

// NewFeedContentErrorTypeFromValue returns a pointer to a valid FeedContentErrorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeedContentErrorTypeFromValue(v string) (*FeedContentErrorType, error) {
	ev := FeedContentErrorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeedContentErrorType: valid values are %v", v, AllowedFeedContentErrorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeedContentErrorType) IsValid() bool {
	for _, existing := range AllowedFeedContentErrorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeedContentErrorType value
func (v FeedContentErrorType) Ptr() *FeedContentErrorType {
	return &v
}

type NullableFeedContentErrorType struct {
	value *FeedContentErrorType
	isSet bool
}

func (v NullableFeedContentErrorType) Get() *FeedContentErrorType {
	return v.value
}

func (v *NullableFeedContentErrorType) Set(val *FeedContentErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedContentErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedContentErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedContentErrorType(val *FeedContentErrorType) *NullableFeedContentErrorType {
	return &NullableFeedContentErrorType{value: val, isSet: true}
}

func (v NullableFeedContentErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedContentErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

