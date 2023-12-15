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

// checks if the ApiClientDataErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiClientDataErrorResponse{}

// ApiClientDataErrorResponse Ошибка в данных переданных от клиента.
type ApiClientDataErrorResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	// Список ошибок.
	Errors []ApiErrorDTO `json:"errors,omitempty"`
}

// NewApiClientDataErrorResponse instantiates a new ApiClientDataErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiClientDataErrorResponse() *ApiClientDataErrorResponse {
	this := ApiClientDataErrorResponse{}
	return &this
}

// NewApiClientDataErrorResponseWithDefaults instantiates a new ApiClientDataErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiClientDataErrorResponseWithDefaults() *ApiClientDataErrorResponse {
	this := ApiClientDataErrorResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiClientDataErrorResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClientDataErrorResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiClientDataErrorResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *ApiClientDataErrorResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ApiClientDataErrorResponse) GetErrors() []ApiErrorDTO {
	if o == nil || IsNil(o.Errors) {
		var ret []ApiErrorDTO
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClientDataErrorResponse) GetErrorsOk() ([]ApiErrorDTO, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ApiClientDataErrorResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ApiErrorDTO and assigns it to the Errors field.
func (o *ApiClientDataErrorResponse) SetErrors(v []ApiErrorDTO) {
	o.Errors = v
}

func (o ApiClientDataErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiClientDataErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableApiClientDataErrorResponse struct {
	value *ApiClientDataErrorResponse
	isSet bool
}

func (v NullableApiClientDataErrorResponse) Get() *ApiClientDataErrorResponse {
	return v.value
}

func (v *NullableApiClientDataErrorResponse) Set(val *ApiClientDataErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiClientDataErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiClientDataErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiClientDataErrorResponse(val *ApiClientDataErrorResponse) *NullableApiClientDataErrorResponse {
	return &NullableApiClientDataErrorResponse{value: val, isSet: true}
}

func (v NullableApiClientDataErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiClientDataErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


