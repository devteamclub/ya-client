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

// checks if the OrdersShipmentInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrdersShipmentInfoDTO{}

// OrdersShipmentInfoDTO Годные/негодные ярлыки по заказам в отгрузке.
type OrdersShipmentInfoDTO struct {
	// Идентификаторы заказов в отгрузке, для которых можно распечатать ярлыки.
	OrderIdsWithLabels []int64 `json:"orderIdsWithLabels"`
	// Идентификаторы заказов в отгрузке, для которых нельзя распечатать ярлыки.
	OrderIdsWithoutLabels []int64 `json:"orderIdsWithoutLabels"`
}

type _OrdersShipmentInfoDTO OrdersShipmentInfoDTO

// NewOrdersShipmentInfoDTO instantiates a new OrdersShipmentInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdersShipmentInfoDTO(orderIdsWithLabels []int64, orderIdsWithoutLabels []int64) *OrdersShipmentInfoDTO {
	this := OrdersShipmentInfoDTO{}
	this.OrderIdsWithLabels = orderIdsWithLabels
	this.OrderIdsWithoutLabels = orderIdsWithoutLabels
	return &this
}

// NewOrdersShipmentInfoDTOWithDefaults instantiates a new OrdersShipmentInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdersShipmentInfoDTOWithDefaults() *OrdersShipmentInfoDTO {
	this := OrdersShipmentInfoDTO{}
	return &this
}

// GetOrderIdsWithLabels returns the OrderIdsWithLabels field value
func (o *OrdersShipmentInfoDTO) GetOrderIdsWithLabels() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.OrderIdsWithLabels
}

// GetOrderIdsWithLabelsOk returns a tuple with the OrderIdsWithLabels field value
// and a boolean to check if the value has been set.
func (o *OrdersShipmentInfoDTO) GetOrderIdsWithLabelsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderIdsWithLabels, true
}

// SetOrderIdsWithLabels sets field value
func (o *OrdersShipmentInfoDTO) SetOrderIdsWithLabels(v []int64) {
	o.OrderIdsWithLabels = v
}

// GetOrderIdsWithoutLabels returns the OrderIdsWithoutLabels field value
func (o *OrdersShipmentInfoDTO) GetOrderIdsWithoutLabels() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.OrderIdsWithoutLabels
}

// GetOrderIdsWithoutLabelsOk returns a tuple with the OrderIdsWithoutLabels field value
// and a boolean to check if the value has been set.
func (o *OrdersShipmentInfoDTO) GetOrderIdsWithoutLabelsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderIdsWithoutLabels, true
}

// SetOrderIdsWithoutLabels sets field value
func (o *OrdersShipmentInfoDTO) SetOrderIdsWithoutLabels(v []int64) {
	o.OrderIdsWithoutLabels = v
}

func (o OrdersShipmentInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrdersShipmentInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderIdsWithLabels"] = o.OrderIdsWithLabels
	toSerialize["orderIdsWithoutLabels"] = o.OrderIdsWithoutLabels
	return toSerialize, nil
}

func (o *OrdersShipmentInfoDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orderIdsWithLabels",
		"orderIdsWithoutLabels",
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

	varOrdersShipmentInfoDTO := _OrdersShipmentInfoDTO{}

	err = json.Unmarshal(bytes, &varOrdersShipmentInfoDTO)

	if err != nil {
		return err
	}

	*o = OrdersShipmentInfoDTO(varOrdersShipmentInfoDTO)

	return err
}

type NullableOrdersShipmentInfoDTO struct {
	value *OrdersShipmentInfoDTO
	isSet bool
}

func (v NullableOrdersShipmentInfoDTO) Get() *OrdersShipmentInfoDTO {
	return v.value
}

func (v *NullableOrdersShipmentInfoDTO) Set(val *OrdersShipmentInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersShipmentInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersShipmentInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersShipmentInfoDTO(val *OrdersShipmentInfoDTO) *NullableOrdersShipmentInfoDTO {
	return &NullableOrdersShipmentInfoDTO{value: val, isSet: true}
}

func (v NullableOrdersShipmentInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersShipmentInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

