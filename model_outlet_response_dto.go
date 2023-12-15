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

// checks if the OutletResponseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutletResponseDTO{}

// OutletResponseDTO Результат выполнения запроса. Выводится, если `status=\"OK\"`. 
type OutletResponseDTO struct {
	// Идентификатор точки продаж, присвоенный Яндекс Маркетом.
	Id *int64 `json:"id,omitempty"`
}

// NewOutletResponseDTO instantiates a new OutletResponseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutletResponseDTO() *OutletResponseDTO {
	this := OutletResponseDTO{}
	return &this
}

// NewOutletResponseDTOWithDefaults instantiates a new OutletResponseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutletResponseDTOWithDefaults() *OutletResponseDTO {
	this := OutletResponseDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OutletResponseDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutletResponseDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OutletResponseDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OutletResponseDTO) SetId(v int64) {
	o.Id = &v
}

func (o OutletResponseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutletResponseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableOutletResponseDTO struct {
	value *OutletResponseDTO
	isSet bool
}

func (v NullableOutletResponseDTO) Get() *OutletResponseDTO {
	return v.value
}

func (v *NullableOutletResponseDTO) Set(val *OutletResponseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOutletResponseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOutletResponseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutletResponseDTO(val *OutletResponseDTO) *NullableOutletResponseDTO {
	return &NullableOutletResponseDTO{value: val, isSet: true}
}

func (v NullableOutletResponseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutletResponseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


