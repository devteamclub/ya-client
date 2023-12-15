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

// FeedbackStateType Статус отзыва:  * `LAST` — актуален. * `PREVIOUS` — устарел. * `DELETED` — удален. 
type FeedbackStateType string

// List of FeedbackStateType
const (
	LAST FeedbackStateType = "LAST"
	PREVIOUS FeedbackStateType = "PREVIOUS"
	DELETED FeedbackStateType = "DELETED"
)

// All allowed values of FeedbackStateType enum
var AllowedFeedbackStateTypeEnumValues = []FeedbackStateType{
	"LAST",
	"PREVIOUS",
	"DELETED",
}

func (v *FeedbackStateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeedbackStateType(value)
	for _, existing := range AllowedFeedbackStateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeedbackStateType", value)
}

// NewFeedbackStateTypeFromValue returns a pointer to a valid FeedbackStateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeedbackStateTypeFromValue(v string) (*FeedbackStateType, error) {
	ev := FeedbackStateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeedbackStateType: valid values are %v", v, AllowedFeedbackStateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeedbackStateType) IsValid() bool {
	for _, existing := range AllowedFeedbackStateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeedbackStateType value
func (v FeedbackStateType) Ptr() *FeedbackStateType {
	return &v
}

type NullableFeedbackStateType struct {
	value *FeedbackStateType
	isSet bool
}

func (v NullableFeedbackStateType) Get() *FeedbackStateType {
	return v.value
}

func (v *NullableFeedbackStateType) Set(val *FeedbackStateType) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackStateType) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackStateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackStateType(val *FeedbackStateType) *NullableFeedbackStateType {
	return &NullableFeedbackStateType{value: val, isSet: true}
}

func (v NullableFeedbackStateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackStateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
