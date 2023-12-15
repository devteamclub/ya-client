/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ya-client

import (
	"encoding/json"
	"time"
)

// checks if the FullOutletLicenseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FullOutletLicenseDTO{}

// FullOutletLicenseDTO Информация о лицензии.
type FullOutletLicenseDTO struct {
	// Идентификатор лицензии.  Параметр указывается, только если нужно изменить информацию о существующей лицензии. Ее идентификатор можно узнать с помощью запроса `GET /campaigns/{campaignId}/outlets/licenses`. При передаче информации о новой лицензии указывать идентификатор не нужно.  Идентификатор лицензии присваивается Маркетом. Не путайте его с номером, указанным на лицензии: он передается в параметре `number`. 
	Id *int64 `json:"id,omitempty"`
	// Идентификатор точки продаж, для которой действительна лицензия.
	OutletId *int64 `json:"outletId,omitempty"`
	LicenseType *LicenseType `json:"licenseType,omitempty"`
	// Номер лицензии.
	Number *string `json:"number,omitempty"`
	// Дата выдачи лицензии.  Формат даты: ISO 8601 со смещением относительно UTC. Нужно передать дату, указанную на лицензии, время `00:00:00` и часовой пояс, соответствующий региону точки продаж. Например, если лицензия для точки продаж в Москве выдана 13 ноября 2017 года, то параметр должен иметь значение `2017-11-13T00:00:00+03:00`.  Обязательный параметр.  Не может быть позже даты окончания срока действия, указанной в параметре `dateOfExpiry`. 
	DateOfIssue *time.Time `json:"dateOfIssue,omitempty"`
	// Дата окончания действия лицензии.  Формат даты: ISO 8601 со смещением относительно UTC. Нужно передать дату, указанную на лицензии, время `00:00:00` и часовой пояс, соответствующий региону точки продаж. Например, если действие лицензии для точки продаж в Москве заканчивается 20 ноября 2022 года, то параметр должен иметь значение `2022-11-20T00:00:00+03:00`.  Обязательный параметр.  Не может быть раньше даты выдачи, указанной в параметре `dateOfIssue`. 
	DateOfExpiry *time.Time `json:"dateOfExpiry,omitempty"`
	CheckStatus *LicenseCheckStatusType `json:"checkStatus,omitempty"`
	// Причина, по которой лицензия не прошла проверку. Параметр возвращается, только если параметр `checkStatus` имеет значение `FAIL`. 
	CheckComment *string `json:"checkComment,omitempty"`
}

// NewFullOutletLicenseDTO instantiates a new FullOutletLicenseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullOutletLicenseDTO() *FullOutletLicenseDTO {
	this := FullOutletLicenseDTO{}
	return &this
}

// NewFullOutletLicenseDTOWithDefaults instantiates a new FullOutletLicenseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullOutletLicenseDTOWithDefaults() *FullOutletLicenseDTO {
	this := FullOutletLicenseDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullOutletLicenseDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOutletLicenseDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullOutletLicenseDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *FullOutletLicenseDTO) SetId(v int64) {
	o.Id = &v
}

// GetOutletId returns the OutletId field value if set, zero value otherwise.
func (o *FullOutletLicenseDTO) GetOutletId() int64 {
	if o == nil || IsNil(o.OutletId) {
		var ret int64
		return ret
	}
	return *o.OutletId
}

// GetOutletIdOk returns a tuple with the OutletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOutletLicenseDTO) GetOutletIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OutletId) {
		return nil, false
	}
	return o.OutletId, true
}

// HasOutletId returns a boolean if a field has been set.
func (o *FullOutletLicenseDTO) HasOutletId() bool {
	if o != nil && !IsNil(o.OutletId) {
		return true
	}

	return false
}

// SetOutletId gets a reference to the given int64 and assigns it to the OutletId field.
func (o *FullOutletLicenseDTO) SetOutletId(v int64) {
	o.OutletId = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *FullOutletLicenseDTO) GetLicenseType() LicenseType {
	if o == nil || IsNil(o.LicenseType) {
		var ret LicenseType
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOutletLicenseDTO) GetLicenseTypeOk() (*LicenseType, bool) {
	if o == nil || IsNil(o.LicenseType) {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *FullOutletLicenseDTO) HasLicenseType() bool {
	if o != nil && !IsNil(o.LicenseType) {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given LicenseType and assigns it to the LicenseType field.
func (o *FullOutletLicenseDTO) SetLicenseType(v LicenseType) {
	o.LicenseType = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *FullOutletLicenseDTO) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOutletLicenseDTO) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *FullOutletLicenseDTO) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *FullOutletLicenseDTO) SetNumber(v string) {
	o.Number = &v
}

// GetDateOfIssue returns the DateOfIssue field value if set, zero value otherwise.
func (o *FullOutletLicenseDTO) GetDateOfIssue() time.Time {
	if o == nil || IsNil(o.DateOfIssue) {
		var ret time.Time
		return ret
	}
	return *o.DateOfIssue
}

// GetDateOfIssueOk returns a tuple with the DateOfIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOutletLicenseDTO) GetDateOfIssueOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateOfIssue) {
		return nil, false
	}
	return o.DateOfIssue, true
}

