# OrderItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Идентификатор товара в заказе. | [optional] 
**FeedId** | Pointer to **int64** | Идентификатор каталога товаров. | [optional] 
**OfferId** | Pointer to **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | [optional] 
**FeedCategoryId** | Pointer to **string** | Идентификатор категории, указанный в каталоге. | [optional] 
**OfferName** | Pointer to **string** | Название товара. | [optional] 
**Price** | Pointer to **float32** | Цена товара в валюте заказа без учета вознаграждения партнеру за скидки по промокодам, купонам и акциям (параметр &#x60;subsidy&#x60;).  Для отделения целой части от дробной используется точка.  | [optional] 
**BuyerPrice** | Pointer to **float32** | Цена товара в валюте покупателя. В цене уже учтены скидки по:  * акциям; * купонам; * промокодам.  Для отделения целой части от дробной используется точка.  | [optional] 
**BuyerPriceBeforeDiscount** | Pointer to **float32** | Стоимость товара в валюте покупателя до применения скидок.  Для отделения целой части от дробной используется точка.  | [optional] 
**PriceBeforeDiscount** | Pointer to **float32** | Стоимость товара в валюте магазина до применения скидок.  Для отделения целой части от дробной используется точка.  | [optional] 
**Count** | Pointer to **int32** | Количество товара. | [optional] 
**Vat** | Pointer to [**OrderVatType**](OrderVatType.md) |  | [optional] 
**ShopSku** | Pointer to **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | [optional] 
**Subsidy** | Pointer to **float32** | Общее вознаграждение партнеру за все скидки на товар:  * по промокодам; * по купонам; * по баллам кешбэка по подписке Яндекс Плюс; * по акциям.  Передается в валюте заказа, для отделения целой части от дробной используется точка.  | [optional] 
**PartnerWarehouseId** | Pointer to **string** | Идентификатор склада в системе партнера, на который сформирован заказ.  {% note alert %}  Параметр устарел, временно поддерживается, но не доступен для ввода и редактирования.  {% endnote %}  | [optional] 
**Promos** | Pointer to [**[]OrderItemPromoDTO**](OrderItemPromoDTO.md) | Информация о вознаграждениях партнеру за скидки на товар по промокодам, купонам и акциям. | [optional] 
**Instances** | Pointer to [**[]OrderItemInstanceDTO**](OrderItemInstanceDTO.md) | Информация о маркировке единиц товара.  Возвращаются данные для маркировки, переданные в запросе &#x60;PUT /campaigns/{campaignId}/orders/{orderId}/cis&#x60;.  Если магазин еще не передавал коды для этого заказа, &#x60;instances&#x60; отсутствует.  | [optional] 
**Details** | Pointer to [**[]OrderItemDetailDTO**](OrderItemDetailDTO.md) | Информация об удалении товара из заказа.  | [optional] 
**Subsidies** | Pointer to [**[]OrderItemSubsidyDTO**](OrderItemSubsidyDTO.md) | Список субсидий по типам. | [optional] 
**RequiredInstanceTypes** | Pointer to [**[]OrderItemInstanceType**](OrderItemInstanceType.md) | Список необходимых маркировок товара. | [optional] 

## Methods

### NewOrderItemDTO

`func NewOrderItemDTO() *OrderItemDTO`

NewOrderItemDTO instantiates a new OrderItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemDTOWithDefaults

`func NewOrderItemDTOWithDefaults() *OrderItemDTO`

NewOrderItemDTOWithDefaults instantiates a new OrderItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderItemDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderItemDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderItemDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OrderItemDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFeedId

`func (o *OrderItemDTO) GetFeedId() int64`

GetFeedId returns the FeedId field if non-nil, zero value otherwise.

### GetFeedIdOk

`func (o *OrderItemDTO) GetFeedIdOk() (*int64, bool)`

GetFeedIdOk returns a tuple with the FeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedId

`func (o *OrderItemDTO) SetFeedId(v int64)`

SetFeedId sets FeedId field to given value.

### HasFeedId

`func (o *OrderItemDTO) HasFeedId() bool`

HasFeedId returns a boolean if a field has been set.

### GetOfferId

`func (o *OrderItemDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *OrderItemDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *OrderItemDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *OrderItemDTO) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetFeedCategoryId

`func (o *OrderItemDTO) GetFeedCategoryId() string`

GetFeedCategoryId returns the FeedCategoryId field if non-nil, zero value otherwise.

### GetFeedCategoryIdOk

`func (o *OrderItemDTO) GetFeedCategoryIdOk() (*string, bool)`

GetFeedCategoryIdOk returns a tuple with the FeedCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedCategoryId

`func (o *OrderItemDTO) SetFeedCategoryId(v string)`

SetFeedCategoryId sets FeedCategoryId field to given value.

### HasFeedCategoryId

`func (o *OrderItemDTO) HasFeedCategoryId() bool`

HasFeedCategoryId returns a boolean if a field has been set.

### GetOfferName

`func (o *OrderItemDTO) GetOfferName() string`

GetOfferName returns the OfferName field if non-nil, zero value otherwise.

### GetOfferNameOk

`func (o *OrderItemDTO) GetOfferNameOk() (*string, bool)`

GetOfferNameOk returns a tuple with the OfferName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferName

`func (o *OrderItemDTO) SetOfferName(v string)`

SetOfferName sets OfferName field to given value.

### HasOfferName

`func (o *OrderItemDTO) HasOfferName() bool`

HasOfferName returns a boolean if a field has been set.

### GetPrice

`func (o *OrderItemDTO) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderItemDTO) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderItemDTO) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderItemDTO) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBuyerPrice

