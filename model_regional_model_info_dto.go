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

// checks if the RegionalModelInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionalModelInfoDTO{}

// RegionalModelInfoDTO Региональная информация.
type RegionalModelInfoDTO struct {
	Currency *CurrencyType `json:"currency,omitempty"`
	// Идентификатор региона, для которого выводится информация о предложениях модели (доставляемых в этот регион).  Информацию о регионе по идентификатору можно получить с помощью запроса `GET /regions/{regionId}`. 
	RegionId *int64 `json:"regionId,omitempty"`
}

// NewRegionalModelInfoDTO instantiates a new RegionalModelInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionalModelInfoDTO() *RegionalModelInfoDTO {
	this := RegionalModelInfoDTO{}
	return &this
}

// NewRegionalModelInfoDTOWithDefaults instantiates a new RegionalModelInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionalModelInfoDTOWithDefaults() *RegionalModelInfoDTO {
	this := RegionalModelInfoDTO{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *RegionalModelInfoDTO) GetCurrency() CurrencyType {
	if o == nil || IsNil(o.Currency) {
		var ret CurrencyType
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalModelInfoDTO) GetCurrencyOk() (*CurrencyType, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *RegionalModelInfoDTO) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CurrencyType and assigns it to the Currency field.
func (o *RegionalModelInfoDTO) SetCurrency(v CurrencyType) {
	o.Currency = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *RegionalModelInfoDTO) GetRegionId() int64 {
	if o == nil || IsNil(o.RegionId) {
		var ret int64
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalModelInfoDTO) GetRegionIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RegionId) {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *RegionalModelInfoDTO) HasRegionId() bool {
	if o != nil && !IsNil(o.RegionId) {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given int64 and assigns it to the RegionId field.
func (o *RegionalModelInfoDTO) SetRegionId(v int64) {
	o.RegionId = &v
}

func (o RegionalModelInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionalModelInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.RegionId) {
		toSerialize["regionId"] = o.RegionId
	}
	return toSerialize, nil
}

type NullableRegionalModelInfoDTO struct {
	value *RegionalModelInfoDTO
	isSet bool
}

func (v NullableRegionalModelInfoDTO) Get() *RegionalModelInfoDTO {
	return v.value
}

func (v *NullableRegionalModelInfoDTO) Set(val *RegionalModelInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionalModelInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionalModelInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionalModelInfoDTO(val *RegionalModelInfoDTO) *NullableRegionalModelInfoDTO {
	return &NullableRegionalModelInfoDTO{value: val, isSet: true}
}

func (v NullableRegionalModelInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionalModelInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


