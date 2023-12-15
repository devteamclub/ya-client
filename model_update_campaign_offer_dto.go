/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ya-client

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateCampaignOfferDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCampaignOfferDTO{}

// UpdateCampaignOfferDTO Параметры размещения товара в магазине.
type UpdateCampaignOfferDTO struct {
	//   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы `. , / \\ ( ) [ ] - = _`  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields). 
	OfferId string `json:"offerId"`
	Quantum *QuantumDTO `json:"quantum,omitempty"`
	// Есть ли товар в продаже. 
	Available *bool `json:"available,omitempty"`
	// Ставка НДС, применяемая для товара. Задается цифрой:  * 2 — 10%. * 5 — 0%. * 6 — не облагается НДС. * 7 — 20%.  Если параметр не указан, используется ставка НДС, установленная в личном кабинете магазина. 
	Vat *int32 `json:"vat,omitempty"`
}

type _UpdateCampaignOfferDTO UpdateCampaignOfferDTO

// NewUpdateCampaignOfferDTO instantiates a new UpdateCampaignOfferDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCampaignOfferDTO(offerId string) *UpdateCampaignOfferDTO {
	this := UpdateCampaignOfferDTO{}
	this.OfferId = offerId
	return &this
}

// NewUpdateCampaignOfferDTOWithDefaults instantiates a new UpdateCampaignOfferDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCampaignOfferDTOWithDefaults() *UpdateCampaignOfferDTO {
	this := UpdateCampaignOfferDTO{}
	return &this
}

// GetOfferId returns the OfferId field value
func (o *UpdateCampaignOfferDTO) GetOfferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferId
}

// GetOfferIdOk returns a tuple with the OfferId field value
// and a boolean to check if the value has been set.
func (o *UpdateCampaignOfferDTO) GetOfferIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferId, true
}

// SetOfferId sets field value
func (o *UpdateCampaignOfferDTO) SetOfferId(v string) {
	o.OfferId = v
}

// GetQuantum returns the Quantum field value if set, zero value otherwise.
func (o *UpdateCampaignOfferDTO) GetQuantum() QuantumDTO {
	if o == nil || IsNil(o.Quantum) {
		var ret QuantumDTO
		return ret
	}
	return *o.Quantum
}

// GetQuantumOk returns a tuple with the Quantum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignOfferDTO) GetQuantumOk() (*QuantumDTO, bool) {
	if o == nil || IsNil(o.Quantum) {
		return nil, false
	}
	return o.Quantum, true
}

// HasQuantum returns a boolean if a field has been set.
func (o *UpdateCampaignOfferDTO) HasQuantum() bool {
	if o != nil && !IsNil(o.Quantum) {
		return true
	}

	return false
}

// SetQuantum gets a reference to the given QuantumDTO and assigns it to the Quantum field.
func (o *UpdateCampaignOfferDTO) SetQuantum(v QuantumDTO) {
	o.Quantum = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *UpdateCampaignOfferDTO) GetAvailable() bool {
	if o == nil || IsNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignOfferDTO) GetAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *UpdateCampaignOfferDTO) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *UpdateCampaignOfferDTO) SetAvailable(v bool) {
	o.Available = &v
}

// GetVat returns the Vat field value if set, zero value otherwise.
func (o *UpdateCampaignOfferDTO) GetVat() int32 {
	if o == nil || IsNil(o.Vat) {
		var ret int32
		return ret
	}
	return *o.Vat
}

// GetVatOk returns a tuple with the Vat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignOfferDTO) GetVatOk() (*int32, bool) {
	if o == nil || IsNil(o.Vat) {
		return nil, false
	}
	return o.Vat, true
}

// HasVat returns a boolean if a field has been set.
func (o *UpdateCampaignOfferDTO) HasVat() bool {
	if o != nil && !IsNil(o.Vat) {
		return true
	}

	return false
}

// SetVat gets a reference to the given int32 and assigns it to the Vat field.
func (o *UpdateCampaignOfferDTO) SetVat(v int32) {
	o.Vat = &v
}

func (o UpdateCampaignOfferDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCampaignOfferDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offerId"] = o.OfferId
	if !IsNil(o.Quantum) {
		toSerialize["quantum"] = o.Quantum
	}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.Vat) {
		toSerialize["vat"] = o.Vat
	}
	return toSerialize, nil
}

func (o *UpdateCampaignOfferDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"offerId",
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

	varUpdateCampaignOfferDTO := _UpdateCampaignOfferDTO{}

	err = json.Unmarshal(bytes, &varUpdateCampaignOfferDTO)

	if err != nil {
		return err
	}

	*o = UpdateCampaignOfferDTO(varUpdateCampaignOfferDTO)

	return err
}

type NullableUpdateCampaignOfferDTO struct {
	value *UpdateCampaignOfferDTO
	isSet bool
}

func (v NullableUpdateCampaignOfferDTO) Get() *UpdateCampaignOfferDTO {
	return v.value
}

func (v *NullableUpdateCampaignOfferDTO) Set(val *UpdateCampaignOfferDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCampaignOfferDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCampaignOfferDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCampaignOfferDTO(val *UpdateCampaignOfferDTO) *NullableUpdateCampaignOfferDTO {
	return &NullableUpdateCampaignOfferDTO{value: val, isSet: true}
}

func (v NullableUpdateCampaignOfferDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCampaignOfferDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


