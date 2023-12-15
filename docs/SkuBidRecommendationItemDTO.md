# SkuBidRecommendationItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | 
**Bid** | **int32** | Значение ставки. | 
**BidRecommendations** | Pointer to [**[]BidRecommendationItemDTO**](BidRecommendationItemDTO.md) | Список рекомендованных ставок с соответствующими долями показов. Чем больше ставка, тем большую долю показов она помогает получить.  | [optional] 
**PriceRecommendations** | Pointer to [**[]PriceRecommendationItemDTO**](PriceRecommendationItemDTO.md) | Рекомендованные цены. | [optional] 

## Methods

### NewSkuBidRecommendationItemDTO

`func NewSkuBidRecommendationItemDTO(sku string, bid int32, ) *SkuBidRecommendationItemDTO`

NewSkuBidRecommendationItemDTO instantiates a new SkuBidRecommendationItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuBidRecommendationItemDTOWithDefaults

`func NewSkuBidRecommendationItemDTOWithDefaults() *SkuBidRecommendationItemDTO`

NewSkuBidRecommendationItemDTOWithDefaults instantiates a new SkuBidRecommendationItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *SkuBidRecommendationItemDTO) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *SkuBidRecommendationItemDTO) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *SkuBidRecommendationItemDTO) SetSku(v string)`

SetSku sets Sku field to given value.


### GetBid

`func (o *SkuBidRecommendationItemDTO) GetBid() int32`

GetBid returns the Bid field if non-nil, zero value otherwise.

### GetBidOk

`func (o *SkuBidRecommendationItemDTO) GetBidOk() (*int32, bool)`

GetBidOk returns a tuple with the Bid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBid

`func (o *SkuBidRecommendationItemDTO) SetBid(v int32)`

SetBid sets Bid field to given value.


### GetBidRecommendations

`func (o *SkuBidRecommendationItemDTO) GetBidRecommendations() []BidRecommendationItemDTO`

GetBidRecommendations returns the BidRecommendations field if non-nil, zero value otherwise.

### GetBidRecommendationsOk

`func (o *SkuBidRecommendationItemDTO) GetBidRecommendationsOk() (*[]BidRecommendationItemDTO, bool)`

GetBidRecommendationsOk returns a tuple with the BidRecommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidRecommendations

`func (o *SkuBidRecommendationItemDTO) SetBidRecommendations(v []BidRecommendationItemDTO)`

SetBidRecommendations sets BidRecommendations field to given value.

### HasBidRecommendations

`func (o *SkuBidRecommendationItemDTO) HasBidRecommendations() bool`

HasBidRecommendations returns a boolean if a field has been set.

### GetPriceRecommendations

`func (o *SkuBidRecommendationItemDTO) GetPriceRecommendations() []PriceRecommendationItemDTO`

GetPriceRecommendations returns the PriceRecommendations field if non-nil, zero value otherwise.

### GetPriceRecommendationsOk

`func (o *SkuBidRecommendationItemDTO) GetPriceRecommendationsOk() (*[]PriceRecommendationItemDTO, bool)`

GetPriceRecommendationsOk returns a tuple with the PriceRecommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRecommendations

`func (o *SkuBidRecommendationItemDTO) SetPriceRecommendations(v []PriceRecommendationItemDTO)`

SetPriceRecommendations sets PriceRecommendations field to given value.

### HasPriceRecommendations

`func (o *SkuBidRecommendationItemDTO) HasPriceRecommendations() bool`

HasPriceRecommendations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


