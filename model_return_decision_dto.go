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

// checks if the ReturnDecisionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnDecisionDTO{}

// ReturnDecisionDTO Решения по возвратам.
type ReturnDecisionDTO struct {
	// Идентификатор товара в возврате.
	ReturnItemId *int64 `json:"returnItemId,omitempty"`
	// Количество товаров.
	Count *int32 `json:"count,omitempty"`
	// Комментарий.
	Comment *string `json:"comment,omitempty"`
	ReasonType *ReturnDecisionReasonType `json:"reasonType,omitempty"`
	SubreasonType *ReturnDecisionSubreasonType `json:"subreasonType,omitempty"`
	DecisionType *ReturnDecisionType `json:"decisionType,omitempty"`
	// Сумма возврата.
	RefundAmount *int64 `json:"refundAmount,omitempty"`
	// Компенсация за обратную доставку.
	PartnerCompensation *int64 `json:"partnerCompensation,omitempty"`
	// Список хеш-кодов фотографий товара от покупателя.
	Images []string `json:"images,omitempty"`
}

// NewReturnDecisionDTO instantiates a new ReturnDecisionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnDecisionDTO() *ReturnDecisionDTO {
	this := ReturnDecisionDTO{}
	return &this
}

// NewReturnDecisionDTOWithDefaults instantiates a new ReturnDecisionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnDecisionDTOWithDefaults() *ReturnDecisionDTO {
	this := ReturnDecisionDTO{}
	return &this
}

// GetReturnItemId returns the ReturnItemId field value if set, zero value otherwise.
func (o *ReturnDecisionDTO) GetReturnItemId() int64 {
	if o == nil || IsNil(o.ReturnItemId) {
		var ret int64
		return ret
	}
	return *o.ReturnItemId
}

// GetReturnItemIdOk returns a tuple with the ReturnItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDecisionDTO) GetReturnItemIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ReturnItemId) {
		return nil, false
	}
	return o.ReturnItemId, true
}

// HasReturnItemId returns a boolean if a field has been set.
func (o *ReturnDecisionDTO) HasReturnItemId() bool {
	if o != nil && !IsNil(o.ReturnItemId) {
		return true
	}

	return false
}

// SetReturnItemId gets a reference to the given int64 and assigns it to the ReturnItemId field.
func (o *ReturnDecisionDTO) SetReturnItemId(v int64) {
	o.ReturnItemId = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ReturnDecisionDTO) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDecisionDTO) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ReturnDecisionDTO) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ReturnDecisionDTO) SetCount(v int32) {
	o.Count = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ReturnDecisionDTO) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDecisionDTO) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ReturnDecisionDTO) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ReturnDecisionDTO) SetComment(v string) {
	o.Comment = &v
}

// GetReasonType returns the ReasonType field value if set, zero value otherwise.
func (o *ReturnDecisionDTO) GetReasonType() ReturnDecisionReasonType {
	if o == nil || IsNil(o.ReasonType) {
		var ret ReturnDecisionReasonType
		return ret
	}
	return *o.ReasonType
}

// GetReasonTypeOk returns a tuple with the ReasonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDecisionDTO) GetReasonTypeOk() (*ReturnDecisionReasonType, bool) {
	if o == nil || IsNil(o.ReasonType) {
		return nil, false
	}
	return o.ReasonType, true
}

// HasReasonType returns a boolean if a field has been set.
func (o *ReturnDecisionDTO) HasReasonType() bool {
	if o != nil && !IsNil(o.ReasonType) {
		return true
	}

	return false
}

// SetReasonType gets a reference to the given ReturnDecisionReasonType and assigns it to the ReasonType field.
func (o *ReturnDecisionDTO) SetReasonType(v ReturnDecisionReasonType) {
	o.ReasonType = &v
}

// GetSubreasonType returns the SubreasonType field value if set, zero value otherwise.
func (o *ReturnDecisionDTO) GetSubreasonType() ReturnDecisionSubreasonType {
	if o == nil || IsNil(o.SubreasonType) {
		var ret ReturnDecisionSubreasonType
		return ret
	}
	return *o.SubreasonType
}

// GetSubreasonTypeOk returns a tuple with the SubreasonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDecisionDTO) GetSubreasonTypeOk() (*ReturnDecisionSubreasonType, bool) {
	if o == nil || IsNil(o.SubreasonType) {
		return nil, false
	}
	return o.SubreasonType, true
}

// HasSubreasonType returns a boolean if a field has been set.
func (o *ReturnDecisionDTO) HasSubreasonType() bool {
	if o != nil && !IsNil(o.SubreasonType) {
		return true
	}

	return false
}

