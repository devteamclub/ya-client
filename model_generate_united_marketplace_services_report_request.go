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
	"fmt"
)

// checks if the GenerateUnitedMarketplaceServicesReportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateUnitedMarketplaceServicesReportRequest{}

// GenerateUnitedMarketplaceServicesReportRequest Данные, необходимые для генерации отчета: идентификатор магазина, период, за который нужен отчет, а также фильтры. 
type GenerateUnitedMarketplaceServicesReportRequest struct {
	// Идентификатор бизнеса.
	BusinessId int64 `json:"businessId"`
	// Начало периода, включительно.
	DateTimeFrom *time.Time `json:"dateTimeFrom,omitempty"`
	// Конец периода, включительно. Максимальный период — 1 год.
	DateTimeTo *time.Time `json:"dateTimeTo,omitempty"`
	// Начальный год формирования акта.
	YearFrom *int32 `json:"yearFrom,omitempty"`
	// Начальный номер месяца формирования акта.
	MonthFrom *int32 `json:"monthFrom,omitempty"`
	// Конечный год формирования акта.
	YearTo *int32 `json:"yearTo,omitempty"`
	// Конечный номер месяца формирования акта.
	MonthTo *int32 `json:"monthTo,omitempty"`
	// Список моделей, которые нужны в отчете. 
	PlacementPrograms []PlacementType `json:"placementPrograms,omitempty"`
	// Список ИНН, которые нужны в отчете.
	Inns []string `json:"inns,omitempty"`
	// Список магазинов, которые нужны в отчете.
	CampaignIds []int64 `json:"campaignIds,omitempty"`
}

type _GenerateUnitedMarketplaceServicesReportRequest GenerateUnitedMarketplaceServicesReportRequest

// NewGenerateUnitedMarketplaceServicesReportRequest instantiates a new GenerateUnitedMarketplaceServicesReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateUnitedMarketplaceServicesReportRequest(businessId int64) *GenerateUnitedMarketplaceServicesReportRequest {
	this := GenerateUnitedMarketplaceServicesReportRequest{}
	this.BusinessId = businessId
	return &this
}

// NewGenerateUnitedMarketplaceServicesReportRequestWithDefaults instantiates a new GenerateUnitedMarketplaceServicesReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateUnitedMarketplaceServicesReportRequestWithDefaults() *GenerateUnitedMarketplaceServicesReportRequest {
	this := GenerateUnitedMarketplaceServicesReportRequest{}
	return &this
}

// GetBusinessId returns the BusinessId field value
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetBusinessId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value
// and a boolean to check if the value has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetBusinessIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessId, true
}

// SetBusinessId sets field value
func (o *GenerateUnitedMarketplaceServicesReportRequest) SetBusinessId(v int64) {
	o.BusinessId = v
}

// GetDateTimeFrom returns the DateTimeFrom field value if set, zero value otherwise.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetDateTimeFrom() time.Time {
	if o == nil || IsNil(o.DateTimeFrom) {
		var ret time.Time
		return ret
	}
	return *o.DateTimeFrom
}

// GetDateTimeFromOk returns a tuple with the DateTimeFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetDateTimeFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateTimeFrom) {
		return nil, false
	}
	return o.DateTimeFrom, true
}

// HasDateTimeFrom returns a boolean if a field has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) HasDateTimeFrom() bool {
	if o != nil && !IsNil(o.DateTimeFrom) {
		return true
	}

	return false
}

// SetDateTimeFrom gets a reference to the given time.Time and assigns it to the DateTimeFrom field.
func (o *GenerateUnitedMarketplaceServicesReportRequest) SetDateTimeFrom(v time.Time) {
	o.DateTimeFrom = &v
}

// GetDateTimeTo returns the DateTimeTo field value if set, zero value otherwise.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetDateTimeTo() time.Time {
	if o == nil || IsNil(o.DateTimeTo) {
		var ret time.Time
		return ret
	}
	return *o.DateTimeTo
}

// GetDateTimeToOk returns a tuple with the DateTimeTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetDateTimeToOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateTimeTo) {
		return nil, false
	}
	return o.DateTimeTo, true
}

// HasDateTimeTo returns a boolean if a field has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) HasDateTimeTo() bool {
	if o != nil && !IsNil(o.DateTimeTo) {
		return true
	}

	return false
}

