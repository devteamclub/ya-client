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

// checks if the ConfirmPricesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmPricesRequest{}

// ConfirmPricesRequest Запрос на подтверждение цены. 
type ConfirmPricesRequest struct {
	// Идентификаторы товаров, у которых подтверждается цена.
	OfferIds []string `json:"offerIds"`
}

type _ConfirmPricesRequest ConfirmPricesRequest

// NewConfirmPricesRequest instantiates a new ConfirmPricesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmPricesRequest(offerIds []string) *ConfirmPricesRequest {
	this := ConfirmPricesRequest{}
	this.OfferIds = offerIds
	return &this
}

// NewConfirmPricesRequestWithDefaults instantiates a new ConfirmPricesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmPricesRequestWithDefaults() *ConfirmPricesRequest {
	this := ConfirmPricesRequest{}
	return &this
}

// GetOfferIds returns the OfferIds field value
func (o *ConfirmPricesRequest) GetOfferIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OfferIds
}

// GetOfferIdsOk returns a tuple with the OfferIds field value
// and a boolean to check if the value has been set.
func (o *ConfirmPricesRequest) GetOfferIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OfferIds, true
}

// SetOfferIds sets field value
func (o *ConfirmPricesRequest) SetOfferIds(v []string) {
	o.OfferIds = v
}

func (o ConfirmPricesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfirmPricesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offerIds"] = o.OfferIds
	return toSerialize, nil
}

func (o *ConfirmPricesRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"offerIds",
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

	varConfirmPricesRequest := _ConfirmPricesRequest{}

	err = json.Unmarshal(bytes, &varConfirmPricesRequest)

	if err != nil {
		return err
	}

	*o = ConfirmPricesRequest(varConfirmPricesRequest)

	return err
}

type NullableConfirmPricesRequest struct {
	value *ConfirmPricesRequest
	isSet bool
}

func (v NullableConfirmPricesRequest) Get() *ConfirmPricesRequest {
	return v.value
}

func (v *NullableConfirmPricesRequest) Set(val *ConfirmPricesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmPricesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmPricesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmPricesRequest(val *ConfirmPricesRequest) *NullableConfirmPricesRequest {
	return &NullableConfirmPricesRequest{value: val, isSet: true}
}

func (v NullableConfirmPricesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfirmPricesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


