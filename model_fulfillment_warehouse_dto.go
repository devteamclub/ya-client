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

// checks if the FulfillmentWarehouseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FulfillmentWarehouseDTO{}

// FulfillmentWarehouseDTO Склад Маркета (FBY).
type FulfillmentWarehouseDTO struct {
	// Идентификатор склада.
	Id int64 `json:"id"`
	// Название склада.
	Name string `json:"name"`
}

type _FulfillmentWarehouseDTO FulfillmentWarehouseDTO

// NewFulfillmentWarehouseDTO instantiates a new FulfillmentWarehouseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentWarehouseDTO(id int64, name string) *FulfillmentWarehouseDTO {
	this := FulfillmentWarehouseDTO{}
	this.Id = id
	this.Name = name
	return &this
}

// NewFulfillmentWarehouseDTOWithDefaults instantiates a new FulfillmentWarehouseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentWarehouseDTOWithDefaults() *FulfillmentWarehouseDTO {
	this := FulfillmentWarehouseDTO{}
	return &this
}

// GetId returns the Id field value
func (o *FulfillmentWarehouseDTO) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FulfillmentWarehouseDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FulfillmentWarehouseDTO) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FulfillmentWarehouseDTO) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FulfillmentWarehouseDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FulfillmentWarehouseDTO) SetName(v string) {
	o.Name = v
}

func (o FulfillmentWarehouseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FulfillmentWarehouseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *FulfillmentWarehouseDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varFulfillmentWarehouseDTO := _FulfillmentWarehouseDTO{}

	err = json.Unmarshal(bytes, &varFulfillmentWarehouseDTO)

	if err != nil {
		return err
	}

	*o = FulfillmentWarehouseDTO(varFulfillmentWarehouseDTO)

	return err
}

type NullableFulfillmentWarehouseDTO struct {
	value *FulfillmentWarehouseDTO
	isSet bool
}

func (v NullableFulfillmentWarehouseDTO) Get() *FulfillmentWarehouseDTO {
	return v.value
}

func (v *NullableFulfillmentWarehouseDTO) Set(val *FulfillmentWarehouseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentWarehouseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentWarehouseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentWarehouseDTO(val *FulfillmentWarehouseDTO) *NullableFulfillmentWarehouseDTO {
	return &NullableFulfillmentWarehouseDTO{value: val, isSet: true}
}

func (v NullableFulfillmentWarehouseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillmentWarehouseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

