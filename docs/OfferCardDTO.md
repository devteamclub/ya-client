# OfferCardDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | 
**Mapping** | Pointer to [**GetMappingDTO**](GetMappingDTO.md) |  | [optional] 
**CardStatus** | Pointer to [**OfferCardStatusType**](OfferCardStatusType.md) |  | [optional] 
**ContentRating** | Pointer to **int32** | Процент заполненности карточки. | [optional] 
**Recommendations** | Pointer to [**[]OfferCardRecommendationDTO**](OfferCardRecommendationDTO.md) | Список рекомендаций к заполнению карточки.  Рекомендации Маркета помогают заполнять карточку так, чтобы покупателям было проще найти ваш товар и решиться на покупку.  | [optional] 
**Errors** | Pointer to [**[]OfferErrorDTO**](OfferErrorDTO.md) | Ошибки в контенте, препятствующие размещению товара на витрине. | [optional] 
**Warnings** | Pointer to [**[]OfferErrorDTO**](OfferErrorDTO.md) | Связанные с контентом предупреждения, не препятствующие размещению товара на витрине. | [optional] 

## Methods

### NewOfferCardDTO

`func NewOfferCardDTO(offerId string, ) *OfferCardDTO`

NewOfferCardDTO instantiates a new OfferCardDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferCardDTOWithDefaults

`func NewOfferCardDTOWithDefaults() *OfferCardDTO`

NewOfferCardDTOWithDefaults instantiates a new OfferCardDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *OfferCardDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *OfferCardDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *OfferCardDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetMapping

`func (o *OfferCardDTO) GetMapping() GetMappingDTO`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *OfferCardDTO) GetMappingOk() (*GetMappingDTO, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *OfferCardDTO) SetMapping(v GetMappingDTO)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *OfferCardDTO) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetCardStatus

`func (o *OfferCardDTO) GetCardStatus() OfferCardStatusType`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *OfferCardDTO) GetCardStatusOk() (*OfferCardStatusType, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *OfferCardDTO) SetCardStatus(v OfferCardStatusType)`

SetCardStatus sets CardStatus field to given value.

### HasCardStatus

`func (o *OfferCardDTO) HasCardStatus() bool`

HasCardStatus returns a boolean if a field has been set.

### GetContentRating

`func (o *OfferCardDTO) GetContentRating() int32`

GetContentRating returns the ContentRating field if non-nil, zero value otherwise.

### GetContentRatingOk

`func (o *OfferCardDTO) GetContentRatingOk() (*int32, bool)`

GetContentRatingOk returns a tuple with the ContentRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRating

`func (o *OfferCardDTO) SetContentRating(v int32)`

SetContentRating sets ContentRating field to given value.

### HasContentRating

`func (o *OfferCardDTO) HasContentRating() bool`

HasContentRating returns a boolean if a field has been set.

### GetRecommendations

`func (o *OfferCardDTO) GetRecommendations() []OfferCardRecommendationDTO`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *OfferCardDTO) GetRecommendationsOk() (*[]OfferCardRecommendationDTO, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *OfferCardDTO) SetRecommendations(v []OfferCardRecommendationDTO)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *OfferCardDTO) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### GetErrors

`func (o *OfferCardDTO) GetErrors() []OfferErrorDTO`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *OfferCardDTO) GetErrorsOk() (*[]OfferErrorDTO, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *OfferCardDTO) SetErrors(v []OfferErrorDTO)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *OfferCardDTO) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *OfferCardDTO) GetWarnings() []OfferErrorDTO`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *OfferCardDTO) GetWarningsOk() (*[]OfferErrorDTO, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *OfferCardDTO) SetWarnings(v []OfferErrorDTO)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *OfferCardDTO) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


