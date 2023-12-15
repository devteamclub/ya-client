# OrderBuyerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Идентификатор покупателя. | [optional] 
**LastName** | Pointer to **string** | Фамилия покупателя. | [optional] 
**FirstName** | Pointer to **string** | Имя покупателя. | [optional] 
**MiddleName** | Pointer to **string** | Отчество покупателя. | [optional] 
**Phone** | Pointer to **string** | Номер телефона покупателя.  Формат номера: &#x60;+&lt;код_страны&gt;&lt;код_региона&gt;&lt;номер_телефона&gt;&#x60;.  | [optional] 
**Email** | Pointer to **string** | Адрес электронной почты покупателя.  Допускается любой адрес электронной почты, соответствующий стандарту RFC 2822.  | [optional] 
**Type** | Pointer to [**OrderBuyerType**](OrderBuyerType.md) |  | [optional] 

## Methods

### NewOrderBuyerDTO

`func NewOrderBuyerDTO() *OrderBuyerDTO`

NewOrderBuyerDTO instantiates a new OrderBuyerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderBuyerDTOWithDefaults

`func NewOrderBuyerDTOWithDefaults() *OrderBuyerDTO`

NewOrderBuyerDTOWithDefaults instantiates a new OrderBuyerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderBuyerDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderBuyerDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderBuyerDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderBuyerDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastName

`func (o *OrderBuyerDTO) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OrderBuyerDTO) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OrderBuyerDTO) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *OrderBuyerDTO) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetFirstName

`func (o *OrderBuyerDTO) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OrderBuyerDTO) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OrderBuyerDTO) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *OrderBuyerDTO) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *OrderBuyerDTO) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *OrderBuyerDTO) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *OrderBuyerDTO) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *OrderBuyerDTO) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhone

`func (o *OrderBuyerDTO) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrderBuyerDTO) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrderBuyerDTO) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *OrderBuyerDTO) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *OrderBuyerDTO) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderBuyerDTO) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderBuyerDTO) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderBuyerDTO) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetType

`func (o *OrderBuyerDTO) GetType() OrderBuyerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderBuyerDTO) GetTypeOk() (*OrderBuyerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderBuyerDTO) SetType(v OrderBuyerType)`

SetType sets Type field to given value.

### HasType

`func (o *OrderBuyerDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