// HasDateOfIssue returns a boolean if a field has been set.
func (o *FullOutletLicenseDTO) HasDateOfIssue() bool {
	if o != nil && !IsNil(o.DateOfIssue) {
		return true
	}

	return false
}

// SetDateOfIssue gets a reference to the given time.Time and assigns it to the DateOfIssue field.
func (o *FullOutletLicenseDTO) SetDateOfIssue(v time.Time) {
	o.DateOfIssue = &v
}

// GetDateOfExpiry returns the DateOfExpiry field value if set, zero value otherwise.
func (o *FullOutletLicenseDTO) GetDateOfExpiry() time.Time {
	if o == nil || IsNil(o.DateOfExpiry) {
		var ret time.Time
		return ret
	}
	return *o.DateOfExpiry
}

// GetDateOfExpiryOk returns a tuple with the DateOfExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOutletLicenseDTO) GetDateOfExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateOfExpiry) {
		return nil, false
	}
	return o.DateOfExpiry, true
}

// HasDateOfExpiry returns a boolean if a field has been set.
func (o *FullOutletLicenseDTO) HasDateOfExpiry() bool {
	if o != nil && !IsNil(o.DateOfExpiry) {
		return true
	}

	return false
}

// SetDateOfExpiry gets a reference to the given time.Time and assigns it to the DateOfExpiry field.
func (o *FullOutletLicenseDTO) SetDateOfExpiry(v time.Time) {
	o.DateOfExpiry = &v
}

// GetCheckStatus returns the CheckStatus field value if set, zero value otherwise.
func (o *FullOutletLicenseDTO) GetCheckStatus() LicenseCheckStatusType {
	if o == nil || IsNil(o.CheckStatus) {
		var ret LicenseCheckStatusType
		return ret
	}
	return *o.CheckStatus
}

// GetCheckStatusOk returns a tuple with the CheckStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOutletLicenseDTO) GetCheckStatusOk() (*LicenseCheckStatusType, bool) {
	if o == nil || IsNil(o.CheckStatus) {
		return nil, false
	}
	return o.CheckStatus, true
}

// HasCheckStatus returns a boolean if a field has been set.
func (o *FullOutletLicenseDTO) HasCheckStatus() bool {
	if o != nil && !IsNil(o.CheckStatus) {
		return true
	}

	return false
}

// SetCheckStatus gets a reference to the given LicenseCheckStatusType and assigns it to the CheckStatus field.
func (o *FullOutletLicenseDTO) SetCheckStatus(v LicenseCheckStatusType) {
	o.CheckStatus = &v
}

// GetCheckComment returns the CheckComment field value if set, zero value otherwise.
func (o *FullOutletLicenseDTO) GetCheckComment() string {
	if o == nil || IsNil(o.CheckComment) {
		var ret string
		return ret
	}
	return *o.CheckComment
}

// GetCheckCommentOk returns a tuple with the CheckComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullOutletLicenseDTO) GetCheckCommentOk() (*string, bool) {
	if o == nil || IsNil(o.CheckComment) {
		return nil, false
	}
	return o.CheckComment, true
}

// HasCheckComment returns a boolean if a field has been set.
func (o *FullOutletLicenseDTO) HasCheckComment() bool {
	if o != nil && !IsNil(o.CheckComment) {
		return true
	}

	return false
}

// SetCheckComment gets a reference to the given string and assigns it to the CheckComment field.
func (o *FullOutletLicenseDTO) SetCheckComment(v string) {
	o.CheckComment = &v
}

func (o FullOutletLicenseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FullOutletLicenseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OutletId) {
		toSerialize["outletId"] = o.OutletId
	}
	if !IsNil(o.LicenseType) {
		toSerialize["licenseType"] = o.LicenseType
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.DateOfIssue) {
		toSerialize["dateOfIssue"] = o.DateOfIssue
	}
	if !IsNil(o.DateOfExpiry) {
		toSerialize["dateOfExpiry"] = o.DateOfExpiry
	}
	if !IsNil(o.CheckStatus) {
		toSerialize["checkStatus"] = o.CheckStatus
	}
	if !IsNil(o.CheckComment) {
		toSerialize["checkComment"] = o.CheckComment
	}
	return toSerialize, nil
}

type NullableFullOutletLicenseDTO struct {
	value *FullOutletLicenseDTO
	isSet bool
}

func (v NullableFullOutletLicenseDTO) Get() *FullOutletLicenseDTO {
	return v.value
}

func (v *NullableFullOutletLicenseDTO) Set(val *FullOutletLicenseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFullOutletLicenseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFullOutletLicenseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullOutletLicenseDTO(val *FullOutletLicenseDTO) *NullableFullOutletLicenseDTO {
	return &NullableFullOutletLicenseDTO{value: val, isSet: true}
}

func (v NullableFullOutletLicenseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullOutletLicenseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

