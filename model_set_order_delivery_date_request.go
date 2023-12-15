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

// checks if the SetOrderDeliveryDateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetOrderDeliveryDateRequest{}

// SetOrderDeliveryDateRequest struct for SetOrderDeliveryDateRequest
type SetOrderDeliveryDateRequest struct {
	Dates OrderDeliveryDateDTO `json:"dates"`
	Reason OrderDeliveryDateReasonType `json:"reason"`
}

type _SetOrderDeliveryDateRequest SetOrderDeliveryDateRequest

// NewSetOrderDeliveryDateRequest instantiates a new SetOrderDeliveryDateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetOrderDeliveryDateRequest(dates OrderDeliveryDateDTO, reason OrderDeliveryDateReasonType) *SetOrderDeliveryDateRequest {
	this := SetOrderDeliveryDateRequest{}
	this.Dates = dates
	this.Reason = reason
	return &this
}

// NewSetOrderDeliveryDateRequestWithDefaults instantiates a new SetOrderDeliveryDateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetOrderDeliveryDateRequestWithDefaults() *SetOrderDeliveryDateRequest {
	this := SetOrderDeliveryDateRequest{}
	return &this
}

// GetDates returns the Dates field value
func (o *SetOrderDeliveryDateRequest) GetDates() OrderDeliveryDateDTO {
	if o == nil {
		var ret OrderDeliveryDateDTO
		return ret
	}

	return o.Dates
}

// GetDatesOk returns a tuple with the Dates field value
// and a boolean to check if the value has been set.
func (o *SetOrderDeliveryDateRequest) GetDatesOk() (*OrderDeliveryDateDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dates, true
}

// SetDates sets field value
func (o *SetOrderDeliveryDateRequest) SetDates(v OrderDeliveryDateDTO) {
	o.Dates = v
}

// GetReason returns the Reason field value
func (o *SetOrderDeliveryDateRequest) GetReason() OrderDeliveryDateReasonType {
	if o == nil {
		var ret OrderDeliveryDateReasonType
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SetOrderDeliveryDateRequest) GetReasonOk() (*OrderDeliveryDateReasonType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SetOrderDeliveryDateRequest) SetReason(v OrderDeliveryDateReasonType) {
	o.Reason = v
}

func (o SetOrderDeliveryDateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetOrderDeliveryDateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dates"] = o.Dates
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

func (o *SetOrderDeliveryDateRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dates",
		"reason",
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

	varSetOrderDeliveryDateRequest := _SetOrderDeliveryDateRequest{}

	err = json.Unmarshal(bytes, &varSetOrderDeliveryDateRequest)

	if err != nil {
		return err
	}

	*o = SetOrderDeliveryDateRequest(varSetOrderDeliveryDateRequest)

	return err
}

type NullableSetOrderDeliveryDateRequest struct {
	value *SetOrderDeliveryDateRequest
	isSet bool
}

func (v NullableSetOrderDeliveryDateRequest) Get() *SetOrderDeliveryDateRequest {
	return v.value
}

func (v *NullableSetOrderDeliveryDateRequest) Set(val *SetOrderDeliveryDateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetOrderDeliveryDateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetOrderDeliveryDateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetOrderDeliveryDateRequest(val *SetOrderDeliveryDateRequest) *NullableSetOrderDeliveryDateRequest {
	return &NullableSetOrderDeliveryDateRequest{value: val, isSet: true}
}

func (v NullableSetOrderDeliveryDateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetOrderDeliveryDateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

