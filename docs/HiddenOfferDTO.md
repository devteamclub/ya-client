# HiddenOfferDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedId** | Pointer to **int64** |  | [optional] 
**OfferId** | Pointer to **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | [optional] 
**TtlInHours** | Pointer to **int32** | {% note alert \&quot;Это поле устарело\&quot; %}  Не используйте его — это может привести к ошибкам.  {% endnote %}  Количество часов до возобновления показа товара.  Минимальное значение — &#x60;1&#x60;.  Максимальное значение — &#x60;720&#x60;.  Значение по умолчанию — &#x60;48&#x60;.  Используется только совместно с параметром &#x60;priority&#x3D;\&quot;true\&quot;&#x60;.  | [optional] 
**Comment** | Pointer to **string** | {% note alert \&quot;Это поле устарело\&quot; %}  Не используйте его — это может привести к ошибкам.  {% endnote %}  Комментарий магазина.  Максимальная длина комментария — 100 символов.  Используется только совместно с параметром &#x60;priority&#x3D;\&quot;true\&quot;&#x60;. Если скрытие было без приоритета, комментарий не вернется в ответе.  | [optional] 
**MarketSku** | Pointer to **int64** | SKU на Маркете. | [optional] 
**Priority** | Pointer to **bool** | {% note alert \&quot;Это поле устарело\&quot; %}  Не используйте его — это может привести к ошибкам.  {% endnote %}  Приоритет скрытия предложений через API над скрытием в личном кабинете. Параметр принимает только значение true.  | [optional] 

## Methods

### NewHiddenOfferDTO

`func NewHiddenOfferDTO() *HiddenOfferDTO`

NewHiddenOfferDTO instantiates a new HiddenOfferDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiddenOfferDTOWithDefaults

`func NewHiddenOfferDTOWithDefaults() *HiddenOfferDTO`

NewHiddenOfferDTOWithDefaults instantiates a new HiddenOfferDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedId

`func (o *HiddenOfferDTO) GetFeedId() int64`

GetFeedId returns the FeedId field if non-nil, zero value otherwise.

### GetFeedIdOk

`func (o *HiddenOfferDTO) GetFeedIdOk() (*int64, bool)`

GetFeedIdOk returns a tuple with the FeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedId

`func (o *HiddenOfferDTO) SetFeedId(v int64)`

SetFeedId sets FeedId field to given value.

### HasFeedId

`func (o *HiddenOfferDTO) HasFeedId() bool`

HasFeedId returns a boolean if a field has been set.

### GetOfferId

`func (o *HiddenOfferDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *HiddenOfferDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *HiddenOfferDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *HiddenOfferDTO) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetTtlInHours

`func (o *HiddenOfferDTO) GetTtlInHours() int32`

GetTtlInHours returns the TtlInHours field if non-nil, zero value otherwise.

### GetTtlInHoursOk

`func (o *HiddenOfferDTO) GetTtlInHoursOk() (*int32, bool)`

GetTtlInHoursOk returns a tuple with the TtlInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlInHours

`func (o *HiddenOfferDTO) SetTtlInHours(v int32)`

SetTtlInHours sets TtlInHours field to given value.

### HasTtlInHours

`func (o *HiddenOfferDTO) HasTtlInHours() bool`

HasTtlInHours returns a boolean if a field has been set.

### GetComment

`func (o *HiddenOfferDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *HiddenOfferDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *HiddenOfferDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *HiddenOfferDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetMarketSku

`func (o *HiddenOfferDTO) GetMarketSku() int64`

GetMarketSku returns the MarketSku field if non-nil, zero value otherwise.

### GetMarketSkuOk

`func (o *HiddenOfferDTO) GetMarketSkuOk() (*int64, bool)`

GetMarketSkuOk returns a tuple with the MarketSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketSku

`func (o *HiddenOfferDTO) SetMarketSku(v int64)`

SetMarketSku sets MarketSku field to given value.

### HasMarketSku

`func (o *HiddenOfferDTO) HasMarketSku() bool`

HasMarketSku returns a boolean if a field has been set.

### GetPriority

`func (o *HiddenOfferDTO) GetPriority() bool`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *HiddenOfferDTO) GetPriorityOk() (*bool, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *HiddenOfferDTO) SetPriority(v bool)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *HiddenOfferDTO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


