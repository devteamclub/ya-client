# GenerateStocksOnWarehousesReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int64** | Идентификатор магазина. | 
**WarehouseIds** | Pointer to **[]int64** | Фильтр по идентификаторам складов (только FBY). Чтобы узнать идентификатор, воспользуйтесь запросом [GET warehouses](../../reference/warehouses/getFulfillmentWarehouses.md). | [optional] 
**ReportDate** | Pointer to **string** | Фильтр по дате (для FBY). В отчет попадут данные за **предшествующий** дате день. | [optional] 

## Methods

### NewGenerateStocksOnWarehousesReportRequest

`func NewGenerateStocksOnWarehousesReportRequest(campaignId int64, ) *GenerateStocksOnWarehousesReportRequest`

NewGenerateStocksOnWarehousesReportRequest instantiates a new GenerateStocksOnWarehousesReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateStocksOnWarehousesReportRequestWithDefaults

`func NewGenerateStocksOnWarehousesReportRequestWithDefaults() *GenerateStocksOnWarehousesReportRequest`

NewGenerateStocksOnWarehousesReportRequestWithDefaults instantiates a new GenerateStocksOnWarehousesReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *GenerateStocksOnWarehousesReportRequest) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *GenerateStocksOnWarehousesReportRequest) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *GenerateStocksOnWarehousesReportRequest) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetWarehouseIds

`func (o *GenerateStocksOnWarehousesReportRequest) GetWarehouseIds() []int64`

GetWarehouseIds returns the WarehouseIds field if non-nil, zero value otherwise.

### GetWarehouseIdsOk

`func (o *GenerateStocksOnWarehousesReportRequest) GetWarehouseIdsOk() (*[]int64, bool)`

GetWarehouseIdsOk returns a tuple with the WarehouseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseIds

`func (o *GenerateStocksOnWarehousesReportRequest) SetWarehouseIds(v []int64)`

SetWarehouseIds sets WarehouseIds field to given value.

### HasWarehouseIds

`func (o *GenerateStocksOnWarehousesReportRequest) HasWarehouseIds() bool`

HasWarehouseIds returns a boolean if a field has been set.

### GetReportDate

`func (o *GenerateStocksOnWarehousesReportRequest) GetReportDate() string`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *GenerateStocksOnWarehousesReportRequest) GetReportDateOk() (*string, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *GenerateStocksOnWarehousesReportRequest) SetReportDate(v string)`

SetReportDate sets ReportDate field to given value.

### HasReportDate

`func (o *GenerateStocksOnWarehousesReportRequest) HasReportDate() bool`

HasReportDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


