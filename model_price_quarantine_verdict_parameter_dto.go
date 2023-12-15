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

// checks if the PriceQuarantineVerdictParameterDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceQuarantineVerdictParameterDTO{}

// PriceQuarantineVerdictParameterDTO Параметр карантина.
type PriceQuarantineVerdictParameterDTO struct {
	Name PriceQuarantineVerdictParamNameType `json:"name"`
	// Значение параметра.
	Value string `json:"value"`
}

type _PriceQuarantineVerdictParameterDTO PriceQuarantineVerdictParameterDTO

// NewPriceQuarantineVerdictParameterDTO instantiates a new PriceQuarantineVerdictParameterDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceQuarantineVerdictParameterDTO(name PriceQuarantineVerdictParamNameType, value string) *PriceQuarantineVerdictParameterDTO {
	this := PriceQuarantineVerdictParameterDTO{}
	this.Name = name
	this.Value = value
	return &this
}

// NewPriceQuarantineVerdictParameterDTOWithDefaults instantiates a new PriceQuarantineVerdictParameterDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceQuarantineVerdictParameterDTOWithDefaults() *PriceQuarantineVerdictParameterDTO {
	this := PriceQuarantineVerdictParameterDTO{}
	return &this
}

// GetName returns the Name field value
func (o *PriceQuarantineVerdictParameterDTO) GetName() PriceQuarantineVerdictParamNameType {
	if o == nil {
		var ret PriceQuarantineVerdictParamNameType
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PriceQuarantineVerdictParameterDTO) GetNameOk() (*PriceQuarantineVerdictParamNameType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PriceQuarantineVerdictParameterDTO) SetName(v PriceQuarantineVerdictParamNameType) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *PriceQuarantineVerdictParameterDTO) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PriceQuarantineVerdictParameterDTO) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PriceQuarantineVerdictParameterDTO) SetValue(v string) {
	o.Value = v
}

func (o PriceQuarantineVerdictParameterDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceQuarantineVerdictParameterDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *PriceQuarantineVerdictParameterDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varPriceQuarantineVerdictParameterDTO := _PriceQuarantineVerdictParameterDTO{}

	err = json.Unmarshal(bytes, &varPriceQuarantineVerdictParameterDTO)

	if err != nil {
		return err
	}

	*o = PriceQuarantineVerdictParameterDTO(varPriceQuarantineVerdictParameterDTO)

	return err
}

type NullablePriceQuarantineVerdictParameterDTO struct {
	value *PriceQuarantineVerdictParameterDTO
	isSet bool
}

func (v NullablePriceQuarantineVerdictParameterDTO) Get() *PriceQuarantineVerdictParameterDTO {
	return v.value
}

func (v *NullablePriceQuarantineVerdictParameterDTO) Set(val *PriceQuarantineVerdictParameterDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceQuarantineVerdictParameterDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceQuarantineVerdictParameterDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceQuarantineVerdictParameterDTO(val *PriceQuarantineVerdictParameterDTO) *NullablePriceQuarantineVerdictParameterDTO {
	return &NullablePriceQuarantineVerdictParameterDTO{value: val, isSet: true}
}

func (v NullablePriceQuarantineVerdictParameterDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceQuarantineVerdictParameterDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


