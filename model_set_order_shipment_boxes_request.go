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

// checks if the SetOrderShipmentBoxesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetOrderShipmentBoxesRequest{}

// SetOrderShipmentBoxesRequest struct for SetOrderShipmentBoxesRequest
type SetOrderShipmentBoxesRequest struct {
	// Список грузовых мест. Маркет определяет количество мест по длине этого списка.
	Boxes []ParcelBoxDTO `json:"boxes"`
}

type _SetOrderShipmentBoxesRequest SetOrderShipmentBoxesRequest

// NewSetOrderShipmentBoxesRequest instantiates a new SetOrderShipmentBoxesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetOrderShipmentBoxesRequest(boxes []ParcelBoxDTO) *SetOrderShipmentBoxesRequest {
	this := SetOrderShipmentBoxesRequest{}
	this.Boxes = boxes
	return &this
}

// NewSetOrderShipmentBoxesRequestWithDefaults instantiates a new SetOrderShipmentBoxesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetOrderShipmentBoxesRequestWithDefaults() *SetOrderShipmentBoxesRequest {
	this := SetOrderShipmentBoxesRequest{}
	return &this
}

// GetBoxes returns the Boxes field value
func (o *SetOrderShipmentBoxesRequest) GetBoxes() []ParcelBoxDTO {
	if o == nil {
		var ret []ParcelBoxDTO
		return ret
	}

	return o.Boxes
}

// GetBoxesOk returns a tuple with the Boxes field value
// and a boolean to check if the value has been set.
func (o *SetOrderShipmentBoxesRequest) GetBoxesOk() ([]ParcelBoxDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Boxes, true
}

// SetBoxes sets field value
func (o *SetOrderShipmentBoxesRequest) SetBoxes(v []ParcelBoxDTO) {
	o.Boxes = v
}

func (o SetOrderShipmentBoxesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetOrderShipmentBoxesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boxes"] = o.Boxes
	return toSerialize, nil
}

func (o *SetOrderShipmentBoxesRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"boxes",
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

	varSetOrderShipmentBoxesRequest := _SetOrderShipmentBoxesRequest{}

	err = json.Unmarshal(bytes, &varSetOrderShipmentBoxesRequest)

	if err != nil {
		return err
	}

	*o = SetOrderShipmentBoxesRequest(varSetOrderShipmentBoxesRequest)

	return err
}

type NullableSetOrderShipmentBoxesRequest struct {
	value *SetOrderShipmentBoxesRequest
	isSet bool
}

func (v NullableSetOrderShipmentBoxesRequest) Get() *SetOrderShipmentBoxesRequest {
	return v.value
}

func (v *NullableSetOrderShipmentBoxesRequest) Set(val *SetOrderShipmentBoxesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetOrderShipmentBoxesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetOrderShipmentBoxesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetOrderShipmentBoxesRequest(val *SetOrderShipmentBoxesRequest) *NullableSetOrderShipmentBoxesRequest {
	return &NullableSetOrderShipmentBoxesRequest{value: val, isSet: true}
}

func (v NullableSetOrderShipmentBoxesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetOrderShipmentBoxesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


