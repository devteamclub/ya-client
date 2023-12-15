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

// checks if the CreateOutletResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOutletResponse{}

// CreateOutletResponse Ответ на запрос о создании точки продаж.
type CreateOutletResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *OutletResponseDTO `json:"result,omitempty"`
}

// NewCreateOutletResponse instantiates a new CreateOutletResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOutletResponse() *CreateOutletResponse {
	this := CreateOutletResponse{}
	return &this
}

// NewCreateOutletResponseWithDefaults instantiates a new CreateOutletResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOutletResponseWithDefaults() *CreateOutletResponse {
	this := CreateOutletResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateOutletResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutletResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateOutletResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *CreateOutletResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateOutletResponse) GetResult() OutletResponseDTO {
	if o == nil || IsNil(o.Result) {
		var ret OutletResponseDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutletResponse) GetResultOk() (*OutletResponseDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateOutletResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OutletResponseDTO and assigns it to the Result field.
func (o *CreateOutletResponse) SetResult(v OutletResponseDTO) {
	o.Result = &v
}

func (o CreateOutletResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOutletResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableCreateOutletResponse struct {
	value *CreateOutletResponse
	isSet bool
}

func (v NullableCreateOutletResponse) Get() *CreateOutletResponse {
	return v.value
}

func (v *NullableCreateOutletResponse) Set(val *CreateOutletResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOutletResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOutletResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOutletResponse(val *CreateOutletResponse) *NullableCreateOutletResponse {
	return &NullableCreateOutletResponse{value: val, isSet: true}
}

func (v NullableCreateOutletResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOutletResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


