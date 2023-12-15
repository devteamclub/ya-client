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

// RefundStatusType Cтатус возврата средств:  * STARTED_BY_USER — создан клиентом из личного кабинета.  * REFUND_IN_PROGRESS — ждет решение о возврате средств.  * REFUNDED — по возврату проведены все возвратные денежные транзакции.  * FAILED — невозможно провести возврат покупателю.  * WAITING_FOR_DECISION — ожидает решения.  * DECISION_MADE — по возврату принято решение.  * REFUNDED_WITH_BONUSES — возврат осуществлен баллами Плюса или промокодом.  * REFUNDED_BY_SHOP — магазин сделал самостоятельно возврат средств.  * CANCELLED — возврат отменен. 
type RefundStatusType string

// List of RefundStatusType
const (
	STARTED_BY_USER RefundStatusType = "STARTED_BY_USER"
	REFUND_IN_PROGRESS RefundStatusType = "REFUND_IN_PROGRESS"
	REFUNDED RefundStatusType = "REFUNDED"
	FAILED RefundStatusType = "FAILED"
	WAITING_FOR_DECISION RefundStatusType = "WAITING_FOR_DECISION"
	DECISION_MADE RefundStatusType = "DECISION_MADE"
	REFUNDED_WITH_BONUSES RefundStatusType = "REFUNDED_WITH_BONUSES"
	REFUNDED_BY_SHOP RefundStatusType = "REFUNDED_BY_SHOP"
	CANCELLED RefundStatusType = "CANCELLED"
	UNKNOWN RefundStatusType = "UNKNOWN"
)

// All allowed values of RefundStatusType enum
var AllowedRefundStatusTypeEnumValues = []RefundStatusType{
	"STARTED_BY_USER",
	"REFUND_IN_PROGRESS",
	"REFUNDED",
	"FAILED",
	"WAITING_FOR_DECISION",
	"DECISION_MADE",
	"REFUNDED_WITH_BONUSES",
	"REFUNDED_BY_SHOP",
	"CANCELLED",
	"UNKNOWN",
}

func (v *RefundStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RefundStatusType(value)
	for _, existing := range AllowedRefundStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RefundStatusType", value)
}

// NewRefundStatusTypeFromValue returns a pointer to a valid RefundStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRefundStatusTypeFromValue(v string) (*RefundStatusType, error) {
	ev := RefundStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RefundStatusType: valid values are %v", v, AllowedRefundStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RefundStatusType) IsValid() bool {
	for _, existing := range AllowedRefundStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RefundStatusType value
func (v RefundStatusType) Ptr() *RefundStatusType {
	return &v
}

type NullableRefundStatusType struct {
	value *RefundStatusType
	isSet bool
}

func (v NullableRefundStatusType) Get() *RefundStatusType {
	return v.value
}

func (v *NullableRefundStatusType) Set(val *RefundStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundStatusType(val *RefundStatusType) *NullableRefundStatusType {
	return &NullableRefundStatusType{value: val, isSet: true}
}

func (v NullableRefundStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

