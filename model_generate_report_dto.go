/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the GenerateReportDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateReportDTO{}

// GenerateReportDTO Идентификатор, который понадобится для отслеживания статуса генерации и получения готового отчета.
type GenerateReportDTO struct {
	// Идентификатор, который понадобится для отслеживания статуса генерации и получения готового отчета.
	ReportId string `json:"reportId"`
	// Ожидаемая продолжительность генерации в миллисекундах.
	EstimatedGenerationTime int64 `json:"estimatedGenerationTime"`
}

type _GenerateReportDTO GenerateReportDTO

// NewGenerateReportDTO instantiates a new GenerateReportDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateReportDTO(reportId string, estimatedGenerationTime int64) *GenerateReportDTO {
	this := GenerateReportDTO{}
	this.ReportId = reportId
	this.EstimatedGenerationTime = estimatedGenerationTime
	return &this
}

// NewGenerateReportDTOWithDefaults instantiates a new GenerateReportDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateReportDTOWithDefaults() *GenerateReportDTO {
	this := GenerateReportDTO{}
	return &this
}

// GetReportId returns the ReportId field value
func (o *GenerateReportDTO) GetReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value
// and a boolean to check if the value has been set.
func (o *GenerateReportDTO) GetReportIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportId, true
}

// SetReportId sets field value
func (o *GenerateReportDTO) SetReportId(v string) {
	o.ReportId = v
}

// GetEstimatedGenerationTime returns the EstimatedGenerationTime field value
func (o *GenerateReportDTO) GetEstimatedGenerationTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EstimatedGenerationTime
}

// GetEstimatedGenerationTimeOk returns a tuple with the EstimatedGenerationTime field value
// and a boolean to check if the value has been set.
func (o *GenerateReportDTO) GetEstimatedGenerationTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedGenerationTime, true
}

// SetEstimatedGenerationTime sets field value
func (o *GenerateReportDTO) SetEstimatedGenerationTime(v int64) {
	o.EstimatedGenerationTime = v
}

func (o GenerateReportDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateReportDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reportId"] = o.ReportId
	toSerialize["estimatedGenerationTime"] = o.EstimatedGenerationTime
	return toSerialize, nil
}

func (o *GenerateReportDTO) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reportId",
		"estimatedGenerationTime",
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

	varGenerateReportDTO := _GenerateReportDTO{}

	err = json.Unmarshal(bytes, &varGenerateReportDTO)

	if err != nil {
		return err
	}

	*o = GenerateReportDTO(varGenerateReportDTO)

	return err
}

type NullableGenerateReportDTO struct {
	value *GenerateReportDTO
	isSet bool
}

func (v NullableGenerateReportDTO) Get() *GenerateReportDTO {
	return v.value
}

func (v *NullableGenerateReportDTO) Set(val *GenerateReportDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateReportDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateReportDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateReportDTO(val *GenerateReportDTO) *NullableGenerateReportDTO {
	return &NullableGenerateReportDTO{value: val, isSet: true}
}

func (v NullableGenerateReportDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateReportDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


