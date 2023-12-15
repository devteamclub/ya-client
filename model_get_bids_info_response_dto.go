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

// checks if the GetBidsInfoResponseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBidsInfoResponseDTO{}

// GetBidsInfoResponseDTO Список товаров с указанными ставками.
type GetBidsInfoResponseDTO struct {
	// Страница списка товаров.
	Bids []SkuBidItemDTO `json:"bids"`
	Paging *ForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type _GetBidsInfoResponseDTO GetBidsInfoResponseDTO

// NewGetBidsInfoResponseDTO instantiates a new GetBidsInfoResponseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBidsInfoResponseDTO(bids []SkuBidItemDTO) *GetBidsInfoResponseDTO {
	this := GetBidsInfoResponseDTO{}
	this.Bids = bids
	return &this
}

// NewGetBidsInfoResponseDTOWithDefaults instantiates a new GetBidsInfoResponseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBidsInfoResponseDTOWithDefaults() *GetBidsInfoResponseDTO {
	this := GetBidsInfoResponseDTO{}
	return &this
}

// GetBids returns the Bids field value
func (o *GetBidsInfoResponseDTO) GetBids() []SkuBidItemDTO {
	if o == nil {
		var ret []SkuBidItemDTO
		return ret
	}

	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value
// and a boolean to check if the value has been set.
func (o *GetBidsInfoResponseDTO) GetBidsOk() ([]SkuBidItemDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bids, true
}

// SetBids sets field value
func (o *GetBidsInfoResponseDTO) SetBids(v []SkuBidItemDTO) {
	o.Bids = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *GetBidsInfoResponseDTO) GetPaging() ForwardScrollingPagerDTO {
	if o == nil || IsNil(o.Paging) {
		var ret ForwardScrollingPagerDTO
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBidsInfoResponseDTO) GetPagingOk() (*ForwardScrollingPagerDTO, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *GetBidsInfoResponseDTO) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ForwardScrollingPagerDTO and assigns it to the Paging field.
func (o *GetBidsInfoResponseDTO) SetPaging(v ForwardScrollingPagerDTO) {
	o.Paging = &v
}

func (o GetBidsInfoResponseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBidsInfoResponseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bids"] = o.Bids
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

func (o *GetBidsInfoResponseDTO) UnmarshalJSON(bytes []byte) (err error) {
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

	varGetBidsInfoResponseDTO := _GetBidsInfoResponseDTO{}

	err = json.Unmarshal(bytes, &varGetBidsInfoResponseDTO)

	if err != nil {
		return err
	}

	*o = GetBidsInfoResponseDTO(varGetBidsInfoResponseDTO)

	return err
}

type NullableGetBidsInfoResponseDTO struct {
	value *GetBidsInfoResponseDTO
	isSet bool
}

func (v NullableGetBidsInfoResponseDTO) Get() *GetBidsInfoResponseDTO {
	return v.value
}

func (v *NullableGetBidsInfoResponseDTO) Set(val *GetBidsInfoResponseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBidsInfoResponseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBidsInfoResponseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBidsInfoResponseDTO(val *GetBidsInfoResponseDTO) *NullableGetBidsInfoResponseDTO {
	return &NullableGetBidsInfoResponseDTO{value: val, isSet: true}
}

func (v NullableGetBidsInfoResponseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBidsInfoResponseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


