/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetQuarantineOffersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetQuarantineOffersRequest{}

// GetQuarantineOffersRequest Фильтрации товаров  В запросе можно указать либо фильтр offerIds, либо любые другие фильтры товаров. Совместное использование фильтра offerIds с другими фильтрациями приведет к ошибке. 
type GetQuarantineOffersRequest struct {
	// Идентификаторы товаров, информация о которых нужна. ⚠️ Не используйте это поле одновременно с фильтрами по статусам карточек, категориям, брендам или тегам. Если вы хотите воспользоваться фильтрами, оставьте поле пустым.
	OfferIds []string `json:"offerIds,omitempty"`
	// Фильтр по статусам карточек.  [Что такое карточка товара](https://yandex.ru/support/marketplace/assortment/content/index.html) 
	CardStatuses []OfferCardStatusType `json:"cardStatuses,omitempty"`
	// Фильтр по категориям на Маркете.
	CategoryIds []int32 `json:"categoryIds,omitempty"`
	// Фильтр по брендам.
	VendorNames []string `json:"vendorNames,omitempty"`
	// Фильтр по тегам.
	Tags []string `json:"tags,omitempty"`
}

// NewGetQuarantineOffersRequest instantiates a new GetQuarantineOffersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetQuarantineOffersRequest() *GetQuarantineOffersRequest {
	this := GetQuarantineOffersRequest{}
	return &this
}

// NewGetQuarantineOffersRequestWithDefaults instantiates a new GetQuarantineOffersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetQuarantineOffersRequestWithDefaults() *GetQuarantineOffersRequest {
	this := GetQuarantineOffersRequest{}
	return &this
}

// GetOfferIds returns the OfferIds field value if set, zero value otherwise.
func (o *GetQuarantineOffersRequest) GetOfferIds() []string {
	if o == nil || IsNil(o.OfferIds) {
		var ret []string
		return ret
	}
	return o.OfferIds
}

// GetOfferIdsOk returns a tuple with the OfferIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuarantineOffersRequest) GetOfferIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OfferIds) {
		return nil, false
	}
	return o.OfferIds, true
}

// HasOfferIds returns a boolean if a field has been set.
func (o *GetQuarantineOffersRequest) HasOfferIds() bool {
	if o != nil && !IsNil(o.OfferIds) {
		return true
	}

	return false
}

// SetOfferIds gets a reference to the given []string and assigns it to the OfferIds field.
func (o *GetQuarantineOffersRequest) SetOfferIds(v []string) {
	o.OfferIds = v
}

// GetCardStatuses returns the CardStatuses field value if set, zero value otherwise.
func (o *GetQuarantineOffersRequest) GetCardStatuses() []OfferCardStatusType {
	if o == nil || IsNil(o.CardStatuses) {
		var ret []OfferCardStatusType
		return ret
	}
	return o.CardStatuses
}

// GetCardStatusesOk returns a tuple with the CardStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuarantineOffersRequest) GetCardStatusesOk() ([]OfferCardStatusType, bool) {
	if o == nil || IsNil(o.CardStatuses) {
		return nil, false
	}
	return o.CardStatuses, true
}

// HasCardStatuses returns a boolean if a field has been set.
func (o *GetQuarantineOffersRequest) HasCardStatuses() bool {
	if o != nil && !IsNil(o.CardStatuses) {
		return true
	}

	return false
}

// SetCardStatuses gets a reference to the given []OfferCardStatusType and assigns it to the CardStatuses field.
func (o *GetQuarantineOffersRequest) SetCardStatuses(v []OfferCardStatusType) {
	o.CardStatuses = v
}

// GetCategoryIds returns the CategoryIds field value if set, zero value otherwise.
func (o *GetQuarantineOffersRequest) GetCategoryIds() []int32 {
	if o == nil || IsNil(o.CategoryIds) {
		var ret []int32
		return ret
	}
	return o.CategoryIds
}

// GetCategoryIdsOk returns a tuple with the CategoryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuarantineOffersRequest) GetCategoryIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.CategoryIds) {
		return nil, false
	}
	return o.CategoryIds, true
}

// HasCategoryIds returns a boolean if a field has been set.
func (o *GetQuarantineOffersRequest) HasCategoryIds() bool {
	if o != nil && !IsNil(o.CategoryIds) {
		return true
	}

	return false
}

// SetCategoryIds gets a reference to the given []int32 and assigns it to the CategoryIds field.
func (o *GetQuarantineOffersRequest) SetCategoryIds(v []int32) {
	o.CategoryIds = v
}

// GetVendorNames returns the VendorNames field value if set, zero value otherwise.
func (o *GetQuarantineOffersRequest) GetVendorNames() []string {
	if o == nil || IsNil(o.VendorNames) {
		var ret []string
		return ret
	}
	return o.VendorNames
}

// GetVendorNamesOk returns a tuple with the VendorNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuarantineOffersRequest) GetVendorNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.VendorNames) {
		return nil, false
	}
	return o.VendorNames, true
}

// HasVendorNames returns a boolean if a field has been set.
func (o *GetQuarantineOffersRequest) HasVendorNames() bool {
	if o != nil && !IsNil(o.VendorNames) {
		return true
	}

	return false
}

// SetVendorNames gets a reference to the given []string and assigns it to the VendorNames field.
func (o *GetQuarantineOffersRequest) SetVendorNames(v []string) {
	o.VendorNames = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetQuarantineOffersRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuarantineOffersRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetQuarantineOffersRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetQuarantineOffersRequest) SetTags(v []string) {
	o.Tags = v
}

func (o GetQuarantineOffersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetQuarantineOffersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OfferIds) {
		toSerialize["offerIds"] = o.OfferIds
	}
	if !IsNil(o.CardStatuses) {
		toSerialize["cardStatuses"] = o.CardStatuses
	}
	if !IsNil(o.CategoryIds) {
		toSerialize["categoryIds"] = o.CategoryIds
	}
	if !IsNil(o.VendorNames) {
		toSerialize["vendorNames"] = o.VendorNames
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableGetQuarantineOffersRequest struct {
	value *GetQuarantineOffersRequest
	isSet bool
}

func (v NullableGetQuarantineOffersRequest) Get() *GetQuarantineOffersRequest {
	return v.value
}

func (v *NullableGetQuarantineOffersRequest) Set(val *GetQuarantineOffersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQuarantineOffersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQuarantineOffersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQuarantineOffersRequest(val *GetQuarantineOffersRequest) *NullableGetQuarantineOffersRequest {
	return &NullableGetQuarantineOffersRequest{value: val, isSet: true}
}

func (v NullableGetQuarantineOffersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQuarantineOffersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