`func (o *OrderItemDTO) GetBuyerPrice() float32`

GetBuyerPrice returns the BuyerPrice field if non-nil, zero value otherwise.

### GetBuyerPriceOk

`func (o *OrderItemDTO) GetBuyerPriceOk() (*float32, bool)`

GetBuyerPriceOk returns a tuple with the BuyerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerPrice

`func (o *OrderItemDTO) SetBuyerPrice(v float32)`

SetBuyerPrice sets BuyerPrice field to given value.

### HasBuyerPrice

`func (o *OrderItemDTO) HasBuyerPrice() bool`

HasBuyerPrice returns a boolean if a field has been set.

### GetBuyerPriceBeforeDiscount

`func (o *OrderItemDTO) GetBuyerPriceBeforeDiscount() float32`

GetBuyerPriceBeforeDiscount returns the BuyerPriceBeforeDiscount field if non-nil, zero value otherwise.

### GetBuyerPriceBeforeDiscountOk

`func (o *OrderItemDTO) GetBuyerPriceBeforeDiscountOk() (*float32, bool)`

GetBuyerPriceBeforeDiscountOk returns a tuple with the BuyerPriceBeforeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerPriceBeforeDiscount

`func (o *OrderItemDTO) SetBuyerPriceBeforeDiscount(v float32)`

SetBuyerPriceBeforeDiscount sets BuyerPriceBeforeDiscount field to given value.

### HasBuyerPriceBeforeDiscount

`func (o *OrderItemDTO) HasBuyerPriceBeforeDiscount() bool`

HasBuyerPriceBeforeDiscount returns a boolean if a field has been set.

### GetPriceBeforeDiscount

`func (o *OrderItemDTO) GetPriceBeforeDiscount() float32`

GetPriceBeforeDiscount returns the PriceBeforeDiscount field if non-nil, zero value otherwise.

### GetPriceBeforeDiscountOk

`func (o *OrderItemDTO) GetPriceBeforeDiscountOk() (*float32, bool)`

GetPriceBeforeDiscountOk returns a tuple with the PriceBeforeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceBeforeDiscount

`func (o *OrderItemDTO) SetPriceBeforeDiscount(v float32)`

SetPriceBeforeDiscount sets PriceBeforeDiscount field to given value.

### HasPriceBeforeDiscount

`func (o *OrderItemDTO) HasPriceBeforeDiscount() bool`

HasPriceBeforeDiscount returns a boolean if a field has been set.

### GetCount

`func (o *OrderItemDTO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OrderItemDTO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OrderItemDTO) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *OrderItemDTO) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetVat

`func (o *OrderItemDTO) GetVat() OrderVatType`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *OrderItemDTO) GetVatOk() (*OrderVatType, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *OrderItemDTO) SetVat(v OrderVatType)`

SetVat sets Vat field to given value.

### HasVat

`func (o *OrderItemDTO) HasVat() bool`

HasVat returns a boolean if a field has been set.

### GetShopSku

`func (o *OrderItemDTO) GetShopSku() string`

GetShopSku returns the ShopSku field if non-nil, zero value otherwise.

### GetShopSkuOk

`func (o *OrderItemDTO) GetShopSkuOk() (*string, bool)`

