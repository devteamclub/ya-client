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

// checks if the OrdersStatsOrderDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrdersStatsOrderDTO{}

// OrdersStatsOrderDTO Информация о заказе.
type OrdersStatsOrderDTO struct {
	// Идентификатор заказа.
	Id *int64 `json:"id,omitempty"`
	// Дата создания заказа. Формат даты: `ГГГГ-ММ-ДД`. 
	CreationDate *string `json:"creationDate,omitempty"`
	// Дата и время, когда статус заказа был изменен в последний раз. Формат даты и времени: ISO 8601. Например, `2017-11-21T00:00:00`. Часовой пояс — UTC+03:00 (Москва). 
	StatusUpdateDate *time.Time `json:"statusUpdateDate,omitempty"`
	Status *OrderStatsStatusType `json:"status,omitempty"`
	// Идентификатор заказа в информационной системе магазина.
	PartnerOrderId *string `json:"partnerOrderId,omitempty"`
	PaymentType *OrdersStatsOrderPaymentType `json:"paymentType,omitempty"`
	// Тип заказа:  * false — заказ покупателя.  * true — тестовый заказ Маркета. 
	Fake *bool `json:"fake,omitempty"`
	DeliveryRegion *OrdersStatsDeliveryRegionDTO `json:"deliveryRegion,omitempty"`
	// Список товаров в заказе после возможных изменений.
	Items []OrdersStatsItemDTO `json:"items,omitempty"`
	// Список товаров в заказе до изменений.
	InitialItems []OrdersStatsItemDTO `json:"initialItems,omitempty"`
	// Информация о денежных переводах по заказу.
	Payments []OrdersStatsPaymentDTO `json:"payments,omitempty"`
	// Информация о комиссиях за заказ.
	Commissions []OrdersStatsCommissionDTO `json:"commissions,omitempty"`
}

// NewOrdersStatsOrderDTO instantiates a new OrdersStatsOrderDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdersStatsOrderDTO() *OrdersStatsOrderDTO {
	this := OrdersStatsOrderDTO{}
	return &this
}

// NewOrdersStatsOrderDTOWithDefaults instantiates a new OrdersStatsOrderDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdersStatsOrderDTOWithDefaults() *OrdersStatsOrderDTO {
	this := OrdersStatsOrderDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrdersStatsOrderDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsOrderDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrdersStatsOrderDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OrdersStatsOrderDTO) SetId(v int64) {
	o.Id = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *OrdersStatsOrderDTO) GetCreationDate() string {
	if o == nil || IsNil(o.CreationDate) {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsOrderDTO) GetCreationDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *OrdersStatsOrderDTO) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *OrdersStatsOrderDTO) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetStatusUpdateDate returns the StatusUpdateDate field value if set, zero value otherwise.
func (o *OrdersStatsOrderDTO) GetStatusUpdateDate() time.Time {
	if o == nil || IsNil(o.StatusUpdateDate) {
		var ret time.Time
		return ret
	}
	return *o.StatusUpdateDate
}

// GetStatusUpdateDateOk returns a tuple with the StatusUpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsOrderDTO) GetStatusUpdateDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StatusUpdateDate) {
		return nil, false
	}
	return o.StatusUpdateDate, true
}

// HasStatusUpdateDate returns a boolean if a field has been set.
func (o *OrdersStatsOrderDTO) HasStatusUpdateDate() bool {
	if o != nil && !IsNil(o.StatusUpdateDate) {
		return true
	}

	return false
}

// SetStatusUpdateDate gets a reference to the given time.Time and assigns it to the StatusUpdateDate field.
func (o *OrdersStatsOrderDTO) SetStatusUpdateDate(v time.Time) {
	o.StatusUpdateDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrdersStatsOrderDTO) GetStatus() OrderStatsStatusType {
	if o == nil || IsNil(o.Status) {
		var ret OrderStatsStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsOrderDTO) GetStatusOk() (*OrderStatsStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrdersStatsOrderDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OrderStatsStatusType and assigns it to the Status field.
func (o *OrdersStatsOrderDTO) SetStatus(v OrderStatsStatusType) {
	o.Status = &v
}

// GetPartnerOrderId returns the PartnerOrderId field value if set, zero value otherwise.
func (o *OrdersStatsOrderDTO) GetPartnerOrderId() string {
	if o == nil || IsNil(o.PartnerOrderId) {
		var ret string
		return ret
	}
	return *o.PartnerOrderId
}

// GetPartnerOrderIdOk returns a tuple with the PartnerOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsOrderDTO) GetPartnerOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerOrderId) {
		return nil, false
	}
	return o.PartnerOrderId, true
}

