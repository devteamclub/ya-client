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

// checks if the GetFeedIndexLogsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFeedIndexLogsResponse{}

// GetFeedIndexLogsResponse struct for GetFeedIndexLogsResponse
type GetFeedIndexLogsResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *FeedIndexLogsResultDTO `json:"result,omitempty"`
}

// NewGetFeedIndexLogsResponse instantiates a new GetFeedIndexLogsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFeedIndexLogsResponse() *GetFeedIndexLogsResponse {
	this := GetFeedIndexLogsResponse{}
	return &this
}

// NewGetFeedIndexLogsResponseWithDefaults instantiates a new GetFeedIndexLogsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFeedIndexLogsResponseWithDefaults() *GetFeedIndexLogsResponse {
	this := GetFeedIndexLogsResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetFeedIndexLogsResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFeedIndexLogsResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetFeedIndexLogsResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *GetFeedIndexLogsResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetFeedIndexLogsResponse) GetResult() FeedIndexLogsResultDTO {
	if o == nil || IsNil(o.Result) {
		var ret FeedIndexLogsResultDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFeedIndexLogsResponse) GetResultOk() (*FeedIndexLogsResultDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetFeedIndexLogsResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given FeedIndexLogsResultDTO and assigns it to the Result field.
func (o *GetFeedIndexLogsResponse) SetResult(v FeedIndexLogsResultDTO) {
	o.Result = &v
}

func (o GetFeedIndexLogsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFeedIndexLogsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetFeedIndexLogsResponse struct {
	value *GetFeedIndexLogsResponse
	isSet bool
}

func (v NullableGetFeedIndexLogsResponse) Get() *GetFeedIndexLogsResponse {
	return v.value
}

func (v *NullableGetFeedIndexLogsResponse) Set(val *GetFeedIndexLogsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFeedIndexLogsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFeedIndexLogsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFeedIndexLogsResponse(val *GetFeedIndexLogsResponse) *NullableGetFeedIndexLogsResponse {
	return &NullableGetFeedIndexLogsResponse{value: val, isSet: true}
}

func (v NullableGetFeedIndexLogsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFeedIndexLogsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


