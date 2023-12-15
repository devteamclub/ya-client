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

// checks if the ModelOfferDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelOfferDTO{}

// ModelOfferDTO Информация о предложении.
type ModelOfferDTO struct {
	// Скидка на предложение, в %.
	Discount *int32 `json:"discount,omitempty"`
	// Наименование предложения.
	Name *string `json:"name,omitempty"`
	// Позиция предложения в выдаче Маркета на карточке модели.
	Pos *int32 `json:"pos,omitempty"`
	// Цена предложения без скидки.
	PreDiscountPrice *float32 `json:"preDiscountPrice,omitempty"`
	// Цена предложения.  До версии 2.0 партнерского API у параметра был тип String. 
	Price *float32 `json:"price,omitempty"`
	// Идентификатор региона предложения (регион, откуда доставляется товар).  Сначала показываются предложения, доставляемые из региона, указанного в запросе в параметре `regionId`. Предложения, доставляемые из других регионов, показываются после них. 
	RegionId *int64 `json:"regionId,omitempty"`
	// Стоимость доставки товара в регион.  Если значение параметра — `0`, доставка осуществляется бесплатно. Если значение параметра — `-1`, магазин не осуществляет доставку этого товара (самовывоз). Если стоимость доставки неизвестна, параметр не выводится. 
	ShippingCost *float32 `json:"shippingCost,omitempty"`
	// Название магазина (в том виде, в котором отображается на Маркете).
	ShopName *string `json:"shopName,omitempty"`
	// Рейтинг магазина.  Возможные значения: * `-1` — у магазинов, недавно появившихся на Маркете, рейтинг появляется не сразу. До момента появления рейтинга для таких магазинов возвращается значение `-1`. * `1`. * `2`. * `3`. * `4`. * `5`. 
	ShopRating *int32 `json:"shopRating,omitempty"`
	// {% note alert %}  Параметр устарел и не рекомендуется к использованию.  {% endnote %} 
	// Deprecated
	InStock *int32 `json:"inStock,omitempty"`
}

// NewModelOfferDTO instantiates a new ModelOfferDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelOfferDTO() *ModelOfferDTO {
	this := ModelOfferDTO{}
	return &this
}

// NewModelOfferDTOWithDefaults instantiates a new ModelOfferDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelOfferDTOWithDefaults() *ModelOfferDTO {
	this := ModelOfferDTO{}
	return &this
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *ModelOfferDTO) GetDiscount() int32 {
	if o == nil || IsNil(o.Discount) {
		var ret int32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOfferDTO) GetDiscountOk() (*int32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *ModelOfferDTO) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given int32 and assigns it to the Discount field.
func (o *ModelOfferDTO) SetDiscount(v int32) {
	o.Discount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelOfferDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOfferDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelOfferDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelOfferDTO) SetName(v string) {
	o.Name = &v
}

// GetPos returns the Pos field value if set, zero value otherwise.
func (o *ModelOfferDTO) GetPos() int32 {
	if o == nil || IsNil(o.Pos) {
		var ret int32
		return ret
	}
	return *o.Pos
}

// GetPosOk returns a tuple with the Pos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOfferDTO) GetPosOk() (*int32, bool) {
	if o == nil || IsNil(o.Pos) {
		return nil, false
	}
	return o.Pos, true
}

// HasPos returns a boolean if a field has been set.
func (o *ModelOfferDTO) HasPos() bool {
	if o != nil && !IsNil(o.Pos) {
		return true
	}

	return false
}

// SetPos gets a reference to the given int32 and assigns it to the Pos field.
func (o *ModelOfferDTO) SetPos(v int32) {
	o.Pos = &v
}

// GetPreDiscountPrice returns the PreDiscountPrice field value if set, zero value otherwise.
func (o *ModelOfferDTO) GetPreDiscountPrice() float32 {
	if o == nil || IsNil(o.PreDiscountPrice) {
		var ret float32
		return ret
	}
	return *o.PreDiscountPrice
}

// GetPreDiscountPriceOk returns a tuple with the PreDiscountPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOfferDTO) GetPreDiscountPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.PreDiscountPrice) {
		return nil, false
	}
	return o.PreDiscountPrice, true
}

// HasPreDiscountPrice returns a boolean if a field has been set.
func (o *ModelOfferDTO) HasPreDiscountPrice() bool {
	if o != nil && !IsNil(o.PreDiscountPrice) {
		return true
	}

	return false
}

