/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ya-client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the StockItemDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockItemDTO{}

// StockItemDTO Информация об остатках товара. 
type StockItemDTO struct {
	// Количество доступного товара. 
	Count int64 `json:"count"`
	Type StockType `json:"type"`
	// Дата и время последнего обновления информации об остатках указанного типа.  Формат даты и времени: ISO 8601 со смещением относительно UTC. Например, `2017-11-21T00:42:42+03:00`. 
	UpdatedAt time.Time `json:"updatedAt"`
}

type _StockItemDTO StockItemDTO

// NewStockItemDTO instantiates a new StockItemDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockItemDTO(count int64, type_ StockType, updatedAt time.Time) *StockItemDTO {
	this := StockItemDTO{}
	this.Count = count
	this.Type = type_
	this.UpdatedAt = updatedAt
	return &this
}

// NewStockItemDTOWithDefaults instantiates a new StockItemDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockItemDTOWithDefaults() *StockItemDTO {
	this := StockItemDTO{}
	return &this
}

// GetCount returns the Count field value
func (o *StockItemDTO) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *StockItemDTO) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *StockItemDTO) SetCount(v int64) {
	o.Count = v
}

// GetType returns the Type field value
func (o *StockItemDTO) GetType() StockType {
	if o == nil {
		var ret StockType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StockItemDTO) GetTypeOk() (*StockType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StockItemDTO) SetType(v StockType) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *StockItemDTO) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *StockItemDTO) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *StockItemDTO) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o StockItemDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockItemDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["type"] = o.Type
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *StockItemDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"type",
		"updatedAt",
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

	varStockItemDTO := _StockItemDTO{}

	err = json.Unmarshal(bytes, &varStockItemDTO)

	if err != nil {
		return err
	}

	*o = StockItemDTO(varStockItemDTO)

	return err
}

type NullableStockItemDTO struct {
	value *StockItemDTO
	isSet bool
}

func (v NullableStockItemDTO) Get() *StockItemDTO {
	return v.value
}

func (v *NullableStockItemDTO) Set(val *StockItemDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableStockItemDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableStockItemDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockItemDTO(val *StockItemDTO) *NullableStockItemDTO {
	return &NullableStockItemDTO{value: val, isSet: true}
}

func (v NullableStockItemDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockItemDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

