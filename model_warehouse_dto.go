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

// checks if the WarehouseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarehouseDTO{}

// WarehouseDTO Информация о складе.
type WarehouseDTO struct {
	// Идентификатор склада.
	Id int64 `json:"id"`
	// Название склада.
	Name string `json:"name"`
	// Идентификатор кампании в API и идентификатор магазина.
	CampaignId int64 `json:"campaignId"`
}

type _WarehouseDTO WarehouseDTO

// NewWarehouseDTO instantiates a new WarehouseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouseDTO(id int64, name string, campaignId int64) *WarehouseDTO {
	this := WarehouseDTO{}
	this.Id = id
	this.Name = name
	this.CampaignId = campaignId
	return &this
}

// NewWarehouseDTOWithDefaults instantiates a new WarehouseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseDTOWithDefaults() *WarehouseDTO {
	this := WarehouseDTO{}
	return &this
}

// GetId returns the Id field value
func (o *WarehouseDTO) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WarehouseDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WarehouseDTO) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *WarehouseDTO) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WarehouseDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WarehouseDTO) SetName(v string) {
	o.Name = v
}

// GetCampaignId returns the CampaignId field value
func (o *WarehouseDTO) GetCampaignId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *WarehouseDTO) GetCampaignIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *WarehouseDTO) SetCampaignId(v int64) {
	o.CampaignId = v
}

func (o WarehouseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarehouseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["campaignId"] = o.CampaignId
	return toSerialize, nil
}

func (o *WarehouseDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"campaignId",
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

	varWarehouseDTO := _WarehouseDTO{}

	err = json.Unmarshal(bytes, &varWarehouseDTO)

	if err != nil {
		return err
	}

	*o = WarehouseDTO(varWarehouseDTO)

	return err
}

type NullableWarehouseDTO struct {
	value *WarehouseDTO
	isSet bool
}

func (v NullableWarehouseDTO) Get() *WarehouseDTO {
	return v.value
}

func (v *NullableWarehouseDTO) Set(val *WarehouseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouseDTO(val *WarehouseDTO) *NullableWarehouseDTO {
	return &NullableWarehouseDTO{value: val, isSet: true}
}

func (v NullableWarehouseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

