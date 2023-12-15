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

// checks if the GetShipmentOrdersInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShipmentOrdersInfoResponse{}

// GetShipmentOrdersInfoResponse struct for GetShipmentOrdersInfoResponse
type GetShipmentOrdersInfoResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *OrdersShipmentInfoDTO `json:"result,omitempty"`
}

// NewGetShipmentOrdersInfoResponse instantiates a new GetShipmentOrdersInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShipmentOrdersInfoResponse() *GetShipmentOrdersInfoResponse {
	this := GetShipmentOrdersInfoResponse{}
	return &this
}

// NewGetShipmentOrdersInfoResponseWithDefaults instantiates a new GetShipmentOrdersInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShipmentOrdersInfoResponseWithDefaults() *GetShipmentOrdersInfoResponse {
	this := GetShipmentOrdersInfoResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetShipmentOrdersInfoResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentOrdersInfoResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetShipmentOrdersInfoResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *GetShipmentOrdersInfoResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetShipmentOrdersInfoResponse) GetResult() OrdersShipmentInfoDTO {
	if o == nil || IsNil(o.Result) {
		var ret OrdersShipmentInfoDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShipmentOrdersInfoResponse) GetResultOk() (*OrdersShipmentInfoDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetShipmentOrdersInfoResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OrdersShipmentInfoDTO and assigns it to the Result field.
func (o *GetShipmentOrdersInfoResponse) SetResult(v OrdersShipmentInfoDTO) {
	o.Result = &v
}

func (o GetShipmentOrdersInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetShipmentOrdersInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetShipmentOrdersInfoResponse struct {
	value *GetShipmentOrdersInfoResponse
	isSet bool
}

func (v NullableGetShipmentOrdersInfoResponse) Get() *GetShipmentOrdersInfoResponse {
	return v.value
}

func (v *NullableGetShipmentOrdersInfoResponse) Set(val *GetShipmentOrdersInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShipmentOrdersInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShipmentOrdersInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShipmentOrdersInfoResponse(val *GetShipmentOrdersInfoResponse) *NullableGetShipmentOrdersInfoResponse {
	return &NullableGetShipmentOrdersInfoResponse{value: val, isSet: true}
}

func (v NullableGetShipmentOrdersInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetShipmentOrdersInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


