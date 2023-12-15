# StockDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | **string** |   **Ваш SKU**  Идентификатор товара в магазине. Разрешены английские и русские буквы, цифры и символы &#x60;. , / \\ ( ) [ ] - &#x3D; _&#x60;  Максимальная длина — 80 знаков.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields).  | 
**WarehouseId** | **int64** | Идентификатор склада.  Узнать идентификатор склада вы можете в личном кабинете в разделе **Логистика → Склады**. Он указан в поле ID склада.  Если вы работаете с общими остатками, вы можете посмотреть идентификатор склада в личном кабинете в разделе **Настройки → Настройки API** в блоке **Обновление данных об остатках товаров** или с помощью запроса [GET businesses/{businessId}/warehouses](../../reference/warehouses/getWarehouses.md).  Если вы отвечаете на запрос Маркета, указывайте тот идентификатор, что пришел в запросе.  | 
**Items** | [**[]StockItemDTO**](StockItemDTO.md) | Информация об остатках товара на данном складе.  | 

## Methods

### NewStockDTO

`func NewStockDTO(sku string, warehouseId int64, items []StockItemDTO, ) *StockDTO`

NewStockDTO instantiates a new StockDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockDTOWithDefaults

`func NewStockDTOWithDefaults() *StockDTO`

NewStockDTOWithDefaults instantiates a new StockDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *StockDTO) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StockDTO) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *StockDTO) SetSku(v string)`

SetSku sets Sku field to given value.


### GetWarehouseId

`func (o *StockDTO) GetWarehouseId() int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *StockDTO) GetWarehouseIdOk() (*int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *StockDTO) SetWarehouseId(v int64)`

SetWarehouseId sets WarehouseId field to given value.


### GetItems

`func (o *StockDTO) GetItems() []StockItemDTO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *StockDTO) GetItemsOk() (*[]StockItemDTO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *StockDTO) SetItems(v []StockItemDTO)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


