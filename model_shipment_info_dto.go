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

// checks if the ShipmentInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentInfoDTO{}

// ShipmentInfoDTO Список с информацией об отгрузках.
type ShipmentInfoDTO struct {
	// Идентификатор отгрузки.
	Id *int64 `json:"id,omitempty"`
	// Начало планового интервала отгрузки.
	PlanIntervalFrom *time.Time `json:"planIntervalFrom,omitempty"`
	// Конец планового интервала отгрузки.
	PlanIntervalTo *time.Time `json:"planIntervalTo,omitempty"`
	ShipmentType *ShipmentType `json:"shipmentType,omitempty"`
	// Идентификатор отгрузки в вашей системе. Если вы еще не передавали идентификатор, вернется идентификатор из параметра `id`.
	ExternalId *string `json:"externalId,omitempty"`
	Status *ShipmentStatusType `json:"status,omitempty"`
	// Описание статуса отгрузки.
	StatusDescription *string `json:"statusDescription,omitempty"`
	DeliveryService *DeliveryServiceDTO `json:"deliveryService,omitempty"`
	// Количество заказов, запланированных к отгрузке.
	DraftCount *int32 `json:"draftCount,omitempty"`
	// Количество отгруженных заказов.
	PlannedCount *int32 `json:"plannedCount,omitempty"`
	// Количество заказов, принятых в сортировочном центре или пункте приема.
	FactCount *int32 `json:"factCount,omitempty"`
}

// NewShipmentInfoDTO instantiates a new ShipmentInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentInfoDTO() *ShipmentInfoDTO {
	this := ShipmentInfoDTO{}
	return &this
}

// NewShipmentInfoDTOWithDefaults instantiates a new ShipmentInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentInfoDTOWithDefaults() *ShipmentInfoDTO {
	this := ShipmentInfoDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShipmentInfoDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInfoDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShipmentInfoDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ShipmentInfoDTO) SetId(v int64) {
	o.Id = &v
}

// GetPlanIntervalFrom returns the PlanIntervalFrom field value if set, zero value otherwise.
func (o *ShipmentInfoDTO) GetPlanIntervalFrom() time.Time {
	if o == nil || IsNil(o.PlanIntervalFrom) {
		var ret time.Time
		return ret
	}
	return *o.PlanIntervalFrom
}

// GetPlanIntervalFromOk returns a tuple with the PlanIntervalFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInfoDTO) GetPlanIntervalFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PlanIntervalFrom) {
		return nil, false
	}
	return o.PlanIntervalFrom, true
}

// HasPlanIntervalFrom returns a boolean if a field has been set.
func (o *ShipmentInfoDTO) HasPlanIntervalFrom() bool {
	if o != nil && !IsNil(o.PlanIntervalFrom) {
		return true
	}

	return false
}

// SetPlanIntervalFrom gets a reference to the given time.Time and assigns it to the PlanIntervalFrom field.
func (o *ShipmentInfoDTO) SetPlanIntervalFrom(v time.Time) {
	o.PlanIntervalFrom = &v
}

// GetPlanIntervalTo returns the PlanIntervalTo field value if set, zero value otherwise.
func (o *ShipmentInfoDTO) GetPlanIntervalTo() time.Time {
	if o == nil || IsNil(o.PlanIntervalTo) {
		var ret time.Time
		return ret
	}
	return *o.PlanIntervalTo
}

// GetPlanIntervalToOk returns a tuple with the PlanIntervalTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInfoDTO) GetPlanIntervalToOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PlanIntervalTo) {
		return nil, false
	}
	return o.PlanIntervalTo, true
}

// HasPlanIntervalTo returns a boolean if a field has been set.
func (o *ShipmentInfoDTO) HasPlanIntervalTo() bool {
	if o != nil && !IsNil(o.PlanIntervalTo) {
		return true
	}

	return false
}

// SetPlanIntervalTo gets a reference to the given time.Time and assigns it to the PlanIntervalTo field.
func (o *ShipmentInfoDTO) SetPlanIntervalTo(v time.Time) {
	o.PlanIntervalTo = &v
}

// GetShipmentType returns the ShipmentType field value if set, zero value otherwise.
func (o *ShipmentInfoDTO) GetShipmentType() ShipmentType {
	if o == nil || IsNil(o.ShipmentType) {
		var ret ShipmentType
		return ret
	}
	return *o.ShipmentType
}

// GetShipmentTypeOk returns a tuple with the ShipmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInfoDTO) GetShipmentTypeOk() (*ShipmentType, bool) {
	if o == nil || IsNil(o.ShipmentType) {
		return nil, false
	}
	return o.ShipmentType, true
}

// HasShipmentType returns a boolean if a field has been set.
func (o *ShipmentInfoDTO) HasShipmentType() bool {
	if o != nil && !IsNil(o.ShipmentType) {
		return true
	}

	return false
}

// SetShipmentType gets a reference to the given ShipmentType and assigns it to the ShipmentType field.
func (o *ShipmentInfoDTO) SetShipmentType(v ShipmentType) {
	o.ShipmentType = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ShipmentInfoDTO) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInfoDTO) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ShipmentInfoDTO) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ShipmentInfoDTO) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ShipmentInfoDTO) GetStatus() ShipmentStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ShipmentStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInfoDTO) GetStatusOk() (*ShipmentStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ShipmentInfoDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ShipmentStatusType and assigns it to the Status field.
func (o *ShipmentInfoDTO) SetStatus(v ShipmentStatusType) {
	o.Status = &v
}

// GetStatusDescription returns the StatusDescription field value if set, zero value otherwise.
func (o *ShipmentInfoDTO) GetStatusDescription() string {
	if o == nil || IsNil(o.StatusDescription) {
		var ret string
		return ret
	}
	return *o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInfoDTO) GetStatusDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDescription) {
		return nil, false
	}
	return o.StatusDescription, true
}

