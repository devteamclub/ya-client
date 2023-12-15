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

// checks if the GetBidsRecommendationsResponseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBidsRecommendationsResponseDTO{}

// GetBidsRecommendationsResponseDTO Список товаров с рекомендованными ставками.
type GetBidsRecommendationsResponseDTO struct {
	// Список товаров с рекомендованными ставками.
	Recommendations []SkuBidRecommendationItemDTO `json:"recommendations"`
}

type _GetBidsRecommendationsResponseDTO GetBidsRecommendationsResponseDTO

// NewGetBidsRecommendationsResponseDTO instantiates a new GetBidsRecommendationsResponseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBidsRecommendationsResponseDTO(recommendations []SkuBidRecommendationItemDTO) *GetBidsRecommendationsResponseDTO {
	this := GetBidsRecommendationsResponseDTO{}
	this.Recommendations = recommendations
	return &this
}

// NewGetBidsRecommendationsResponseDTOWithDefaults instantiates a new GetBidsRecommendationsResponseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBidsRecommendationsResponseDTOWithDefaults() *GetBidsRecommendationsResponseDTO {
	this := GetBidsRecommendationsResponseDTO{}
	return &this
}

// GetRecommendations returns the Recommendations field value
func (o *GetBidsRecommendationsResponseDTO) GetRecommendations() []SkuBidRecommendationItemDTO {
	if o == nil {
		var ret []SkuBidRecommendationItemDTO
		return ret
	}

	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value
// and a boolean to check if the value has been set.
func (o *GetBidsRecommendationsResponseDTO) GetRecommendationsOk() ([]SkuBidRecommendationItemDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recommendations, true
}

// SetRecommendations sets field value
func (o *GetBidsRecommendationsResponseDTO) SetRecommendations(v []SkuBidRecommendationItemDTO) {
	o.Recommendations = v
}

func (o GetBidsRecommendationsResponseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBidsRecommendationsResponseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recommendations"] = o.Recommendations
	return toSerialize, nil
}

func (o *GetBidsRecommendationsResponseDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recommendations",
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

	varGetBidsRecommendationsResponseDTO := _GetBidsRecommendationsResponseDTO{}

	err = json.Unmarshal(bytes, &varGetBidsRecommendationsResponseDTO)

	if err != nil {
		return err
	}

	*o = GetBidsRecommendationsResponseDTO(varGetBidsRecommendationsResponseDTO)

	return err
}

type NullableGetBidsRecommendationsResponseDTO struct {
	value *GetBidsRecommendationsResponseDTO
	isSet bool
}

func (v NullableGetBidsRecommendationsResponseDTO) Get() *GetBidsRecommendationsResponseDTO {
	return v.value
}

func (v *NullableGetBidsRecommendationsResponseDTO) Set(val *GetBidsRecommendationsResponseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBidsRecommendationsResponseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBidsRecommendationsResponseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBidsRecommendationsResponseDTO(val *GetBidsRecommendationsResponseDTO) *NullableGetBidsRecommendationsResponseDTO {
	return &NullableGetBidsRecommendationsResponseDTO{value: val, isSet: true}
}

func (v NullableGetBidsRecommendationsResponseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBidsRecommendationsResponseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


