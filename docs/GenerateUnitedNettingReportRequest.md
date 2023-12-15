# GenerateUnitedNettingReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | **int64** | Идентификатор бизнеса. | 
**DateTimeFrom** | Pointer to **time.Time** | Начало периода, включительно. | [optional] 
**DateTimeTo** | Pointer to **time.Time** | Конец периода, включительно. Максимальный период — 1 год. | [optional] 
**BankOrderId** | Pointer to **int64** | Номер платежного поручения. | [optional] 
**BankOrderDateTime** | Pointer to **time.Time** | Дата платежного поручения. | [optional] 
**PlacementPrograms** | Pointer to [**[]PlacementType**](PlacementType.md) | Список моделей, которые нужны в отчете.  | [optional] 
**Inns** | Pointer to **[]string** | Список ИНН, которые нужны в отчете. | [optional] 
**CampaignIds** | Pointer to **[]int64** | Список магазинов, которые нужны в отчете. | [optional] 

## Methods

### NewGenerateUnitedNettingReportRequest

`func NewGenerateUnitedNettingReportRequest(businessId int64, ) *GenerateUnitedNettingReportRequest`

NewGenerateUnitedNettingReportRequest instantiates a new GenerateUnitedNettingReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateUnitedNettingReportRequestWithDefaults

`func NewGenerateUnitedNettingReportRequestWithDefaults() *GenerateUnitedNettingReportRequest`

NewGenerateUnitedNettingReportRequestWithDefaults instantiates a new GenerateUnitedNettingReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *GenerateUnitedNettingReportRequest) GetBusinessId() int64`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *GenerateUnitedNettingReportRequest) GetBusinessIdOk() (*int64, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *GenerateUnitedNettingReportRequest) SetBusinessId(v int64)`

SetBusinessId sets BusinessId field to given value.


### GetDateTimeFrom

`func (o *GenerateUnitedNettingReportRequest) GetDateTimeFrom() time.Time`

GetDateTimeFrom returns the DateTimeFrom field if non-nil, zero value otherwise.

### GetDateTimeFromOk

`func (o *GenerateUnitedNettingReportRequest) GetDateTimeFromOk() (*time.Time, bool)`

GetDateTimeFromOk returns a tuple with the DateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFrom

`func (o *GenerateUnitedNettingReportRequest) SetDateTimeFrom(v time.Time)`

SetDateTimeFrom sets DateTimeFrom field to given value.

### HasDateTimeFrom

`func (o *GenerateUnitedNettingReportRequest) HasDateTimeFrom() bool`

HasDateTimeFrom returns a boolean if a field has been set.

### GetDateTimeTo

`func (o *GenerateUnitedNettingReportRequest) GetDateTimeTo() time.Time`

GetDateTimeTo returns the DateTimeTo field if non-nil, zero value otherwise.

### GetDateTimeToOk

`func (o *GenerateUnitedNettingReportRequest) GetDateTimeToOk() (*time.Time, bool)`

GetDateTimeToOk returns a tuple with the DateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTo

`func (o *GenerateUnitedNettingReportRequest) SetDateTimeTo(v time.Time)`

SetDateTimeTo sets DateTimeTo field to given value.

### HasDateTimeTo

`func (o *GenerateUnitedNettingReportRequest) HasDateTimeTo() bool`

HasDateTimeTo returns a boolean if a field has been set.

### GetBankOrderId

`func (o *GenerateUnitedNettingReportRequest) GetBankOrderId() int64`

GetBankOrderId returns the BankOrderId field if non-nil, zero value otherwise.

### GetBankOrderIdOk

`func (o *GenerateUnitedNettingReportRequest) GetBankOrderIdOk() (*int64, bool)`

GetBankOrderIdOk returns a tuple with the BankOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankOrderId

`func (o *GenerateUnitedNettingReportRequest) SetBankOrderId(v int64)`

SetBankOrderId sets BankOrderId field to given value.

### HasBankOrderId

`func (o *GenerateUnitedNettingReportRequest) HasBankOrderId() bool`

HasBankOrderId returns a boolean if a field has been set.

### GetBankOrderDateTime

`func (o *GenerateUnitedNettingReportRequest) GetBankOrderDateTime() time.Time`

GetBankOrderDateTime returns the BankOrderDateTime field if non-nil, zero value otherwise.

### GetBankOrderDateTimeOk

`func (o *GenerateUnitedNettingReportRequest) GetBankOrderDateTimeOk() (*time.Time, bool)`

GetBankOrderDateTimeOk returns a tuple with the BankOrderDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankOrderDateTime

`func (o *GenerateUnitedNettingReportRequest) SetBankOrderDateTime(v time.Time)`

SetBankOrderDateTime sets BankOrderDateTime field to given value.

### HasBankOrderDateTime

`func (o *GenerateUnitedNettingReportRequest) HasBankOrderDateTime() bool`

HasBankOrderDateTime returns a boolean if a field has been set.

### GetPlacementPrograms

`func (o *GenerateUnitedNettingReportRequest) GetPlacementPrograms() []PlacementType`

GetPlacementPrograms returns the PlacementPrograms field if non-nil, zero value otherwise.

### GetPlacementProgramsOk

`func (o *GenerateUnitedNettingReportRequest) GetPlacementProgramsOk() (*[]PlacementType, bool)`

GetPlacementProgramsOk returns a tuple with the PlacementPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementPrograms

`func (o *GenerateUnitedNettingReportRequest) SetPlacementPrograms(v []PlacementType)`

SetPlacementPrograms sets PlacementPrograms field to given value.

### HasPlacementPrograms

`func (o *GenerateUnitedNettingReportRequest) HasPlacementPrograms() bool`

HasPlacementPrograms returns a boolean if a field has been set.

### GetInns

`func (o *GenerateUnitedNettingReportRequest) GetInns() []string`

GetInns returns the Inns field if non-nil, zero value otherwise.

### GetInnsOk

`func (o *GenerateUnitedNettingReportRequest) GetInnsOk() (*[]string, bool)`

GetInnsOk returns a tuple with the Inns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInns

`func (o *GenerateUnitedNettingReportRequest) SetInns(v []string)`

SetInns sets Inns field to given value.

### HasInns

`func (o *GenerateUnitedNettingReportRequest) HasInns() bool`

HasInns returns a boolean if a field has been set.

### GetCampaignIds

`func (o *GenerateUnitedNettingReportRequest) GetCampaignIds() []int64`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *GenerateUnitedNettingReportRequest) GetCampaignIdsOk() (*[]int64, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *GenerateUnitedNettingReportRequest) SetCampaignIds(v []int64)`

SetCampaignIds sets CampaignIds field to given value.

### HasCampaignIds

`func (o *GenerateUnitedNettingReportRequest) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


