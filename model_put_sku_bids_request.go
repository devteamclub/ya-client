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

// checks if the PutSkuBidsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutSkuBidsRequest{}

// PutSkuBidsRequest description.
type PutSkuBidsRequest struct {
	// Список товаров и ставки для продвижения, которые на них нужно установить.
	Bids []SkuBidItemDTO `json:"bids"`
}

type _PutSkuBidsRequest PutSkuBidsRequest

// NewPutSkuBidsRequest instantiates a new PutSkuBidsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutSkuBidsRequest(bids []SkuBidItemDTO) *PutSkuBidsRequest {
	this := PutSkuBidsRequest{}
	this.Bids = bids
	return &this
}

// NewPutSkuBidsRequestWithDefaults instantiates a new PutSkuBidsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutSkuBidsRequestWithDefaults() *PutSkuBidsRequest {
	this := PutSkuBidsRequest{}
	return &this
}

// GetBids returns the Bids field value
func (o *PutSkuBidsRequest) GetBids() []SkuBidItemDTO {
	if o == nil {
		var ret []SkuBidItemDTO
		return ret
	}

	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value
// and a boolean to check if the value has been set.
func (o *PutSkuBidsRequest) GetBidsOk() ([]SkuBidItemDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bids, true
}

// SetBids sets field value
func (o *PutSkuBidsRequest) SetBids(v []SkuBidItemDTO) {
	o.Bids = v
}

func (o PutSkuBidsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutSkuBidsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bids"] = o.Bids
	return toSerialize, nil
}

func (o *PutSkuBidsRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bids",
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

	varPutSkuBidsRequest := _PutSkuBidsRequest{}

	err = json.Unmarshal(bytes, &varPutSkuBidsRequest)

	if err != nil {
		return err
	}

	*o = PutSkuBidsRequest(varPutSkuBidsRequest)

	return err
}

type NullablePutSkuBidsRequest struct {
	value *PutSkuBidsRequest
	isSet bool
}

func (v NullablePutSkuBidsRequest) Get() *PutSkuBidsRequest {
	return v.value
}

func (v *NullablePutSkuBidsRequest) Set(val *PutSkuBidsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutSkuBidsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutSkuBidsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutSkuBidsRequest(val *PutSkuBidsRequest) *NullablePutSkuBidsRequest {
	return &NullablePutSkuBidsRequest{value: val, isSet: true}
}

func (v NullablePutSkuBidsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutSkuBidsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


