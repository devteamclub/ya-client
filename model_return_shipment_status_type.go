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

// ReturnShipmentStatusType Статус передачи возврата:  * CREATED — возврат создан.  * RECEIVED — принят у покупателя.  * IN_TRANSIT — возврат в пути.  * READY_FOR_PICKUP — возврат готов к выдаче магазину.  * PICKED — возврат выдан магазину.  * LOST — возврат утерян при транспортировке.  * CANCELLED — возврат отменен.  * FULFILMENT_RECEIVED — возврат принят на складе Маркета.  * PREPARED_FOR_UTILIZATION — возврат передан в утилизацию.  * UTILIZED — возврат утилизирован. 
type ReturnShipmentStatusType string

// List of ReturnShipmentStatusType
const (
	CREATED ReturnShipmentStatusType = "CREATED"
	RECEIVED ReturnShipmentStatusType = "RECEIVED"
	IN_TRANSIT ReturnShipmentStatusType = "IN_TRANSIT"
	READY_FOR_PICKUP ReturnShipmentStatusType = "READY_FOR_PICKUP"
	PICKED ReturnShipmentStatusType = "PICKED"
	LOST ReturnShipmentStatusType = "LOST"
	EXPIRED ReturnShipmentStatusType = "EXPIRED"
	CANCELLED ReturnShipmentStatusType = "CANCELLED"
	FULFILMENT_RECEIVED ReturnShipmentStatusType = "FULFILMENT_RECEIVED"
	PREPARED_FOR_UTILIZATION ReturnShipmentStatusType = "PREPARED_FOR_UTILIZATION"
	NOT_IN_DEMAND ReturnShipmentStatusType = "NOT_IN_DEMAND"
	UTILIZED ReturnShipmentStatusType = "UTILIZED"
)

// All allowed values of ReturnShipmentStatusType enum
var AllowedReturnShipmentStatusTypeEnumValues = []ReturnShipmentStatusType{
	"CREATED",
	"RECEIVED",
	"IN_TRANSIT",
	"READY_FOR_PICKUP",
	"PICKED",
	"LOST",
	"EXPIRED",
	"CANCELLED",
	"FULFILMENT_RECEIVED",
	"PREPARED_FOR_UTILIZATION",
	"NOT_IN_DEMAND",
	"UTILIZED",
}

func (v *ReturnShipmentStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReturnShipmentStatusType(value)
	for _, existing := range AllowedReturnShipmentStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReturnShipmentStatusType", value)
}

// NewReturnShipmentStatusTypeFromValue returns a pointer to a valid ReturnShipmentStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReturnShipmentStatusTypeFromValue(v string) (*ReturnShipmentStatusType, error) {
	ev := ReturnShipmentStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReturnShipmentStatusType: valid values are %v", v, AllowedReturnShipmentStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReturnShipmentStatusType) IsValid() bool {
	for _, existing := range AllowedReturnShipmentStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReturnShipmentStatusType value
func (v ReturnShipmentStatusType) Ptr() *ReturnShipmentStatusType {
	return &v
}

type NullableReturnShipmentStatusType struct {
	value *ReturnShipmentStatusType
	isSet bool
}

func (v NullableReturnShipmentStatusType) Get() *ReturnShipmentStatusType {
	return v.value
}

func (v *NullableReturnShipmentStatusType) Set(val *ReturnShipmentStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnShipmentStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnShipmentStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnShipmentStatusType(val *ReturnShipmentStatusType) *NullableReturnShipmentStatusType {
	return &NullableReturnShipmentStatusType{value: val, isSet: true}
}

func (v NullableReturnShipmentStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnShipmentStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

