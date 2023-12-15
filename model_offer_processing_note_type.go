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

// OfferProcessingNoteType Тип причины, по которой товар не прошел модерацию:  * `ASSORTMENT` — товар производится в разных вариантах. Каждый из них нужно описать как отдельный товар (входной параметр `offer-mapping-entry` запроса POST /campaigns/{campaignId}/offer-mapping-entries/updates или строка в каталоге, если вы загружаете товары через личный кабинет магазина). * `CANCELLED` — товар отозван с модерации по вашей инициативе. * `CONFLICTING_INFORMATION` _(ранее ошибочно `CONFLICTING`)_ — вы предоставили противоречивую информацию о товаре. Параметры, которые нужно исправить, указаны в параметре `payload`. * `DEPARTMENT_FROZEN` — правила размещения товаров в данной категории перерабатываются, поэтому товар пока не может пройти модерацию. * `INCORRECT_INFORMATION` — информация о товаре, которую вы предоставили, противоречит описанию от производителя. Параметры, которые нужно исправить, указаны в параметре `payload`. * `LEGAL_CONFLICT` — товар не прошел модерацию по юридическим причинам. Например, он официально не продается в России или у вас нет разрешения на его продажу. * `NEED_CLASSIFICATION_INFORMATION` — информации о товаре, которую вы предоставили, не хватает, чтобы отнести его к категории. Проверьте, что правильно указали название, категорию, производителя и страны производства товара, а также URL изображений или страниц с описанием, по которым можно идентифицировать товар. * `NEED_INFORMATION` — товар раньше не продавался в России и пока не размещается на Маркете. Для него можно создать карточку. Подробнее см. в разделе [Работа с карточкой товара](https://yandex.ru/support/marketplace/assortment/content/index.html) Справки Маркета для продавцов. * `NEED_PICTURES` — для идентификации товара нужны его изображения. Отправьте URL изображений товара в запросе POST /campaigns/{campaignId}/offer-mapping-entries/updates или загрузите обновленный каталог через личный кабинет магазина. * `NEED_VENDOR` — неверно указан производитель товара. * `NO_CATEGORY`, `NO_KNOWLEDGE` — товары из указанной категории пока не размещаются на Маркете. Если категория появится, товар будет снова отправлен на модерацию. * `NO_PARAMETERS_IN_SHOP_TITLE` — товар производится в разных вариантах, и из указанного названия непонятно, о каком идет речь. Параметры, которые нужно добавить в название товара, указаны в параметре `payload`. * `NO_SIZE_MEASURE` — для этого товара нужна размерная сетка. Отправьте ее в службу поддержки или вашему менеджеру. Требования к размерной сетке указаны в параметре `payload`. * `UNKNOWN` — товар не прошел модерацию по другой причине. Обратитесь в службу поддержки или к вашему менеджеру. 
type OfferProcessingNoteType string

// List of OfferProcessingNoteType
const (
	ASSORTMENT OfferProcessingNoteType = "ASSORTMENT"
	CANCELLED OfferProcessingNoteType = "CANCELLED"
	CONFLICTING_INFORMATION OfferProcessingNoteType = "CONFLICTING_INFORMATION"
	OTHER OfferProcessingNoteType = "OTHER"
	DEPARTMENT_FROZEN OfferProcessingNoteType = "DEPARTMENT_FROZEN"
	INCORRECT_INFORMATION OfferProcessingNoteType = "INCORRECT_INFORMATION"
	LEGAL_CONFLICT OfferProcessingNoteType = "LEGAL_CONFLICT"
	NEED_CLASSIFICATION_INFORMATION OfferProcessingNoteType = "NEED_CLASSIFICATION_INFORMATION"
	NEED_INFORMATION OfferProcessingNoteType = "NEED_INFORMATION"
	NEED_PICTURES OfferProcessingNoteType = "NEED_PICTURES"
	NEED_VENDOR OfferProcessingNoteType = "NEED_VENDOR"
	NO_CATEGORY OfferProcessingNoteType = "NO_CATEGORY"
	NO_KNOWLEDGE OfferProcessingNoteType = "NO_KNOWLEDGE"
	NO_PARAMETERS_IN_SHOP_TITLE OfferProcessingNoteType = "NO_PARAMETERS_IN_SHOP_TITLE"
	NO_SIZE_MEASURE OfferProcessingNoteType = "NO_SIZE_MEASURE"
	SAMPLE_LINE OfferProcessingNoteType = "SAMPLE_LINE"
)

// All allowed values of OfferProcessingNoteType enum
var AllowedOfferProcessingNoteTypeEnumValues = []OfferProcessingNoteType{
	"ASSORTMENT",
	"CANCELLED",
	"CONFLICTING_INFORMATION",
	"OTHER",
	"DEPARTMENT_FROZEN",
	"INCORRECT_INFORMATION",
	"LEGAL_CONFLICT",
	"NEED_CLASSIFICATION_INFORMATION",
	"NEED_INFORMATION",
	"NEED_PICTURES",
	"NEED_VENDOR",
	"NO_CATEGORY",
	"NO_KNOWLEDGE",
	"NO_PARAMETERS_IN_SHOP_TITLE",
	"NO_SIZE_MEASURE",
	"SAMPLE_LINE",
}

func (v *OfferProcessingNoteType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OfferProcessingNoteType(value)
	for _, existing := range AllowedOfferProcessingNoteTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OfferProcessingNoteType", value)
}

// NewOfferProcessingNoteTypeFromValue returns a pointer to a valid OfferProcessingNoteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOfferProcessingNoteTypeFromValue(v string) (*OfferProcessingNoteType, error) {
	ev := OfferProcessingNoteType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OfferProcessingNoteType: valid values are %v", v, AllowedOfferProcessingNoteTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OfferProcessingNoteType) IsValid() bool {
	for _, existing := range AllowedOfferProcessingNoteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OfferProcessingNoteType value
func (v OfferProcessingNoteType) Ptr() *OfferProcessingNoteType {
	return &v
}

type NullableOfferProcessingNoteType struct {
	value *OfferProcessingNoteType
	isSet bool
}

func (v NullableOfferProcessingNoteType) Get() *OfferProcessingNoteType {
	return v.value
}

func (v *NullableOfferProcessingNoteType) Set(val *OfferProcessingNoteType) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferProcessingNoteType) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferProcessingNoteType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferProcessingNoteType(val *OfferProcessingNoteType) *NullableOfferProcessingNoteType {
	return &NullableOfferProcessingNoteType{value: val, isSet: true}
}

func (v NullableOfferProcessingNoteType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferProcessingNoteType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
