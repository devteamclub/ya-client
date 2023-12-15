# OfferPriceByOfferIdsResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | Pointer to **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | [optional] 
**Price** | Pointer to [**PriceDTO**](PriceDTO.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Дата и время последнего обновления цены. | [optional] 

## Methods

### NewOfferPriceByOfferIdsResponseDTO

`func NewOfferPriceByOfferIdsResponseDTO() *OfferPriceByOfferIdsResponseDTO`

NewOfferPriceByOfferIdsResponseDTO instantiates a new OfferPriceByOfferIdsResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferPriceByOfferIdsResponseDTOWithDefaults

`func NewOfferPriceByOfferIdsResponseDTOWithDefaults() *OfferPriceByOfferIdsResponseDTO`

NewOfferPriceByOfferIdsResponseDTOWithDefaults instantiates a new OfferPriceByOfferIdsResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *OfferPriceByOfferIdsResponseDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *OfferPriceByOfferIdsResponseDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *OfferPriceByOfferIdsResponseDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *OfferPriceByOfferIdsResponseDTO) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetPrice

`func (o *OfferPriceByOfferIdsResponseDTO) GetPrice() PriceDTO`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OfferPriceByOfferIdsResponseDTO) GetPriceOk() (*PriceDTO, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OfferPriceByOfferIdsResponseDTO) SetPrice(v PriceDTO)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OfferPriceByOfferIdsResponseDTO) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OfferPriceByOfferIdsResponseDTO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OfferPriceByOfferIdsResponseDTO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OfferPriceByOfferIdsResponseDTO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OfferPriceByOfferIdsResponseDTO) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


