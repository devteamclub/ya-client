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

// checks if the OrderParcelBoxDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderParcelBoxDTO{}

// OrderParcelBoxDTO Информация о грузоместе.
type OrderParcelBoxDTO struct {
	// Идентификатор грузоместа.
	Id *int64 `json:"id,omitempty"`
	// Идентификатор грузового места в информационной системе магазина.
	FulfilmentId *string `json:"fulfilmentId,omitempty"`
}

// NewOrderParcelBoxDTO instantiates a new OrderParcelBoxDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderParcelBoxDTO() *OrderParcelBoxDTO {
	this := OrderParcelBoxDTO{}
	return &this
}

// NewOrderParcelBoxDTOWithDefaults instantiates a new OrderParcelBoxDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderParcelBoxDTOWithDefaults() *OrderParcelBoxDTO {
	this := OrderParcelBoxDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderParcelBoxDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderParcelBoxDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderParcelBoxDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OrderParcelBoxDTO) SetId(v int64) {
	o.Id = &v
}

// GetFulfilmentId returns the FulfilmentId field value if set, zero value otherwise.
func (o *OrderParcelBoxDTO) GetFulfilmentId() string {
	if o == nil || IsNil(o.FulfilmentId) {
		var ret string
		return ret
	}
	return *o.FulfilmentId
}

// GetFulfilmentIdOk returns a tuple with the FulfilmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderParcelBoxDTO) GetFulfilmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.FulfilmentId) {
		return nil, false
	}
	return o.FulfilmentId, true
}

// HasFulfilmentId returns a boolean if a field has been set.
func (o *OrderParcelBoxDTO) HasFulfilmentId() bool {
	if o != nil && !IsNil(o.FulfilmentId) {
		return true
	}

	return false
}

// SetFulfilmentId gets a reference to the given string and assigns it to the FulfilmentId field.
func (o *OrderParcelBoxDTO) SetFulfilmentId(v string) {
	o.FulfilmentId = &v
}

func (o OrderParcelBoxDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderParcelBoxDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.FulfilmentId) {
		toSerialize["fulfilmentId"] = o.FulfilmentId
	}
	return toSerialize, nil
}

type NullableOrderParcelBoxDTO struct {
	value *OrderParcelBoxDTO
	isSet bool
}

func (v NullableOrderParcelBoxDTO) Get() *OrderParcelBoxDTO {
	return v.value
}

func (v *NullableOrderParcelBoxDTO) Set(val *OrderParcelBoxDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderParcelBoxDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderParcelBoxDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderParcelBoxDTO(val *OrderParcelBoxDTO) *NullableOrderParcelBoxDTO {
	return &NullableOrderParcelBoxDTO{value: val, isSet: true}
}

func (v NullableOrderParcelBoxDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderParcelBoxDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

