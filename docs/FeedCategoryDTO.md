# FeedCategoryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedId** | Pointer to **int64** | Идентификатор прайс-листа. | [optional] 
**Id** | Pointer to **string** | Идентификатор категории. | [optional] 
**Name** | Pointer to **string** | Название категории. | [optional] 
**ParentId** | Pointer to **string** | Идентификатор родительской категории. Не выводится, если категория — корневая.  | [optional] 

## Methods

### NewFeedCategoryDTO

`func NewFeedCategoryDTO() *FeedCategoryDTO`

NewFeedCategoryDTO instantiates a new FeedCategoryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedCategoryDTOWithDefaults

`func NewFeedCategoryDTOWithDefaults() *FeedCategoryDTO`

NewFeedCategoryDTOWithDefaults instantiates a new FeedCategoryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedId

`func (o *FeedCategoryDTO) GetFeedId() int64`

GetFeedId returns the FeedId field if non-nil, zero value otherwise.

### GetFeedIdOk

`func (o *FeedCategoryDTO) GetFeedIdOk() (*int64, bool)`

GetFeedIdOk returns a tuple with the FeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedId

`func (o *FeedCategoryDTO) SetFeedId(v int64)`

SetFeedId sets FeedId field to given value.

### HasFeedId

`func (o *FeedCategoryDTO) HasFeedId() bool`

HasFeedId returns a boolean if a field has been set.

### GetId

`func (o *FeedCategoryDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeedCategoryDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeedCategoryDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FeedCategoryDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FeedCategoryDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeedCategoryDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeedCategoryDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FeedCategoryDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentId

`func (o *FeedCategoryDTO) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *FeedCategoryDTO) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *FeedCategoryDTO) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *FeedCategoryDTO) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


