# StockItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Количество доступного товара.  | 
**Type** | [**StockType**](StockType.md) |  | 
**UpdatedAt** | **time.Time** | Дата и время последнего обновления информации об остатках указанного типа.  Формат даты и времени: ISO 8601 со смещением относительно UTC. Например, &#x60;2017-11-21T00:42:42+03:00&#x60;.  | 

## Methods

### NewStockItemDTO

`func NewStockItemDTO(count int64, type_ StockType, updatedAt time.Time, ) *StockItemDTO`

NewStockItemDTO instantiates a new StockItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockItemDTOWithDefaults

`func NewStockItemDTOWithDefaults() *StockItemDTO`

NewStockItemDTOWithDefaults instantiates a new StockItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *StockItemDTO) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StockItemDTO) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StockItemDTO) SetCount(v int64)`

SetCount sets Count field to given value.


### GetType

`func (o *StockItemDTO) GetType() StockType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockItemDTO) GetTypeOk() (*StockType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockItemDTO) SetType(v StockType)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *StockItemDTO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StockItemDTO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StockItemDTO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


