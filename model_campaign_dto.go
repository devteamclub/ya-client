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

// checks if the CampaignDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignDTO{}

// CampaignDTO Информация о магазине.
type CampaignDTO struct {
	// URL магазина.
	Domain *string `json:"domain,omitempty"`
	// Идентификатор кампании.
	Id *int64 `json:"id,omitempty"`
	// Идентификатор плательщика в Яндекс Балансе.
	ClientId *int64 `json:"clientId,omitempty"`
	Business *BusinessDTO `json:"business,omitempty"`
	PlacementType *PlacementType `json:"placementType,omitempty"`
}

// NewCampaignDTO instantiates a new CampaignDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignDTO() *CampaignDTO {
	this := CampaignDTO{}
	return &this
}

// NewCampaignDTOWithDefaults instantiates a new CampaignDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignDTOWithDefaults() *CampaignDTO {
	this := CampaignDTO{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *CampaignDTO) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDTO) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *CampaignDTO) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *CampaignDTO) SetDomain(v string) {
	o.Domain = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CampaignDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CampaignDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CampaignDTO) SetId(v int64) {
	o.Id = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CampaignDTO) GetClientId() int64 {
	if o == nil || IsNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDTO) GetClientIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CampaignDTO) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *CampaignDTO) SetClientId(v int64) {
	o.ClientId = &v
}

// GetBusiness returns the Business field value if set, zero value otherwise.
func (o *CampaignDTO) GetBusiness() BusinessDTO {
	if o == nil || IsNil(o.Business) {
		var ret BusinessDTO
		return ret
	}
	return *o.Business
}

// GetBusinessOk returns a tuple with the Business field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDTO) GetBusinessOk() (*BusinessDTO, bool) {
	if o == nil || IsNil(o.Business) {
		return nil, false
	}
	return o.Business, true
}

// HasBusiness returns a boolean if a field has been set.
func (o *CampaignDTO) HasBusiness() bool {
	if o != nil && !IsNil(o.Business) {
		return true
	}

	return false
}

// SetBusiness gets a reference to the given BusinessDTO and assigns it to the Business field.
func (o *CampaignDTO) SetBusiness(v BusinessDTO) {
	o.Business = &v
}

// GetPlacementType returns the PlacementType field value if set, zero value otherwise.
func (o *CampaignDTO) GetPlacementType() PlacementType {
	if o == nil || IsNil(o.PlacementType) {
		var ret PlacementType
		return ret
	}
	return *o.PlacementType
}

// GetPlacementTypeOk returns a tuple with the PlacementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDTO) GetPlacementTypeOk() (*PlacementType, bool) {
	if o == nil || IsNil(o.PlacementType) {
		return nil, false
	}
	return o.PlacementType, true
}

// HasPlacementType returns a boolean if a field has been set.
func (o *CampaignDTO) HasPlacementType() bool {
	if o != nil && !IsNil(o.PlacementType) {
		return true
	}

	return false
}

// SetPlacementType gets a reference to the given PlacementType and assigns it to the PlacementType field.
func (o *CampaignDTO) SetPlacementType(v PlacementType) {
	o.PlacementType = &v
}

func (o CampaignDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.Business) {
		toSerialize["business"] = o.Business
	}
	if !IsNil(o.PlacementType) {
		toSerialize["placementType"] = o.PlacementType
	}
	return toSerialize, nil
}

type NullableCampaignDTO struct {
	value *CampaignDTO
	isSet bool
}

func (v NullableCampaignDTO) Get() *CampaignDTO {
	return v.value
}

func (v *NullableCampaignDTO) Set(val *CampaignDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignDTO(val *CampaignDTO) *NullableCampaignDTO {
	return &NullableCampaignDTO{value: val, isSet: true}
}

func (v NullableCampaignDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


