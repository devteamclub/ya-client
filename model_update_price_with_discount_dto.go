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

// checks if the UpdatePriceWithDiscountDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePriceWithDiscountDTO{}

// UpdatePriceWithDiscountDTO Цена с указанием скидки.
type UpdatePriceWithDiscountDTO struct {
	// Значение.
	Value float32 `json:"value"`
	// Валюта.  Если `BasePriceDTO` присутствует в запросе, указывайте `RUR` — российский рубль. 
	CurrencyId string `json:"currencyId"`
	// Цена до скидки.
	DiscountBase *float32 `json:"discountBase,omitempty"`
}

type _UpdatePriceWithDiscountDTO UpdatePriceWithDiscountDTO

// NewUpdatePriceWithDiscountDTO instantiates a new UpdatePriceWithDiscountDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePriceWithDiscountDTO(value float32, currencyId string) *UpdatePriceWithDiscountDTO {
	this := UpdatePriceWithDiscountDTO{}
	this.Value = value
	this.CurrencyId = currencyId
	return &this
}

// NewUpdatePriceWithDiscountDTOWithDefaults instantiates a new UpdatePriceWithDiscountDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePriceWithDiscountDTOWithDefaults() *UpdatePriceWithDiscountDTO {
	this := UpdatePriceWithDiscountDTO{}
	return &this
}

// GetValue returns the Value field value
func (o *UpdatePriceWithDiscountDTO) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UpdatePriceWithDiscountDTO) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UpdatePriceWithDiscountDTO) SetValue(v float32) {
	o.Value = v
}

// GetCurrencyId returns the CurrencyId field value
func (o *UpdatePriceWithDiscountDTO) GetCurrencyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value
// and a boolean to check if the value has been set.
func (o *UpdatePriceWithDiscountDTO) GetCurrencyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyId, true
}

// SetCurrencyId sets field value
func (o *UpdatePriceWithDiscountDTO) SetCurrencyId(v string) {
	o.CurrencyId = v
}

// GetDiscountBase returns the DiscountBase field value if set, zero value otherwise.
func (o *UpdatePriceWithDiscountDTO) GetDiscountBase() float32 {
	if o == nil || IsNil(o.DiscountBase) {
		var ret float32
		return ret
	}
	return *o.DiscountBase
}

// GetDiscountBaseOk returns a tuple with the DiscountBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePriceWithDiscountDTO) GetDiscountBaseOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountBase) {
		return nil, false
	}
	return o.DiscountBase, true
}

// HasDiscountBase returns a boolean if a field has been set.
func (o *UpdatePriceWithDiscountDTO) HasDiscountBase() bool {
	if o != nil && !IsNil(o.DiscountBase) {
		return true
	}

	return false
}

// SetDiscountBase gets a reference to the given float32 and assigns it to the DiscountBase field.
func (o *UpdatePriceWithDiscountDTO) SetDiscountBase(v float32) {
	o.DiscountBase = &v
}

func (o UpdatePriceWithDiscountDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePriceWithDiscountDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["currencyId"] = o.CurrencyId
	if !IsNil(o.DiscountBase) {
		toSerialize["discountBase"] = o.DiscountBase
	}
	return toSerialize, nil
}

func (o *UpdatePriceWithDiscountDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"currencyId",
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

	varUpdatePriceWithDiscountDTO := _UpdatePriceWithDiscountDTO{}

	err = json.Unmarshal(bytes, &varUpdatePriceWithDiscountDTO)

	if err != nil {
		return err
	}

	*o = UpdatePriceWithDiscountDTO(varUpdatePriceWithDiscountDTO)

	return err
}

type NullableUpdatePriceWithDiscountDTO struct {
	value *UpdatePriceWithDiscountDTO
	isSet bool
}

func (v NullableUpdatePriceWithDiscountDTO) Get() *UpdatePriceWithDiscountDTO {
	return v.value
}

func (v *NullableUpdatePriceWithDiscountDTO) Set(val *UpdatePriceWithDiscountDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePriceWithDiscountDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePriceWithDiscountDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePriceWithDiscountDTO(val *UpdatePriceWithDiscountDTO) *NullableUpdatePriceWithDiscountDTO {
	return &NullableUpdatePriceWithDiscountDTO{value: val, isSet: true}
}

func (v NullableUpdatePriceWithDiscountDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePriceWithDiscountDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


