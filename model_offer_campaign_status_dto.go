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

// checks if the OfferCampaignStatusDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferCampaignStatusDTO{}

// OfferCampaignStatusDTO Статус товара в магазине.
type OfferCampaignStatusDTO struct {
	// Идентификатор кампании. 
	CampaignId int64 `json:"campaignId"`
	Status OfferCampaignStatusType `json:"status"`
}

type _OfferCampaignStatusDTO OfferCampaignStatusDTO

// NewOfferCampaignStatusDTO instantiates a new OfferCampaignStatusDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferCampaignStatusDTO(campaignId int64, status OfferCampaignStatusType) *OfferCampaignStatusDTO {
	this := OfferCampaignStatusDTO{}
	this.CampaignId = campaignId
	this.Status = status
	return &this
}

// NewOfferCampaignStatusDTOWithDefaults instantiates a new OfferCampaignStatusDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferCampaignStatusDTOWithDefaults() *OfferCampaignStatusDTO {
	this := OfferCampaignStatusDTO{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *OfferCampaignStatusDTO) GetCampaignId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *OfferCampaignStatusDTO) GetCampaignIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *OfferCampaignStatusDTO) SetCampaignId(v int64) {
	o.CampaignId = v
}

// GetStatus returns the Status field value
func (o *OfferCampaignStatusDTO) GetStatus() OfferCampaignStatusType {
	if o == nil {
		var ret OfferCampaignStatusType
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OfferCampaignStatusDTO) GetStatusOk() (*OfferCampaignStatusType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OfferCampaignStatusDTO) SetStatus(v OfferCampaignStatusType) {
	o.Status = v
}

func (o OfferCampaignStatusDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferCampaignStatusDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *OfferCampaignStatusDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"campaignId",
		"status",
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

	varOfferCampaignStatusDTO := _OfferCampaignStatusDTO{}

	err = json.Unmarshal(bytes, &varOfferCampaignStatusDTO)

	if err != nil {
		return err
	}

	*o = OfferCampaignStatusDTO(varOfferCampaignStatusDTO)

	return err
}

type NullableOfferCampaignStatusDTO struct {
	value *OfferCampaignStatusDTO
	isSet bool
}

func (v NullableOfferCampaignStatusDTO) Get() *OfferCampaignStatusDTO {
	return v.value
}

func (v *NullableOfferCampaignStatusDTO) Set(val *OfferCampaignStatusDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferCampaignStatusDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferCampaignStatusDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferCampaignStatusDTO(val *OfferCampaignStatusDTO) *NullableOfferCampaignStatusDTO {
	return &NullableOfferCampaignStatusDTO{value: val, isSet: true}
}

func (v NullableOfferCampaignStatusDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferCampaignStatusDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