// HasStatusDescription returns a boolean if a field has been set.
func (o *ShipmentInfoDTO) HasStatusDescription() bool {
	if o != nil && !IsNil(o.StatusDescription) {
		return true
	}

	return false
}

// SetStatusDescription gets a reference to the given string and assigns it to the StatusDescription field.
func (o *ShipmentInfoDTO) SetStatusDescription(v string) {
	o.StatusDescription = &v
}

// GetDeliveryService returns the DeliveryService field value if set, zero value otherwise.
func (o *ShipmentInfoDTO) GetDeliveryService() DeliveryServiceDTO {
	if o == nil || IsNil(o.DeliveryService) {
		var ret DeliveryServiceDTO
		return ret
	}
	return *o.DeliveryService
}

// GetDeliveryServiceOk returns a tuple with the DeliveryService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInfoDTO) GetDeliveryServiceOk() (*DeliveryServiceDTO, bool) {
	if o == nil || IsNil(o.DeliveryService) {
		return nil, false
	}
	return o.DeliveryService, true
}

// HasDeliveryService returns a boolean if a field has been set.
func (o *ShipmentInfoDTO) HasDeliveryService() bool {
	if o != nil && !IsNil(o.DeliveryService) {
		return true
	}

	return false
}

// SetDeliveryService gets a reference to the given DeliveryServiceDTO and assigns it to the DeliveryService field.
func (o *ShipmentInfoDTO) SetDeliveryService(v DeliveryServiceDTO) {
	o.DeliveryService = &v
}

// GetDraftCount returns the DraftCount field value if set, zero value otherwise.
func (o *ShipmentInfoDTO) GetDraftCount() int32 {
	if o == nil || IsNil(o.DraftCount) {
		var ret int32
		return ret
	}
	return *o.DraftCount
}

// GetDraftCountOk returns a tuple with the DraftCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInfoDTO) GetDraftCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DraftCount) {
		return nil, false
	}
	return o.DraftCount, true
}

// HasDraftCount returns a boolean if a field has been set.
func (o *ShipmentInfoDTO) HasDraftCount() bool {
	if o != nil && !IsNil(o.DraftCount) {
		return true
	}

	return false
}

// SetDraftCount gets a reference to the given int32 and assigns it to the DraftCount field.
func (o *ShipmentInfoDTO) SetDraftCount(v int32) {
	o.DraftCount = &v
}

// GetPlannedCount returns the PlannedCount field value if set, zero value otherwise.
func (o *ShipmentInfoDTO) GetPlannedCount() int32 {
	if o == nil || IsNil(o.PlannedCount) {
		var ret int32
		return ret
	}
	return *o.PlannedCount
}

// GetPlannedCountOk returns a tuple with the PlannedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInfoDTO) GetPlannedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PlannedCount) {
		return nil, false
	}
	return o.PlannedCount, true
}

// HasPlannedCount returns a boolean if a field has been set.
func (o *ShipmentInfoDTO) HasPlannedCount() bool {
	if o != nil && !IsNil(o.PlannedCount) {
		return true
	}

	return false
}

// SetPlannedCount gets a reference to the given int32 and assigns it to the PlannedCount field.
func (o *ShipmentInfoDTO) SetPlannedCount(v int32) {
	o.PlannedCount = &v
}

// GetFactCount returns the FactCount field value if set, zero value otherwise.
func (o *ShipmentInfoDTO) GetFactCount() int32 {
	if o == nil || IsNil(o.FactCount) {
		var ret int32
		return ret
	}
	return *o.FactCount
}

// GetFactCountOk returns a tuple with the FactCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInfoDTO) GetFactCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FactCount) {
		return nil, false
	}
	return o.FactCount, true
}

// HasFactCount returns a boolean if a field has been set.
func (o *ShipmentInfoDTO) HasFactCount() bool {
	if o != nil && !IsNil(o.FactCount) {
		return true
	}

	return false
}

// SetFactCount gets a reference to the given int32 and assigns it to the FactCount field.
func (o *ShipmentInfoDTO) SetFactCount(v int32) {
	o.FactCount = &v
}

func (o ShipmentInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PlanIntervalFrom) {
		toSerialize["planIntervalFrom"] = o.PlanIntervalFrom
	}
	if !IsNil(o.PlanIntervalTo) {
		toSerialize["planIntervalTo"] = o.PlanIntervalTo
	}
	if !IsNil(o.ShipmentType) {
		toSerialize["shipmentType"] = o.ShipmentType
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusDescription) {
		toSerialize["statusDescription"] = o.StatusDescription
	}
	if !IsNil(o.DeliveryService) {
		toSerialize["deliveryService"] = o.DeliveryService
	}
	if !IsNil(o.DraftCount) {
		toSerialize["draftCount"] = o.DraftCount
	}
	if !IsNil(o.PlannedCount) {
		toSerialize["plannedCount"] = o.PlannedCount
	}
	if !IsNil(o.FactCount) {
		toSerialize["factCount"] = o.FactCount
	}
	return toSerialize, nil
}

type NullableShipmentInfoDTO struct {
	value *ShipmentInfoDTO
	isSet bool
}

func (v NullableShipmentInfoDTO) Get() *ShipmentInfoDTO {
	return v.value
}

func (v *NullableShipmentInfoDTO) Set(val *ShipmentInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentInfoDTO(val *ShipmentInfoDTO) *NullableShipmentInfoDTO {
	return &NullableShipmentInfoDTO{value: val, isSet: true}
}

func (v NullableShipmentInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


