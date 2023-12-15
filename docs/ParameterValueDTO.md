# ParameterValueDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParameterId** | **int64** | Идентификатор характеристики. | 
**ValueId** | Pointer to **int64** | Идентификатор значения.  Обязательно передавайте идентификатор, если передаете значение из перечня допустимых значений, полученного от Маркета.  Только для характеристик типа &#x60;ENUM&#x60;.  | [optional] 
**Value** | **string** | Значение. | 

## Methods

### NewParameterValueDTO

`func NewParameterValueDTO(parameterId int64, value string, ) *ParameterValueDTO`

NewParameterValueDTO instantiates a new ParameterValueDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterValueDTOWithDefaults

`func NewParameterValueDTOWithDefaults() *ParameterValueDTO`

NewParameterValueDTOWithDefaults instantiates a new ParameterValueDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameterId

`func (o *ParameterValueDTO) GetParameterId() int64`

GetParameterId returns the ParameterId field if non-nil, zero value otherwise.

### GetParameterIdOk

`func (o *ParameterValueDTO) GetParameterIdOk() (*int64, bool)`

GetParameterIdOk returns a tuple with the ParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterId

`func (o *ParameterValueDTO) SetParameterId(v int64)`

SetParameterId sets ParameterId field to given value.


### GetValueId

`func (o *ParameterValueDTO) GetValueId() int64`

GetValueId returns the ValueId field if non-nil, zero value otherwise.

### GetValueIdOk

`func (o *ParameterValueDTO) GetValueIdOk() (*int64, bool)`

GetValueIdOk returns a tuple with the ValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueId

`func (o *ParameterValueDTO) SetValueId(v int64)`

SetValueId sets ValueId field to given value.

### HasValueId

`func (o *ParameterValueDTO) HasValueId() bool`

HasValueId returns a boolean if a field has been set.

### GetValue

`func (o *ParameterValueDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ParameterValueDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ParameterValueDTO) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


