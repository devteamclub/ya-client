# OfferContentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | 
**CategoryId** | **int32** | Идентификатор категории на Маркете. | 
**ParameterValues** | [**[]ParameterValueDTO**](ParameterValueDTO.md) | Список характеристик с их значениями. Обновляется всегда целиком. Максимальное количество — 300.  | 

## Methods

### NewOfferContentDTO

`func NewOfferContentDTO(offerId string, categoryId int32, parameterValues []ParameterValueDTO, ) *OfferContentDTO`

NewOfferContentDTO instantiates a new OfferContentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferContentDTOWithDefaults

`func NewOfferContentDTOWithDefaults() *OfferContentDTO`

NewOfferContentDTOWithDefaults instantiates a new OfferContentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *OfferContentDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *OfferContentDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *OfferContentDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetCategoryId

`func (o *OfferContentDTO) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *OfferContentDTO) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *OfferContentDTO) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetParameterValues

`func (o *OfferContentDTO) GetParameterValues() []ParameterValueDTO`

GetParameterValues returns the ParameterValues field if non-nil, zero value otherwise.

### GetParameterValuesOk

`func (o *OfferContentDTO) GetParameterValuesOk() (*[]ParameterValueDTO, bool)`

GetParameterValuesOk returns a tuple with the ParameterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterValues

`func (o *OfferContentDTO) SetParameterValues(v []ParameterValueDTO)`

SetParameterValues sets ParameterValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


