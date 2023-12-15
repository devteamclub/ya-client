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

// checks if the ScrollingPagerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScrollingPagerDTO{}

// ScrollingPagerDTO Информация о страницах результатов.
type ScrollingPagerDTO struct {
	// Идентификатор следующей страницы результатов.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// Идентификатор предыдущей страницы результатов.
	PrevPageToken *string `json:"prevPageToken,omitempty"`
}

// NewScrollingPagerDTO instantiates a new ScrollingPagerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScrollingPagerDTO() *ScrollingPagerDTO {
	this := ScrollingPagerDTO{}
	return &this
}

// NewScrollingPagerDTOWithDefaults instantiates a new ScrollingPagerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScrollingPagerDTOWithDefaults() *ScrollingPagerDTO {
	this := ScrollingPagerDTO{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ScrollingPagerDTO) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScrollingPagerDTO) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ScrollingPagerDTO) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ScrollingPagerDTO) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetPrevPageToken returns the PrevPageToken field value if set, zero value otherwise.
func (o *ScrollingPagerDTO) GetPrevPageToken() string {
	if o == nil || IsNil(o.PrevPageToken) {
		var ret string
		return ret
	}
	return *o.PrevPageToken
}

// GetPrevPageTokenOk returns a tuple with the PrevPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScrollingPagerDTO) GetPrevPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PrevPageToken) {
		return nil, false
	}
	return o.PrevPageToken, true
}

// HasPrevPageToken returns a boolean if a field has been set.
func (o *ScrollingPagerDTO) HasPrevPageToken() bool {
	if o != nil && !IsNil(o.PrevPageToken) {
		return true
	}

	return false
}

// SetPrevPageToken gets a reference to the given string and assigns it to the PrevPageToken field.
func (o *ScrollingPagerDTO) SetPrevPageToken(v string) {
	o.PrevPageToken = &v
}

func (o ScrollingPagerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScrollingPagerDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	if !IsNil(o.PrevPageToken) {
		toSerialize["prevPageToken"] = o.PrevPageToken
	}
	return toSerialize, nil
}

type NullableScrollingPagerDTO struct {
	value *ScrollingPagerDTO
	isSet bool
}

func (v NullableScrollingPagerDTO) Get() *ScrollingPagerDTO {
	return v.value
}

func (v *NullableScrollingPagerDTO) Set(val *ScrollingPagerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableScrollingPagerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableScrollingPagerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScrollingPagerDTO(val *ScrollingPagerDTO) *NullableScrollingPagerDTO {
	return &NullableScrollingPagerDTO{value: val, isSet: true}
}

func (v NullableScrollingPagerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScrollingPagerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