// SetSubreasonType gets a reference to the given ReturnDecisionSubreasonType and assigns it to the SubreasonType field.
func (o *ReturnDecisionDTO) SetSubreasonType(v ReturnDecisionSubreasonType) {
	o.SubreasonType = &v
}

// GetDecisionType returns the DecisionType field value if set, zero value otherwise.
func (o *ReturnDecisionDTO) GetDecisionType() ReturnDecisionType {
	if o == nil || IsNil(o.DecisionType) {
		var ret ReturnDecisionType
		return ret
	}
	return *o.DecisionType
}

// GetDecisionTypeOk returns a tuple with the DecisionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDecisionDTO) GetDecisionTypeOk() (*ReturnDecisionType, bool) {
	if o == nil || IsNil(o.DecisionType) {
		return nil, false
	}
	return o.DecisionType, true
}

// HasDecisionType returns a boolean if a field has been set.
func (o *ReturnDecisionDTO) HasDecisionType() bool {
	if o != nil && !IsNil(o.DecisionType) {
		return true
	}

	return false
}

// SetDecisionType gets a reference to the given ReturnDecisionType and assigns it to the DecisionType field.
func (o *ReturnDecisionDTO) SetDecisionType(v ReturnDecisionType) {
	o.DecisionType = &v
}

// GetRefundAmount returns the RefundAmount field value if set, zero value otherwise.
func (o *ReturnDecisionDTO) GetRefundAmount() int64 {
	if o == nil || IsNil(o.RefundAmount) {
		var ret int64
		return ret
	}
	return *o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDecisionDTO) GetRefundAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.RefundAmount) {
		return nil, false
	}
	return o.RefundAmount, true
}

// HasRefundAmount returns a boolean if a field has been set.
func (o *ReturnDecisionDTO) HasRefundAmount() bool {
	if o != nil && !IsNil(o.RefundAmount) {
		return true
	}

	return false
}

// SetRefundAmount gets a reference to the given int64 and assigns it to the RefundAmount field.
func (o *ReturnDecisionDTO) SetRefundAmount(v int64) {
	o.RefundAmount = &v
}

// GetPartnerCompensation returns the PartnerCompensation field value if set, zero value otherwise.
func (o *ReturnDecisionDTO) GetPartnerCompensation() int64 {
	if o == nil || IsNil(o.PartnerCompensation) {
		var ret int64
		return ret
	}
	return *o.PartnerCompensation
}

// GetPartnerCompensationOk returns a tuple with the PartnerCompensation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDecisionDTO) GetPartnerCompensationOk() (*int64, bool) {
	if o == nil || IsNil(o.PartnerCompensation) {
		return nil, false
	}
	return o.PartnerCompensation, true
}

// HasPartnerCompensation returns a boolean if a field has been set.
func (o *ReturnDecisionDTO) HasPartnerCompensation() bool {
	if o != nil && !IsNil(o.PartnerCompensation) {
		return true
	}

	return false
}

// SetPartnerCompensation gets a reference to the given int64 and assigns it to the PartnerCompensation field.
func (o *ReturnDecisionDTO) SetPartnerCompensation(v int64) {
	o.PartnerCompensation = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *ReturnDecisionDTO) GetImages() []string {
	if o == nil || IsNil(o.Images) {
		var ret []string
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDecisionDTO) GetImagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *ReturnDecisionDTO) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []string and assigns it to the Images field.
func (o *ReturnDecisionDTO) SetImages(v []string) {
	o.Images = v
}

func (o ReturnDecisionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnDecisionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReturnItemId) {
		toSerialize["returnItemId"] = o.ReturnItemId
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ReasonType) {
		toSerialize["reasonType"] = o.ReasonType
	}
	if !IsNil(o.SubreasonType) {
		toSerialize["subreasonType"] = o.SubreasonType
	}
	if !IsNil(o.DecisionType) {
		toSerialize["decisionType"] = o.DecisionType
	}
	if !IsNil(o.RefundAmount) {
		toSerialize["refundAmount"] = o.RefundAmount
	}
	if !IsNil(o.PartnerCompensation) {
		toSerialize["partnerCompensation"] = o.PartnerCompensation
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	return toSerialize, nil
}

type NullableReturnDecisionDTO struct {
	value *ReturnDecisionDTO
	isSet bool
}

func (v NullableReturnDecisionDTO) Get() *ReturnDecisionDTO {
	return v.value
}

func (v *NullableReturnDecisionDTO) Set(val *ReturnDecisionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnDecisionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnDecisionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnDecisionDTO(val *ReturnDecisionDTO) *NullableReturnDecisionDTO {
	return &NullableReturnDecisionDTO{value: val, isSet: true}
}

func (v NullableReturnDecisionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnDecisionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

