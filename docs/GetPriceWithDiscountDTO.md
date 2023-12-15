# GetPriceWithDiscountDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **float32** | Значение. | 
**CurrencyId** | **string** | Валюта.  Если &#x60;BasePriceDTO&#x60; присутствует в запросе, указывайте &#x60;RUR&#x60; — российский рубль.  | 
**DiscountBase** | Pointer to **float32** | Цена до скидки. | [optional] 
**UpdatedAt** | **time.Time** | Время последнего обновления. | 

## Methods

### NewGetPriceWithDiscountDTO

`func NewGetPriceWithDiscountDTO(value float32, currencyId string, updatedAt time.Time, ) *GetPriceWithDiscountDTO`

NewGetPriceWithDiscountDTO instantiates a new GetPriceWithDiscountDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPriceWithDiscountDTOWithDefaults

`func NewGetPriceWithDiscountDTOWithDefaults() *GetPriceWithDiscountDTO`

NewGetPriceWithDiscountDTOWithDefaults instantiates a new GetPriceWithDiscountDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *GetPriceWithDiscountDTO) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetPriceWithDiscountDTO) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetPriceWithDiscountDTO) SetValue(v float32)`

SetValue sets Value field to given value.


### GetCurrencyId

`func (o *GetPriceWithDiscountDTO) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *GetPriceWithDiscountDTO) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *GetPriceWithDiscountDTO) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.


### GetDiscountBase

`func (o *GetPriceWithDiscountDTO) GetDiscountBase() float32`

GetDiscountBase returns the DiscountBase field if non-nil, zero value otherwise.

### GetDiscountBaseOk

`func (o *GetPriceWithDiscountDTO) GetDiscountBaseOk() (*float32, bool)`

GetDiscountBaseOk returns a tuple with the DiscountBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountBase

`func (o *GetPriceWithDiscountDTO) SetDiscountBase(v float32)`

SetDiscountBase sets DiscountBase field to given value.

### HasDiscountBase

`func (o *GetPriceWithDiscountDTO) HasDiscountBase() bool`

HasDiscountBase returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetPriceWithDiscountDTO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetPriceWithDiscountDTO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetPriceWithDiscountDTO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