// HasPartnerOrderId returns a boolean if a field has been set.
func (o *OrdersStatsOrderDTO) HasPartnerOrderId() bool {
	if o != nil && !IsNil(o.PartnerOrderId) {
		return true
	}

	return false
}

// SetPartnerOrderId gets a reference to the given string and assigns it to the PartnerOrderId field.
func (o *OrdersStatsOrderDTO) SetPartnerOrderId(v string) {
	o.PartnerOrderId = &v
}

// GetPaymentType returns the PaymentType field value if set, zero value otherwise.
func (o *OrdersStatsOrderDTO) GetPaymentType() OrdersStatsOrderPaymentType {
	if o == nil || IsNil(o.PaymentType) {
		var ret OrdersStatsOrderPaymentType
		return ret
	}
	return *o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsOrderDTO) GetPaymentTypeOk() (*OrdersStatsOrderPaymentType, bool) {
	if o == nil || IsNil(o.PaymentType) {
		return nil, false
	}
	return o.PaymentType, true
}

// HasPaymentType returns a boolean if a field has been set.
func (o *OrdersStatsOrderDTO) HasPaymentType() bool {
	if o != nil && !IsNil(o.PaymentType) {
		return true
	}

	return false
}

// SetPaymentType gets a reference to the given OrdersStatsOrderPaymentType and assigns it to the PaymentType field.
func (o *OrdersStatsOrderDTO) SetPaymentType(v OrdersStatsOrderPaymentType) {
	o.PaymentType = &v
}

// GetFake returns the Fake field value if set, zero value otherwise.
func (o *OrdersStatsOrderDTO) GetFake() bool {
	if o == nil || IsNil(o.Fake) {
		var ret bool
		return ret
	}
	return *o.Fake
}

// GetFakeOk returns a tuple with the Fake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsOrderDTO) GetFakeOk() (*bool, bool) {
	if o == nil || IsNil(o.Fake) {
		return nil, false
	}
	return o.Fake, true
}

// HasFake returns a boolean if a field has been set.
func (o *OrdersStatsOrderDTO) HasFake() bool {
	if o != nil && !IsNil(o.Fake) {
		return true
	}

	return false
}

// SetFake gets a reference to the given bool and assigns it to the Fake field.
func (o *OrdersStatsOrderDTO) SetFake(v bool) {
	o.Fake = &v
}

// GetDeliveryRegion returns the DeliveryRegion field value if set, zero value otherwise.
func (o *OrdersStatsOrderDTO) GetDeliveryRegion() OrdersStatsDeliveryRegionDTO {
	if o == nil || IsNil(o.DeliveryRegion) {
		var ret OrdersStatsDeliveryRegionDTO
		return ret
	}
	return *o.DeliveryRegion
}

// GetDeliveryRegionOk returns a tuple with the DeliveryRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsOrderDTO) GetDeliveryRegionOk() (*OrdersStatsDeliveryRegionDTO, bool) {
	if o == nil || IsNil(o.DeliveryRegion) {
		return nil, false
	}
	return o.DeliveryRegion, true
}

// HasDeliveryRegion returns a boolean if a field has been set.
func (o *OrdersStatsOrderDTO) HasDeliveryRegion() bool {
	if o != nil && !IsNil(o.DeliveryRegion) {
		return true
	}

	return false
}

// SetDeliveryRegion gets a reference to the given OrdersStatsDeliveryRegionDTO and assigns it to the DeliveryRegion field.
func (o *OrdersStatsOrderDTO) SetDeliveryRegion(v OrdersStatsDeliveryRegionDTO) {
	o.DeliveryRegion = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *OrdersStatsOrderDTO) GetItems() []OrdersStatsItemDTO {
	if o == nil || IsNil(o.Items) {
		var ret []OrdersStatsItemDTO
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsOrderDTO) GetItemsOk() ([]OrdersStatsItemDTO, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *OrdersStatsOrderDTO) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrdersStatsItemDTO and assigns it to the Items field.
func (o *OrdersStatsOrderDTO) SetItems(v []OrdersStatsItemDTO) {
	o.Items = v
}

// GetInitialItems returns the InitialItems field value if set, zero value otherwise.
func (o *OrdersStatsOrderDTO) GetInitialItems() []OrdersStatsItemDTO {
	if o == nil || IsNil(o.InitialItems) {
		var ret []OrdersStatsItemDTO
		return ret
	}
	return o.InitialItems
}

// GetInitialItemsOk returns a tuple with the InitialItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsOrderDTO) GetInitialItemsOk() ([]OrdersStatsItemDTO, bool) {
	if o == nil || IsNil(o.InitialItems) {
		return nil, false
	}
	return o.InitialItems, true
}

