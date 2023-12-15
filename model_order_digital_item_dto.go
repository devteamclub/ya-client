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

// checks if the OrderDigitalItemDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDigitalItemDTO{}

// OrderDigitalItemDTO Ключ цифрового товара.
type OrderDigitalItemDTO struct {
	// Идентификатор товара в заказе.  Он приходит в ответе на запрос [GET campaigns/{campaignId}/orders/{orderId}](../../reference/orders/getOrder.md) и в запросе Маркета [POST order/accept](../../pushapi/reference/post-order-accept.md) — параметр `id` в `items`. 
	Id int64 `json:"id"`
	// Сам ключ.
	Code string `json:"code"`
	// Инструкция по активации.
	Slip string `json:"slip"`
	// Дата, до которой нужно активировать ключ. Если ключ действует бессрочно, укажите любую дату в отдаленном будущем.  Формат даты: `ГГГГ-ММ-ДД`. 
	ActivateTill string `json:"activate_till"`
}

type _OrderDigitalItemDTO OrderDigitalItemDTO

// NewOrderDigitalItemDTO instantiates a new OrderDigitalItemDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDigitalItemDTO(id int64, code string, slip string, activateTill string) *OrderDigitalItemDTO {
	this := OrderDigitalItemDTO{}
	this.Id = id
	this.Code = code
	this.Slip = slip
	this.ActivateTill = activateTill
	return &this
}

// NewOrderDigitalItemDTOWithDefaults instantiates a new OrderDigitalItemDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDigitalItemDTOWithDefaults() *OrderDigitalItemDTO {
	this := OrderDigitalItemDTO{}
	return &this
}

// GetId returns the Id field value
func (o *OrderDigitalItemDTO) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderDigitalItemDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderDigitalItemDTO) SetId(v int64) {
	o.Id = v
}

// GetCode returns the Code field value
func (o *OrderDigitalItemDTO) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OrderDigitalItemDTO) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OrderDigitalItemDTO) SetCode(v string) {
	o.Code = v
}

// GetSlip returns the Slip field value
func (o *OrderDigitalItemDTO) GetSlip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slip
}

// GetSlipOk returns a tuple with the Slip field value
// and a boolean to check if the value has been set.
func (o *OrderDigitalItemDTO) GetSlipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slip, true
}

// SetSlip sets field value
func (o *OrderDigitalItemDTO) SetSlip(v string) {
	o.Slip = v
}

// GetActivateTill returns the ActivateTill field value
func (o *OrderDigitalItemDTO) GetActivateTill() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActivateTill
}

// GetActivateTillOk returns a tuple with the ActivateTill field value
// and a boolean to check if the value has been set.
func (o *OrderDigitalItemDTO) GetActivateTillOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivateTill, true
}

// SetActivateTill sets field value
func (o *OrderDigitalItemDTO) SetActivateTill(v string) {
	o.ActivateTill = v
}

func (o OrderDigitalItemDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDigitalItemDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["code"] = o.Code
	toSerialize["slip"] = o.Slip
	toSerialize["activate_till"] = o.ActivateTill
	return toSerialize, nil
}

func (o *OrderDigitalItemDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"code",
		"slip",
		"activate_till",
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

	varOrderDigitalItemDTO := _OrderDigitalItemDTO{}

	err = json.Unmarshal(bytes, &varOrderDigitalItemDTO)

	if err != nil {
		return err
	}

	*o = OrderDigitalItemDTO(varOrderDigitalItemDTO)

	return err
}

type NullableOrderDigitalItemDTO struct {
	value *OrderDigitalItemDTO
	isSet bool
}

func (v NullableOrderDigitalItemDTO) Get() *OrderDigitalItemDTO {
	return v.value
}

func (v *NullableOrderDigitalItemDTO) Set(val *OrderDigitalItemDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDigitalItemDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDigitalItemDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDigitalItemDTO(val *OrderDigitalItemDTO) *NullableOrderDigitalItemDTO {
	return &NullableOrderDigitalItemDTO{value: val, isSet: true}
}

func (v NullableOrderDigitalItemDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDigitalItemDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

