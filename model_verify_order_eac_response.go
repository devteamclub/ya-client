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

// checks if the VerifyOrderEacResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyOrderEacResponse{}

// VerifyOrderEacResponse struct for VerifyOrderEacResponse
type VerifyOrderEacResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *EacVerificationResultDTO `json:"result,omitempty"`
}

// NewVerifyOrderEacResponse instantiates a new VerifyOrderEacResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyOrderEacResponse() *VerifyOrderEacResponse {
	this := VerifyOrderEacResponse{}
	return &this
}

// NewVerifyOrderEacResponseWithDefaults instantiates a new VerifyOrderEacResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyOrderEacResponseWithDefaults() *VerifyOrderEacResponse {
	this := VerifyOrderEacResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VerifyOrderEacResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyOrderEacResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VerifyOrderEacResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *VerifyOrderEacResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *VerifyOrderEacResponse) GetResult() EacVerificationResultDTO {
	if o == nil || IsNil(o.Result) {
		var ret EacVerificationResultDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyOrderEacResponse) GetResultOk() (*EacVerificationResultDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *VerifyOrderEacResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given EacVerificationResultDTO and assigns it to the Result field.
func (o *VerifyOrderEacResponse) SetResult(v EacVerificationResultDTO) {
	o.Result = &v
}

func (o VerifyOrderEacResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyOrderEacResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableVerifyOrderEacResponse struct {
	value *VerifyOrderEacResponse
	isSet bool
}

func (v NullableVerifyOrderEacResponse) Get() *VerifyOrderEacResponse {
	return v.value
}

func (v *NullableVerifyOrderEacResponse) Set(val *VerifyOrderEacResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyOrderEacResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyOrderEacResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyOrderEacResponse(val *VerifyOrderEacResponse) *NullableVerifyOrderEacResponse {
	return &NullableVerifyOrderEacResponse{value: val, isSet: true}
}

func (v NullableVerifyOrderEacResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyOrderEacResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


