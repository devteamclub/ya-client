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

// checks if the OfferProcessingNoteDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferProcessingNoteDTO{}

// OfferProcessingNoteDTO Причины, по которым товар не прошел модерацию.
type OfferProcessingNoteDTO struct {
	Type *OfferProcessingNoteType `json:"type,omitempty"`
	// Дополнительная информация о причине отклонения товара. 
	Payload *string `json:"payload,omitempty"`
}

// NewOfferProcessingNoteDTO instantiates a new OfferProcessingNoteDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferProcessingNoteDTO() *OfferProcessingNoteDTO {
	this := OfferProcessingNoteDTO{}
	return &this
}

// NewOfferProcessingNoteDTOWithDefaults instantiates a new OfferProcessingNoteDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferProcessingNoteDTOWithDefaults() *OfferProcessingNoteDTO {
	this := OfferProcessingNoteDTO{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OfferProcessingNoteDTO) GetType() OfferProcessingNoteType {
	if o == nil || IsNil(o.Type) {
		var ret OfferProcessingNoteType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferProcessingNoteDTO) GetTypeOk() (*OfferProcessingNoteType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OfferProcessingNoteDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given OfferProcessingNoteType and assigns it to the Type field.
func (o *OfferProcessingNoteDTO) SetType(v OfferProcessingNoteType) {
	o.Type = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *OfferProcessingNoteDTO) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferProcessingNoteDTO) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *OfferProcessingNoteDTO) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *OfferProcessingNoteDTO) SetPayload(v string) {
	o.Payload = &v
}

func (o OfferProcessingNoteDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferProcessingNoteDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableOfferProcessingNoteDTO struct {
	value *OfferProcessingNoteDTO
	isSet bool
}

func (v NullableOfferProcessingNoteDTO) Get() *OfferProcessingNoteDTO {
	return v.value
}

func (v *NullableOfferProcessingNoteDTO) Set(val *OfferProcessingNoteDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferProcessingNoteDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferProcessingNoteDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferProcessingNoteDTO(val *OfferProcessingNoteDTO) *NullableOfferProcessingNoteDTO {
	return &NullableOfferProcessingNoteDTO{value: val, isSet: true}
}

func (v NullableOfferProcessingNoteDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferProcessingNoteDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


