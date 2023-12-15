# ReturnItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketSku** | Pointer to **int64** | SKU на Маркете. | [optional] 
**ShopSku** | **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | 
**Count** | **int64** | Количество товаров. | 
**Decisions** | Pointer to [**[]ReturnDecisionDTO**](ReturnDecisionDTO.md) | Список решений по возврату. | [optional] 
**Instances** | Pointer to [**[]ReturnInstanceDTO**](ReturnInstanceDTO.md) | Список логистических позиций возврата. | [optional] 
**Tracks** | Pointer to [**[]TrackDTO**](TrackDTO.md) | Список трек-кодов для почтовых отправлений. | [optional] 

## Methods

### NewReturnItemDTO

`func NewReturnItemDTO(shopSku string, count int64, ) *ReturnItemDTO`

NewReturnItemDTO instantiates a new ReturnItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnItemDTOWithDefaults

`func NewReturnItemDTOWithDefaults() *ReturnItemDTO`

NewReturnItemDTOWithDefaults instantiates a new ReturnItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketSku

`func (o *ReturnItemDTO) GetMarketSku() int64`

GetMarketSku returns the MarketSku field if non-nil, zero value otherwise.

### GetMarketSkuOk

`func (o *ReturnItemDTO) GetMarketSkuOk() (*int64, bool)`

GetMarketSkuOk returns a tuple with the MarketSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketSku

`func (o *ReturnItemDTO) SetMarketSku(v int64)`

SetMarketSku sets MarketSku field to given value.

### HasMarketSku

`func (o *ReturnItemDTO) HasMarketSku() bool`

HasMarketSku returns a boolean if a field has been set.

### GetShopSku

`func (o *ReturnItemDTO) GetShopSku() string`

GetShopSku returns the ShopSku field if non-nil, zero value otherwise.

### GetShopSkuOk

`func (o *ReturnItemDTO) GetShopSkuOk() (*string, bool)`

GetShopSkuOk returns a tuple with the ShopSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopSku

`func (o *ReturnItemDTO) SetShopSku(v string)`

SetShopSku sets ShopSku field to given value.


### GetCount

`func (o *ReturnItemDTO) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ReturnItemDTO) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ReturnItemDTO) SetCount(v int64)`

SetCount sets Count field to given value.


### GetDecisions

`func (o *ReturnItemDTO) GetDecisions() []ReturnDecisionDTO`

GetDecisions returns the Decisions field if non-nil, zero value otherwise.

### GetDecisionsOk

`func (o *ReturnItemDTO) GetDecisionsOk() (*[]ReturnDecisionDTO, bool)`

GetDecisionsOk returns a tuple with the Decisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisions

`func (o *ReturnItemDTO) SetDecisions(v []ReturnDecisionDTO)`

SetDecisions sets Decisions field to given value.

### HasDecisions

`func (o *ReturnItemDTO) HasDecisions() bool`

HasDecisions returns a boolean if a field has been set.

### GetInstances

`func (o *ReturnItemDTO) GetInstances() []ReturnInstanceDTO`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ReturnItemDTO) GetInstancesOk() (*[]ReturnInstanceDTO, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ReturnItemDTO) SetInstances(v []ReturnInstanceDTO)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ReturnItemDTO) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetTracks

`func (o *ReturnItemDTO) GetTracks() []TrackDTO`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *ReturnItemDTO) GetTracksOk() (*[]TrackDTO, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *ReturnItemDTO) SetTracks(v []TrackDTO)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *ReturnItemDTO) HasTracks() bool`

HasTracks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


