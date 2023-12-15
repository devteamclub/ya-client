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

// checks if the OrderItemInstanceModificationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderItemInstanceModificationDTO{}

// OrderItemInstanceModificationDTO Позиция в корзине, требующая маркировки.
type OrderItemInstanceModificationDTO struct {
	// Идентификатор товара в заказе.  Он приходит в ответе на запрос [GET campaigns/{campaignId}/orders/{orderId}](../../reference/orders/getOrder.md) и в запросе Маркета [POST order/accept](../../pushapi/reference/post-order-accept.md) — параметр `id` в `items`. 
	Id int64 `json:"id"`
	// Список кодов маркировки единиц товара. 
	Instances []BriefOrderItemInstanceDTO `json:"instances"`
}

type _OrderItemInstanceModificationDTO OrderItemInstanceModificationDTO

// NewOrderItemInstanceModificationDTO instantiates a new OrderItemInstanceModificationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItemInstanceModificationDTO(id int64, instances []BriefOrderItemInstanceDTO) *OrderItemInstanceModificationDTO {
	this := OrderItemInstanceModificationDTO{}
	this.Id = id
	this.Instances = instances
	return &this
}

// NewOrderItemInstanceModificationDTOWithDefaults instantiates a new OrderItemInstanceModificationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemInstanceModificationDTOWithDefaults() *OrderItemInstanceModificationDTO {
	this := OrderItemInstanceModificationDTO{}
	return &this
}

// GetId returns the Id field value
func (o *OrderItemInstanceModificationDTO) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderItemInstanceModificationDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderItemInstanceModificationDTO) SetId(v int64) {
	o.Id = v
}

// GetInstances returns the Instances field value
func (o *OrderItemInstanceModificationDTO) GetInstances() []BriefOrderItemInstanceDTO {
	if o == nil {
		var ret []BriefOrderItemInstanceDTO
		return ret
	}

	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *OrderItemInstanceModificationDTO) GetInstancesOk() ([]BriefOrderItemInstanceDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instances, true
}

// SetInstances sets field value
func (o *OrderItemInstanceModificationDTO) SetInstances(v []BriefOrderItemInstanceDTO) {
	o.Instances = v
}

func (o OrderItemInstanceModificationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderItemInstanceModificationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["instances"] = o.Instances
	return toSerialize, nil
}

func (o *OrderItemInstanceModificationDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"instances",
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

	varOrderItemInstanceModificationDTO := _OrderItemInstanceModificationDTO{}

	err = json.Unmarshal(bytes, &varOrderItemInstanceModificationDTO)

	if err != nil {
		return err
	}

	*o = OrderItemInstanceModificationDTO(varOrderItemInstanceModificationDTO)

	return err
}

type NullableOrderItemInstanceModificationDTO struct {
	value *OrderItemInstanceModificationDTO
	isSet bool
}

func (v NullableOrderItemInstanceModificationDTO) Get() *OrderItemInstanceModificationDTO {
	return v.value
}

func (v *NullableOrderItemInstanceModificationDTO) Set(val *OrderItemInstanceModificationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItemInstanceModificationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItemInstanceModificationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItemInstanceModificationDTO(val *OrderItemInstanceModificationDTO) *NullableOrderItemInstanceModificationDTO {
	return &NullableOrderItemInstanceModificationDTO{value: val, isSet: true}
}

func (v NullableOrderItemInstanceModificationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderItemInstanceModificationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


