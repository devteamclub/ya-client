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

// checks if the OrdersStatsItemDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrdersStatsItemDTO{}

// OrdersStatsItemDTO Список товаров в заказе после возможных изменений.  В ходе обработки заказа Маркет может удалить из него единицы товаров — при проблемах на складе или по инициативе пользователя.  * Если из заказа удалены все единицы товара, его не будет в списке `items` — только в списке `initialItems`.  * Если в заказе осталась хотя бы одна единица товара, он будет и в списке `items` (с уменьшенным количеством единиц `count`), и в списке `initialItems` (с первоначальным количеством единиц `initialCount`). 
type OrdersStatsItemDTO struct {
	// Название товара.
	OfferName *string `json:"offerName,omitempty"`
	// SKU на Маркете.
	MarketSku *int64 `json:"marketSku,omitempty"`
	//   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы `. , / \\ ( ) [ ] - = _`  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields). 
	ShopSku *string `json:"shopSku,omitempty"`
	// Количество единиц товара с учетом удаленных единиц.  Если из заказа удалены все единицы товара, он попадет только в список `initialItems`. 
	Count *int32 `json:"count,omitempty"`
	// Цена или скидки на товар.
	Prices []OrdersStatsPriceDTO `json:"prices,omitempty"`
	Warehouse *OrdersStatsWarehouseDTO `json:"warehouse,omitempty"`
	// Информация об удалении товара из заказа.
	Details []OrdersStatsDetailsDTO `json:"details,omitempty"`
	// Список кодов идентификации товара в системе «Честный ЗНАК».
	CisList []string `json:"cisList,omitempty"`
	// Первоначальное количество единиц товара.
	InitialCount *int32 `json:"initialCount,omitempty"`
	// Списанная ставка ближайшего конкурента.  Указывается в процентах от стоимости товара и умножается на 100. Например, ставка 5% обозначается как 500. 
	BidFee *int32 `json:"bidFee,omitempty"`
	// Порог для скидок с Маркетом на момент оформления заказа. [Что это такое?](https://yandex.ru/support/marketplace/marketing/smart-pricing.html#sponsored-discounts)  Указан в рублях. Точность — два знака после запятой. 
	CofinanceThreshold *float32 `json:"cofinanceThreshold,omitempty"`
	// Скидка с Маркетом. [Что это такое?](https://yandex.ru/support/marketplace/marketing/smart-pricing.html#sponsored-discounts)  Указана в рублях. Точность — два знака после запятой. 
	CofinanceValue *float32 `json:"cofinanceValue,omitempty"`
}

// NewOrdersStatsItemDTO instantiates a new OrdersStatsItemDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdersStatsItemDTO() *OrdersStatsItemDTO {
	this := OrdersStatsItemDTO{}
	return &this
}

// NewOrdersStatsItemDTOWithDefaults instantiates a new OrdersStatsItemDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdersStatsItemDTOWithDefaults() *OrdersStatsItemDTO {
	this := OrdersStatsItemDTO{}
	return &this
}

// GetOfferName returns the OfferName field value if set, zero value otherwise.
func (o *OrdersStatsItemDTO) GetOfferName() string {
	if o == nil || IsNil(o.OfferName) {
		var ret string
		return ret
	}
	return *o.OfferName
}

// GetOfferNameOk returns a tuple with the OfferName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsItemDTO) GetOfferNameOk() (*string, bool) {
	if o == nil || IsNil(o.OfferName) {
		return nil, false
	}
	return o.OfferName, true
}

// HasOfferName returns a boolean if a field has been set.
func (o *OrdersStatsItemDTO) HasOfferName() bool {
	if o != nil && !IsNil(o.OfferName) {
		return true
	}

	return false
}

// SetOfferName gets a reference to the given string and assigns it to the OfferName field.
func (o *OrdersStatsItemDTO) SetOfferName(v string) {
	o.OfferName = &v
}

// GetMarketSku returns the MarketSku field value if set, zero value otherwise.
func (o *OrdersStatsItemDTO) GetMarketSku() int64 {
	if o == nil || IsNil(o.MarketSku) {
		var ret int64
		return ret
	}
	return *o.MarketSku
}

// GetMarketSkuOk returns a tuple with the MarketSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsItemDTO) GetMarketSkuOk() (*int64, bool) {
	if o == nil || IsNil(o.MarketSku) {
		return nil, false
	}
	return o.MarketSku, true
}

// HasMarketSku returns a boolean if a field has been set.
func (o *OrdersStatsItemDTO) HasMarketSku() bool {
	if o != nil && !IsNil(o.MarketSku) {
		return true
	}

	return false
}

// SetMarketSku gets a reference to the given int64 and assigns it to the MarketSku field.
func (o *OrdersStatsItemDTO) SetMarketSku(v int64) {
	o.MarketSku = &v
}

