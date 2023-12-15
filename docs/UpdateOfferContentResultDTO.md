# UpdateOfferContentResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | 
**Errors** | Pointer to [**[]OfferContentErrorDTO**](OfferContentErrorDTO.md) | Ошибки, препятствующие отправке контента в каталог. | [optional] 
**Warnings** | Pointer to [**[]OfferContentErrorDTO**](OfferContentErrorDTO.md) | Предупреждения, не препятствующие отправке контента в каталог. | [optional] 

## Methods

### NewUpdateOfferContentResultDTO

`func NewUpdateOfferContentResultDTO(offerId string, ) *UpdateOfferContentResultDTO`

NewUpdateOfferContentResultDTO instantiates a new UpdateOfferContentResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOfferContentResultDTOWithDefaults

`func NewUpdateOfferContentResultDTOWithDefaults() *UpdateOfferContentResultDTO`

NewUpdateOfferContentResultDTOWithDefaults instantiates a new UpdateOfferContentResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *UpdateOfferContentResultDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *UpdateOfferContentResultDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *UpdateOfferContentResultDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetErrors

`func (o *UpdateOfferContentResultDTO) GetErrors() []OfferContentErrorDTO`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UpdateOfferContentResultDTO) GetErrorsOk() (*[]OfferContentErrorDTO, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UpdateOfferContentResultDTO) SetErrors(v []OfferContentErrorDTO)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *UpdateOfferContentResultDTO) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *UpdateOfferContentResultDTO) GetWarnings() []OfferContentErrorDTO`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *UpdateOfferContentResultDTO) GetWarningsOk() (*[]OfferContentErrorDTO, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *UpdateOfferContentResultDTO) SetWarnings(v []OfferContentErrorDTO)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *UpdateOfferContentResultDTO) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


