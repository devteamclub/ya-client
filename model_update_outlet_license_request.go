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

// checks if the UpdateOutletLicenseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOutletLicenseRequest{}

// UpdateOutletLicenseRequest Запрос на создание или изменение лицензий для точек продаж.
type UpdateOutletLicenseRequest struct {
	// Список лицензий. Обязательный параметр, должен содержать информацию хотя бы об одной лицензии. 
	Licenses []OutletLicenseDTO `json:"licenses"`
}

type _UpdateOutletLicenseRequest UpdateOutletLicenseRequest

// NewUpdateOutletLicenseRequest instantiates a new UpdateOutletLicenseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOutletLicenseRequest(licenses []OutletLicenseDTO) *UpdateOutletLicenseRequest {
	this := UpdateOutletLicenseRequest{}
	this.Licenses = licenses
	return &this
}

// NewUpdateOutletLicenseRequestWithDefaults instantiates a new UpdateOutletLicenseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOutletLicenseRequestWithDefaults() *UpdateOutletLicenseRequest {
	this := UpdateOutletLicenseRequest{}
	return &this
}

// GetLicenses returns the Licenses field value
func (o *UpdateOutletLicenseRequest) GetLicenses() []OutletLicenseDTO {
	if o == nil {
		var ret []OutletLicenseDTO
		return ret
	}

	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *UpdateOutletLicenseRequest) GetLicensesOk() ([]OutletLicenseDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Licenses, true
}

// SetLicenses sets field value
func (o *UpdateOutletLicenseRequest) SetLicenses(v []OutletLicenseDTO) {
	o.Licenses = v
}

func (o UpdateOutletLicenseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOutletLicenseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["licenses"] = o.Licenses
	return toSerialize, nil
}

func (o *UpdateOutletLicenseRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"licenses",
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

	varUpdateOutletLicenseRequest := _UpdateOutletLicenseRequest{}

	err = json.Unmarshal(bytes, &varUpdateOutletLicenseRequest)

	if err != nil {
		return err
	}

	*o = UpdateOutletLicenseRequest(varUpdateOutletLicenseRequest)

	return err
}

type NullableUpdateOutletLicenseRequest struct {
	value *UpdateOutletLicenseRequest
	isSet bool
}

func (v NullableUpdateOutletLicenseRequest) Get() *UpdateOutletLicenseRequest {
	return v.value
}

func (v *NullableUpdateOutletLicenseRequest) Set(val *UpdateOutletLicenseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOutletLicenseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOutletLicenseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOutletLicenseRequest(val *UpdateOutletLicenseRequest) *NullableUpdateOutletLicenseRequest {
	return &NullableUpdateOutletLicenseRequest{value: val, isSet: true}
}

func (v NullableUpdateOutletLicenseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOutletLicenseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

