# GetPriceWithVatDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **float32** | Цена на товар. | [optional] 
**DiscountBase** | Pointer to **float32** | Цена на товар без скидки. | [optional] 
**CurrencyId** | Pointer to **string** | Валюта, в которой указана цена на товар.  Возможно только значение RUR — российский рубль.  | [optional] 
**Vat** | Pointer to **int32** | Идентификатор ставки НДС, применяемой для товара:  * 2 — 10%. * 5 — 0%. * 6 — не облагается НДС. * 7 — 20%.  Если параметр не указан, используется ставка НДС, установленная в личном кабинете магазина.  | [optional] 
**UpdatedAt** | **time.Time** | Время последнего обновления. | 

## Methods

### NewGetPriceWithVatDTO

`func NewGetPriceWithVatDTO(updatedAt time.Time, ) *GetPriceWithVatDTO`

NewGetPriceWithVatDTO instantiates a new GetPriceWithVatDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPriceWithVatDTOWithDefaults

`func NewGetPriceWithVatDTOWithDefaults() *GetPriceWithVatDTO`

NewGetPriceWithVatDTOWithDefaults instantiates a new GetPriceWithVatDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *GetPriceWithVatDTO) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetPriceWithVatDTO) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetPriceWithVatDTO) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetPriceWithVatDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDiscountBase

`func (o *GetPriceWithVatDTO) GetDiscountBase() float32`

GetDiscountBase returns the DiscountBase field if non-nil, zero value otherwise.

### GetDiscountBaseOk

`func (o *GetPriceWithVatDTO) GetDiscountBaseOk() (*float32, bool)`

GetDiscountBaseOk returns a tuple with the DiscountBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountBase

`func (o *GetPriceWithVatDTO) SetDiscountBase(v float32)`

SetDiscountBase sets DiscountBase field to given value.

### HasDiscountBase

`func (o *GetPriceWithVatDTO) HasDiscountBase() bool`

HasDiscountBase returns a boolean if a field has been set.

### GetCurrencyId

`func (o *GetPriceWithVatDTO) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *GetPriceWithVatDTO) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *GetPriceWithVatDTO) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *GetPriceWithVatDTO) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetVat

`func (o *GetPriceWithVatDTO) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *GetPriceWithVatDTO) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *GetPriceWithVatDTO) SetVat(v int32)`

SetVat sets Vat field to given value.

### HasVat

`func (o *GetPriceWithVatDTO) HasVat() bool`

HasVat returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetPriceWithVatDTO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetPriceWithVatDTO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetPriceWithVatDTO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


