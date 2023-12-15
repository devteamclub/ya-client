# BriefOrderItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Идентификатор товара в заказе. | [optional] 
**Vat** | Pointer to [**OrderVatType**](OrderVatType.md) |  | [optional] 
**Count** | Pointer to **int32** | Количество единиц товара. | [optional] 
**Price** | Pointer to **float32** | Цена на товар. Указана в той валюте, которая была задана в каталоге. Разделитель целой и дробной части — точка.  | [optional] 
**FeedId** | Pointer to **int64** | Идентификатор каталога, в котором указан товар. | [optional] 
**FeedCategoryId** | Pointer to **string** | Идентификатор категории, в которую входит товар. | [optional] 
**OfferName** | Pointer to **string** | Название товара. | [optional] 
**OfferId** | Pointer to **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | [optional] 
**Instances** | Pointer to [**[]OrderItemInstanceDTO**](OrderItemInstanceDTO.md) | Переданные вами коды маркировки. | [optional] 

## Methods

### NewBriefOrderItemDTO

`func NewBriefOrderItemDTO() *BriefOrderItemDTO`

NewBriefOrderItemDTO instantiates a new BriefOrderItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefOrderItemDTOWithDefaults

`func NewBriefOrderItemDTOWithDefaults() *BriefOrderItemDTO`

NewBriefOrderItemDTOWithDefaults instantiates a new BriefOrderItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefOrderItemDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefOrderItemDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefOrderItemDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BriefOrderItemDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVat

`func (o *BriefOrderItemDTO) GetVat() OrderVatType`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *BriefOrderItemDTO) GetVatOk() (*OrderVatType, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *BriefOrderItemDTO) SetVat(v OrderVatType)`

SetVat sets Vat field to given value.

### HasVat

`func (o *BriefOrderItemDTO) HasVat() bool`

HasVat returns a boolean if a field has been set.

### GetCount

`func (o *BriefOrderItemDTO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BriefOrderItemDTO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BriefOrderItemDTO) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BriefOrderItemDTO) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPrice

`func (o *BriefOrderItemDTO) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BriefOrderItemDTO) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BriefOrderItemDTO) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BriefOrderItemDTO) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetFeedId

`func (o *BriefOrderItemDTO) GetFeedId() int64`

GetFeedId returns the FeedId field if non-nil, zero value otherwise.

### GetFeedIdOk

`func (o *BriefOrderItemDTO) GetFeedIdOk() (*int64, bool)`

GetFeedIdOk returns a tuple with the FeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedId

`func (o *BriefOrderItemDTO) SetFeedId(v int64)`

SetFeedId sets FeedId field to given value.

### HasFeedId

`func (o *BriefOrderItemDTO) HasFeedId() bool`

HasFeedId returns a boolean if a field has been set.

### GetFeedCategoryId

`func (o *BriefOrderItemDTO) GetFeedCategoryId() string`

GetFeedCategoryId returns the FeedCategoryId field if non-nil, zero value otherwise.

### GetFeedCategoryIdOk

`func (o *BriefOrderItemDTO) GetFeedCategoryIdOk() (*string, bool)`

GetFeedCategoryIdOk returns a tuple with the FeedCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedCategoryId

`func (o *BriefOrderItemDTO) SetFeedCategoryId(v string)`

SetFeedCategoryId sets FeedCategoryId field to given value.

### HasFeedCategoryId

`func (o *BriefOrderItemDTO) HasFeedCategoryId() bool`

HasFeedCategoryId returns a boolean if a field has been set.

### GetOfferName

`func (o *BriefOrderItemDTO) GetOfferName() string`

GetOfferName returns the OfferName field if non-nil, zero value otherwise.

### GetOfferNameOk

`func (o *BriefOrderItemDTO) GetOfferNameOk() (*string, bool)`

GetOfferNameOk returns a tuple with the OfferName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferName

`func (o *BriefOrderItemDTO) SetOfferName(v string)`

SetOfferName sets OfferName field to given value.

### HasOfferName

`func (o *BriefOrderItemDTO) HasOfferName() bool`

HasOfferName returns a boolean if a field has been set.

### GetOfferId

`func (o *BriefOrderItemDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *BriefOrderItemDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *BriefOrderItemDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *BriefOrderItemDTO) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetInstances

`func (o *BriefOrderItemDTO) GetInstances() []OrderItemInstanceDTO`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *BriefOrderItemDTO) GetInstancesOk() (*[]OrderItemInstanceDTO, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *BriefOrderItemDTO) SetInstances(v []OrderItemInstanceDTO)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *BriefOrderItemDTO) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


