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

// OrderPromoType Тип скидки:  * `MARKET_COUPON` — скидка по промокоду от Маркета. * `MARKET_DEAL` — скидка в рамках соглашения на оказание услуг по продвижению сервиса между маркетплейсом Яндекс Маркета и партнером. * `MARKET_COIN` — скидка по купонам. 
type OrderPromoType string

// List of OrderPromoType
const (
	DIRECT_DISCOUNT OrderPromoType = "DIRECT_DISCOUNT"
	BLUE_SET OrderPromoType = "BLUE_SET"
	BLUE_FLASH OrderPromoType = "BLUE_FLASH"
	GENERIC_BUNDLE OrderPromoType = "GENERIC_BUNDLE"
	MARKET_COUPON OrderPromoType = "MARKET_COUPON"
	MARKET_PROMOCODE OrderPromoType = "MARKET_PROMOCODE"
	MARKET_DEAL OrderPromoType = "MARKET_DEAL"
	MARKET_BLUE OrderPromoType = "MARKET_BLUE"
	MARKET_PRIME OrderPromoType = "MARKET_PRIME"
	YANDEX_PLUS OrderPromoType = "YANDEX_PLUS"
	BERU_PLUS OrderPromoType = "BERU_PLUS"
	MARKET_COIN OrderPromoType = "MARKET_COIN"
	YANDEX_EMPLOYEE OrderPromoType = "YANDEX_EMPLOYEE"
	LIMITED_FREE_DELIVERY_PROMO OrderPromoType = "LIMITED_FREE_DELIVERY_PROMO"
	FREE_DELIVERY_THRESHOLD OrderPromoType = "FREE_DELIVERY_THRESHOLD"
	MULTICART_DISCOUNT OrderPromoType = "MULTICART_DISCOUNT"
	PRICE_DROP_AS_YOU_SHOP OrderPromoType = "PRICE_DROP_AS_YOU_SHOP"
	FREE_DELIVERY_FOR_LDI OrderPromoType = "FREE_DELIVERY_FOR_LDI"
	FREE_DELIVERY_FOR_LSC OrderPromoType = "FREE_DELIVERY_FOR_LSC"
	SECRET_SALE OrderPromoType = "SECRET_SALE"
	FREE_PICKUP OrderPromoType = "FREE_PICKUP"
	CHEAPEST_AS_GIFT OrderPromoType = "CHEAPEST_AS_GIFT"
	CASHBACK OrderPromoType = "CASHBACK"
	SUPPLIER_MULTICART_DISCOUNT OrderPromoType = "SUPPLIER_MULTICART_DISCOUNT"
	SPREAD_DISCOUNT_COUNT OrderPromoType = "SPREAD_DISCOUNT_COUNT"
	SPREAD_DISCOUNT_RECEIPT OrderPromoType = "SPREAD_DISCOUNT_RECEIPT"
	ANNOUNCEMENT_PROMO OrderPromoType = "ANNOUNCEMENT_PROMO"
	DISCOUNT_BY_PAYMENT_TYPE OrderPromoType = "DISCOUNT_BY_PAYMENT_TYPE"
	PERCENT_DISCOUNT OrderPromoType = "PERCENT_DISCOUNT"
	EMPTY_PROMO OrderPromoType = "EMPTY_PROMO"
	UNKNOWN OrderPromoType = "UNKNOWN"
)

// All allowed values of OrderPromoType enum
var AllowedOrderPromoTypeEnumValues = []OrderPromoType{
	"DIRECT_DISCOUNT",
	"BLUE_SET",
	"BLUE_FLASH",
	"GENERIC_BUNDLE",
	"MARKET_COUPON",
	"MARKET_PROMOCODE",
	"MARKET_DEAL",
	"MARKET_BLUE",
	"MARKET_PRIME",
	"YANDEX_PLUS",
	"BERU_PLUS",
	"MARKET_COIN",
	"YANDEX_EMPLOYEE",
	"LIMITED_FREE_DELIVERY_PROMO",
	"FREE_DELIVERY_THRESHOLD",
	"MULTICART_DISCOUNT",
	"PRICE_DROP_AS_YOU_SHOP",
	"FREE_DELIVERY_FOR_LDI",
	"FREE_DELIVERY_FOR_LSC",
	"SECRET_SALE",
	"FREE_PICKUP",
	"CHEAPEST_AS_GIFT",
	"CASHBACK",
	"SUPPLIER_MULTICART_DISCOUNT",
	"SPREAD_DISCOUNT_COUNT",
	"SPREAD_DISCOUNT_RECEIPT",
	"ANNOUNCEMENT_PROMO",
	"DISCOUNT_BY_PAYMENT_TYPE",
	"PERCENT_DISCOUNT",
	"EMPTY_PROMO",
	"UNKNOWN",
}

func (v *OrderPromoType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderPromoType(value)
	for _, existing := range AllowedOrderPromoTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderPromoType", value)
}

// NewOrderPromoTypeFromValue returns a pointer to a valid OrderPromoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderPromoTypeFromValue(v string) (*OrderPromoType, error) {
	ev := OrderPromoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderPromoType: valid values are %v", v, AllowedOrderPromoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderPromoType) IsValid() bool {
	for _, existing := range AllowedOrderPromoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderPromoType value
func (v OrderPromoType) Ptr() *OrderPromoType {
	return &v
}

type NullableOrderPromoType struct {
	value *OrderPromoType
	isSet bool
}

func (v NullableOrderPromoType) Get() *OrderPromoType {
	return v.value
}

func (v *NullableOrderPromoType) Set(val *OrderPromoType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderPromoType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderPromoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderPromoType(val *OrderPromoType) *NullableOrderPromoType {
	return &NullableOrderPromoType{value: val, isSet: true}
}

func (v NullableOrderPromoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderPromoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