// GetShopSku returns the ShopSku field value if set, zero value otherwise.
func (o *OrdersStatsItemDTO) GetShopSku() string {
	if o == nil || IsNil(o.ShopSku) {
		var ret string
		return ret
	}
	return *o.ShopSku
}

// GetShopSkuOk returns a tuple with the ShopSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsItemDTO) GetShopSkuOk() (*string, bool) {
	if o == nil || IsNil(o.ShopSku) {
		return nil, false
	}
	return o.ShopSku, true
}

// HasShopSku returns a boolean if a field has been set.
func (o *OrdersStatsItemDTO) HasShopSku() bool {
	if o != nil && !IsNil(o.ShopSku) {
		return true
	}

	return false
}

// SetShopSku gets a reference to the given string and assigns it to the ShopSku field.
func (o *OrdersStatsItemDTO) SetShopSku(v string) {
	o.ShopSku = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *OrdersStatsItemDTO) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsItemDTO) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *OrdersStatsItemDTO) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *OrdersStatsItemDTO) SetCount(v int32) {
	o.Count = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *OrdersStatsItemDTO) GetPrices() []OrdersStatsPriceDTO {
	if o == nil || IsNil(o.Prices) {
		var ret []OrdersStatsPriceDTO
		return ret
	}
	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsItemDTO) GetPricesOk() ([]OrdersStatsPriceDTO, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *OrdersStatsItemDTO) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []OrdersStatsPriceDTO and assigns it to the Prices field.
func (o *OrdersStatsItemDTO) SetPrices(v []OrdersStatsPriceDTO) {
	o.Prices = v
}

// GetWarehouse returns the Warehouse field value if set, zero value otherwise.
func (o *OrdersStatsItemDTO) GetWarehouse() OrdersStatsWarehouseDTO {
	if o == nil || IsNil(o.Warehouse) {
		var ret OrdersStatsWarehouseDTO
		return ret
	}
	return *o.Warehouse
}

// GetWarehouseOk returns a tuple with the Warehouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsItemDTO) GetWarehouseOk() (*OrdersStatsWarehouseDTO, bool) {
	if o == nil || IsNil(o.Warehouse) {
		return nil, false
	}
	return o.Warehouse, true
}

// HasWarehouse returns a boolean if a field has been set.
func (o *OrdersStatsItemDTO) HasWarehouse() bool {
	if o != nil && !IsNil(o.Warehouse) {
		return true
	}

	return false
}

// SetWarehouse gets a reference to the given OrdersStatsWarehouseDTO and assigns it to the Warehouse field.
func (o *OrdersStatsItemDTO) SetWarehouse(v OrdersStatsWarehouseDTO) {
	o.Warehouse = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *OrdersStatsItemDTO) GetDetails() []OrdersStatsDetailsDTO {
	if o == nil || IsNil(o.Details) {
		var ret []OrdersStatsDetailsDTO
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsItemDTO) GetDetailsOk() ([]OrdersStatsDetailsDTO, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *OrdersStatsItemDTO) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []OrdersStatsDetailsDTO and assigns it to the Details field.
func (o *OrdersStatsItemDTO) SetDetails(v []OrdersStatsDetailsDTO) {
	o.Details = v
}

// GetCisList returns the CisList field value if set, zero value otherwise.
func (o *OrdersStatsItemDTO) GetCisList() []string {
	if o == nil || IsNil(o.CisList) {
		var ret []string
		return ret
	}
	return o.CisList
}

// GetCisListOk returns a tuple with the CisList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsItemDTO) GetCisListOk() ([]string, bool) {
	if o == nil || IsNil(o.CisList) {
		return nil, false
	}
	return o.CisList, true
}

// HasCisList returns a boolean if a field has been set.
func (o *OrdersStatsItemDTO) HasCisList() bool {
	if o != nil && !IsNil(o.CisList) {
		return true
	}

	return false
}

// SetCisList gets a reference to the given []string and assigns it to the CisList field.
func (o *OrdersStatsItemDTO) SetCisList(v []string) {
	o.CisList = v
}

// GetInitialCount returns the InitialCount field value if set, zero value otherwise.
func (o *OrdersStatsItemDTO) GetInitialCount() int32 {
	if o == nil || IsNil(o.InitialCount) {
		var ret int32
		return ret
	}
	return *o.InitialCount
}

// GetInitialCountOk returns a tuple with the InitialCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsItemDTO) GetInitialCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InitialCount) {
		return nil, false
	}
	return o.InitialCount, true
}

// HasInitialCount returns a boolean if a field has been set.
func (o *OrdersStatsItemDTO) HasInitialCount() bool {
	if o != nil && !IsNil(o.InitialCount) {
		return true
	}

	return false
}

