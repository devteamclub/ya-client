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

// checks if the GoodsStatsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoodsStatsDTO{}

// GoodsStatsDTO Отчет по товарам.
type GoodsStatsDTO struct {
	// Список товаров.
	ShopSkus []GoodsStatsGoodsDTO `json:"shopSkus,omitempty"`
}

// NewGoodsStatsDTO instantiates a new GoodsStatsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoodsStatsDTO() *GoodsStatsDTO {
	this := GoodsStatsDTO{}
	return &this
}

// NewGoodsStatsDTOWithDefaults instantiates a new GoodsStatsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoodsStatsDTOWithDefaults() *GoodsStatsDTO {
	this := GoodsStatsDTO{}
	return &this
}

// GetShopSkus returns the ShopSkus field value if set, zero value otherwise.
func (o *GoodsStatsDTO) GetShopSkus() []GoodsStatsGoodsDTO {
	if o == nil || IsNil(o.ShopSkus) {
		var ret []GoodsStatsGoodsDTO
		return ret
	}
	return o.ShopSkus
}

// GetShopSkusOk returns a tuple with the ShopSkus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoodsStatsDTO) GetShopSkusOk() ([]GoodsStatsGoodsDTO, bool) {
	if o == nil || IsNil(o.ShopSkus) {
		return nil, false
	}
	return o.ShopSkus, true
}

// HasShopSkus returns a boolean if a field has been set.
func (o *GoodsStatsDTO) HasShopSkus() bool {
	if o != nil && !IsNil(o.ShopSkus) {
		return true
	}

	return false
}

// SetShopSkus gets a reference to the given []GoodsStatsGoodsDTO and assigns it to the ShopSkus field.
func (o *GoodsStatsDTO) SetShopSkus(v []GoodsStatsGoodsDTO) {
	o.ShopSkus = v
}

func (o GoodsStatsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoodsStatsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShopSkus) {
		toSerialize["shopSkus"] = o.ShopSkus
	}
	return toSerialize, nil
}

type NullableGoodsStatsDTO struct {
	value *GoodsStatsDTO
	isSet bool
}

func (v NullableGoodsStatsDTO) Get() *GoodsStatsDTO {
	return v.value
}

func (v *NullableGoodsStatsDTO) Set(val *GoodsStatsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGoodsStatsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGoodsStatsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoodsStatsDTO(val *GoodsStatsDTO) *NullableGoodsStatsDTO {
	return &NullableGoodsStatsDTO{value: val, isSet: true}
}

func (v NullableGoodsStatsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoodsStatsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


