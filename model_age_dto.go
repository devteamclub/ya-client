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

// checks if the AgeDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgeDTO{}

// AgeDTO Возраст в заданных единицах измерения.
type AgeDTO struct {
	// Значение. 
	Value float32 `json:"value"`
	AgeUnit AgeUnitType `json:"ageUnit"`
}

type _AgeDTO AgeDTO

// NewAgeDTO instantiates a new AgeDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgeDTO(value float32, ageUnit AgeUnitType) *AgeDTO {
	this := AgeDTO{}
	this.Value = value
	this.AgeUnit = ageUnit
	return &this
}

// NewAgeDTOWithDefaults instantiates a new AgeDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgeDTOWithDefaults() *AgeDTO {
	this := AgeDTO{}
	return &this
}

// GetValue returns the Value field value
func (o *AgeDTO) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AgeDTO) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AgeDTO) SetValue(v float32) {
	o.Value = v
}

// GetAgeUnit returns the AgeUnit field value
func (o *AgeDTO) GetAgeUnit() AgeUnitType {
	if o == nil {
		var ret AgeUnitType
		return ret
	}

	return o.AgeUnit
}

// GetAgeUnitOk returns a tuple with the AgeUnit field value
// and a boolean to check if the value has been set.
func (o *AgeDTO) GetAgeUnitOk() (*AgeUnitType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgeUnit, true
}

// SetAgeUnit sets field value
func (o *AgeDTO) SetAgeUnit(v AgeUnitType) {
	o.AgeUnit = v
}

func (o AgeDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgeDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["ageUnit"] = o.AgeUnit
	return toSerialize, nil
}

func (o *AgeDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"ageUnit",
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

	varAgeDTO := _AgeDTO{}

	err = json.Unmarshal(bytes, &varAgeDTO)

	if err != nil {
		return err
	}

	*o = AgeDTO(varAgeDTO)

	return err
}

type NullableAgeDTO struct {
	value *AgeDTO
	isSet bool
}

func (v NullableAgeDTO) Get() *AgeDTO {
	return v.value
}

func (v *NullableAgeDTO) Set(val *AgeDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAgeDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAgeDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgeDTO(val *AgeDTO) *NullableAgeDTO {
	return &NullableAgeDTO{value: val, isSet: true}
}

func (v NullableAgeDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgeDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


