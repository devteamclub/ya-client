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

// checks if the StockDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockDTO{}

// StockDTO Информация об остатках одного товара на одном из складов.
type StockDTO struct {
	//   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы `. , / \\ ( ) [ ] - = _`  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields). 
	Sku string `json:"sku"`
	// Идентификатор склада.  Узнать идентификатор склада вы можете в личном кабинете в разделе **Логистика → Склады**. Он указан в поле ID склада.  Если вы работаете с общими остатками, вы можете посмотреть идентификатор склада в личном кабинете в разделе **Настройки → Настройки API** в блоке **Обновление данных об остатках товаров** или с помощью запроса [GET businesses/{businessId}/warehouses](../../reference/warehouses/getWarehouses.md).  Если вы отвечаете на запрос Маркета, указывайте тот идентификатор, что пришел в запросе. 
	WarehouseId int64 `json:"warehouseId"`
	// Информация об остатках товара на данном складе. 
	Items []StockItemDTO `json:"items"`
}

type _StockDTO StockDTO

// NewStockDTO instantiates a new StockDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockDTO(sku string, warehouseId int64, items []StockItemDTO) *StockDTO {
	this := StockDTO{}
	this.Sku = sku
	this.WarehouseId = warehouseId
	this.Items = items
	return &this
}

// NewStockDTOWithDefaults instantiates a new StockDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockDTOWithDefaults() *StockDTO {
	this := StockDTO{}
	return &this
}

// GetSku returns the Sku field value
func (o *StockDTO) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *StockDTO) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *StockDTO) SetSku(v string) {
	o.Sku = v
}

// GetWarehouseId returns the WarehouseId field value
func (o *StockDTO) GetWarehouseId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value
// and a boolean to check if the value has been set.
func (o *StockDTO) GetWarehouseIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarehouseId, true
}

// SetWarehouseId sets field value
func (o *StockDTO) SetWarehouseId(v int64) {
	o.WarehouseId = v
}

// GetItems returns the Items field value
func (o *StockDTO) GetItems() []StockItemDTO {
	if o == nil {
		var ret []StockItemDTO
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *StockDTO) GetItemsOk() ([]StockItemDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *StockDTO) SetItems(v []StockItemDTO) {
	o.Items = v
}

func (o StockDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sku"] = o.Sku
	toSerialize["warehouseId"] = o.WarehouseId
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *StockDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sku",
		"warehouseId",
		"items",
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

	varStockDTO := _StockDTO{}

	err = json.Unmarshal(bytes, &varStockDTO)

	if err != nil {
		return err
	}

	*o = StockDTO(varStockDTO)

	return err
}

type NullableStockDTO struct {
	value *StockDTO
	isSet bool
}

func (v NullableStockDTO) Get() *StockDTO {
	return v.value
}

func (v *NullableStockDTO) Set(val *StockDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableStockDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableStockDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockDTO(val *StockDTO) *NullableStockDTO {
	return &NullableStockDTO{value: val, isSet: true}
}

func (v NullableStockDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


