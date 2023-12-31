/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// OfferCardRecommendationType Рекомендация по дополнению или замене контента.  Часть рекомендаций относятся к **основным параметрам**, которые есть у товаров любых категорий. Другие — к тем **характеристикам**, которые есть у товара потому, что он относится к определенной категории.  **1. Рекомендации, относящиеся к основным параметрам**  Каждая такая рекомендация относится к **единственному параметру**. Чтобы заполнить этот параметр, пользуйтесь запросом [POST businesses/{businessId}/offer-mappings/update](../../reference/business-assortment/updateOfferMappings.md).  |Рекомендация        |Содержание                                                           |Параметр в updateOfferMappings| |--------------------|---------------------------------------------------------------------|------------------------------| |`HAS_VIDEO`         |Добавьте на карточку видео.                                          |`videos`                      | |`RECOGNIZED_VENDOR` |Напишите название производителя так, как его пишет сам производитель.|`vendor`                      | |`PICTURE_COUNT`     |Добавьте изображения.                                                |`pictures`                    | |`HAS_DESCRIPTION`   |Заполните описание.                                                  |`description`                 | |`HAS_BARCODE`       |Укажите штрихкод.                                                    |`barcode`                     | |`FIRST_PICTURE_SIZE`|Замените первое изображение более крупным.                           |`pictures`                    |  **2. Рекомендации, относящиеся к характеристикам по категориям**  Каждая такая рекомендация предполагает заполнение **одной или нескольких характеристик**. Чтобы узнать, какие именно характеристики нужно заполнить, воспользуйтесь запросом POST category/{categoryId}/parameters. Например, если вы получили рекомендацию `MAIN`, нужно заполнить характеристики, имеющие `MAIN` в массиве `recommendationTypes`.  |Рекомендация |Содержание                                                                  | |-------------|----------------------------------------------------------------------------| |`MAIN`       |Заполните ключевые характеристики товара.                                   | |`ADDITIONAL` |Заполните дополнительные характеристики товара.                             | |`FILTERABLE` |Заполните характеристики, используемые в фильтрах.                          | |`DISTINCTIVE`|Заполните характеристики, которыми отличаются друг от друга варианты товара.| 
type OfferCardRecommendationType string

// List of OfferCardRecommendationType
const (
	HAS_VIDEO OfferCardRecommendationType = "HAS_VIDEO"
	RECOGNIZED_VENDOR OfferCardRecommendationType = "RECOGNIZED_VENDOR"
	MAIN OfferCardRecommendationType = "MAIN"
	ADDITIONAL OfferCardRecommendationType = "ADDITIONAL"
	DISTINCTIVE OfferCardRecommendationType = "DISTINCTIVE"
	FILTERABLE OfferCardRecommendationType = "FILTERABLE"
	PICTURE_COUNT OfferCardRecommendationType = "PICTURE_COUNT"
	HAS_DESCRIPTION OfferCardRecommendationType = "HAS_DESCRIPTION"
	HAS_BARCODE OfferCardRecommendationType = "HAS_BARCODE"
	FIRST_PICTURE_SIZE OfferCardRecommendationType = "FIRST_PICTURE_SIZE"
)

// All allowed values of OfferCardRecommendationType enum
var AllowedOfferCardRecommendationTypeEnumValues = []OfferCardRecommendationType{
	"HAS_VIDEO",
	"RECOGNIZED_VENDOR",
	"MAIN",
	"ADDITIONAL",
	"DISTINCTIVE",
	"FILTERABLE",
	"PICTURE_COUNT",
	"HAS_DESCRIPTION",
	"HAS_BARCODE",
	"FIRST_PICTURE_SIZE",
}

func (v *OfferCardRecommendationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OfferCardRecommendationType(value)
	for _, existing := range AllowedOfferCardRecommendationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OfferCardRecommendationType", value)
}

// NewOfferCardRecommendationTypeFromValue returns a pointer to a valid OfferCardRecommendationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOfferCardRecommendationTypeFromValue(v string) (*OfferCardRecommendationType, error) {
	ev := OfferCardRecommendationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OfferCardRecommendationType: valid values are %v", v, AllowedOfferCardRecommendationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OfferCardRecommendationType) IsValid() bool {
	for _, existing := range AllowedOfferCardRecommendationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OfferCardRecommendationType value
func (v OfferCardRecommendationType) Ptr() *OfferCardRecommendationType {
	return &v
}

type NullableOfferCardRecommendationType struct {
	value *OfferCardRecommendationType
	isSet bool
}

func (v NullableOfferCardRecommendationType) Get() *OfferCardRecommendationType {
	return v.value
}

func (v *NullableOfferCardRecommendationType) Set(val *OfferCardRecommendationType) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferCardRecommendationType) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferCardRecommendationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferCardRecommendationType(val *OfferCardRecommendationType) *NullableOfferCardRecommendationType {
	return &NullableOfferCardRecommendationType{value: val, isSet: true}
}

func (v NullableOfferCardRecommendationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferCardRecommendationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

