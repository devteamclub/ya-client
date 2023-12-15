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

// checks if the GoodsStatsWarehouseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoodsStatsWarehouseDTO{}

// GoodsStatsWarehouseDTO Информация о складе.
type GoodsStatsWarehouseDTO struct {
	// Идентификатор склада.
	Id *int64 `json:"id,omitempty"`
	// Название склада.
	Name *string `json:"name,omitempty"`
	// Информация об остатках товаров на складе.
	Stocks []WarehouseStockDTO `json:"stocks,omitempty"`
}

// NewGoodsStatsWarehouseDTO instantiates a new GoodsStatsWarehouseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoodsStatsWarehouseDTO() *GoodsStatsWarehouseDTO {
	this := GoodsStatsWarehouseDTO{}
	return &this
}

// NewGoodsStatsWarehouseDTOWithDefaults instantiates a new GoodsStatsWarehouseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoodsStatsWarehouseDTOWithDefaults() *GoodsStatsWarehouseDTO {
	this := GoodsStatsWarehouseDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GoodsStatsWarehouseDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoodsStatsWarehouseDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GoodsStatsWarehouseDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GoodsStatsWarehouseDTO) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GoodsStatsWarehouseDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoodsStatsWarehouseDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GoodsStatsWarehouseDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GoodsStatsWarehouseDTO) SetName(v string) {
	o.Name = &v
}

// GetStocks returns the Stocks field value if set, zero value otherwise.
func (o *GoodsStatsWarehouseDTO) GetStocks() []WarehouseStockDTO {
	if o == nil || IsNil(o.Stocks) {
		var ret []WarehouseStockDTO
		return ret
	}
	return o.Stocks
}

// GetStocksOk returns a tuple with the Stocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoodsStatsWarehouseDTO) GetStocksOk() ([]WarehouseStockDTO, bool) {
	if o == nil || IsNil(o.Stocks) {
		return nil, false
	}
	return o.Stocks, true
}

// HasStocks returns a boolean if a field has been set.
func (o *GoodsStatsWarehouseDTO) HasStocks() bool {
	if o != nil && !IsNil(o.Stocks) {
		return true
	}

	return false
}

// SetStocks gets a reference to the given []WarehouseStockDTO and assigns it to the Stocks field.
func (o *GoodsStatsWarehouseDTO) SetStocks(v []WarehouseStockDTO) {
	o.Stocks = v
}

func (o GoodsStatsWarehouseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoodsStatsWarehouseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Stocks) {
		toSerialize["stocks"] = o.Stocks
	}
	return toSerialize, nil
}

type NullableGoodsStatsWarehouseDTO struct {
	value *GoodsStatsWarehouseDTO
	isSet bool
}

func (v NullableGoodsStatsWarehouseDTO) Get() *GoodsStatsWarehouseDTO {
	return v.value
}

func (v *NullableGoodsStatsWarehouseDTO) Set(val *GoodsStatsWarehouseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGoodsStatsWarehouseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGoodsStatsWarehouseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoodsStatsWarehouseDTO(val *GoodsStatsWarehouseDTO) *NullableGoodsStatsWarehouseDTO {
	return &NullableGoodsStatsWarehouseDTO{value: val, isSet: true}
}

func (v NullableGoodsStatsWarehouseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoodsStatsWarehouseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


