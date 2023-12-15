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

// checks if the GenerateGoodsRealizationReportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateGoodsRealizationReportRequest{}

// GenerateGoodsRealizationReportRequest Данные, необходимые для генерации отчета: идентификатор магазина и период, за который нужен отчет. 
type GenerateGoodsRealizationReportRequest struct {
	// Идентификатор кампании.
	CampaignId int64 `json:"campaignId"`
	// Год.
	Year int32 `json:"year"`
	// Номер месяца.
	Month int32 `json:"month"`
}

type _GenerateGoodsRealizationReportRequest GenerateGoodsRealizationReportRequest

// NewGenerateGoodsRealizationReportRequest instantiates a new GenerateGoodsRealizationReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateGoodsRealizationReportRequest(campaignId int64, year int32, month int32) *GenerateGoodsRealizationReportRequest {
	this := GenerateGoodsRealizationReportRequest{}
	this.CampaignId = campaignId
	this.Year = year
	this.Month = month
	return &this
}

// NewGenerateGoodsRealizationReportRequestWithDefaults instantiates a new GenerateGoodsRealizationReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateGoodsRealizationReportRequestWithDefaults() *GenerateGoodsRealizationReportRequest {
	this := GenerateGoodsRealizationReportRequest{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *GenerateGoodsRealizationReportRequest) GetCampaignId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *GenerateGoodsRealizationReportRequest) GetCampaignIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *GenerateGoodsRealizationReportRequest) SetCampaignId(v int64) {
	o.CampaignId = v
}

// GetYear returns the Year field value
func (o *GenerateGoodsRealizationReportRequest) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *GenerateGoodsRealizationReportRequest) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *GenerateGoodsRealizationReportRequest) SetYear(v int32) {
	o.Year = v
}

// GetMonth returns the Month field value
func (o *GenerateGoodsRealizationReportRequest) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *GenerateGoodsRealizationReportRequest) GetMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *GenerateGoodsRealizationReportRequest) SetMonth(v int32) {
	o.Month = v
}

func (o GenerateGoodsRealizationReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateGoodsRealizationReportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["year"] = o.Year
	toSerialize["month"] = o.Month
	return toSerialize, nil
}

func (o *GenerateGoodsRealizationReportRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"campaignId",
		"year",
		"month",
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

	varGenerateGoodsRealizationReportRequest := _GenerateGoodsRealizationReportRequest{}

	err = json.Unmarshal(bytes, &varGenerateGoodsRealizationReportRequest)

	if err != nil {
		return err
	}

	*o = GenerateGoodsRealizationReportRequest(varGenerateGoodsRealizationReportRequest)

	return err
}

type NullableGenerateGoodsRealizationReportRequest struct {
	value *GenerateGoodsRealizationReportRequest
	isSet bool
}

func (v NullableGenerateGoodsRealizationReportRequest) Get() *GenerateGoodsRealizationReportRequest {
	return v.value
}

func (v *NullableGenerateGoodsRealizationReportRequest) Set(val *GenerateGoodsRealizationReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateGoodsRealizationReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateGoodsRealizationReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateGoodsRealizationReportRequest(val *GenerateGoodsRealizationReportRequest) *NullableGenerateGoodsRealizationReportRequest {
	return &NullableGenerateGoodsRealizationReportRequest{value: val, isSet: true}
}

func (v NullableGenerateGoodsRealizationReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateGoodsRealizationReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


