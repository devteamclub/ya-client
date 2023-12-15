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

// checks if the UpdateOrderStorageLimitRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrderStorageLimitRequest{}

// UpdateOrderStorageLimitRequest Запрос на обновление срока хранения заказа в ПВЗ.
type UpdateOrderStorageLimitRequest struct {
	// Новая дата, до которой заказ будет храниться в пункте выдачи.  Срок хранения можно увеличить не больше, чем на 30 дней.  Формат даты: `ГГГГ-ММ-ДД`. 
	NewDate string `json:"newDate"`
}

type _UpdateOrderStorageLimitRequest UpdateOrderStorageLimitRequest

// NewUpdateOrderStorageLimitRequest instantiates a new UpdateOrderStorageLimitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrderStorageLimitRequest(newDate string) *UpdateOrderStorageLimitRequest {
	this := UpdateOrderStorageLimitRequest{}
	this.NewDate = newDate
	return &this
}

// NewUpdateOrderStorageLimitRequestWithDefaults instantiates a new UpdateOrderStorageLimitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrderStorageLimitRequestWithDefaults() *UpdateOrderStorageLimitRequest {
	this := UpdateOrderStorageLimitRequest{}
	return &this
}

// GetNewDate returns the NewDate field value
func (o *UpdateOrderStorageLimitRequest) GetNewDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewDate
}

// GetNewDateOk returns a tuple with the NewDate field value
// and a boolean to check if the value has been set.
func (o *UpdateOrderStorageLimitRequest) GetNewDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewDate, true
}

// SetNewDate sets field value
func (o *UpdateOrderStorageLimitRequest) SetNewDate(v string) {
	o.NewDate = v
}

func (o UpdateOrderStorageLimitRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrderStorageLimitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["newDate"] = o.NewDate
	return toSerialize, nil
}

func (o *UpdateOrderStorageLimitRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"newDate",
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

	varUpdateOrderStorageLimitRequest := _UpdateOrderStorageLimitRequest{}

	err = json.Unmarshal(bytes, &varUpdateOrderStorageLimitRequest)

	if err != nil {
		return err
	}

	*o = UpdateOrderStorageLimitRequest(varUpdateOrderStorageLimitRequest)

	return err
}

type NullableUpdateOrderStorageLimitRequest struct {
	value *UpdateOrderStorageLimitRequest
	isSet bool
}

func (v NullableUpdateOrderStorageLimitRequest) Get() *UpdateOrderStorageLimitRequest {
	return v.value
}

func (v *NullableUpdateOrderStorageLimitRequest) Set(val *UpdateOrderStorageLimitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrderStorageLimitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrderStorageLimitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrderStorageLimitRequest(val *UpdateOrderStorageLimitRequest) *NullableUpdateOrderStorageLimitRequest {
	return &NullableUpdateOrderStorageLimitRequest{value: val, isSet: true}
}

func (v NullableUpdateOrderStorageLimitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrderStorageLimitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

