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

// checks if the FeedDownloadErrorDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedDownloadErrorDTO{}

// FeedDownloadErrorDTO Информация об ошибке при загрузке прайс-листа. Выводится, если параметр `download status=ERROR`. 
type FeedDownloadErrorDTO struct {
	// HTTP-код ошибки индексации прайс-листа. Выводится, если `type=DOWNLOAD_HTTP_ERROR`. 
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty"`
	Type *FeedDownloadErrorType `json:"type,omitempty"`
	// Описание ошибки. Выводится, если `type=DOWNLOAD_ERROR`. 
	Description *string `json:"description,omitempty"`
}

// NewFeedDownloadErrorDTO instantiates a new FeedDownloadErrorDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedDownloadErrorDTO() *FeedDownloadErrorDTO {
	this := FeedDownloadErrorDTO{}
	return &this
}

// NewFeedDownloadErrorDTOWithDefaults instantiates a new FeedDownloadErrorDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedDownloadErrorDTOWithDefaults() *FeedDownloadErrorDTO {
	this := FeedDownloadErrorDTO{}
	return &this
}

// GetHttpStatusCode returns the HttpStatusCode field value if set, zero value otherwise.
func (o *FeedDownloadErrorDTO) GetHttpStatusCode() int32 {
	if o == nil || IsNil(o.HttpStatusCode) {
		var ret int32
		return ret
	}
	return *o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedDownloadErrorDTO) GetHttpStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.HttpStatusCode) {
		return nil, false
	}
	return o.HttpStatusCode, true
}

// HasHttpStatusCode returns a boolean if a field has been set.
func (o *FeedDownloadErrorDTO) HasHttpStatusCode() bool {
	if o != nil && !IsNil(o.HttpStatusCode) {
		return true
	}

	return false
}

// SetHttpStatusCode gets a reference to the given int32 and assigns it to the HttpStatusCode field.
func (o *FeedDownloadErrorDTO) SetHttpStatusCode(v int32) {
	o.HttpStatusCode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FeedDownloadErrorDTO) GetType() FeedDownloadErrorType {
	if o == nil || IsNil(o.Type) {
		var ret FeedDownloadErrorType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedDownloadErrorDTO) GetTypeOk() (*FeedDownloadErrorType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FeedDownloadErrorDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FeedDownloadErrorType and assigns it to the Type field.
func (o *FeedDownloadErrorDTO) SetType(v FeedDownloadErrorType) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FeedDownloadErrorDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedDownloadErrorDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FeedDownloadErrorDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FeedDownloadErrorDTO) SetDescription(v string) {
	o.Description = &v
}

func (o FeedDownloadErrorDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedDownloadErrorDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HttpStatusCode) {
		toSerialize["httpStatusCode"] = o.HttpStatusCode
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableFeedDownloadErrorDTO struct {
	value *FeedDownloadErrorDTO
	isSet bool
}

func (v NullableFeedDownloadErrorDTO) Get() *FeedDownloadErrorDTO {
	return v.value
}

func (v *NullableFeedDownloadErrorDTO) Set(val *FeedDownloadErrorDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedDownloadErrorDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedDownloadErrorDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedDownloadErrorDTO(val *FeedDownloadErrorDTO) *NullableFeedDownloadErrorDTO {
	return &NullableFeedDownloadErrorDTO{value: val, isSet: true}
}

func (v NullableFeedDownloadErrorDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedDownloadErrorDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


