# GetCampaignCategoriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]FeedCategoryDTO**](FeedCategoryDTO.md) | Список категорий. | [optional] 
**Pager** | Pointer to [**FlippingPagerDTO**](FlippingPagerDTO.md) |  | [optional] 

## Methods

### NewGetCampaignCategoriesResponse

`func NewGetCampaignCategoriesResponse() *GetCampaignCategoriesResponse`

NewGetCampaignCategoriesResponse instantiates a new GetCampaignCategoriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignCategoriesResponseWithDefaults

`func NewGetCampaignCategoriesResponseWithDefaults() *GetCampaignCategoriesResponse`

NewGetCampaignCategoriesResponseWithDefaults instantiates a new GetCampaignCategoriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *GetCampaignCategoriesResponse) GetCategories() []FeedCategoryDTO`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *GetCampaignCategoriesResponse) GetCategoriesOk() (*[]FeedCategoryDTO, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *GetCampaignCategoriesResponse) SetCategories(v []FeedCategoryDTO)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *GetCampaignCategoriesResponse) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetPager

`func (o *GetCampaignCategoriesResponse) GetPager() FlippingPagerDTO`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *GetCampaignCategoriesResponse) GetPagerOk() (*FlippingPagerDTO, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *GetCampaignCategoriesResponse) SetPager(v FlippingPagerDTO)`

SetPager sets Pager field to given value.

### HasPager

`func (o *GetCampaignCategoriesResponse) HasPager() bool`

HasPager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


