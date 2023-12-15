/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the GetGoodsStatsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGoodsStatsRequest{}

// GetGoodsStatsRequest Запрос отчета по товарам.
type GetGoodsStatsRequest struct {
	// Список ваших идентификаторов SKU. Максимальное количество идентификаторов: 500. Обязательный параметр. Должен содержать хотя бы один SKU. 
	ShopSkus []string `json:"shopSkus"`
}

type _GetGoodsStatsRequest GetGoodsStatsRequest

// NewGetGoodsStatsRequest instantiates a new GetGoodsStatsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGoodsStatsRequest(shopSkus []string) *GetGoodsStatsRequest {
	this := GetGoodsStatsRequest{}
	this.ShopSkus = shopSkus
	return &this
}

// NewGetGoodsStatsRequestWithDefaults instantiates a new GetGoodsStatsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGoodsStatsRequestWithDefaults() *GetGoodsStatsRequest {
	this := GetGoodsStatsRequest{}
	return &this
}

// GetShopSkus returns the ShopSkus field value
func (o *GetGoodsStatsRequest) GetShopSkus() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ShopSkus
}

// GetShopSkusOk returns a tuple with the ShopSkus field value
// and a boolean to check if the value has been set.
func (o *GetGoodsStatsRequest) GetShopSkusOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShopSkus, true
}

// SetShopSkus sets field value
func (o *GetGoodsStatsRequest) SetShopSkus(v []string) {
	o.ShopSkus = v
}

func (o GetGoodsStatsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGoodsStatsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shopSkus"] = o.ShopSkus
	return toSerialize, nil
}

func (o *GetGoodsStatsRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"shopSkus",
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

	varGetGoodsStatsRequest := _GetGoodsStatsRequest{}

	err = json.Unmarshal(bytes, &varGetGoodsStatsRequest)

	if err != nil {
		return err
	}

	*o = GetGoodsStatsRequest(varGetGoodsStatsRequest)

	return err
}

type NullableGetGoodsStatsRequest struct {
	value *GetGoodsStatsRequest
	isSet bool
}

func (v NullableGetGoodsStatsRequest) Get() *GetGoodsStatsRequest {
	return v.value
}

func (v *NullableGetGoodsStatsRequest) Set(val *GetGoodsStatsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGoodsStatsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGoodsStatsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGoodsStatsRequest(val *GetGoodsStatsRequest) *NullableGetGoodsStatsRequest {
	return &NullableGetGoodsStatsRequest{value: val, isSet: true}
}

func (v NullableGetGoodsStatsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGoodsStatsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


