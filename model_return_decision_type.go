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

// ReturnDecisionType Решение по возврату:  * REFUND_MONEY — вернуть деньги покупателю.  * REFUND_MONEY_INCLUDING_SHIPMENT — вернуть деньги за товар и пересылку.  * REPAIR — отремонтировать товар.  * REPLACE — заменить товар.  * SEND_TO_EXAMINATION — взять товар на экспертизу.  * DECLINE_REFUND — отказать в возврате.  * OTHER_DECISION — другое решение. 
type ReturnDecisionType string

// List of ReturnDecisionType
const (
	REFUND_MONEY ReturnDecisionType = "REFUND_MONEY"
	REFUND_MONEY_INCLUDING_SHIPMENT ReturnDecisionType = "REFUND_MONEY_INCLUDING_SHIPMENT"
	REPAIR ReturnDecisionType = "REPAIR"
	REPLACE ReturnDecisionType = "REPLACE"
	SEND_TO_EXAMINATION ReturnDecisionType = "SEND_TO_EXAMINATION"
	DECLINE_REFUND ReturnDecisionType = "DECLINE_REFUND"
	OTHER_DECISION ReturnDecisionType = "OTHER_DECISION"
	UNKNOWN ReturnDecisionType = "UNKNOWN"
)

// All allowed values of ReturnDecisionType enum
var AllowedReturnDecisionTypeEnumValues = []ReturnDecisionType{
	"REFUND_MONEY",
	"REFUND_MONEY_INCLUDING_SHIPMENT",
	"REPAIR",
	"REPLACE",
	"SEND_TO_EXAMINATION",
	"DECLINE_REFUND",
	"OTHER_DECISION",
	"UNKNOWN",
}

func (v *ReturnDecisionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReturnDecisionType(value)
	for _, existing := range AllowedReturnDecisionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReturnDecisionType", value)
}

// NewReturnDecisionTypeFromValue returns a pointer to a valid ReturnDecisionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReturnDecisionTypeFromValue(v string) (*ReturnDecisionType, error) {
	ev := ReturnDecisionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReturnDecisionType: valid values are %v", v, AllowedReturnDecisionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReturnDecisionType) IsValid() bool {
	for _, existing := range AllowedReturnDecisionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReturnDecisionType value
func (v ReturnDecisionType) Ptr() *ReturnDecisionType {
	return &v
}

type NullableReturnDecisionType struct {
	value *ReturnDecisionType
	isSet bool
}

func (v NullableReturnDecisionType) Get() *ReturnDecisionType {
	return v.value
}

func (v *NullableReturnDecisionType) Set(val *ReturnDecisionType) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnDecisionType) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnDecisionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnDecisionType(val *ReturnDecisionType) *NullableReturnDecisionType {
	return &NullableReturnDecisionType{value: val, isSet: true}
}

func (v NullableReturnDecisionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnDecisionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

