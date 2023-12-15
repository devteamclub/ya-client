# SearchShipmentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateFrom** | Pointer to **string** | Начальная дата для фильтрации по дате отгрузки (включительно).  Формат даты: &#x60;ДД-ММ-ГГГГ&#x60;.  | [optional] 
**DateTo** | Pointer to **string** | Конечная дата для фильтрации по дате отгрузки (включительно).  Формат даты: &#x60;ДД-ММ-ГГГГ&#x60;.  | [optional] 
**Statuses** | Pointer to [**[]ShipmentStatusType**](ShipmentStatusType.md) | Список статусов отгрузок. | [optional] 
**OrderIds** | Pointer to **[]int64** | Список идентификаторов заказов из отгрузок. | [optional] 

## Methods

### NewSearchShipmentsRequest

`func NewSearchShipmentsRequest() *SearchShipmentsRequest`

NewSearchShipmentsRequest instantiates a new SearchShipmentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchShipmentsRequestWithDefaults

`func NewSearchShipmentsRequestWithDefaults() *SearchShipmentsRequest`

NewSearchShipmentsRequestWithDefaults instantiates a new SearchShipmentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateFrom

`func (o *SearchShipmentsRequest) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *SearchShipmentsRequest) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *SearchShipmentsRequest) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *SearchShipmentsRequest) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### GetDateTo

`func (o *SearchShipmentsRequest) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *SearchShipmentsRequest) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *SearchShipmentsRequest) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *SearchShipmentsRequest) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### GetStatuses

`func (o *SearchShipmentsRequest) GetStatuses() []ShipmentStatusType`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *SearchShipmentsRequest) GetStatusesOk() (*[]ShipmentStatusType, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *SearchShipmentsRequest) SetStatuses(v []ShipmentStatusType)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *SearchShipmentsRequest) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetOrderIds

`func (o *SearchShipmentsRequest) GetOrderIds() []int64`

GetOrderIds returns the OrderIds field if non-nil, zero value otherwise.

### GetOrderIdsOk

`func (o *SearchShipmentsRequest) GetOrderIdsOk() (*[]int64, bool)`

GetOrderIdsOk returns a tuple with the OrderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIds

`func (o *SearchShipmentsRequest) SetOrderIds(v []int64)`

SetOrderIds sets OrderIds field to given value.

### HasOrderIds

`func (o *SearchShipmentsRequest) HasOrderIds() bool`

HasOrderIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


