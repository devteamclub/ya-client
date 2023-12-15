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

// checks if the OrdersStatsDetailsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrdersStatsDetailsDTO{}

// OrdersStatsDetailsDTO Информация об удалении товара из заказа.
type OrdersStatsDetailsDTO struct {
	ItemStatus *OrdersStatsItemStatusType `json:"itemStatus,omitempty"`
	// Количество товара со статусом, указанном в параметре `itemStatus`.
	ItemCount *int64 `json:"itemCount,omitempty"`
	// Дата, когда товар получил статус, указанный в параметре `itemStatus`.  Формат даты: `ГГГГ-ММ-ДД`. 
	UpdateDate *string `json:"updateDate,omitempty"`
	StockType *OrdersStatsStockType `json:"stockType,omitempty"`
}

// NewOrdersStatsDetailsDTO instantiates a new OrdersStatsDetailsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdersStatsDetailsDTO() *OrdersStatsDetailsDTO {
	this := OrdersStatsDetailsDTO{}
	return &this
}

// NewOrdersStatsDetailsDTOWithDefaults instantiates a new OrdersStatsDetailsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdersStatsDetailsDTOWithDefaults() *OrdersStatsDetailsDTO {
	this := OrdersStatsDetailsDTO{}
	return &this
}

// GetItemStatus returns the ItemStatus field value if set, zero value otherwise.
func (o *OrdersStatsDetailsDTO) GetItemStatus() OrdersStatsItemStatusType {
	if o == nil || IsNil(o.ItemStatus) {
		var ret OrdersStatsItemStatusType
		return ret
	}
	return *o.ItemStatus
}

// GetItemStatusOk returns a tuple with the ItemStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsDetailsDTO) GetItemStatusOk() (*OrdersStatsItemStatusType, bool) {
	if o == nil || IsNil(o.ItemStatus) {
		return nil, false
	}
	return o.ItemStatus, true
}

// HasItemStatus returns a boolean if a field has been set.
func (o *OrdersStatsDetailsDTO) HasItemStatus() bool {
	if o != nil && !IsNil(o.ItemStatus) {
		return true
	}

	return false
}

// SetItemStatus gets a reference to the given OrdersStatsItemStatusType and assigns it to the ItemStatus field.
func (o *OrdersStatsDetailsDTO) SetItemStatus(v OrdersStatsItemStatusType) {
	o.ItemStatus = &v
}

// GetItemCount returns the ItemCount field value if set, zero value otherwise.
func (o *OrdersStatsDetailsDTO) GetItemCount() int64 {
	if o == nil || IsNil(o.ItemCount) {
		var ret int64
		return ret
	}
	return *o.ItemCount
}

// GetItemCountOk returns a tuple with the ItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsDetailsDTO) GetItemCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ItemCount) {
		return nil, false
	}
	return o.ItemCount, true
}

// HasItemCount returns a boolean if a field has been set.
func (o *OrdersStatsDetailsDTO) HasItemCount() bool {
	if o != nil && !IsNil(o.ItemCount) {
		return true
	}

	return false
}

// SetItemCount gets a reference to the given int64 and assigns it to the ItemCount field.
func (o *OrdersStatsDetailsDTO) SetItemCount(v int64) {
	o.ItemCount = &v
}

// GetUpdateDate returns the UpdateDate field value if set, zero value otherwise.
func (o *OrdersStatsDetailsDTO) GetUpdateDate() string {
	if o == nil || IsNil(o.UpdateDate) {
		var ret string
		return ret
	}
	return *o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsDetailsDTO) GetUpdateDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateDate) {
		return nil, false
	}
	return o.UpdateDate, true
}

// HasUpdateDate returns a boolean if a field has been set.
func (o *OrdersStatsDetailsDTO) HasUpdateDate() bool {
	if o != nil && !IsNil(o.UpdateDate) {
		return true
	}

	return false
}

// SetUpdateDate gets a reference to the given string and assigns it to the UpdateDate field.
func (o *OrdersStatsDetailsDTO) SetUpdateDate(v string) {
	o.UpdateDate = &v
}

// GetStockType returns the StockType field value if set, zero value otherwise.
func (o *OrdersStatsDetailsDTO) GetStockType() OrdersStatsStockType {
	if o == nil || IsNil(o.StockType) {
		var ret OrdersStatsStockType
		return ret
	}
	return *o.StockType
}

// GetStockTypeOk returns a tuple with the StockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsDetailsDTO) GetStockTypeOk() (*OrdersStatsStockType, bool) {
	if o == nil || IsNil(o.StockType) {
		return nil, false
	}
	return o.StockType, true
}

// HasStockType returns a boolean if a field has been set.
func (o *OrdersStatsDetailsDTO) HasStockType() bool {
	if o != nil && !IsNil(o.StockType) {
		return true
	}

	return false
}

// SetStockType gets a reference to the given OrdersStatsStockType and assigns it to the StockType field.
func (o *OrdersStatsDetailsDTO) SetStockType(v OrdersStatsStockType) {
	o.StockType = &v
}

func (o OrdersStatsDetailsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrdersStatsDetailsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemStatus) {
		toSerialize["itemStatus"] = o.ItemStatus
	}
	if !IsNil(o.ItemCount) {
		toSerialize["itemCount"] = o.ItemCount
	}
	if !IsNil(o.UpdateDate) {
		toSerialize["updateDate"] = o.UpdateDate
	}
	if !IsNil(o.StockType) {
		toSerialize["stockType"] = o.StockType
	}
	return toSerialize, nil
}

type NullableOrdersStatsDetailsDTO struct {
	value *OrdersStatsDetailsDTO
	isSet bool
}

func (v NullableOrdersStatsDetailsDTO) Get() *OrdersStatsDetailsDTO {
	return v.value
}

func (v *NullableOrdersStatsDetailsDTO) Set(val *OrdersStatsDetailsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersStatsDetailsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersStatsDetailsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersStatsDetailsDTO(val *OrdersStatsDetailsDTO) *NullableOrdersStatsDetailsDTO {
	return &NullableOrdersStatsDetailsDTO{value: val, isSet: true}
}

func (v NullableOrdersStatsDetailsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersStatsDetailsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


