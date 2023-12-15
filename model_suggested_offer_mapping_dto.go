/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ya-client

import (
	"encoding/json"
)

// checks if the SuggestedOfferMappingDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuggestedOfferMappingDTO{}

// SuggestedOfferMappingDTO Товар с соответствующей карточкой на Маркете.
type SuggestedOfferMappingDTO struct {
	Offer *SuggestedOfferDTO `json:"offer,omitempty"`
	Mapping *GetMappingDTO `json:"mapping,omitempty"`
}

// NewSuggestedOfferMappingDTO instantiates a new SuggestedOfferMappingDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestedOfferMappingDTO() *SuggestedOfferMappingDTO {
	this := SuggestedOfferMappingDTO{}
	return &this
}

// NewSuggestedOfferMappingDTOWithDefaults instantiates a new SuggestedOfferMappingDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestedOfferMappingDTOWithDefaults() *SuggestedOfferMappingDTO {
	this := SuggestedOfferMappingDTO{}
	return &this
}

// GetOffer returns the Offer field value if set, zero value otherwise.
func (o *SuggestedOfferMappingDTO) GetOffer() SuggestedOfferDTO {
	if o == nil || IsNil(o.Offer) {
		var ret SuggestedOfferDTO
		return ret
	}
	return *o.Offer
}

// GetOfferOk returns a tuple with the Offer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedOfferMappingDTO) GetOfferOk() (*SuggestedOfferDTO, bool) {
	if o == nil || IsNil(o.Offer) {
		return nil, false
	}
	return o.Offer, true
}

// HasOffer returns a boolean if a field has been set.
func (o *SuggestedOfferMappingDTO) HasOffer() bool {
	if o != nil && !IsNil(o.Offer) {
		return true
	}

	return false
}

// SetOffer gets a reference to the given SuggestedOfferDTO and assigns it to the Offer field.
func (o *SuggestedOfferMappingDTO) SetOffer(v SuggestedOfferDTO) {
	o.Offer = &v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *SuggestedOfferMappingDTO) GetMapping() GetMappingDTO {
	if o == nil || IsNil(o.Mapping) {
		var ret GetMappingDTO
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedOfferMappingDTO) GetMappingOk() (*GetMappingDTO, bool) {
	if o == nil || IsNil(o.Mapping) {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *SuggestedOfferMappingDTO) HasMapping() bool {
	if o != nil && !IsNil(o.Mapping) {
		return true
	}

	return false
}

// SetMapping gets a reference to the given GetMappingDTO and assigns it to the Mapping field.
func (o *SuggestedOfferMappingDTO) SetMapping(v GetMappingDTO) {
	o.Mapping = &v
}

func (o SuggestedOfferMappingDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuggestedOfferMappingDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Offer) {
		toSerialize["offer"] = o.Offer
	}
	if !IsNil(o.Mapping) {
		toSerialize["mapping"] = o.Mapping
	}
	return toSerialize, nil
}

type NullableSuggestedOfferMappingDTO struct {
	value *SuggestedOfferMappingDTO
	isSet bool
}

func (v NullableSuggestedOfferMappingDTO) Get() *SuggestedOfferMappingDTO {
	return v.value
}

func (v *NullableSuggestedOfferMappingDTO) Set(val *SuggestedOfferMappingDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestedOfferMappingDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestedOfferMappingDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestedOfferMappingDTO(val *SuggestedOfferMappingDTO) *NullableSuggestedOfferMappingDTO {
	return &NullableSuggestedOfferMappingDTO{value: val, isSet: true}
}

func (v NullableSuggestedOfferMappingDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuggestedOfferMappingDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


