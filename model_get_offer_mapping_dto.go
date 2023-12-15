/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOfferMappingDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOfferMappingDTO{}

// GetOfferMappingDTO Информация о товаре.
type GetOfferMappingDTO struct {
	Offer *GetOfferDTO `json:"offer,omitempty"`
	Mapping *GetMappingDTO `json:"mapping,omitempty"`
}

// NewGetOfferMappingDTO instantiates a new GetOfferMappingDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOfferMappingDTO() *GetOfferMappingDTO {
	this := GetOfferMappingDTO{}
	return &this
}

// NewGetOfferMappingDTOWithDefaults instantiates a new GetOfferMappingDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOfferMappingDTOWithDefaults() *GetOfferMappingDTO {
	this := GetOfferMappingDTO{}
	return &this
}

// GetOffer returns the Offer field value if set, zero value otherwise.
func (o *GetOfferMappingDTO) GetOffer() GetOfferDTO {
	if o == nil || IsNil(o.Offer) {
		var ret GetOfferDTO
		return ret
	}
	return *o.Offer
}

// GetOfferOk returns a tuple with the Offer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOfferMappingDTO) GetOfferOk() (*GetOfferDTO, bool) {
	if o == nil || IsNil(o.Offer) {
		return nil, false
	}
	return o.Offer, true
}

// HasOffer returns a boolean if a field has been set.
func (o *GetOfferMappingDTO) HasOffer() bool {
	if o != nil && !IsNil(o.Offer) {
		return true
	}

	return false
}

// SetOffer gets a reference to the given GetOfferDTO and assigns it to the Offer field.
func (o *GetOfferMappingDTO) SetOffer(v GetOfferDTO) {
	o.Offer = &v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *GetOfferMappingDTO) GetMapping() GetMappingDTO {
	if o == nil || IsNil(o.Mapping) {
		var ret GetMappingDTO
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOfferMappingDTO) GetMappingOk() (*GetMappingDTO, bool) {
	if o == nil || IsNil(o.Mapping) {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *GetOfferMappingDTO) HasMapping() bool {
	if o != nil && !IsNil(o.Mapping) {
		return true
	}

	return false
}

// SetMapping gets a reference to the given GetMappingDTO and assigns it to the Mapping field.
func (o *GetOfferMappingDTO) SetMapping(v GetMappingDTO) {
	o.Mapping = &v
}

func (o GetOfferMappingDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOfferMappingDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Offer) {
		toSerialize["offer"] = o.Offer
	}
	if !IsNil(o.Mapping) {
		toSerialize["mapping"] = o.Mapping
	}
	return toSerialize, nil
}

type NullableGetOfferMappingDTO struct {
	value *GetOfferMappingDTO
	isSet bool
}

func (v NullableGetOfferMappingDTO) Get() *GetOfferMappingDTO {
	return v.value
}

func (v *NullableGetOfferMappingDTO) Set(val *GetOfferMappingDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOfferMappingDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOfferMappingDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOfferMappingDTO(val *GetOfferMappingDTO) *NullableGetOfferMappingDTO {
	return &NullableGetOfferMappingDTO{value: val, isSet: true}
}

func (v NullableGetOfferMappingDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOfferMappingDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


