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

// checks if the ProvideOrderDigitalCodesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvideOrderDigitalCodesRequest{}

// ProvideOrderDigitalCodesRequest Запрос на передачу ключей цифровых товаров.
type ProvideOrderDigitalCodesRequest struct {
	// Список проданных ключей.  Если в заказе есть несколько **одинаковых** товаров (например, несколько ключей к одной и той же подписке), передайте каждый в виде отдельного элемента массива. `id` у этих элементов должен быть один и тот же. 
	Items []OrderDigitalItemDTO `json:"items"`
}

type _ProvideOrderDigitalCodesRequest ProvideOrderDigitalCodesRequest

// NewProvideOrderDigitalCodesRequest instantiates a new ProvideOrderDigitalCodesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvideOrderDigitalCodesRequest(items []OrderDigitalItemDTO) *ProvideOrderDigitalCodesRequest {
	this := ProvideOrderDigitalCodesRequest{}
	this.Items = items
	return &this
}

// NewProvideOrderDigitalCodesRequestWithDefaults instantiates a new ProvideOrderDigitalCodesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvideOrderDigitalCodesRequestWithDefaults() *ProvideOrderDigitalCodesRequest {
	this := ProvideOrderDigitalCodesRequest{}
	return &this
}

// GetItems returns the Items field value
func (o *ProvideOrderDigitalCodesRequest) GetItems() []OrderDigitalItemDTO {
	if o == nil {
		var ret []OrderDigitalItemDTO
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ProvideOrderDigitalCodesRequest) GetItemsOk() ([]OrderDigitalItemDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ProvideOrderDigitalCodesRequest) SetItems(v []OrderDigitalItemDTO) {
	o.Items = v
}

func (o ProvideOrderDigitalCodesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvideOrderDigitalCodesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *ProvideOrderDigitalCodesRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varProvideOrderDigitalCodesRequest := _ProvideOrderDigitalCodesRequest{}

	err = json.Unmarshal(bytes, &varProvideOrderDigitalCodesRequest)

	if err != nil {
		return err
	}

	*o = ProvideOrderDigitalCodesRequest(varProvideOrderDigitalCodesRequest)

	return err
}

type NullableProvideOrderDigitalCodesRequest struct {
	value *ProvideOrderDigitalCodesRequest
	isSet bool
}

func (v NullableProvideOrderDigitalCodesRequest) Get() *ProvideOrderDigitalCodesRequest {
	return v.value
}

func (v *NullableProvideOrderDigitalCodesRequest) Set(val *ProvideOrderDigitalCodesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProvideOrderDigitalCodesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProvideOrderDigitalCodesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvideOrderDigitalCodesRequest(val *ProvideOrderDigitalCodesRequest) *NullableProvideOrderDigitalCodesRequest {
	return &NullableProvideOrderDigitalCodesRequest{value: val, isSet: true}
}

func (v NullableProvideOrderDigitalCodesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvideOrderDigitalCodesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