// SetInitialCount gets a reference to the given int32 and assigns it to the InitialCount field.
func (o *OrdersStatsItemDTO) SetInitialCount(v int32) {
	o.InitialCount = &v
}

// GetBidFee returns the BidFee field value if set, zero value otherwise.
func (o *OrdersStatsItemDTO) GetBidFee() int32 {
	if o == nil || IsNil(o.BidFee) {
		var ret int32
		return ret
	}
	return *o.BidFee
}

// GetBidFeeOk returns a tuple with the BidFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsItemDTO) GetBidFeeOk() (*int32, bool) {
	if o == nil || IsNil(o.BidFee) {
		return nil, false
	}
	return o.BidFee, true
}

// HasBidFee returns a boolean if a field has been set.
func (o *OrdersStatsItemDTO) HasBidFee() bool {
	if o != nil && !IsNil(o.BidFee) {
		return true
	}

	return false
}

// SetBidFee gets a reference to the given int32 and assigns it to the BidFee field.
func (o *OrdersStatsItemDTO) SetBidFee(v int32) {
	o.BidFee = &v
}

// GetCofinanceThreshold returns the CofinanceThreshold field value if set, zero value otherwise.
func (o *OrdersStatsItemDTO) GetCofinanceThreshold() float32 {
	if o == nil || IsNil(o.CofinanceThreshold) {
		var ret float32
		return ret
	}
	return *o.CofinanceThreshold
}

// GetCofinanceThresholdOk returns a tuple with the CofinanceThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsItemDTO) GetCofinanceThresholdOk() (*float32, bool) {
	if o == nil || IsNil(o.CofinanceThreshold) {
		return nil, false
	}
	return o.CofinanceThreshold, true
}

// HasCofinanceThreshold returns a boolean if a field has been set.
func (o *OrdersStatsItemDTO) HasCofinanceThreshold() bool {
	if o != nil && !IsNil(o.CofinanceThreshold) {
		return true
	}

	return false
}

// SetCofinanceThreshold gets a reference to the given float32 and assigns it to the CofinanceThreshold field.
func (o *OrdersStatsItemDTO) SetCofinanceThreshold(v float32) {
	o.CofinanceThreshold = &v
}

// GetCofinanceValue returns the CofinanceValue field value if set, zero value otherwise.
func (o *OrdersStatsItemDTO) GetCofinanceValue() float32 {
	if o == nil || IsNil(o.CofinanceValue) {
		var ret float32
		return ret
	}
	return *o.CofinanceValue
}

// GetCofinanceValueOk returns a tuple with the CofinanceValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsItemDTO) GetCofinanceValueOk() (*float32, bool) {
	if o == nil || IsNil(o.CofinanceValue) {
		return nil, false
	}
	return o.CofinanceValue, true
}

// HasCofinanceValue returns a boolean if a field has been set.
func (o *OrdersStatsItemDTO) HasCofinanceValue() bool {
	if o != nil && !IsNil(o.CofinanceValue) {
		return true
	}

	return false
}

// SetCofinanceValue gets a reference to the given float32 and assigns it to the CofinanceValue field.
func (o *OrdersStatsItemDTO) SetCofinanceValue(v float32) {
	o.CofinanceValue = &v
}

func (o OrdersStatsItemDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrdersStatsItemDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OfferName) {
		toSerialize["offerName"] = o.OfferName
	}
	if !IsNil(o.MarketSku) {
		toSerialize["marketSku"] = o.MarketSku
	}
	if !IsNil(o.ShopSku) {
		toSerialize["shopSku"] = o.ShopSku
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	if !IsNil(o.Warehouse) {
		toSerialize["warehouse"] = o.Warehouse
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.CisList) {
		toSerialize["cisList"] = o.CisList
	}
	if !IsNil(o.InitialCount) {
		toSerialize["initialCount"] = o.InitialCount
	}
	if !IsNil(o.BidFee) {
		toSerialize["bidFee"] = o.BidFee
	}
	if !IsNil(o.CofinanceThreshold) {
		toSerialize["cofinanceThreshold"] = o.CofinanceThreshold
	}
	if !IsNil(o.CofinanceValue) {
		toSerialize["cofinanceValue"] = o.CofinanceValue
	}
	return toSerialize, nil
}

type NullableOrdersStatsItemDTO struct {
	value *OrdersStatsItemDTO
	isSet bool
}

func (v NullableOrdersStatsItemDTO) Get() *OrdersStatsItemDTO {
	return v.value
}

func (v *NullableOrdersStatsItemDTO) Set(val *OrdersStatsItemDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersStatsItemDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersStatsItemDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersStatsItemDTO(val *OrdersStatsItemDTO) *NullableOrdersStatsItemDTO {
	return &NullableOrdersStatsItemDTO{value: val, isSet: true}
}

func (v NullableOrdersStatsItemDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersStatsItemDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