// HasInitialItems returns a boolean if a field has been set.
func (o *OrdersStatsOrderDTO) HasInitialItems() bool {
	if o != nil && !IsNil(o.InitialItems) {
		return true
	}

	return false
}

// SetInitialItems gets a reference to the given []OrdersStatsItemDTO and assigns it to the InitialItems field.
func (o *OrdersStatsOrderDTO) SetInitialItems(v []OrdersStatsItemDTO) {
	o.InitialItems = v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *OrdersStatsOrderDTO) GetPayments() []OrdersStatsPaymentDTO {
	if o == nil || IsNil(o.Payments) {
		var ret []OrdersStatsPaymentDTO
		return ret
	}
	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsOrderDTO) GetPaymentsOk() ([]OrdersStatsPaymentDTO, bool) {
	if o == nil || IsNil(o.Payments) {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *OrdersStatsOrderDTO) HasPayments() bool {
	if o != nil && !IsNil(o.Payments) {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []OrdersStatsPaymentDTO and assigns it to the Payments field.
func (o *OrdersStatsOrderDTO) SetPayments(v []OrdersStatsPaymentDTO) {
	o.Payments = v
}

// GetCommissions returns the Commissions field value if set, zero value otherwise.
func (o *OrdersStatsOrderDTO) GetCommissions() []OrdersStatsCommissionDTO {
	if o == nil || IsNil(o.Commissions) {
		var ret []OrdersStatsCommissionDTO
		return ret
	}
	return o.Commissions
}

// GetCommissionsOk returns a tuple with the Commissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsOrderDTO) GetCommissionsOk() ([]OrdersStatsCommissionDTO, bool) {
	if o == nil || IsNil(o.Commissions) {
		return nil, false
	}
	return o.Commissions, true
}

// HasCommissions returns a boolean if a field has been set.
func (o *OrdersStatsOrderDTO) HasCommissions() bool {
	if o != nil && !IsNil(o.Commissions) {
		return true
	}

	return false
}

// SetCommissions gets a reference to the given []OrdersStatsCommissionDTO and assigns it to the Commissions field.
func (o *OrdersStatsOrderDTO) SetCommissions(v []OrdersStatsCommissionDTO) {
	o.Commissions = v
}

func (o OrdersStatsOrderDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrdersStatsOrderDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !IsNil(o.StatusUpdateDate) {
		toSerialize["statusUpdateDate"] = o.StatusUpdateDate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.PartnerOrderId) {
		toSerialize["partnerOrderId"] = o.PartnerOrderId
	}
	if !IsNil(o.PaymentType) {
		toSerialize["paymentType"] = o.PaymentType
	}
	if !IsNil(o.Fake) {
		toSerialize["fake"] = o.Fake
	}
	if !IsNil(o.DeliveryRegion) {
		toSerialize["deliveryRegion"] = o.DeliveryRegion
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.InitialItems) {
		toSerialize["initialItems"] = o.InitialItems
	}
	if !IsNil(o.Payments) {
		toSerialize["payments"] = o.Payments
	}
	if !IsNil(o.Commissions) {
		toSerialize["commissions"] = o.Commissions
	}
	return toSerialize, nil
}

type NullableOrdersStatsOrderDTO struct {
	value *OrdersStatsOrderDTO
	isSet bool
}

func (v NullableOrdersStatsOrderDTO) Get() *OrdersStatsOrderDTO {
	return v.value
}

func (v *NullableOrdersStatsOrderDTO) Set(val *OrdersStatsOrderDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersStatsOrderDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersStatsOrderDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersStatsOrderDTO(val *OrdersStatsOrderDTO) *NullableOrdersStatsOrderDTO {
	return &NullableOrdersStatsOrderDTO{value: val, isSet: true}
}

func (v NullableOrdersStatsOrderDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersStatsOrderDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

