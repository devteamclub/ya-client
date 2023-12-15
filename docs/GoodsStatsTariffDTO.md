# GoodsStatsTariffDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**GoodsStatsTariffType**](GoodsStatsTariffType.md) |  | [optional] 
**Percent** | Pointer to **float32** | Значение тарифа в процентах. | [optional] 
**Amount** | Pointer to **float32** | Значение тарифа в рублях. | [optional] 

## Methods

### NewGoodsStatsTariffDTO

`func NewGoodsStatsTariffDTO() *GoodsStatsTariffDTO`

NewGoodsStatsTariffDTO instantiates a new GoodsStatsTariffDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoodsStatsTariffDTOWithDefaults

`func NewGoodsStatsTariffDTOWithDefaults() *GoodsStatsTariffDTO`

NewGoodsStatsTariffDTOWithDefaults instantiates a new GoodsStatsTariffDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GoodsStatsTariffDTO) GetType() GoodsStatsTariffType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GoodsStatsTariffDTO) GetTypeOk() (*GoodsStatsTariffType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GoodsStatsTariffDTO) SetType(v GoodsStatsTariffType)`

SetType sets Type field to given value.

### HasType

`func (o *GoodsStatsTariffDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPercent

`func (o *GoodsStatsTariffDTO) GetPercent() float32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *GoodsStatsTariffDTO) GetPercentOk() (*float32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *GoodsStatsTariffDTO) SetPercent(v float32)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *GoodsStatsTariffDTO) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetAmount

`func (o *GoodsStatsTariffDTO) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GoodsStatsTariffDTO) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GoodsStatsTariffDTO) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GoodsStatsTariffDTO) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


