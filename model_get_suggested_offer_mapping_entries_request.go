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

// checks if the GetSuggestedOfferMappingEntriesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSuggestedOfferMappingEntriesRequest{}

// GetSuggestedOfferMappingEntriesRequest Запрос рекомендованных карточек товара.
type GetSuggestedOfferMappingEntriesRequest struct {
	// Список товаров.
	Offers []MappingsOfferDTO `json:"offers"`
}

type _GetSuggestedOfferMappingEntriesRequest GetSuggestedOfferMappingEntriesRequest

// NewGetSuggestedOfferMappingEntriesRequest instantiates a new GetSuggestedOfferMappingEntriesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSuggestedOfferMappingEntriesRequest(offers []MappingsOfferDTO) *GetSuggestedOfferMappingEntriesRequest {
	this := GetSuggestedOfferMappingEntriesRequest{}
	this.Offers = offers
	return &this
}

// NewGetSuggestedOfferMappingEntriesRequestWithDefaults instantiates a new GetSuggestedOfferMappingEntriesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSuggestedOfferMappingEntriesRequestWithDefaults() *GetSuggestedOfferMappingEntriesRequest {
	this := GetSuggestedOfferMappingEntriesRequest{}
	return &this
}

// GetOffers returns the Offers field value
func (o *GetSuggestedOfferMappingEntriesRequest) GetOffers() []MappingsOfferDTO {
	if o == nil {
		var ret []MappingsOfferDTO
		return ret
	}

	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value
// and a boolean to check if the value has been set.
func (o *GetSuggestedOfferMappingEntriesRequest) GetOffersOk() ([]MappingsOfferDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offers, true
}

// SetOffers sets field value
func (o *GetSuggestedOfferMappingEntriesRequest) SetOffers(v []MappingsOfferDTO) {
	o.Offers = v
}

func (o GetSuggestedOfferMappingEntriesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSuggestedOfferMappingEntriesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offers"] = o.Offers
	return toSerialize, nil
}

func (o *GetSuggestedOfferMappingEntriesRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"offers",
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

	varGetSuggestedOfferMappingEntriesRequest := _GetSuggestedOfferMappingEntriesRequest{}

	err = json.Unmarshal(bytes, &varGetSuggestedOfferMappingEntriesRequest)

	if err != nil {
		return err
	}

	*o = GetSuggestedOfferMappingEntriesRequest(varGetSuggestedOfferMappingEntriesRequest)

	return err
}

type NullableGetSuggestedOfferMappingEntriesRequest struct {
	value *GetSuggestedOfferMappingEntriesRequest
	isSet bool
}

func (v NullableGetSuggestedOfferMappingEntriesRequest) Get() *GetSuggestedOfferMappingEntriesRequest {
	return v.value
}

func (v *NullableGetSuggestedOfferMappingEntriesRequest) Set(val *GetSuggestedOfferMappingEntriesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSuggestedOfferMappingEntriesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSuggestedOfferMappingEntriesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSuggestedOfferMappingEntriesRequest(val *GetSuggestedOfferMappingEntriesRequest) *NullableGetSuggestedOfferMappingEntriesRequest {
	return &NullableGetSuggestedOfferMappingEntriesRequest{value: val, isSet: true}
}

func (v NullableGetSuggestedOfferMappingEntriesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSuggestedOfferMappingEntriesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

