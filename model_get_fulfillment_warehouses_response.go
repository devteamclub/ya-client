/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ya-client

import (
	"encoding/json"
)

// checks if the GetFulfillmentWarehousesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFulfillmentWarehousesResponse{}

// GetFulfillmentWarehousesResponse struct for GetFulfillmentWarehousesResponse
type GetFulfillmentWarehousesResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *FulfillmentWarehousesDTO `json:"result,omitempty"`
}

// NewGetFulfillmentWarehousesResponse instantiates a new GetFulfillmentWarehousesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFulfillmentWarehousesResponse() *GetFulfillmentWarehousesResponse {
	this := GetFulfillmentWarehousesResponse{}
	return &this
}

// NewGetFulfillmentWarehousesResponseWithDefaults instantiates a new GetFulfillmentWarehousesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFulfillmentWarehousesResponseWithDefaults() *GetFulfillmentWarehousesResponse {
	this := GetFulfillmentWarehousesResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetFulfillmentWarehousesResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFulfillmentWarehousesResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetFulfillmentWarehousesResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *GetFulfillmentWarehousesResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetFulfillmentWarehousesResponse) GetResult() FulfillmentWarehousesDTO {
	if o == nil || IsNil(o.Result) {
		var ret FulfillmentWarehousesDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFulfillmentWarehousesResponse) GetResultOk() (*FulfillmentWarehousesDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetFulfillmentWarehousesResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given FulfillmentWarehousesDTO and assigns it to the Result field.
func (o *GetFulfillmentWarehousesResponse) SetResult(v FulfillmentWarehousesDTO) {
	o.Result = &v
}

func (o GetFulfillmentWarehousesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFulfillmentWarehousesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetFulfillmentWarehousesResponse struct {
	value *GetFulfillmentWarehousesResponse
	isSet bool
}

func (v NullableGetFulfillmentWarehousesResponse) Get() *GetFulfillmentWarehousesResponse {
	return v.value
}

func (v *NullableGetFulfillmentWarehousesResponse) Set(val *GetFulfillmentWarehousesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFulfillmentWarehousesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFulfillmentWarehousesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFulfillmentWarehousesResponse(val *GetFulfillmentWarehousesResponse) *NullableGetFulfillmentWarehousesResponse {
	return &NullableGetFulfillmentWarehousesResponse{value: val, isSet: true}
}

func (v NullableGetFulfillmentWarehousesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFulfillmentWarehousesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


