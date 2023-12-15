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

// ShipmentStatusType Статус отгрузки:  * `OUTBOUND_CREATED` — формируется. * `OUTBOUND_READY_FOR_CONFIRMATION` — можно обрабатывать. * `OUTBOUND_CONFIRMED` — подтверждена и готова к отправке. * `OUTBOUND_SIGNED` — по ней подписан электронный акт приема-передачи. * `ACCEPTED` — принята в сортировочном центре или пункте приема. * `ACCEPTED_WITH_DISCREPANCIES` — принята с расхождениями. * `FINISHED` — завершена. * `ERROR` — отменена из-за ошибки. 
type ShipmentStatusType string

// List of ShipmentStatusType
const (
	OUTBOUND_CREATED ShipmentStatusType = "OUTBOUND_CREATED"
	OUTBOUND_READY_FOR_CONFIRMATION ShipmentStatusType = "OUTBOUND_READY_FOR_CONFIRMATION"
	OUTBOUND_CONFIRMED ShipmentStatusType = "OUTBOUND_CONFIRMED"
	OUTBOUND_SIGNED ShipmentStatusType = "OUTBOUND_SIGNED"
	FINISHED ShipmentStatusType = "FINISHED"
	ACCEPTED ShipmentStatusType = "ACCEPTED"
	ACCEPTED_WITH_DISCREPANCIES ShipmentStatusType = "ACCEPTED_WITH_DISCREPANCIES"
	ERROR ShipmentStatusType = "ERROR"
)

// All allowed values of ShipmentStatusType enum
var AllowedShipmentStatusTypeEnumValues = []ShipmentStatusType{
	"OUTBOUND_CREATED",
	"OUTBOUND_READY_FOR_CONFIRMATION",
	"OUTBOUND_CONFIRMED",
	"OUTBOUND_SIGNED",
	"FINISHED",
	"ACCEPTED",
	"ACCEPTED_WITH_DISCREPANCIES",
	"ERROR",
}

func (v *ShipmentStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShipmentStatusType(value)
	for _, existing := range AllowedShipmentStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShipmentStatusType", value)
}

// NewShipmentStatusTypeFromValue returns a pointer to a valid ShipmentStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShipmentStatusTypeFromValue(v string) (*ShipmentStatusType, error) {
	ev := ShipmentStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShipmentStatusType: valid values are %v", v, AllowedShipmentStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShipmentStatusType) IsValid() bool {
	for _, existing := range AllowedShipmentStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShipmentStatusType value
func (v ShipmentStatusType) Ptr() *ShipmentStatusType {
	return &v
}

type NullableShipmentStatusType struct {
	value *ShipmentStatusType
	isSet bool
}

func (v NullableShipmentStatusType) Get() *ShipmentStatusType {
	return v.value
}

func (v *NullableShipmentStatusType) Set(val *ShipmentStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentStatusType(val *ShipmentStatusType) *NullableShipmentStatusType {
	return &NullableShipmentStatusType{value: val, isSet: true}
}

func (v NullableShipmentStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

