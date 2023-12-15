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

// ReportStatusType Статус генерации отчета:  * `PENDING` — отчет ожидает генерации; * `PROCESSING` — отчет генерируется; * `FAILED` — во время генерации произошла ошибка; * `DONE` — отчет готов. 
type ReportStatusType string

// List of ReportStatusType
const (
	PENDING ReportStatusType = "PENDING"
	PROCESSING ReportStatusType = "PROCESSING"
	FAILED ReportStatusType = "FAILED"
	DONE ReportStatusType = "DONE"
)

// All allowed values of ReportStatusType enum
var AllowedReportStatusTypeEnumValues = []ReportStatusType{
	"PENDING",
	"PROCESSING",
	"FAILED",
	"DONE",
}

func (v *ReportStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportStatusType(value)
	for _, existing := range AllowedReportStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportStatusType", value)
}

// NewReportStatusTypeFromValue returns a pointer to a valid ReportStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportStatusTypeFromValue(v string) (*ReportStatusType, error) {
	ev := ReportStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportStatusType: valid values are %v", v, AllowedReportStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportStatusType) IsValid() bool {
	for _, existing := range AllowedReportStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportStatusType value
func (v ReportStatusType) Ptr() *ReportStatusType {
	return &v
}

type NullableReportStatusType struct {
	value *ReportStatusType
	isSet bool
}

func (v NullableReportStatusType) Get() *ReportStatusType {
	return v.value
}

func (v *NullableReportStatusType) Set(val *ReportStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableReportStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableReportStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportStatusType(val *ReportStatusType) *NullableReportStatusType {
	return &NullableReportStatusType{value: val, isSet: true}
}

func (v NullableReportStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

