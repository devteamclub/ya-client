# AddOffersToArchiveErrorDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | 
**Error** | [**AddOffersToArchiveErrorType**](AddOffersToArchiveErrorType.md) |  | 

## Methods

### NewAddOffersToArchiveErrorDTO

`func NewAddOffersToArchiveErrorDTO(offerId string, error_ AddOffersToArchiveErrorType, ) *AddOffersToArchiveErrorDTO`

NewAddOffersToArchiveErrorDTO instantiates a new AddOffersToArchiveErrorDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOffersToArchiveErrorDTOWithDefaults

`func NewAddOffersToArchiveErrorDTOWithDefaults() *AddOffersToArchiveErrorDTO`

NewAddOffersToArchiveErrorDTOWithDefaults instantiates a new AddOffersToArchiveErrorDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *AddOffersToArchiveErrorDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *AddOffersToArchiveErrorDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *AddOffersToArchiveErrorDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetError

`func (o *AddOffersToArchiveErrorDTO) GetError() AddOffersToArchiveErrorType`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AddOffersToArchiveErrorDTO) GetErrorOk() (*AddOffersToArchiveErrorType, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AddOffersToArchiveErrorDTO) SetError(v AddOffersToArchiveErrorType)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