GetShopSkuOk returns a tuple with the ShopSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopSku

`func (o *OrderItemDTO) SetShopSku(v string)`

SetShopSku sets ShopSku field to given value.

### HasShopSku

`func (o *OrderItemDTO) HasShopSku() bool`

HasShopSku returns a boolean if a field has been set.

### GetSubsidy

`func (o *OrderItemDTO) GetSubsidy() float32`

GetSubsidy returns the Subsidy field if non-nil, zero value otherwise.

### GetSubsidyOk

`func (o *OrderItemDTO) GetSubsidyOk() (*float32, bool)`

GetSubsidyOk returns a tuple with the Subsidy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsidy

`func (o *OrderItemDTO) SetSubsidy(v float32)`

SetSubsidy sets Subsidy field to given value.

### HasSubsidy

`func (o *OrderItemDTO) HasSubsidy() bool`

HasSubsidy returns a boolean if a field has been set.

### GetPartnerWarehouseId

`func (o *OrderItemDTO) GetPartnerWarehouseId() string`

GetPartnerWarehouseId returns the PartnerWarehouseId field if non-nil, zero value otherwise.

### GetPartnerWarehouseIdOk

`func (o *OrderItemDTO) GetPartnerWarehouseIdOk() (*string, bool)`

GetPartnerWarehouseIdOk returns a tuple with the PartnerWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerWarehouseId

`func (o *OrderItemDTO) SetPartnerWarehouseId(v string)`

SetPartnerWarehouseId sets PartnerWarehouseId field to given value.

### HasPartnerWarehouseId

`func (o *OrderItemDTO) HasPartnerWarehouseId() bool`

HasPartnerWarehouseId returns a boolean if a field has been set.

### GetPromos

`func (o *OrderItemDTO) GetPromos() []OrderItemPromoDTO`

GetPromos returns the Promos field if non-nil, zero value otherwise.

### GetPromosOk

`func (o *OrderItemDTO) GetPromosOk() (*[]OrderItemPromoDTO, bool)`

GetPromosOk returns a tuple with the Promos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromos

`func (o *OrderItemDTO) SetPromos(v []OrderItemPromoDTO)`

SetPromos sets Promos field to given value.

### HasPromos

`func (o *OrderItemDTO) HasPromos() bool`

HasPromos returns a boolean if a field has been set.

### GetInstances

`func (o *OrderItemDTO) GetInstances() []OrderItemInstanceDTO`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *OrderItemDTO) GetInstancesOk() (*[]OrderItemInstanceDTO, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *OrderItemDTO) SetInstances(v []OrderItemInstanceDTO)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *OrderItemDTO) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetDetails

`func (o *OrderItemDTO) GetDetails() []OrderItemDetailDTO`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *OrderItemDTO) GetDetailsOk() (*[]OrderItemDetailDTO, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *OrderItemDTO) SetDetails(v []OrderItemDetailDTO)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *OrderItemDTO) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetSubsidies

`func (o *OrderItemDTO) GetSubsidies() []OrderItemSubsidyDTO`

GetSubsidies returns the Subsidies field if non-nil, zero value otherwise.

### GetSubsidiesOk

`func (o *OrderItemDTO) GetSubsidiesOk() (*[]OrderItemSubsidyDTO, bool)`

GetSubsidiesOk returns a tuple with the Subsidies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsidies

`func (o *OrderItemDTO) SetSubsidies(v []OrderItemSubsidyDTO)`

SetSubsidies sets Subsidies field to given value.

### HasSubsidies

`func (o *OrderItemDTO) HasSubsidies() bool`

HasSubsidies returns a boolean if a field has been set.

### GetRequiredInstanceTypes

`func (o *OrderItemDTO) GetRequiredInstanceTypes() []OrderItemInstanceType`

GetRequiredInstanceTypes returns the RequiredInstanceTypes field if non-nil, zero value otherwise.

### GetRequiredInstanceTypesOk

`func (o *OrderItemDTO) GetRequiredInstanceTypesOk() (*[]OrderItemInstanceType, bool)`

GetRequiredInstanceTypesOk returns a tuple with the RequiredInstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInstanceTypes

`func (o *OrderItemDTO) SetRequiredInstanceTypes(v []OrderItemInstanceType)`

SetRequiredInstanceTypes sets RequiredInstanceTypes field to given value.

### HasRequiredInstanceTypes

`func (o *OrderItemDTO) HasRequiredInstanceTypes() bool`

HasRequiredInstanceTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


