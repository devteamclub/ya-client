# OfferPriceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | Pointer to **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | [optional] 
**Id** | Pointer to **string** | {% note alert \&quot;Это поле устарело\&quot; %}  Не используйте его — это может привести к ошибкам.  {% endnote %}  Идентификатор предложения из прайс-листа.  | [optional] 
**Feed** | Pointer to [**OfferPriceFeedDTO**](OfferPriceFeedDTO.md) |  | [optional] 
**Price** | Pointer to [**PriceDTO**](PriceDTO.md) |  | [optional] 
**MarketSku** | Pointer to **int64** | SKU на Маркете. | [optional] 
**ShopSku** | Pointer to **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | [optional] 

## Methods

### NewOfferPriceDTO

`func NewOfferPriceDTO() *OfferPriceDTO`

NewOfferPriceDTO instantiates a new OfferPriceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferPriceDTOWithDefaults

`func NewOfferPriceDTOWithDefaults() *OfferPriceDTO`

NewOfferPriceDTOWithDefaults instantiates a new OfferPriceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *OfferPriceDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *OfferPriceDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *OfferPriceDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *OfferPriceDTO) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetId

`func (o *OfferPriceDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OfferPriceDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OfferPriceDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OfferPriceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFeed

`func (o *OfferPriceDTO) GetFeed() OfferPriceFeedDTO`

GetFeed returns the Feed field if non-nil, zero value otherwise.

### GetFeedOk

`func (o *OfferPriceDTO) GetFeedOk() (*OfferPriceFeedDTO, bool)`

GetFeedOk returns a tuple with the Feed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeed

`func (o *OfferPriceDTO) SetFeed(v OfferPriceFeedDTO)`

SetFeed sets Feed field to given value.

### HasFeed

`func (o *OfferPriceDTO) HasFeed() bool`

HasFeed returns a boolean if a field has been set.

### GetPrice

`func (o *OfferPriceDTO) GetPrice() PriceDTO`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OfferPriceDTO) GetPriceOk() (*PriceDTO, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OfferPriceDTO) SetPrice(v PriceDTO)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OfferPriceDTO) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetMarketSku

`func (o *OfferPriceDTO) GetMarketSku() int64`

GetMarketSku returns the MarketSku field if non-nil, zero value otherwise.

### GetMarketSkuOk

`func (o *OfferPriceDTO) GetMarketSkuOk() (*int64, bool)`

GetMarketSkuOk returns a tuple with the MarketSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketSku

`func (o *OfferPriceDTO) SetMarketSku(v int64)`

SetMarketSku sets MarketSku field to given value.

### HasMarketSku

`func (o *OfferPriceDTO) HasMarketSku() bool`

HasMarketSku returns a boolean if a field has been set.

### GetShopSku

`func (o *OfferPriceDTO) GetShopSku() string`

GetShopSku returns the ShopSku field if non-nil, zero value otherwise.

### GetShopSkuOk

`func (o *OfferPriceDTO) GetShopSkuOk() (*string, bool)`

GetShopSkuOk returns a tuple with the ShopSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopSku

`func (o *OfferPriceDTO) SetShopSku(v string)`

SetShopSku sets ShopSku field to given value.

### HasShopSku

`func (o *OfferPriceDTO) HasShopSku() bool`

HasShopSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


