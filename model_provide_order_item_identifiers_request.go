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

// checks if the ProvideOrderItemIdentifiersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvideOrderItemIdentifiersRequest{}

// ProvideOrderItemIdentifiersRequest struct for ProvideOrderItemIdentifiersRequest
type ProvideOrderItemIdentifiersRequest struct {
	// Список позиций, требующих маркировки. 
	Items []OrderItemInstanceModificationDTO `json:"items"`
}

type _ProvideOrderItemIdentifiersRequest ProvideOrderItemIdentifiersRequest

// NewProvideOrderItemIdentifiersRequest instantiates a new ProvideOrderItemIdentifiersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvideOrderItemIdentifiersRequest(items []OrderItemInstanceModificationDTO) *ProvideOrderItemIdentifiersRequest {
	this := ProvideOrderItemIdentifiersRequest{}
	this.Items = items
	return &this
}

// NewProvideOrderItemIdentifiersRequestWithDefaults instantiates a new ProvideOrderItemIdentifiersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvideOrderItemIdentifiersRequestWithDefaults() *ProvideOrderItemIdentifiersRequest {
	this := ProvideOrderItemIdentifiersRequest{}
	return &this
}

// GetItems returns the Items field value
func (o *ProvideOrderItemIdentifiersRequest) GetItems() []OrderItemInstanceModificationDTO {
	if o == nil {
		var ret []OrderItemInstanceModificationDTO
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ProvideOrderItemIdentifiersRequest) GetItemsOk() ([]OrderItemInstanceModificationDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ProvideOrderItemIdentifiersRequest) SetItems(v []OrderItemInstanceModificationDTO) {
	o.Items = v
}

func (o ProvideOrderItemIdentifiersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvideOrderItemIdentifiersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *ProvideOrderItemIdentifiersRequest) UnmarshalJSON(bytes []byte) (err error) {
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

	varProvideOrderItemIdentifiersRequest := _ProvideOrderItemIdentifiersRequest{}

	err = json.Unmarshal(bytes, &varProvideOrderItemIdentifiersRequest)

	if err != nil {
		return err
	}

	*o = ProvideOrderItemIdentifiersRequest(varProvideOrderItemIdentifiersRequest)

	return err
}

type NullableProvideOrderItemIdentifiersRequest struct {
	value *ProvideOrderItemIdentifiersRequest
	isSet bool
}

func (v NullableProvideOrderItemIdentifiersRequest) Get() *ProvideOrderItemIdentifiersRequest {
	return v.value
}

func (v *NullableProvideOrderItemIdentifiersRequest) Set(val *ProvideOrderItemIdentifiersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProvideOrderItemIdentifiersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProvideOrderItemIdentifiersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvideOrderItemIdentifiersRequest(val *ProvideOrderItemIdentifiersRequest) *NullableProvideOrderItemIdentifiersRequest {
	return &NullableProvideOrderItemIdentifiersRequest{value: val, isSet: true}
}

func (v NullableProvideOrderItemIdentifiersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvideOrderItemIdentifiersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


