# ConfirmShipmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalShipmentId** | Pointer to **string** | Идентификатор отгрузки в системе поставщика. | [optional] 
**OrderIds** | **[]int64** | Список идентификаторов заказов в отгрузке. | 

## Methods

### NewConfirmShipmentRequest

`func NewConfirmShipmentRequest(orderIds []int64, ) *ConfirmShipmentRequest`

NewConfirmShipmentRequest instantiates a new ConfirmShipmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmShipmentRequestWithDefaults

`func NewConfirmShipmentRequestWithDefaults() *ConfirmShipmentRequest`

NewConfirmShipmentRequestWithDefaults instantiates a new ConfirmShipmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalShipmentId

`func (o *ConfirmShipmentRequest) GetExternalShipmentId() string`

GetExternalShipmentId returns the ExternalShipmentId field if non-nil, zero value otherwise.

### GetExternalShipmentIdOk

`func (o *ConfirmShipmentRequest) GetExternalShipmentIdOk() (*string, bool)`

GetExternalShipmentIdOk returns a tuple with the ExternalShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShipmentId

`func (o *ConfirmShipmentRequest) SetExternalShipmentId(v string)`

SetExternalShipmentId sets ExternalShipmentId field to given value.

### HasExternalShipmentId

`func (o *ConfirmShipmentRequest) HasExternalShipmentId() bool`

HasExternalShipmentId returns a boolean if a field has been set.

### GetOrderIds

`func (o *ConfirmShipmentRequest) GetOrderIds() []int64`

GetOrderIds returns the OrderIds field if non-nil, zero value otherwise.

### GetOrderIdsOk

`func (o *ConfirmShipmentRequest) GetOrderIdsOk() (*[]int64, bool)`

GetOrderIdsOk returns a tuple with the OrderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIds

`func (o *ConfirmShipmentRequest) SetOrderIds(v []int64)`

SetOrderIds sets OrderIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


