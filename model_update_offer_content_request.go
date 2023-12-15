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

// checks if the UpdateOfferContentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOfferContentRequest{}

// UpdateOfferContentRequest Запрос на установку новых значений для параметров.
type UpdateOfferContentRequest struct {
	// Список товаров с указанными характеристиками.
	OffersContent []OfferContentDTO `json:"offersContent"`
}

type _UpdateOfferContentRequest UpdateOfferContentRequest

// NewUpdateOfferContentRequest instantiates a new UpdateOfferContentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOfferContentRequest(offersContent []OfferContentDTO) *UpdateOfferContentRequest {
	this := UpdateOfferContentRequest{}
	this.OffersContent = offersContent
	return &this
}

// NewUpdateOfferContentRequestWithDefaults instantiates a new UpdateOfferContentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOfferContentRequestWithDefaults() *UpdateOfferContentRequest {
	this := UpdateOfferContentRequest{}
	return &this
}

// GetOffersContent returns the OffersContent field value
func (o *UpdateOfferContentRequest) GetOffersContent() []OfferContentDTO {
	if o == nil {
		var ret []OfferContentDTO
		return ret
	}

	return o.OffersContent
}

// GetOffersContentOk returns a tuple with the OffersContent field value
// and a boolean to check if the value has been set.
func (o *UpdateOfferContentRequest) GetOffersContentOk() ([]OfferContentDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.OffersContent, true
}

// SetOffersContent sets field value
func (o *UpdateOfferContentRequest) SetOffersContent(v []OfferContentDTO) {
	o.OffersContent = v
}

func (o UpdateOfferContentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOfferContentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offersContent"] = o.OffersContent
	return toSerialize, nil
}

func (o *UpdateOfferContentRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"offersContent",
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

	varUpdateOfferContentRequest := _UpdateOfferContentRequest{}

	err = json.Unmarshal(bytes, &varUpdateOfferContentRequest)

	if err != nil {
		return err
	}

	*o = UpdateOfferContentRequest(varUpdateOfferContentRequest)

	return err
}

type NullableUpdateOfferContentRequest struct {
	value *UpdateOfferContentRequest
	isSet bool
}

func (v NullableUpdateOfferContentRequest) Get() *UpdateOfferContentRequest {
	return v.value
}

func (v *NullableUpdateOfferContentRequest) Set(val *UpdateOfferContentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOfferContentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOfferContentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOfferContentRequest(val *UpdateOfferContentRequest) *NullableUpdateOfferContentRequest {
	return &NullableUpdateOfferContentRequest{value: val, isSet: true}
}

func (v NullableUpdateOfferContentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOfferContentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

