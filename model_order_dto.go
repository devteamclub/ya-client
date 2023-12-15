/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the OrderDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDTO{}

// OrderDTO Заказ.
type OrderDTO struct {
	// Идентификатор заказа.
	Id *int64 `json:"id,omitempty"`
	Status *OrderStatusType `json:"status,omitempty"`
	Substatus *OrderSubstatusType `json:"substatus,omitempty"`
	CreationDate *string `json:"creationDate,omitempty"`
	Currency *CurrencyType `json:"currency,omitempty"`
	// Общая сумма заказа в валюте заказа без учета стоимости доставки и вознаграждения партнеру за скидки по промокодам, купонам и акциям (параметр subsidyTotal).  Для отделения целой части от дробной используется точка. 
	ItemsTotal *float32 `json:"itemsTotal,omitempty"`
	// Общая сумма заказа в валюте заказа с учетом стоимости доставки, но без учета вознаграждения партнеру за скидки по промокодам, купонам, кешбэку и акциям (параметр `subsidyTotal`).  Для отделения целой части от дробной используется точка. 
	Total *float32 `json:"total,omitempty"`
	// Стоимость доставки в валюте заказа.
	DeliveryTotal *float32 `json:"deliveryTotal,omitempty"`
	// Общее вознаграждение партнеру за скидки:  * по промокодам; * по купонам; * по баллам кешбэка по подписке Яндекс Плюс; * по акциям.  Передается в валюте, указанной в параметре `currency`.  Для отделения целой части от дробной используется точка. 
	SubsidyTotal *float32 `json:"subsidyTotal,omitempty"`
	// Сумма стоимости всех товаров в заказе и вознаграждения за них в валюте магазина (сумма параметров total и subsidyTotal)
	TotalWithSubsidy *float32 `json:"totalWithSubsidy,omitempty"`
	// Стоимость всех товаров в заказе в валюте покупателя после применения скидок и без учета стоимости доставки.
	BuyerItemsTotal *float32 `json:"buyerItemsTotal,omitempty"`
	// Стоимость всех товаров в заказе в валюте покупателя после применения скидок и с учетом стоимости доставки.
	BuyerTotal *float32 `json:"buyerTotal,omitempty"`
	// Стоимость всех товаров в заказе в валюте покупателя до применения скидок и без учета стоимости доставки.
	BuyerItemsTotalBeforeDiscount *float32 `json:"buyerItemsTotalBeforeDiscount,omitempty"`
	// Стоимость всех товаров в заказе в валюте покупателя до применения скидок и с учетом стоимости доставки.
	BuyerTotalBeforeDiscount *float32 `json:"buyerTotalBeforeDiscount,omitempty"`
	PaymentType *OrderPaymentType `json:"paymentType,omitempty"`
	PaymentMethod *OrderPaymentMethodType `json:"paymentMethod,omitempty"`
	// Тип заказа:  * false — заказ покупателя.  * true — тестовый заказ Маркета. 
	Fake *bool `json:"fake,omitempty"`
	// Список товаров в заказе.
	Items []OrderItemDTO `json:"items,omitempty"`
	// Список субсидий по типам.
	Subsidies []OrderItemSubsidyDTO `json:"subsidies,omitempty"`
	Delivery *OrderDeliveryDTO `json:"delivery,omitempty"`
	Buyer *OrderBuyerDTO `json:"buyer,omitempty"`
	// Комментарий к заказу.
	Notes *string `json:"notes,omitempty"`
	TaxSystem *OrderTaxSystemType `json:"taxSystem,omitempty"`
	// Запрошена ли отмена.
	CancelRequested *bool `json:"cancelRequested,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
}

// NewOrderDTO instantiates a new OrderDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDTO() *OrderDTO {
	this := OrderDTO{}
	return &this
}

// NewOrderDTOWithDefaults instantiates a new OrderDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDTOWithDefaults() *OrderDTO {
	this := OrderDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OrderDTO) SetId(v int64) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderDTO) GetStatus() OrderStatusType {
	if o == nil || IsNil(o.Status) {
		var ret OrderStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetStatusOk() (*OrderStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OrderStatusType and assigns it to the Status field.
func (o *OrderDTO) SetStatus(v OrderStatusType) {
	o.Status = &v
}

// GetSubstatus returns the Substatus field value if set, zero value otherwise.
func (o *OrderDTO) GetSubstatus() OrderSubstatusType {
	if o == nil || IsNil(o.Substatus) {
		var ret OrderSubstatusType
		return ret
	}
	return *o.Substatus
}

// GetSubstatusOk returns a tuple with the Substatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetSubstatusOk() (*OrderSubstatusType, bool) {
	if o == nil || IsNil(o.Substatus) {
		return nil, false
	}
	return o.Substatus, true
}

// HasSubstatus returns a boolean if a field has been set.
func (o *OrderDTO) HasSubstatus() bool {
	if o != nil && !IsNil(o.Substatus) {
		return true
	}

	return false
}

// SetSubstatus gets a reference to the given OrderSubstatusType and assigns it to the Substatus field.
func (o *OrderDTO) SetSubstatus(v OrderSubstatusType) {
	o.Substatus = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *OrderDTO) GetCreationDate() string {
	if o == nil || IsNil(o.CreationDate) {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetCreationDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *OrderDTO) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *OrderDTO) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *OrderDTO) GetCurrency() CurrencyType {
	if o == nil || IsNil(o.Currency) {
		var ret CurrencyType
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetCurrencyOk() (*CurrencyType, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *OrderDTO) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CurrencyType and assigns it to the Currency field.
func (o *OrderDTO) SetCurrency(v CurrencyType) {
	o.Currency = &v
}

// GetItemsTotal returns the ItemsTotal field value if set, zero value otherwise.
func (o *OrderDTO) GetItemsTotal() float32 {
	if o == nil || IsNil(o.ItemsTotal) {
		var ret float32
		return ret
	}
	return *o.ItemsTotal
}

// GetItemsTotalOk returns a tuple with the ItemsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetItemsTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.ItemsTotal) {
		return nil, false
	}
	return o.ItemsTotal, true
}

// HasItemsTotal returns a boolean if a field has been set.
func (o *OrderDTO) HasItemsTotal() bool {
	if o != nil && !IsNil(o.ItemsTotal) {
		return true
	}

	return false
}

// SetItemsTotal gets a reference to the given float32 and assigns it to the ItemsTotal field.
func (o *OrderDTO) SetItemsTotal(v float32) {
	o.ItemsTotal = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrderDTO) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrderDTO) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *OrderDTO) SetTotal(v float32) {
	o.Total = &v
}

// GetDeliveryTotal returns the DeliveryTotal field value if set, zero value otherwise.
func (o *OrderDTO) GetDeliveryTotal() float32 {
	if o == nil || IsNil(o.DeliveryTotal) {
		var ret float32
		return ret
	}
	return *o.DeliveryTotal
}

// GetDeliveryTotalOk returns a tuple with the DeliveryTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetDeliveryTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.DeliveryTotal) {
		return nil, false
	}
	return o.DeliveryTotal, true
}

// HasDeliveryTotal returns a boolean if a field has been set.
func (o *OrderDTO) HasDeliveryTotal() bool {
	if o != nil && !IsNil(o.DeliveryTotal) {
		return true
	}

	return false
}

// SetDeliveryTotal gets a reference to the given float32 and assigns it to the DeliveryTotal field.
func (o *OrderDTO) SetDeliveryTotal(v float32) {
	o.DeliveryTotal = &v
}

// GetSubsidyTotal returns the SubsidyTotal field value if set, zero value otherwise.
func (o *OrderDTO) GetSubsidyTotal() float32 {
	if o == nil || IsNil(o.SubsidyTotal) {
		var ret float32
		return ret
	}
	return *o.SubsidyTotal
}

// GetSubsidyTotalOk returns a tuple with the SubsidyTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetSubsidyTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.SubsidyTotal) {
		return nil, false
	}
	return o.SubsidyTotal, true
}

// HasSubsidyTotal returns a boolean if a field has been set.
func (o *OrderDTO) HasSubsidyTotal() bool {
	if o != nil && !IsNil(o.SubsidyTotal) {
		return true
	}

	return false
}

// SetSubsidyTotal gets a reference to the given float32 and assigns it to the SubsidyTotal field.
func (o *OrderDTO) SetSubsidyTotal(v float32) {
	o.SubsidyTotal = &v
}

// GetTotalWithSubsidy returns the TotalWithSubsidy field value if set, zero value otherwise.
func (o *OrderDTO) GetTotalWithSubsidy() float32 {
	if o == nil || IsNil(o.TotalWithSubsidy) {
		var ret float32
		return ret
	}
	return *o.TotalWithSubsidy
}

// GetTotalWithSubsidyOk returns a tuple with the TotalWithSubsidy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetTotalWithSubsidyOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalWithSubsidy) {
		return nil, false
	}
	return o.TotalWithSubsidy, true
}

// HasTotalWithSubsidy returns a boolean if a field has been set.
func (o *OrderDTO) HasTotalWithSubsidy() bool {
	if o != nil && !IsNil(o.TotalWithSubsidy) {
		return true
	}

	return false
}

// SetTotalWithSubsidy gets a reference to the given float32 and assigns it to the TotalWithSubsidy field.
func (o *OrderDTO) SetTotalWithSubsidy(v float32) {
	o.TotalWithSubsidy = &v
}

// GetBuyerItemsTotal returns the BuyerItemsTotal field value if set, zero value otherwise.
func (o *OrderDTO) GetBuyerItemsTotal() float32 {
	if o == nil || IsNil(o.BuyerItemsTotal) {
		var ret float32
		return ret
	}
	return *o.BuyerItemsTotal
}

// GetBuyerItemsTotalOk returns a tuple with the BuyerItemsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetBuyerItemsTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.BuyerItemsTotal) {
		return nil, false
	}
	return o.BuyerItemsTotal, true
}

// HasBuyerItemsTotal returns a boolean if a field has been set.
func (o *OrderDTO) HasBuyerItemsTotal() bool {
	if o != nil && !IsNil(o.BuyerItemsTotal) {
		return true
	}

	return false
}

// SetBuyerItemsTotal gets a reference to the given float32 and assigns it to the BuyerItemsTotal field.
func (o *OrderDTO) SetBuyerItemsTotal(v float32) {
	o.BuyerItemsTotal = &v
}

// GetBuyerTotal returns the BuyerTotal field value if set, zero value otherwise.
func (o *OrderDTO) GetBuyerTotal() float32 {
	if o == nil || IsNil(o.BuyerTotal) {
		var ret float32
		return ret
	}
	return *o.BuyerTotal
}

// GetBuyerTotalOk returns a tuple with the BuyerTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetBuyerTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.BuyerTotal) {
		return nil, false
	}
	return o.BuyerTotal, true
}

// HasBuyerTotal returns a boolean if a field has been set.
func (o *OrderDTO) HasBuyerTotal() bool {
	if o != nil && !IsNil(o.BuyerTotal) {
		return true
	}

	return false
}

// SetBuyerTotal gets a reference to the given float32 and assigns it to the BuyerTotal field.
func (o *OrderDTO) SetBuyerTotal(v float32) {
	o.BuyerTotal = &v
}

// GetBuyerItemsTotalBeforeDiscount returns the BuyerItemsTotalBeforeDiscount field value if set, zero value otherwise.
func (o *OrderDTO) GetBuyerItemsTotalBeforeDiscount() float32 {
	if o == nil || IsNil(o.BuyerItemsTotalBeforeDiscount) {
		var ret float32
		return ret
	}
	return *o.BuyerItemsTotalBeforeDiscount
}

// GetBuyerItemsTotalBeforeDiscountOk returns a tuple with the BuyerItemsTotalBeforeDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetBuyerItemsTotalBeforeDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.BuyerItemsTotalBeforeDiscount) {
		return nil, false
	}
	return o.BuyerItemsTotalBeforeDiscount, true
}

// HasBuyerItemsTotalBeforeDiscount returns a boolean if a field has been set.
func (o *OrderDTO) HasBuyerItemsTotalBeforeDiscount() bool {
	if o != nil && !IsNil(o.BuyerItemsTotalBeforeDiscount) {
		return true
	}

	return false
}

// SetBuyerItemsTotalBeforeDiscount gets a reference to the given float32 and assigns it to the BuyerItemsTotalBeforeDiscount field.
func (o *OrderDTO) SetBuyerItemsTotalBeforeDiscount(v float32) {
	o.BuyerItemsTotalBeforeDiscount = &v
}

// GetBuyerTotalBeforeDiscount returns the BuyerTotalBeforeDiscount field value if set, zero value otherwise.
func (o *OrderDTO) GetBuyerTotalBeforeDiscount() float32 {
	if o == nil || IsNil(o.BuyerTotalBeforeDiscount) {
		var ret float32
		return ret
	}
	return *o.BuyerTotalBeforeDiscount
}

// GetBuyerTotalBeforeDiscountOk returns a tuple with the BuyerTotalBeforeDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetBuyerTotalBeforeDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.BuyerTotalBeforeDiscount) {
		return nil, false
	}
	return o.BuyerTotalBeforeDiscount, true
}

// HasBuyerTotalBeforeDiscount returns a boolean if a field has been set.
func (o *OrderDTO) HasBuyerTotalBeforeDiscount() bool {
	if o != nil && !IsNil(o.BuyerTotalBeforeDiscount) {
		return true
	}

	return false
}

// SetBuyerTotalBeforeDiscount gets a reference to the given float32 and assigns it to the BuyerTotalBeforeDiscount field.
func (o *OrderDTO) SetBuyerTotalBeforeDiscount(v float32) {
	o.BuyerTotalBeforeDiscount = &v
}

// GetPaymentType returns the PaymentType field value if set, zero value otherwise.
func (o *OrderDTO) GetPaymentType() OrderPaymentType {
	if o == nil || IsNil(o.PaymentType) {
		var ret OrderPaymentType
		return ret
	}
	return *o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetPaymentTypeOk() (*OrderPaymentType, bool) {
	if o == nil || IsNil(o.PaymentType) {
		return nil, false
	}
	return o.PaymentType, true
}

// HasPaymentType returns a boolean if a field has been set.
func (o *OrderDTO) HasPaymentType() bool {
	if o != nil && !IsNil(o.PaymentType) {
		return true
	}

	return false
}

// SetPaymentType gets a reference to the given OrderPaymentType and assigns it to the PaymentType field.
func (o *OrderDTO) SetPaymentType(v OrderPaymentType) {
	o.PaymentType = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *OrderDTO) GetPaymentMethod() OrderPaymentMethodType {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret OrderPaymentMethodType
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetPaymentMethodOk() (*OrderPaymentMethodType, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *OrderDTO) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given OrderPaymentMethodType and assigns it to the PaymentMethod field.
func (o *OrderDTO) SetPaymentMethod(v OrderPaymentMethodType) {
	o.PaymentMethod = &v
}

// GetFake returns the Fake field value if set, zero value otherwise.
func (o *OrderDTO) GetFake() bool {
	if o == nil || IsNil(o.Fake) {
		var ret bool
		return ret
	}
	return *o.Fake
}

// GetFakeOk returns a tuple with the Fake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetFakeOk() (*bool, bool) {
	if o == nil || IsNil(o.Fake) {
		return nil, false
	}
	return o.Fake, true
}

// HasFake returns a boolean if a field has been set.
func (o *OrderDTO) HasFake() bool {
	if o != nil && !IsNil(o.Fake) {
		return true
	}

	return false
}

// SetFake gets a reference to the given bool and assigns it to the Fake field.
func (o *OrderDTO) SetFake(v bool) {
	o.Fake = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *OrderDTO) GetItems() []OrderItemDTO {
	if o == nil || IsNil(o.Items) {
		var ret []OrderItemDTO
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetItemsOk() ([]OrderItemDTO, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *OrderDTO) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrderItemDTO and assigns it to the Items field.
func (o *OrderDTO) SetItems(v []OrderItemDTO) {
	o.Items = v
}

// GetSubsidies returns the Subsidies field value if set, zero value otherwise.
func (o *OrderDTO) GetSubsidies() []OrderItemSubsidyDTO {
	if o == nil || IsNil(o.Subsidies) {
		var ret []OrderItemSubsidyDTO
		return ret
	}
	return o.Subsidies
}

// GetSubsidiesOk returns a tuple with the Subsidies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetSubsidiesOk() ([]OrderItemSubsidyDTO, bool) {
	if o == nil || IsNil(o.Subsidies) {
		return nil, false
	}
	return o.Subsidies, true
}

// HasSubsidies returns a boolean if a field has been set.
func (o *OrderDTO) HasSubsidies() bool {
	if o != nil && !IsNil(o.Subsidies) {
		return true
	}

	return false
}

// SetSubsidies gets a reference to the given []OrderItemSubsidyDTO and assigns it to the Subsidies field.
func (o *OrderDTO) SetSubsidies(v []OrderItemSubsidyDTO) {
	o.Subsidies = v
}

// GetDelivery returns the Delivery field value if set, zero value otherwise.
func (o *OrderDTO) GetDelivery() OrderDeliveryDTO {
	if o == nil || IsNil(o.Delivery) {
		var ret OrderDeliveryDTO
		return ret
	}
	return *o.Delivery
}

// GetDeliveryOk returns a tuple with the Delivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetDeliveryOk() (*OrderDeliveryDTO, bool) {
	if o == nil || IsNil(o.Delivery) {
		return nil, false
	}
	return o.Delivery, true
}

// HasDelivery returns a boolean if a field has been set.
func (o *OrderDTO) HasDelivery() bool {
	if o != nil && !IsNil(o.Delivery) {
		return true
	}

	return false
}

// SetDelivery gets a reference to the given OrderDeliveryDTO and assigns it to the Delivery field.
func (o *OrderDTO) SetDelivery(v OrderDeliveryDTO) {
	o.Delivery = &v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *OrderDTO) GetBuyer() OrderBuyerDTO {
	if o == nil || IsNil(o.Buyer) {
		var ret OrderBuyerDTO
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetBuyerOk() (*OrderBuyerDTO, bool) {
	if o == nil || IsNil(o.Buyer) {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *OrderDTO) HasBuyer() bool {
	if o != nil && !IsNil(o.Buyer) {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given OrderBuyerDTO and assigns it to the Buyer field.
func (o *OrderDTO) SetBuyer(v OrderBuyerDTO) {
	o.Buyer = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OrderDTO) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OrderDTO) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *OrderDTO) SetNotes(v string) {
	o.Notes = &v
}

// GetTaxSystem returns the TaxSystem field value if set, zero value otherwise.
func (o *OrderDTO) GetTaxSystem() OrderTaxSystemType {
	if o == nil || IsNil(o.TaxSystem) {
		var ret OrderTaxSystemType
		return ret
	}
	return *o.TaxSystem
}

// GetTaxSystemOk returns a tuple with the TaxSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetTaxSystemOk() (*OrderTaxSystemType, bool) {
	if o == nil || IsNil(o.TaxSystem) {
		return nil, false
	}
	return o.TaxSystem, true
}

// HasTaxSystem returns a boolean if a field has been set.
func (o *OrderDTO) HasTaxSystem() bool {
	if o != nil && !IsNil(o.TaxSystem) {
		return true
	}

	return false
}

// SetTaxSystem gets a reference to the given OrderTaxSystemType and assigns it to the TaxSystem field.
func (o *OrderDTO) SetTaxSystem(v OrderTaxSystemType) {
	o.TaxSystem = &v
}

// GetCancelRequested returns the CancelRequested field value if set, zero value otherwise.
func (o *OrderDTO) GetCancelRequested() bool {
	if o == nil || IsNil(o.CancelRequested) {
		var ret bool
		return ret
	}
	return *o.CancelRequested
}

// GetCancelRequestedOk returns a tuple with the CancelRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetCancelRequestedOk() (*bool, bool) {
	if o == nil || IsNil(o.CancelRequested) {
		return nil, false
	}
	return o.CancelRequested, true
}

// HasCancelRequested returns a boolean if a field has been set.
func (o *OrderDTO) HasCancelRequested() bool {
	if o != nil && !IsNil(o.CancelRequested) {
		return true
	}

	return false
}

// SetCancelRequested gets a reference to the given bool and assigns it to the CancelRequested field.
func (o *OrderDTO) SetCancelRequested(v bool) {
	o.CancelRequested = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *OrderDTO) GetExpiryDate() string {
	if o == nil || IsNil(o.ExpiryDate) {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDTO) GetExpiryDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *OrderDTO) HasExpiryDate() bool {
	if o != nil && !IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
func (o *OrderDTO) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

func (o OrderDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Substatus) {
		toSerialize["substatus"] = o.Substatus
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ItemsTotal) {
		toSerialize["itemsTotal"] = o.ItemsTotal
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.DeliveryTotal) {
		toSerialize["deliveryTotal"] = o.DeliveryTotal
	}
	if !IsNil(o.SubsidyTotal) {
		toSerialize["subsidyTotal"] = o.SubsidyTotal
	}
	if !IsNil(o.TotalWithSubsidy) {
		toSerialize["totalWithSubsidy"] = o.TotalWithSubsidy
	}
	if !IsNil(o.BuyerItemsTotal) {
		toSerialize["buyerItemsTotal"] = o.BuyerItemsTotal
	}
	if !IsNil(o.BuyerTotal) {
		toSerialize["buyerTotal"] = o.BuyerTotal
	}
	if !IsNil(o.BuyerItemsTotalBeforeDiscount) {
		toSerialize["buyerItemsTotalBeforeDiscount"] = o.BuyerItemsTotalBeforeDiscount
	}
	if !IsNil(o.BuyerTotalBeforeDiscount) {
		toSerialize["buyerTotalBeforeDiscount"] = o.BuyerTotalBeforeDiscount
	}
	if !IsNil(o.PaymentType) {
		toSerialize["paymentType"] = o.PaymentType
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.Fake) {
		toSerialize["fake"] = o.Fake
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Subsidies) {
		toSerialize["subsidies"] = o.Subsidies
	}
	if !IsNil(o.Delivery) {
		toSerialize["delivery"] = o.Delivery
	}
	if !IsNil(o.Buyer) {
		toSerialize["buyer"] = o.Buyer
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.TaxSystem) {
		toSerialize["taxSystem"] = o.TaxSystem
	}
	if !IsNil(o.CancelRequested) {
		toSerialize["cancelRequested"] = o.CancelRequested
	}
	if !IsNil(o.ExpiryDate) {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	return toSerialize, nil
}

type NullableOrderDTO struct {
	value *OrderDTO
	isSet bool
}

func (v NullableOrderDTO) Get() *OrderDTO {
	return v.value
}

func (v *NullableOrderDTO) Set(val *OrderDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDTO(val *OrderDTO) *NullableOrderDTO {
	return &NullableOrderDTO{value: val, isSet: true}
}

func (v NullableOrderDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