// SetPreDiscountPrice gets a reference to the given float32 and assigns it to the PreDiscountPrice field.
func (o *ModelOfferDTO) SetPreDiscountPrice(v float32) {
	o.PreDiscountPrice = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ModelOfferDTO) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOfferDTO) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ModelOfferDTO) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *ModelOfferDTO) SetPrice(v float32) {
	o.Price = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *ModelOfferDTO) GetRegionId() int64 {
	if o == nil || IsNil(o.RegionId) {
		var ret int64
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOfferDTO) GetRegionIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RegionId) {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *ModelOfferDTO) HasRegionId() bool {
	if o != nil && !IsNil(o.RegionId) {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given int64 and assigns it to the RegionId field.
func (o *ModelOfferDTO) SetRegionId(v int64) {
	o.RegionId = &v
}

// GetShippingCost returns the ShippingCost field value if set, zero value otherwise.
func (o *ModelOfferDTO) GetShippingCost() float32 {
	if o == nil || IsNil(o.ShippingCost) {
		var ret float32
		return ret
	}
	return *o.ShippingCost
}

// GetShippingCostOk returns a tuple with the ShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOfferDTO) GetShippingCostOk() (*float32, bool) {
	if o == nil || IsNil(o.ShippingCost) {
		return nil, false
	}
	return o.ShippingCost, true
}

// HasShippingCost returns a boolean if a field has been set.
func (o *ModelOfferDTO) HasShippingCost() bool {
	if o != nil && !IsNil(o.ShippingCost) {
		return true
	}

	return false
}

// SetShippingCost gets a reference to the given float32 and assigns it to the ShippingCost field.
func (o *ModelOfferDTO) SetShippingCost(v float32) {
	o.ShippingCost = &v
}

// GetShopName returns the ShopName field value if set, zero value otherwise.
func (o *ModelOfferDTO) GetShopName() string {
	if o == nil || IsNil(o.ShopName) {
		var ret string
		return ret
	}
	return *o.ShopName
}

// GetShopNameOk returns a tuple with the ShopName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOfferDTO) GetShopNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShopName) {
		return nil, false
	}
	return o.ShopName, true
}

// HasShopName returns a boolean if a field has been set.
func (o *ModelOfferDTO) HasShopName() bool {
	if o != nil && !IsNil(o.ShopName) {
		return true
	}

	return false
}

// SetShopName gets a reference to the given string and assigns it to the ShopName field.
func (o *ModelOfferDTO) SetShopName(v string) {
	o.ShopName = &v
}

// GetShopRating returns the ShopRating field value if set, zero value otherwise.
func (o *ModelOfferDTO) GetShopRating() int32 {
	if o == nil || IsNil(o.ShopRating) {
		var ret int32
		return ret
	}
	return *o.ShopRating
}

// GetShopRatingOk returns a tuple with the ShopRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOfferDTO) GetShopRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.ShopRating) {
		return nil, false
	}
	return o.ShopRating, true
}

// HasShopRating returns a boolean if a field has been set.
func (o *ModelOfferDTO) HasShopRating() bool {
	if o != nil && !IsNil(o.ShopRating) {
		return true
	}

	return false
}

// SetShopRating gets a reference to the given int32 and assigns it to the ShopRating field.
func (o *ModelOfferDTO) SetShopRating(v int32) {
	o.ShopRating = &v
}

// GetInStock returns the InStock field value if set, zero value otherwise.
// Deprecated
func (o *ModelOfferDTO) GetInStock() int32 {
	if o == nil || IsNil(o.InStock) {
		var ret int32
		return ret
	}
	return *o.InStock
}

// GetInStockOk returns a tuple with the InStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ModelOfferDTO) GetInStockOk() (*int32, bool) {
	if o == nil || IsNil(o.InStock) {
		return nil, false
	}
	return o.InStock, true
}

// HasInStock returns a boolean if a field has been set.
func (o *ModelOfferDTO) HasInStock() bool {
	if o != nil && !IsNil(o.InStock) {
		return true
	}

	return false
}

// SetInStock gets a reference to the given int32 and assigns it to the InStock field.
// Deprecated
func (o *ModelOfferDTO) SetInStock(v int32) {
	o.InStock = &v
}

func (o ModelOfferDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelOfferDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Pos) {
		toSerialize["pos"] = o.Pos
	}
	if !IsNil(o.PreDiscountPrice) {
		toSerialize["preDiscountPrice"] = o.PreDiscountPrice
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.RegionId) {
		toSerialize["regionId"] = o.RegionId
	}
	if !IsNil(o.ShippingCost) {
		toSerialize["shippingCost"] = o.ShippingCost
	}
	if !IsNil(o.ShopName) {
		toSerialize["shopName"] = o.ShopName
	}
	if !IsNil(o.ShopRating) {
		toSerialize["shopRating"] = o.ShopRating
	}
	if !IsNil(o.InStock) {
		toSerialize["inStock"] = o.InStock
	}
	return toSerialize, nil
}

type NullableModelOfferDTO struct {
	value *ModelOfferDTO
	isSet bool
}

func (v NullableModelOfferDTO) Get() *ModelOfferDTO {
	return v.value
}

func (v *NullableModelOfferDTO) Set(val *ModelOfferDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableModelOfferDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableModelOfferDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelOfferDTO(val *ModelOfferDTO) *NullableModelOfferDTO {
	return &NullableModelOfferDTO{value: val, isSet: true}
}

func (v NullableModelOfferDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelOfferDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


