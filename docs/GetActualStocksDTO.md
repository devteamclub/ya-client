# GetActualStocksDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skus** | Pointer to [**[]StockDTO**](StockDTO.md) | Информация о стоках. | [optional] 

## Methods

### NewGetActualStocksDTO

`func NewGetActualStocksDTO() *GetActualStocksDTO`

NewGetActualStocksDTO instantiates a new GetActualStocksDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActualStocksDTOWithDefaults

`func NewGetActualStocksDTOWithDefaults() *GetActualStocksDTO`

NewGetActualStocksDTOWithDefaults instantiates a new GetActualStocksDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkus

`func (o *GetActualStocksDTO) GetSkus() []StockDTO`

GetSkus returns the Skus field if non-nil, zero value otherwise.

### GetSkusOk

`func (o *GetActualStocksDTO) GetSkusOk() (*[]StockDTO, bool)`

GetSkusOk returns a tuple with the Skus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkus

`func (o *GetActualStocksDTO) SetSkus(v []StockDTO)`

SetSkus sets Skus field to given value.

### HasSkus

`func (o *GetActualStocksDTO) HasSkus() bool`

HasSkus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


