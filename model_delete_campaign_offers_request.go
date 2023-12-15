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

// checks if the DeleteCampaignOffersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteCampaignOffersRequest{}

// DeleteCampaignOffersRequest Фильтрации удаляемых товаров по offerIds. 
type DeleteCampaignOffersRequest struct {
	// Идентификаторы товаров в каталоге.
	OfferIds []string `json:"offerIds"`
}

type _DeleteCampaignOffersRequest DeleteCampaignOffersRequest

// NewDeleteCampaignOffersRequest instantiates a new DeleteCampaignOffersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCampaignOffersRequest(offerIds []string) *DeleteCampaignOffersRequest {
	this := DeleteCampaignOffersRequest{}
	this.OfferIds = offerIds
	return &this
}

// NewDeleteCampaignOffersRequestWithDefaults instantiates a new DeleteCampaignOffersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCampaignOffersRequestWithDefaults() *DeleteCampaignOffersRequest {
	this := DeleteCampaignOffersRequest{}
	return &this
}

// GetOfferIds returns the OfferIds field value
func (o *DeleteCampaignOffersRequest) GetOfferIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OfferIds
}

// GetOfferIdsOk returns a tuple with the OfferIds field value
// and a boolean to check if the value has been set.
func (o *DeleteCampaignOffersRequest) GetOfferIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OfferIds, true
}

// SetOfferIds sets field value
func (o *DeleteCampaignOffersRequest) SetOfferIds(v []string) {
	o.OfferIds = v
}

func (o DeleteCampaignOffersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteCampaignOffersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offerIds"] = o.OfferIds
	return toSerialize, nil
}

func (o *DeleteCampaignOffersRequest) UnmarshalJSON(bytes []byte) (err error) {
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

	varDeleteCampaignOffersRequest := _DeleteCampaignOffersRequest{}

	err = json.Unmarshal(bytes, &varDeleteCampaignOffersRequest)

	if err != nil {
		return err
	}

	*o = DeleteCampaignOffersRequest(varDeleteCampaignOffersRequest)

	return err
}

type NullableDeleteCampaignOffersRequest struct {
	value *DeleteCampaignOffersRequest
	isSet bool
}

func (v NullableDeleteCampaignOffersRequest) Get() *DeleteCampaignOffersRequest {
	return v.value
}

func (v *NullableDeleteCampaignOffersRequest) Set(val *DeleteCampaignOffersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCampaignOffersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCampaignOffersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCampaignOffersRequest(val *DeleteCampaignOffersRequest) *NullableDeleteCampaignOffersRequest {
	return &NullableDeleteCampaignOffersRequest{value: val, isSet: true}
}

func (v NullableDeleteCampaignOffersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCampaignOffersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


