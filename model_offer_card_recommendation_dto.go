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

// checks if the OfferCardRecommendationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferCardRecommendationDTO{}

// OfferCardRecommendationDTO Рекомендация по заполнению карточки товара.
type OfferCardRecommendationDTO struct {
	Type OfferCardRecommendationType `json:"type"`
	// Процент выполнения рекомендации. Указывается для рекомендаций некоторых типов.
	Percent *int32 `json:"percent,omitempty"`
}

type _OfferCardRecommendationDTO OfferCardRecommendationDTO

// NewOfferCardRecommendationDTO instantiates a new OfferCardRecommendationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferCardRecommendationDTO(type_ OfferCardRecommendationType) *OfferCardRecommendationDTO {
	this := OfferCardRecommendationDTO{}
	this.Type = type_
	return &this
}

// NewOfferCardRecommendationDTOWithDefaults instantiates a new OfferCardRecommendationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferCardRecommendationDTOWithDefaults() *OfferCardRecommendationDTO {
	this := OfferCardRecommendationDTO{}
	return &this
}

// GetType returns the Type field value
func (o *OfferCardRecommendationDTO) GetType() OfferCardRecommendationType {
	if o == nil {
		var ret OfferCardRecommendationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OfferCardRecommendationDTO) GetTypeOk() (*OfferCardRecommendationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OfferCardRecommendationDTO) SetType(v OfferCardRecommendationType) {
	o.Type = v
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *OfferCardRecommendationDTO) GetPercent() int32 {
	if o == nil || IsNil(o.Percent) {
		var ret int32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferCardRecommendationDTO) GetPercentOk() (*int32, bool) {
	if o == nil || IsNil(o.Percent) {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *OfferCardRecommendationDTO) HasPercent() bool {
	if o != nil && !IsNil(o.Percent) {
		return true
	}

	return false
}

// SetPercent gets a reference to the given int32 and assigns it to the Percent field.
func (o *OfferCardRecommendationDTO) SetPercent(v int32) {
	o.Percent = &v
}

func (o OfferCardRecommendationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferCardRecommendationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Percent) {
		toSerialize["percent"] = o.Percent
	}
	return toSerialize, nil
}

func (o *OfferCardRecommendationDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varOfferCardRecommendationDTO := _OfferCardRecommendationDTO{}

	err = json.Unmarshal(bytes, &varOfferCardRecommendationDTO)

	if err != nil {
		return err
	}

	*o = OfferCardRecommendationDTO(varOfferCardRecommendationDTO)

	return err
}

type NullableOfferCardRecommendationDTO struct {
	value *OfferCardRecommendationDTO
	isSet bool
}

func (v NullableOfferCardRecommendationDTO) Get() *OfferCardRecommendationDTO {
	return v.value
}

func (v *NullableOfferCardRecommendationDTO) Set(val *OfferCardRecommendationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferCardRecommendationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferCardRecommendationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferCardRecommendationDTO(val *OfferCardRecommendationDTO) *NullableOfferCardRecommendationDTO {
	return &NullableOfferCardRecommendationDTO{value: val, isSet: true}
}

func (v NullableOfferCardRecommendationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferCardRecommendationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


