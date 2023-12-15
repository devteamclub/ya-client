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

// checks if the WarehousesDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarehousesDTO{}

// WarehousesDTO Информация о складах и группах складов.
type WarehousesDTO struct {
	// Список складов, не входящих в группы.
	Warehouses []WarehouseDTO `json:"warehouses,omitempty"`
	// Список групп складов.
	WarehouseGroups []WarehouseGroupDTO `json:"warehouseGroups,omitempty"`
}

// NewWarehousesDTO instantiates a new WarehousesDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehousesDTO() *WarehousesDTO {
	this := WarehousesDTO{}
	return &this
}

// NewWarehousesDTOWithDefaults instantiates a new WarehousesDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehousesDTOWithDefaults() *WarehousesDTO {
	this := WarehousesDTO{}
	return &this
}

// GetWarehouses returns the Warehouses field value if set, zero value otherwise.
func (o *WarehousesDTO) GetWarehouses() []WarehouseDTO {
	if o == nil || IsNil(o.Warehouses) {
		var ret []WarehouseDTO
		return ret
	}
	return o.Warehouses
}

// GetWarehousesOk returns a tuple with the Warehouses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehousesDTO) GetWarehousesOk() ([]WarehouseDTO, bool) {
	if o == nil || IsNil(o.Warehouses) {
		return nil, false
	}
	return o.Warehouses, true
}

// HasWarehouses returns a boolean if a field has been set.
func (o *WarehousesDTO) HasWarehouses() bool {
	if o != nil && !IsNil(o.Warehouses) {
		return true
	}

	return false
}

// SetWarehouses gets a reference to the given []WarehouseDTO and assigns it to the Warehouses field.
func (o *WarehousesDTO) SetWarehouses(v []WarehouseDTO) {
	o.Warehouses = v
}

// GetWarehouseGroups returns the WarehouseGroups field value if set, zero value otherwise.
func (o *WarehousesDTO) GetWarehouseGroups() []WarehouseGroupDTO {
	if o == nil || IsNil(o.WarehouseGroups) {
		var ret []WarehouseGroupDTO
		return ret
	}
	return o.WarehouseGroups
}

// GetWarehouseGroupsOk returns a tuple with the WarehouseGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehousesDTO) GetWarehouseGroupsOk() ([]WarehouseGroupDTO, bool) {
	if o == nil || IsNil(o.WarehouseGroups) {
		return nil, false
	}
	return o.WarehouseGroups, true
}

// HasWarehouseGroups returns a boolean if a field has been set.
func (o *WarehousesDTO) HasWarehouseGroups() bool {
	if o != nil && !IsNil(o.WarehouseGroups) {
		return true
	}

	return false
}

// SetWarehouseGroups gets a reference to the given []WarehouseGroupDTO and assigns it to the WarehouseGroups field.
func (o *WarehousesDTO) SetWarehouseGroups(v []WarehouseGroupDTO) {
	o.WarehouseGroups = v
}

func (o WarehousesDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarehousesDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warehouses) {
		toSerialize["warehouses"] = o.Warehouses
	}
	if !IsNil(o.WarehouseGroups) {
		toSerialize["warehouseGroups"] = o.WarehouseGroups
	}
	return toSerialize, nil
}

type NullableWarehousesDTO struct {
	value *WarehousesDTO
	isSet bool
}

func (v NullableWarehousesDTO) Get() *WarehousesDTO {
	return v.value
}

func (v *NullableWarehousesDTO) Set(val *WarehousesDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehousesDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehousesDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehousesDTO(val *WarehousesDTO) *NullableWarehousesDTO {
	return &NullableWarehousesDTO{value: val, isSet: true}
}

func (v NullableWarehousesDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehousesDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


