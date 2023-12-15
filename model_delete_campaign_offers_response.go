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

// checks if the DeleteCampaignOffersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteCampaignOffersResponse{}

// DeleteCampaignOffersResponse Результат удаления товаров.
type DeleteCampaignOffersResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *DeleteCampaignOffersDTO `json:"result,omitempty"`
}

// NewDeleteCampaignOffersResponse instantiates a new DeleteCampaignOffersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCampaignOffersResponse() *DeleteCampaignOffersResponse {
	this := DeleteCampaignOffersResponse{}
	return &this
}

// NewDeleteCampaignOffersResponseWithDefaults instantiates a new DeleteCampaignOffersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCampaignOffersResponseWithDefaults() *DeleteCampaignOffersResponse {
	this := DeleteCampaignOffersResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeleteCampaignOffersResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCampaignOffersResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeleteCampaignOffersResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *DeleteCampaignOffersResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DeleteCampaignOffersResponse) GetResult() DeleteCampaignOffersDTO {
	if o == nil || IsNil(o.Result) {
		var ret DeleteCampaignOffersDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCampaignOffersResponse) GetResultOk() (*DeleteCampaignOffersDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DeleteCampaignOffersResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given DeleteCampaignOffersDTO and assigns it to the Result field.
func (o *DeleteCampaignOffersResponse) SetResult(v DeleteCampaignOffersDTO) {
	o.Result = &v
}

func (o DeleteCampaignOffersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteCampaignOffersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableDeleteCampaignOffersResponse struct {
	value *DeleteCampaignOffersResponse
	isSet bool
}

func (v NullableDeleteCampaignOffersResponse) Get() *DeleteCampaignOffersResponse {
	return v.value
}

func (v *NullableDeleteCampaignOffersResponse) Set(val *DeleteCampaignOffersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCampaignOffersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCampaignOffersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCampaignOffersResponse(val *DeleteCampaignOffersResponse) *NullableDeleteCampaignOffersResponse {
	return &NullableDeleteCampaignOffersResponse{value: val, isSet: true}
}

func (v NullableDeleteCampaignOffersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCampaignOffersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

