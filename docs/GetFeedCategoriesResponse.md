# GetFeedCategoriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]FeedCategoryDTO**](FeedCategoryDTO.md) | Список категорий. | [optional] 
**Pager** | Pointer to [**FlippingPagerDTO**](FlippingPagerDTO.md) |  | [optional] 

## Methods

### NewGetFeedCategoriesResponse

`func NewGetFeedCategoriesResponse() *GetFeedCategoriesResponse`

NewGetFeedCategoriesResponse instantiates a new GetFeedCategoriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFeedCategoriesResponseWithDefaults

`func NewGetFeedCategoriesResponseWithDefaults() *GetFeedCategoriesResponse`

NewGetFeedCategoriesResponseWithDefaults instantiates a new GetFeedCategoriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *GetFeedCategoriesResponse) GetCategories() []FeedCategoryDTO`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *GetFeedCategoriesResponse) GetCategoriesOk() (*[]FeedCategoryDTO, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *GetFeedCategoriesResponse) SetCategories(v []FeedCategoryDTO)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *GetFeedCategoriesResponse) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetPager

`func (o *GetFeedCategoriesResponse) GetPager() FlippingPagerDTO`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *GetFeedCategoriesResponse) GetPagerOk() (*FlippingPagerDTO, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *GetFeedCategoriesResponse) SetPager(v FlippingPagerDTO)`

SetPager sets Pager field to given value.

### HasPager

`func (o *GetFeedCategoriesResponse) HasPager() bool`

HasPager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


