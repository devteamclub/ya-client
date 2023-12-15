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

// checks if the AcceptOrderCancellationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcceptOrderCancellationRequest{}

// AcceptOrderCancellationRequest struct for AcceptOrderCancellationRequest
type AcceptOrderCancellationRequest struct {
	// Решение об отмене заказа:  * `true` — заказ отменяется, служба доставки узнала об отмене до передачи заказа покупателю. * `false` — заказ не отменяется, так как он уже доставлен покупателю курьером или передан в пункт выдачи заказов. 
	Accepted bool `json:"accepted"`
	Reason *OrderCancellationReasonType `json:"reason,omitempty"`
}

type _AcceptOrderCancellationRequest AcceptOrderCancellationRequest

// NewAcceptOrderCancellationRequest instantiates a new AcceptOrderCancellationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptOrderCancellationRequest(accepted bool) *AcceptOrderCancellationRequest {
	this := AcceptOrderCancellationRequest{}
	this.Accepted = accepted
	return &this
}

// NewAcceptOrderCancellationRequestWithDefaults instantiates a new AcceptOrderCancellationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptOrderCancellationRequestWithDefaults() *AcceptOrderCancellationRequest {
	this := AcceptOrderCancellationRequest{}
	return &this
}

// GetAccepted returns the Accepted field value
func (o *AcceptOrderCancellationRequest) GetAccepted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Accepted
}

// GetAcceptedOk returns a tuple with the Accepted field value
// and a boolean to check if the value has been set.
func (o *AcceptOrderCancellationRequest) GetAcceptedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accepted, true
}

// SetAccepted sets field value
func (o *AcceptOrderCancellationRequest) SetAccepted(v bool) {
	o.Accepted = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AcceptOrderCancellationRequest) GetReason() OrderCancellationReasonType {
	if o == nil || IsNil(o.Reason) {
		var ret OrderCancellationReasonType
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptOrderCancellationRequest) GetReasonOk() (*OrderCancellationReasonType, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AcceptOrderCancellationRequest) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given OrderCancellationReasonType and assigns it to the Reason field.
func (o *AcceptOrderCancellationRequest) SetReason(v OrderCancellationReasonType) {
	o.Reason = &v
}

func (o AcceptOrderCancellationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptOrderCancellationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accepted"] = o.Accepted
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

func (o *AcceptOrderCancellationRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accepted",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAcceptOrderCancellationRequest := _AcceptOrderCancellationRequest{}

	err = json.Unmarshal(bytes, &varAcceptOrderCancellationRequest)

	if err != nil {
		return err
	}

	*o = AcceptOrderCancellationRequest(varAcceptOrderCancellationRequest)

	return err
}

type NullableAcceptOrderCancellationRequest struct {
	value *AcceptOrderCancellationRequest
	isSet bool
}

func (v NullableAcceptOrderCancellationRequest) Get() *AcceptOrderCancellationRequest {
	return v.value
}

func (v *NullableAcceptOrderCancellationRequest) Set(val *AcceptOrderCancellationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptOrderCancellationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptOrderCancellationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptOrderCancellationRequest(val *AcceptOrderCancellationRequest) *NullableAcceptOrderCancellationRequest {
	return &NullableAcceptOrderCancellationRequest{value: val, isSet: true}
}

func (v NullableAcceptOrderCancellationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptOrderCancellationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

