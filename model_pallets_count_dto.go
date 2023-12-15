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

// checks if the PalletsCountDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalletsCountDTO{}

// PalletsCountDTO Количество палет в отгрузке.
type PalletsCountDTO struct {
	// Количество палет, которое заявил продавец.
	Planned *int32 `json:"planned,omitempty"`
	// Количество палет, которое приняли в сортировочном центре.
	Fact *int32 `json:"fact,omitempty"`
}

// NewPalletsCountDTO instantiates a new PalletsCountDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalletsCountDTO() *PalletsCountDTO {
	this := PalletsCountDTO{}
	return &this
}

// NewPalletsCountDTOWithDefaults instantiates a new PalletsCountDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalletsCountDTOWithDefaults() *PalletsCountDTO {
	this := PalletsCountDTO{}
	return &this
}

// GetPlanned returns the Planned field value if set, zero value otherwise.
func (o *PalletsCountDTO) GetPlanned() int32 {
	if o == nil || IsNil(o.Planned) {
		var ret int32
		return ret
	}
	return *o.Planned
}

// GetPlannedOk returns a tuple with the Planned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PalletsCountDTO) GetPlannedOk() (*int32, bool) {
	if o == nil || IsNil(o.Planned) {
		return nil, false
	}
	return o.Planned, true
}

// HasPlanned returns a boolean if a field has been set.
func (o *PalletsCountDTO) HasPlanned() bool {
	if o != nil && !IsNil(o.Planned) {
		return true
	}

	return false
}

// SetPlanned gets a reference to the given int32 and assigns it to the Planned field.
func (o *PalletsCountDTO) SetPlanned(v int32) {
	o.Planned = &v
}

// GetFact returns the Fact field value if set, zero value otherwise.
func (o *PalletsCountDTO) GetFact() int32 {
	if o == nil || IsNil(o.Fact) {
		var ret int32
		return ret
	}
	return *o.Fact
}

// GetFactOk returns a tuple with the Fact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PalletsCountDTO) GetFactOk() (*int32, bool) {
	if o == nil || IsNil(o.Fact) {
		return nil, false
	}
	return o.Fact, true
}

// HasFact returns a boolean if a field has been set.
func (o *PalletsCountDTO) HasFact() bool {
	if o != nil && !IsNil(o.Fact) {
		return true
	}

	return false
}

// SetFact gets a reference to the given int32 and assigns it to the Fact field.
func (o *PalletsCountDTO) SetFact(v int32) {
	o.Fact = &v
}

func (o PalletsCountDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalletsCountDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Planned) {
		toSerialize["planned"] = o.Planned
	}
	if !IsNil(o.Fact) {
		toSerialize["fact"] = o.Fact
	}
	return toSerialize, nil
}

type NullablePalletsCountDTO struct {
	value *PalletsCountDTO
	isSet bool
}

func (v NullablePalletsCountDTO) Get() *PalletsCountDTO {
	return v.value
}

func (v *NullablePalletsCountDTO) Set(val *PalletsCountDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePalletsCountDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePalletsCountDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalletsCountDTO(val *PalletsCountDTO) *NullablePalletsCountDTO {
	return &NullablePalletsCountDTO{value: val, isSet: true}
}

func (v NullablePalletsCountDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalletsCountDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

