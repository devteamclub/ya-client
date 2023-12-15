# UpdateOfferContentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | [optional] 
**Results** | Pointer to [**[]UpdateOfferContentResultDTO**](UpdateOfferContentResultDTO.md) | Ошибки и предупреждения, возникшие при обработке переданных значений. Каждый элемент списка соответствует одному товару.  Поле не передается, если все в порядке.  | [optional] 

## Methods

### NewUpdateOfferContentResponse

`func NewUpdateOfferContentResponse() *UpdateOfferContentResponse`

NewUpdateOfferContentResponse instantiates a new UpdateOfferContentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOfferContentResponseWithDefaults

`func NewUpdateOfferContentResponseWithDefaults() *UpdateOfferContentResponse`

NewUpdateOfferContentResponseWithDefaults instantiates a new UpdateOfferContentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateOfferContentResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateOfferContentResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateOfferContentResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateOfferContentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *UpdateOfferContentResponse) GetResults() []UpdateOfferContentResultDTO`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UpdateOfferContentResponse) GetResultsOk() (*[]UpdateOfferContentResultDTO, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UpdateOfferContentResponse) SetResults(v []UpdateOfferContentResultDTO)`

SetResults sets Results field to given value.

### HasResults

`func (o *UpdateOfferContentResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


