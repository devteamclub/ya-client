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

// checks if the SetFeedParamsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetFeedParamsRequest{}

// SetFeedParamsRequest Запрос на обновление изменение параметров прайс-листа.
type SetFeedParamsRequest struct {
	// Параметры прайс-листа.  Обязательный параметр. 
	Parameters []FeedParameterDTO `json:"parameters"`
}

type _SetFeedParamsRequest SetFeedParamsRequest

// NewSetFeedParamsRequest instantiates a new SetFeedParamsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetFeedParamsRequest(parameters []FeedParameterDTO) *SetFeedParamsRequest {
	this := SetFeedParamsRequest{}
	this.Parameters = parameters
	return &this
}

// NewSetFeedParamsRequestWithDefaults instantiates a new SetFeedParamsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetFeedParamsRequestWithDefaults() *SetFeedParamsRequest {
	this := SetFeedParamsRequest{}
	return &this
}

// GetParameters returns the Parameters field value
func (o *SetFeedParamsRequest) GetParameters() []FeedParameterDTO {
	if o == nil {
		var ret []FeedParameterDTO
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *SetFeedParamsRequest) GetParametersOk() ([]FeedParameterDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parameters, true
}

// SetParameters sets field value
func (o *SetFeedParamsRequest) SetParameters(v []FeedParameterDTO) {
	o.Parameters = v
}

func (o SetFeedParamsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetFeedParamsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parameters"] = o.Parameters
	return toSerialize, nil
}

func (o *SetFeedParamsRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parameters",
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

	varSetFeedParamsRequest := _SetFeedParamsRequest{}

	err = json.Unmarshal(bytes, &varSetFeedParamsRequest)

	if err != nil {
		return err
	}

	*o = SetFeedParamsRequest(varSetFeedParamsRequest)

	return err
}

type NullableSetFeedParamsRequest struct {
	value *SetFeedParamsRequest
	isSet bool
}

func (v NullableSetFeedParamsRequest) Get() *SetFeedParamsRequest {
	return v.value
}

func (v *NullableSetFeedParamsRequest) Set(val *SetFeedParamsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetFeedParamsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetFeedParamsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetFeedParamsRequest(val *SetFeedParamsRequest) *NullableSetFeedParamsRequest {
	return &NullableSetFeedParamsRequest{value: val, isSet: true}
}

func (v NullableSetFeedParamsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetFeedParamsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