// SetDateTimeTo gets a reference to the given time.Time and assigns it to the DateTimeTo field.
func (o *GenerateUnitedMarketplaceServicesReportRequest) SetDateTimeTo(v time.Time) {
	o.DateTimeTo = &v
}

// GetYearFrom returns the YearFrom field value if set, zero value otherwise.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetYearFrom() int32 {
	if o == nil || IsNil(o.YearFrom) {
		var ret int32
		return ret
	}
	return *o.YearFrom
}

// GetYearFromOk returns a tuple with the YearFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetYearFromOk() (*int32, bool) {
	if o == nil || IsNil(o.YearFrom) {
		return nil, false
	}
	return o.YearFrom, true
}

// HasYearFrom returns a boolean if a field has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) HasYearFrom() bool {
	if o != nil && !IsNil(o.YearFrom) {
		return true
	}

	return false
}

// SetYearFrom gets a reference to the given int32 and assigns it to the YearFrom field.
func (o *GenerateUnitedMarketplaceServicesReportRequest) SetYearFrom(v int32) {
	o.YearFrom = &v
}

// GetMonthFrom returns the MonthFrom field value if set, zero value otherwise.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetMonthFrom() int32 {
	if o == nil || IsNil(o.MonthFrom) {
		var ret int32
		return ret
	}
	return *o.MonthFrom
}

// GetMonthFromOk returns a tuple with the MonthFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetMonthFromOk() (*int32, bool) {
	if o == nil || IsNil(o.MonthFrom) {
		return nil, false
	}
	return o.MonthFrom, true
}

// HasMonthFrom returns a boolean if a field has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) HasMonthFrom() bool {
	if o != nil && !IsNil(o.MonthFrom) {
		return true
	}

	return false
}

// SetMonthFrom gets a reference to the given int32 and assigns it to the MonthFrom field.
func (o *GenerateUnitedMarketplaceServicesReportRequest) SetMonthFrom(v int32) {
	o.MonthFrom = &v
}

// GetYearTo returns the YearTo field value if set, zero value otherwise.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetYearTo() int32 {
	if o == nil || IsNil(o.YearTo) {
		var ret int32
		return ret
	}
	return *o.YearTo
}

// GetYearToOk returns a tuple with the YearTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetYearToOk() (*int32, bool) {
	if o == nil || IsNil(o.YearTo) {
		return nil, false
	}
	return o.YearTo, true
}

// HasYearTo returns a boolean if a field has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) HasYearTo() bool {
	if o != nil && !IsNil(o.YearTo) {
		return true
	}

	return false
}

// SetYearTo gets a reference to the given int32 and assigns it to the YearTo field.
func (o *GenerateUnitedMarketplaceServicesReportRequest) SetYearTo(v int32) {
	o.YearTo = &v
}

// GetMonthTo returns the MonthTo field value if set, zero value otherwise.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetMonthTo() int32 {
	if o == nil || IsNil(o.MonthTo) {
		var ret int32
		return ret
	}
	return *o.MonthTo
}

// GetMonthToOk returns a tuple with the MonthTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetMonthToOk() (*int32, bool) {
	if o == nil || IsNil(o.MonthTo) {
		return nil, false
	}
	return o.MonthTo, true
}

// HasMonthTo returns a boolean if a field has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) HasMonthTo() bool {
	if o != nil && !IsNil(o.MonthTo) {
		return true
	}

	return false
}

// SetMonthTo gets a reference to the given int32 and assigns it to the MonthTo field.
func (o *GenerateUnitedMarketplaceServicesReportRequest) SetMonthTo(v int32) {
	o.MonthTo = &v
}

// GetPlacementPrograms returns the PlacementPrograms field value if set, zero value otherwise.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetPlacementPrograms() []PlacementType {
	if o == nil || IsNil(o.PlacementPrograms) {
		var ret []PlacementType
		return ret
	}
	return o.PlacementPrograms
}

// GetPlacementProgramsOk returns a tuple with the PlacementPrograms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetPlacementProgramsOk() ([]PlacementType, bool) {
	if o == nil || IsNil(o.PlacementPrograms) {
		return nil, false
	}
	return o.PlacementPrograms, true
}

// HasPlacementPrograms returns a boolean if a field has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) HasPlacementPrograms() bool {
	if o != nil && !IsNil(o.PlacementPrograms) {
		return true
	}

	return false
}

