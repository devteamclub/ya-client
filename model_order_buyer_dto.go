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

// checks if the OrderBuyerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderBuyerDTO{}

// OrderBuyerDTO Информация о покупателе.
type OrderBuyerDTO struct {
	// Идентификатор покупателя.
	Id *string `json:"id,omitempty"`
	// Фамилия покупателя.
	LastName *string `json:"lastName,omitempty"`
	// Имя покупателя.
	FirstName *string `json:"firstName,omitempty"`
	// Отчество покупателя.
	MiddleName *string `json:"middleName,omitempty"`
	// Номер телефона покупателя.  Формат номера: `+<код_страны><код_региона><номер_телефона>`. 
	Phone *string `json:"phone,omitempty"`
	// Адрес электронной почты покупателя.  Допускается любой адрес электронной почты, соответствующий стандарту RFC 2822. 
	Email *string `json:"email,omitempty"`
	Type *OrderBuyerType `json:"type,omitempty"`
}

// NewOrderBuyerDTO instantiates a new OrderBuyerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderBuyerDTO() *OrderBuyerDTO {
	this := OrderBuyerDTO{}
	return &this
}

// NewOrderBuyerDTOWithDefaults instantiates a new OrderBuyerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderBuyerDTOWithDefaults() *OrderBuyerDTO {
	this := OrderBuyerDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderBuyerDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBuyerDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderBuyerDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderBuyerDTO) SetId(v string) {
	o.Id = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *OrderBuyerDTO) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBuyerDTO) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *OrderBuyerDTO) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *OrderBuyerDTO) SetLastName(v string) {
	o.LastName = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *OrderBuyerDTO) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBuyerDTO) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *OrderBuyerDTO) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *OrderBuyerDTO) SetFirstName(v string) {
	o.FirstName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *OrderBuyerDTO) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBuyerDTO) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *OrderBuyerDTO) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *OrderBuyerDTO) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *OrderBuyerDTO) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBuyerDTO) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *OrderBuyerDTO) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *OrderBuyerDTO) SetPhone(v string) {
	o.Phone = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrderBuyerDTO) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBuyerDTO) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrderBuyerDTO) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrderBuyerDTO) SetEmail(v string) {
	o.Email = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderBuyerDTO) GetType() OrderBuyerType {
	if o == nil || IsNil(o.Type) {
		var ret OrderBuyerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBuyerDTO) GetTypeOk() (*OrderBuyerType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderBuyerDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given OrderBuyerType and assigns it to the Type field.
func (o *OrderBuyerDTO) SetType(v OrderBuyerType) {
	o.Type = &v
}

func (o OrderBuyerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderBuyerDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middleName"] = o.MiddleName
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableOrderBuyerDTO struct {
	value *OrderBuyerDTO
	isSet bool
}

func (v NullableOrderBuyerDTO) Get() *OrderBuyerDTO {
	return v.value
}

func (v *NullableOrderBuyerDTO) Set(val *OrderBuyerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderBuyerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderBuyerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderBuyerDTO(val *OrderBuyerDTO) *NullableOrderBuyerDTO {
	return &NullableOrderBuyerDTO{value: val, isSet: true}
}

func (v NullableOrderBuyerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderBuyerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

