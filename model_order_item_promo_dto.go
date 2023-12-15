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

// checks if the OrderItemPromoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderItemPromoDTO{}

// OrderItemPromoDTO Информация о вознаграждениях партнеру за скидки на товар по промокодам, купонам и акциям.
type OrderItemPromoDTO struct {
	Type OrderPromoType `json:"type"`
	// Размер пользовательской скидки в валюте покупателя. 
	Discount *float32 `json:"discount,omitempty"`
	// Вознаграждение партнеру за скидку.  Передается в валюте заказа, для отделения целой части от дробной используется точка. 
	Subsidy *float32 `json:"subsidy,omitempty"`
	// Идентификатор акции поставщика. 
	ShopPromoId *string `json:"shopPromoId,omitempty"`
	// Идентификатор акции в рамках соглашения на оказание услуг по продвижению сервиса между маркетплейсом Яндекс Маркета и партнером.  Параметр передается, только если параметр `type=MARKET_DEAL`. 
	MarketPromoId *string `json:"marketPromoId,omitempty"`
}

type _OrderItemPromoDTO OrderItemPromoDTO

// NewOrderItemPromoDTO instantiates a new OrderItemPromoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItemPromoDTO(type_ OrderPromoType) *OrderItemPromoDTO {
	this := OrderItemPromoDTO{}
	this.Type = type_
	return &this
}

// NewOrderItemPromoDTOWithDefaults instantiates a new OrderItemPromoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemPromoDTOWithDefaults() *OrderItemPromoDTO {
	this := OrderItemPromoDTO{}
	return &this
}

// GetType returns the Type field value
func (o *OrderItemPromoDTO) GetType() OrderPromoType {
	if o == nil {
		var ret OrderPromoType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderItemPromoDTO) GetTypeOk() (*OrderPromoType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderItemPromoDTO) SetType(v OrderPromoType) {
	o.Type = v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *OrderItemPromoDTO) GetDiscount() float32 {
	if o == nil || IsNil(o.Discount) {
		var ret float32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemPromoDTO) GetDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *OrderItemPromoDTO) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float32 and assigns it to the Discount field.
func (o *OrderItemPromoDTO) SetDiscount(v float32) {
	o.Discount = &v
}

// GetSubsidy returns the Subsidy field value if set, zero value otherwise.
func (o *OrderItemPromoDTO) GetSubsidy() float32 {
	if o == nil || IsNil(o.Subsidy) {
		var ret float32
		return ret
	}
	return *o.Subsidy
}

// GetSubsidyOk returns a tuple with the Subsidy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemPromoDTO) GetSubsidyOk() (*float32, bool) {
	if o == nil || IsNil(o.Subsidy) {
		return nil, false
	}
	return o.Subsidy, true
}

// HasSubsidy returns a boolean if a field has been set.
func (o *OrderItemPromoDTO) HasSubsidy() bool {
	if o != nil && !IsNil(o.Subsidy) {
		return true
	}

	return false
}

// SetSubsidy gets a reference to the given float32 and assigns it to the Subsidy field.
func (o *OrderItemPromoDTO) SetSubsidy(v float32) {
	o.Subsidy = &v
}

// GetShopPromoId returns the ShopPromoId field value if set, zero value otherwise.
func (o *OrderItemPromoDTO) GetShopPromoId() string {
	if o == nil || IsNil(o.ShopPromoId) {
		var ret string
		return ret
	}
	return *o.ShopPromoId
}

// GetShopPromoIdOk returns a tuple with the ShopPromoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemPromoDTO) GetShopPromoIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShopPromoId) {
		return nil, false
	}
	return o.ShopPromoId, true
}

// HasShopPromoId returns a boolean if a field has been set.
func (o *OrderItemPromoDTO) HasShopPromoId() bool {
	if o != nil && !IsNil(o.ShopPromoId) {
		return true
	}

	return false
}

// SetShopPromoId gets a reference to the given string and assigns it to the ShopPromoId field.
func (o *OrderItemPromoDTO) SetShopPromoId(v string) {
	o.ShopPromoId = &v
}

// GetMarketPromoId returns the MarketPromoId field value if set, zero value otherwise.
func (o *OrderItemPromoDTO) GetMarketPromoId() string {
	if o == nil || IsNil(o.MarketPromoId) {
		var ret string
		return ret
	}
	return *o.MarketPromoId
}

// GetMarketPromoIdOk returns a tuple with the MarketPromoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemPromoDTO) GetMarketPromoIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketPromoId) {
		return nil, false
	}
	return o.MarketPromoId, true
}

// HasMarketPromoId returns a boolean if a field has been set.
func (o *OrderItemPromoDTO) HasMarketPromoId() bool {
	if o != nil && !IsNil(o.MarketPromoId) {
		return true
	}

	return false
}

// SetMarketPromoId gets a reference to the given string and assigns it to the MarketPromoId field.
func (o *OrderItemPromoDTO) SetMarketPromoId(v string) {
	o.MarketPromoId = &v
}

func (o OrderItemPromoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderItemPromoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.Subsidy) {
		toSerialize["subsidy"] = o.Subsidy
	}
	if !IsNil(o.ShopPromoId) {
		toSerialize["shopPromoId"] = o.ShopPromoId
	}
	if !IsNil(o.MarketPromoId) {
		toSerialize["marketPromoId"] = o.MarketPromoId
	}
	return toSerialize, nil
}

func (o *OrderItemPromoDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varOrderItemPromoDTO := _OrderItemPromoDTO{}

	err = json.Unmarshal(bytes, &varOrderItemPromoDTO)

	if err != nil {
		return err
	}

	*o = OrderItemPromoDTO(varOrderItemPromoDTO)

	return err
}

type NullableOrderItemPromoDTO struct {
	value *OrderItemPromoDTO
	isSet bool
}

func (v NullableOrderItemPromoDTO) Get() *OrderItemPromoDTO {
	return v.value
}

func (v *NullableOrderItemPromoDTO) Set(val *OrderItemPromoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItemPromoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItemPromoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItemPromoDTO(val *OrderItemPromoDTO) *NullableOrderItemPromoDTO {
	return &NullableOrderItemPromoDTO{value: val, isSet: true}
}

func (v NullableOrderItemPromoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderItemPromoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


