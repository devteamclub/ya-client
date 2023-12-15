/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOutletLicensesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOutletLicensesResponse{}

// GetOutletLicensesResponse struct for GetOutletLicensesResponse
type GetOutletLicensesResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *OutletLicensesResponseDTO `json:"result,omitempty"`
}

// NewGetOutletLicensesResponse instantiates a new GetOutletLicensesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOutletLicensesResponse() *GetOutletLicensesResponse {
	this := GetOutletLicensesResponse{}
	return &this
}

// NewGetOutletLicensesResponseWithDefaults instantiates a new GetOutletLicensesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOutletLicensesResponseWithDefaults() *GetOutletLicensesResponse {
	this := GetOutletLicensesResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOutletLicensesResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOutletLicensesResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOutletLicensesResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *GetOutletLicensesResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetOutletLicensesResponse) GetResult() OutletLicensesResponseDTO {
	if o == nil || IsNil(o.Result) {
		var ret OutletLicensesResponseDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOutletLicensesResponse) GetResultOk() (*OutletLicensesResponseDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetOutletLicensesResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OutletLicensesResponseDTO and assigns it to the Result field.
func (o *GetOutletLicensesResponse) SetResult(v OutletLicensesResponseDTO) {
	o.Result = &v
}

func (o GetOutletLicensesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOutletLicensesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetOutletLicensesResponse struct {
	value *GetOutletLicensesResponse
	isSet bool
}

func (v NullableGetOutletLicensesResponse) Get() *GetOutletLicensesResponse {
	return v.value
}

func (v *NullableGetOutletLicensesResponse) Set(val *GetOutletLicensesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOutletLicensesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOutletLicensesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOutletLicensesResponse(val *GetOutletLicensesResponse) *NullableGetOutletLicensesResponse {
	return &NullableGetOutletLicensesResponse{value: val, isSet: true}
}

func (v NullableGetOutletLicensesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOutletLicensesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


