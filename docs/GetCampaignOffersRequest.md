# GetCampaignOffersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferIds** | Pointer to **[]string** | Идентификаторы товаров, информация о которых нужна. ⚠️ Не используйте это поле одновременно с фильтрами по статусам карточек, категориям, брендам или тегам. Если вы хотите воспользоваться фильтрами, оставьте поле пустым. | [optional] 
**Statuses** | Pointer to [**[]OfferCampaignStatusType**](OfferCampaignStatusType.md) | Фильтр по статусам товаров.  | [optional] 
**CategoryIds** | Pointer to **[]int32** | Фильтр по категориям на Маркете. | [optional] 
**VendorNames** | Pointer to **[]string** | Фильтр по брендам. | [optional] 
**Tags** | Pointer to **[]string** | Фильтр по тегам. | [optional] 

## Methods

### NewGetCampaignOffersRequest

`func NewGetCampaignOffersRequest() *GetCampaignOffersRequest`

NewGetCampaignOffersRequest instantiates a new GetCampaignOffersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignOffersRequestWithDefaults

`func NewGetCampaignOffersRequestWithDefaults() *GetCampaignOffersRequest`

NewGetCampaignOffersRequestWithDefaults instantiates a new GetCampaignOffersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferIds

`func (o *GetCampaignOffersRequest) GetOfferIds() []string`

GetOfferIds returns the OfferIds field if non-nil, zero value otherwise.

### GetOfferIdsOk

`func (o *GetCampaignOffersRequest) GetOfferIdsOk() (*[]string, bool)`

GetOfferIdsOk returns a tuple with the OfferIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferIds

`func (o *GetCampaignOffersRequest) SetOfferIds(v []string)`

SetOfferIds sets OfferIds field to given value.

### HasOfferIds

`func (o *GetCampaignOffersRequest) HasOfferIds() bool`

HasOfferIds returns a boolean if a field has been set.

### GetStatuses

`func (o *GetCampaignOffersRequest) GetStatuses() []OfferCampaignStatusType`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *GetCampaignOffersRequest) GetStatusesOk() (*[]OfferCampaignStatusType, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *GetCampaignOffersRequest) SetStatuses(v []OfferCampaignStatusType)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *GetCampaignOffersRequest) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetCategoryIds

`func (o *GetCampaignOffersRequest) GetCategoryIds() []int32`

GetCategoryIds returns the CategoryIds field if non-nil, zero value otherwise.

### GetCategoryIdsOk

`func (o *GetCampaignOffersRequest) GetCategoryIdsOk() (*[]int32, bool)`

GetCategoryIdsOk returns a tuple with the CategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryIds

`func (o *GetCampaignOffersRequest) SetCategoryIds(v []int32)`

SetCategoryIds sets CategoryIds field to given value.

### HasCategoryIds

`func (o *GetCampaignOffersRequest) HasCategoryIds() bool`

HasCategoryIds returns a boolean if a field has been set.

### GetVendorNames

`func (o *GetCampaignOffersRequest) GetVendorNames() []string`

GetVendorNames returns the VendorNames field if non-nil, zero value otherwise.

### GetVendorNamesOk

`func (o *GetCampaignOffersRequest) GetVendorNamesOk() (*[]string, bool)`

GetVendorNamesOk returns a tuple with the VendorNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNames

`func (o *GetCampaignOffersRequest) SetVendorNames(v []string)`

SetVendorNames sets VendorNames field to given value.

### HasVendorNames

`func (o *GetCampaignOffersRequest) HasVendorNames() bool`

HasVendorNames returns a boolean if a field has been set.

### GetTags

`func (o *GetCampaignOffersRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetCampaignOffersRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetCampaignOffersRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetCampaignOffersRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