// SetPlacementPrograms gets a reference to the given []PlacementType and assigns it to the PlacementPrograms field.
func (o *GenerateUnitedMarketplaceServicesReportRequest) SetPlacementPrograms(v []PlacementType) {
	o.PlacementPrograms = v
}

// GetInns returns the Inns field value if set, zero value otherwise.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetInns() []string {
	if o == nil || IsNil(o.Inns) {
		var ret []string
		return ret
	}
	return o.Inns
}

// GetInnsOk returns a tuple with the Inns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetInnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Inns) {
		return nil, false
	}
	return o.Inns, true
}

// HasInns returns a boolean if a field has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) HasInns() bool {
	if o != nil && !IsNil(o.Inns) {
		return true
	}

	return false
}

// SetInns gets a reference to the given []string and assigns it to the Inns field.
func (o *GenerateUnitedMarketplaceServicesReportRequest) SetInns(v []string) {
	o.Inns = v
}

// GetCampaignIds returns the CampaignIds field value if set, zero value otherwise.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetCampaignIds() []int64 {
	if o == nil || IsNil(o.CampaignIds) {
		var ret []int64
		return ret
	}
	return o.CampaignIds
}

// GetCampaignIdsOk returns a tuple with the CampaignIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) GetCampaignIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.CampaignIds) {
		return nil, false
	}
	return o.CampaignIds, true
}

// HasCampaignIds returns a boolean if a field has been set.
func (o *GenerateUnitedMarketplaceServicesReportRequest) HasCampaignIds() bool {
	if o != nil && !IsNil(o.CampaignIds) {
		return true
	}

	return false
}

// SetCampaignIds gets a reference to the given []int64 and assigns it to the CampaignIds field.
func (o *GenerateUnitedMarketplaceServicesReportRequest) SetCampaignIds(v []int64) {
	o.CampaignIds = v
}

func (o GenerateUnitedMarketplaceServicesReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateUnitedMarketplaceServicesReportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["businessId"] = o.BusinessId
	if !IsNil(o.DateTimeFrom) {
		toSerialize["dateTimeFrom"] = o.DateTimeFrom
	}
	if !IsNil(o.DateTimeTo) {
		toSerialize["dateTimeTo"] = o.DateTimeTo
	}
	if !IsNil(o.YearFrom) {
		toSerialize["yearFrom"] = o.YearFrom
	}
	if !IsNil(o.MonthFrom) {
		toSerialize["monthFrom"] = o.MonthFrom
	}
	if !IsNil(o.YearTo) {
		toSerialize["yearTo"] = o.YearTo
	}
	if !IsNil(o.MonthTo) {
		toSerialize["monthTo"] = o.MonthTo
	}
	if !IsNil(o.PlacementPrograms) {
		toSerialize["placementPrograms"] = o.PlacementPrograms
	}
	if !IsNil(o.Inns) {
		toSerialize["inns"] = o.Inns
	}
	if !IsNil(o.CampaignIds) {
		toSerialize["campaignIds"] = o.CampaignIds
	}
	return toSerialize, nil
}

func (o *GenerateUnitedMarketplaceServicesReportRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"businessId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGenerateUnitedMarketplaceServicesReportRequest := _GenerateUnitedMarketplaceServicesReportRequest{}

	err = json.Unmarshal(bytes, &varGenerateUnitedMarketplaceServicesReportRequest)

	if err != nil {
		return err
	}

	*o = GenerateUnitedMarketplaceServicesReportRequest(varGenerateUnitedMarketplaceServicesReportRequest)

	return err
}

type NullableGenerateUnitedMarketplaceServicesReportRequest struct {
	value *GenerateUnitedMarketplaceServicesReportRequest
	isSet bool
}

func (v NullableGenerateUnitedMarketplaceServicesReportRequest) Get() *GenerateUnitedMarketplaceServicesReportRequest {
	return v.value
}

func (v *NullableGenerateUnitedMarketplaceServicesReportRequest) Set(val *GenerateUnitedMarketplaceServicesReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateUnitedMarketplaceServicesReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateUnitedMarketplaceServicesReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateUnitedMarketplaceServicesReportRequest(val *GenerateUnitedMarketplaceServicesReportRequest) *NullableGenerateUnitedMarketplaceServicesReportRequest {
	return &NullableGenerateUnitedMarketplaceServicesReportRequest{value: val, isSet: true}
}

func (v NullableGenerateUnitedMarketplaceServicesReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateUnitedMarketplaceServicesReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


