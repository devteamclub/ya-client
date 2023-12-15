# \FbsAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHiddenOffers**](FbsAPI.md#AddHiddenOffers) | **Post** /campaigns/{campaignId}/hidden-offers | Скрытие товаров и настройки скрытия
[**AddOffersToArchive**](FbsAPI.md#AddOffersToArchive) | **Post** /businesses/{businessId}/offer-mappings/archive | Добавление товаров в архив
[**ConfirmBusinessPrices**](FbsAPI.md#ConfirmBusinessPrices) | **Post** /businesses/{businessId}/price-quarantine/confirm | Подтверждение цены товара, попавшего в карантин (основная цена)
[**ConfirmCampaignPrices**](FbsAPI.md#ConfirmCampaignPrices) | **Post** /campaigns/{campaignId}/price-quarantine/confirm | Подтверждение цены товара, попавшего в карантин (цена в магазине)
[**ConfirmShipment**](FbsAPI.md#ConfirmShipment) | **Post** /campaigns/{campaignId}/first-mile/shipments/{shipmentId}/confirm | Подтверждение отгрузки
[**DeleteCampaignOffers**](FbsAPI.md#DeleteCampaignOffers) | **Post** /campaigns/{campaignId}/offers/delete | Удаление товаров из ассортимента магазина
[**DeleteHiddenOffers**](FbsAPI.md#DeleteHiddenOffers) | **Delete** /campaigns/{campaignId}/hidden-offers | Возобновление показа товаров
[**DeleteOffers**](FbsAPI.md#DeleteOffers) | **Post** /businesses/{businessId}/offer-mappings/delete | Удаление товаров из каталога
[**DeleteOffersFromArchive**](FbsAPI.md#DeleteOffersFromArchive) | **Post** /businesses/{businessId}/offer-mappings/unarchive | Восстановление товаров из архива
[**DownloadShipmentAct**](FbsAPI.md#DownloadShipmentAct) | **Get** /campaigns/{campaignId}/first-mile/shipments/{shipmentId}/act | Получение акта приема-передачи
[**DownloadShipmentDiscrepancyAct**](FbsAPI.md#DownloadShipmentDiscrepancyAct) | **Get** /campaigns/{campaignId}/first-mile/shipments/{shipmentId}/discrepancy-act | Получение акта расхождений
[**DownloadShipmentInboundAct**](FbsAPI.md#DownloadShipmentInboundAct) | **Get** /campaigns/{campaignId}/first-mile/shipments/{shipmentId}/inbound-act | Получение фактического акта приема-передачи
[**DownloadShipmentPalletLabels**](FbsAPI.md#DownloadShipmentPalletLabels) | **Get** /campaigns/{campaignId}/first-mile/shipments/{shipmentId}/pallet/labels | Ярлыки на все упаковки в отгрузке
[**DownloadShipmentReceptionTransferAct**](FbsAPI.md#DownloadShipmentReceptionTransferAct) | **Get** /campaigns/{campaignId}/shipments/reception-transfer-act | Подтверждение ближайшей отгрузки и получение акта приема-передачи для нее
[**DownloadShipmentTransportationWaybill**](FbsAPI.md#DownloadShipmentTransportationWaybill) | **Get** /campaigns/{campaignId}/first-mile/shipments/{shipmentId}/transportation-waybill | Получение транспортной накладной
[**GenerateGoodsRealizationReport**](FbsAPI.md#GenerateGoodsRealizationReport) | **Post** /reports/goods-realization/generate | Отчет по реализации
[**GenerateOrderLabel**](FbsAPI.md#GenerateOrderLabel) | **Get** /campaigns/{campaignId}/orders/{orderId}/delivery/shipments/{shipmentId}/boxes/{boxId}/label | Ярлык‑наклейка на отдельное грузовое место в заказе
[**GenerateOrderLabels**](FbsAPI.md#GenerateOrderLabels) | **Get** /campaigns/{campaignId}/orders/{orderId}/delivery/labels | Ярлыки‑наклейки на все грузовые места в заказе
[**GeneratePricesReport**](FbsAPI.md#GeneratePricesReport) | **Post** /reports/prices/generate | Отчет «Цены на рынке»
[**GenerateShowsSalesReport**](FbsAPI.md#GenerateShowsSalesReport) | **Post** /reports/shows-sales/generate | Отчет «Аналитика продаж»
[**GenerateStocksOnWarehousesReport**](FbsAPI.md#GenerateStocksOnWarehousesReport) | **Post** /reports/stocks-on-warehouses/generate | Отчет по остаткам на складах
[**GenerateUnitedMarketplaceServicesReport**](FbsAPI.md#GenerateUnitedMarketplaceServicesReport) | **Post** /reports/united-marketplace-services/generate | Отчет по стоимости услуг
[**GenerateUnitedNettingReport**](FbsAPI.md#GenerateUnitedNettingReport) | **Post** /reports/united-netting/generate | Отчет по платежам
[**GetBidsInfoForBusiness**](FbsAPI.md#GetBidsInfoForBusiness) | **Post** /businesses/{businessId}/bids/info | Информация об установленных ставках
[**GetBidsRecommendations**](FbsAPI.md#GetBidsRecommendations) | **Post** /businesses/{businessId}/bids/recommendations | Рекомендованные ставки для заданных товаров
[**GetBusinessQuarantineOffers**](FbsAPI.md#GetBusinessQuarantineOffers) | **Post** /businesses/{businessId}/price-quarantine | Список товаров, находящихся в карантине (основная цена)
[**GetCampaign**](FbsAPI.md#GetCampaign) | **Get** /campaigns/{campaignId} | Информация о магазине
[**GetCampaignLogins**](FbsAPI.md#GetCampaignLogins) | **Get** /campaigns/{campaignId}/logins | Логины, связанные с магазином
[**GetCampaignOffers**](FbsAPI.md#GetCampaignOffers) | **Post** /campaigns/{campaignId}/offers | Список товаров, размещенных в заданном магазине, с параметрами размещения
[**GetCampaignQuarantineOffers**](FbsAPI.md#GetCampaignQuarantineOffers) | **Post** /campaigns/{campaignId}/price-quarantine | Список товаров, находящихся в карантине (цена в магазине)
[**GetCampaigns**](FbsAPI.md#GetCampaigns) | **Get** /campaigns | Магазины пользователя
[**GetCampaignsByLogin**](FbsAPI.md#GetCampaignsByLogin) | **Get** /campaigns/by_login/{login} | Магазины, доступные логину
[**GetCategoryContentParameters**](FbsAPI.md#GetCategoryContentParameters) | **Post** /category/{categoryId}/parameters | Списки характеристик товаров по категориям
[**GetDeliveryServices**](FbsAPI.md#GetDeliveryServices) | **Get** /delivery/services | Справочник служб доставки
[**GetGoodsStats**](FbsAPI.md#GetGoodsStats) | **Post** /campaigns/{campaignId}/stats/skus | Отчет по товарам
[**GetHiddenOffers**](FbsAPI.md#GetHiddenOffers) | **Get** /campaigns/{campaignId}/hidden-offers | Информация о скрытых вами товарах
[**GetOfferCardsContentStatus**](FbsAPI.md#GetOfferCardsContentStatus) | **Post** /businesses/{businessId}/offer-cards | Получение информации о заполненности карточек
[**GetOfferMappingEntries**](FbsAPI.md#GetOfferMappingEntries) | **Get** /campaigns/{campaignId}/offer-mapping-entries | Список товаров в каталоге
[**GetOfferMappings**](FbsAPI.md#GetOfferMappings) | **Post** /businesses/{businessId}/offer-mappings | Информация о товарах в каталоге
[**GetOfferRecommendations**](FbsAPI.md#GetOfferRecommendations) | **Post** /businesses/{businessId}/offers/recommendations | Рекомендации Маркета, касающиеся цен
[**GetOrder**](FbsAPI.md#GetOrder) | **Get** /campaigns/{campaignId}/orders/{orderId} | Информация о заказе
[**GetOrderLabelsData**](FbsAPI.md#GetOrderLabelsData) | **Get** /campaigns/{campaignId}/orders/{orderId}/delivery/labels/data | Информация на ярлыках‑наклейках
[**GetOrders**](FbsAPI.md#GetOrders) | **Get** /campaigns/{campaignId}/orders | Информация о заказах
[**GetOrdersStats**](FbsAPI.md#GetOrdersStats) | **Post** /campaigns/{campaignId}/stats/orders | Отчет по заказам
[**GetPrices**](FbsAPI.md#GetPrices) | **Get** /campaigns/{campaignId}/offer-prices | Список цен
[**GetPricesByOfferIds**](FbsAPI.md#GetPricesByOfferIds) | **Post** /campaigns/{campaignId}/offer-prices | Просмотр установленных в магазине цен
[**GetReportInfo**](FbsAPI.md#GetReportInfo) | **Get** /reports/info/{reportId} | Статус генерации и скачивание готовых отчетов
[**GetReturn**](FbsAPI.md#GetReturn) | **Get** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId} | Информация о возврате или невыкупе
[**GetReturnApplication**](FbsAPI.md#GetReturnApplication) | **Get** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/application | Получение заявления на возврат
[**GetReturnPhoto**](FbsAPI.md#GetReturnPhoto) | **Get** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/decision/{itemId}/image/{imageHash} | Получение фотографии возврата
[**GetReturns**](FbsAPI.md#GetReturns) | **Get** /campaigns/{campaignId}/returns | Список возвратов и невыкупов
[**GetShipment**](FbsAPI.md#GetShipment) | **Get** /campaigns/{campaignId}/first-mile/shipments/{shipmentId} | Получение информации об отгрузке
[**GetShipmentOrdersInfo**](FbsAPI.md#GetShipmentOrdersInfo) | **Get** /campaigns/{campaignId}/first-mile/shipments/{shipmentId}/orders/info | Получение информации о ярлыках
[**GetStocks**](FbsAPI.md#GetStocks) | **Post** /campaigns/{campaignId}/offers/stocks | Информация об остатках и оборачиваемости
[**GetSuggestedOfferMappingEntries**](FbsAPI.md#GetSuggestedOfferMappingEntries) | **Post** /campaigns/{campaignId}/offer-mapping-entries/suggestions | Рекомендованные карточки для ваших товаров
[**GetSuggestedOfferMappings**](FbsAPI.md#GetSuggestedOfferMappings) | **Post** /businesses/{businessId}/offer-mappings/suggestions | Предварительный просмотр карточек на Маркете, соответствующих вашим товарам
[**GetSuggestedPrices**](FbsAPI.md#GetSuggestedPrices) | **Post** /campaigns/{campaignId}/offer-prices/suggestions | Цены для продвижения товаров
[**GetWarehouses**](FbsAPI.md#GetWarehouses) | **Get** /businesses/{businessId}/warehouses | Список складов и групп складов
[**ProvideOrderItemIdentifiers**](FbsAPI.md#ProvideOrderItemIdentifiers) | **Put** /campaigns/{campaignId}/orders/{orderId}/identifiers | Передача уникальных кодов маркировки единиц товара
[**PutBidsForBusiness**](FbsAPI.md#PutBidsForBusiness) | **Put** /businesses/{businessId}/bids | Включение буста продаж и установка ставок
[**SearchRegionChildren**](FbsAPI.md#SearchRegionChildren) | **Get** /regions/{regionId}/children | Информация о дочерних регионах
[**SearchRegionsById**](FbsAPI.md#SearchRegionsById) | **Get** /regions/{regionId} | Информация о регионе
[**SearchRegionsByName**](FbsAPI.md#SearchRegionsByName) | **Get** /regions | Метод для поиска регионов по их имени
[**SearchShipments**](FbsAPI.md#SearchShipments) | **Put** /campaigns/{campaignId}/first-mile/shipments | Получение информации об отгрузках
[**SetOrderShipmentBoxes**](FbsAPI.md#SetOrderShipmentBoxes) | **Put** /campaigns/{campaignId}/orders/{orderId}/delivery/shipments/{shipmentId}/boxes | Передача количества грузовых мест в заказе
[**SetShipmentPalletsCount**](FbsAPI.md#SetShipmentPalletsCount) | **Put** /campaigns/{campaignId}/first-mile/shipments/{shipmentId}/pallets | Передача количества упаковок в отгрузке
[**TransferOrdersFromShipment**](FbsAPI.md#TransferOrdersFromShipment) | **Post** /campaigns/{campaignId}/first-mile/shipments/{shipmentId}/orders/transfer | Перенос заказов в следующую отгрузку
[**UpdateBusinessPrices**](FbsAPI.md#UpdateBusinessPrices) | **Post** /businesses/{businessId}/offer-prices/updates | Установка цен
[**UpdateCampaignOffers**](FbsAPI.md#UpdateCampaignOffers) | **Post** /campaigns/{campaignId}/offers/update | Настройка размещения товаров в магазине
[**UpdateOfferContent**](FbsAPI.md#UpdateOfferContent) | **Post** /businesses/{businessId}/offer-cards/update | Редактирование категорийных характеристик товара
[**UpdateOfferMappingEntries**](FbsAPI.md#UpdateOfferMappingEntries) | **Post** /campaigns/{campaignId}/offer-mapping-entries/updates | Добавление и редактирование товаров в каталоге
[**UpdateOfferMappings**](FbsAPI.md#UpdateOfferMappings) | **Post** /businesses/{businessId}/offer-mappings/update | Добавление товаров в каталог и редактирование информации о них
[**UpdateOrderItems**](FbsAPI.md#UpdateOrderItems) | **Put** /campaigns/{campaignId}/orders/{orderId}/items | Удаление товара из заказа или уменьшение числа единиц
[**UpdateOrderStatus**](FbsAPI.md#UpdateOrderStatus) | **Put** /campaigns/{campaignId}/orders/{orderId}/status | Изменение статуса заказа
[**UpdateOrderStatuses**](FbsAPI.md#UpdateOrderStatuses) | **Post** /campaigns/{campaignId}/orders/status-update | Изменение статусов заказа
[**UpdatePrices**](FbsAPI.md#UpdatePrices) | **Post** /campaigns/{campaignId}/offer-prices/updates | Установка цен на товары в конкретном магазине
[**UpdateStocks**](FbsAPI.md#UpdateStocks) | **Put** /campaigns/{campaignId}/offers/stocks | Передача информации об остатках



## AddHiddenOffers

> EmptyApiResponse AddHiddenOffers(ctx, campaignId).AddHiddenOffersRequest(addHiddenOffersRequest).Execute()

Скрытие товаров и настройки скрытия



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    addHiddenOffersRequest := *openapiclient.NewAddHiddenOffersRequest([]openapiclient.HiddenOfferDTO{*openapiclient.NewHiddenOfferDTO()}) // AddHiddenOffersRequest | Запрос на скрытие оферов.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.AddHiddenOffers(context.Background(), campaignId).AddHiddenOffersRequest(addHiddenOffersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.AddHiddenOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddHiddenOffers`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.AddHiddenOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddHiddenOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addHiddenOffersRequest** | [**AddHiddenOffersRequest**](AddHiddenOffersRequest.md) | Запрос на скрытие оферов. | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddOffersToArchive

> AddOffersToArchiveResponse AddOffersToArchive(ctx, businessId).AddOffersToArchiveRequest(addOffersToArchiveRequest).Execute()

Добавление товаров в архив



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    addOffersToArchiveRequest := *openapiclient.NewAddOffersToArchiveRequest([]string{"OfferIds_example"}) // AddOffersToArchiveRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.AddOffersToArchive(context.Background(), businessId).AddOffersToArchiveRequest(addOffersToArchiveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.AddOffersToArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOffersToArchive`: AddOffersToArchiveResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.AddOffersToArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOffersToArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addOffersToArchiveRequest** | [**AddOffersToArchiveRequest**](AddOffersToArchiveRequest.md) |  | 

### Return type

[**AddOffersToArchiveResponse**](AddOffersToArchiveResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmBusinessPrices

> EmptyApiResponse ConfirmBusinessPrices(ctx, businessId).ConfirmPricesRequest(confirmPricesRequest).Execute()

Подтверждение цены товара, попавшего в карантин (основная цена)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    confirmPricesRequest := *openapiclient.NewConfirmPricesRequest([]string{"OfferIds_example"}) // ConfirmPricesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.ConfirmBusinessPrices(context.Background(), businessId).ConfirmPricesRequest(confirmPricesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.ConfirmBusinessPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmBusinessPrices`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.ConfirmBusinessPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmBusinessPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confirmPricesRequest** | [**ConfirmPricesRequest**](ConfirmPricesRequest.md) |  | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmCampaignPrices

> EmptyApiResponse ConfirmCampaignPrices(ctx, campaignId).ConfirmPricesRequest(confirmPricesRequest).Execute()

Подтверждение цены товара, попавшего в карантин (цена в магазине)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    confirmPricesRequest := *openapiclient.NewConfirmPricesRequest([]string{"OfferIds_example"}) // ConfirmPricesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.ConfirmCampaignPrices(context.Background(), campaignId).ConfirmPricesRequest(confirmPricesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.ConfirmCampaignPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmCampaignPrices`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.ConfirmCampaignPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmCampaignPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confirmPricesRequest** | [**ConfirmPricesRequest**](ConfirmPricesRequest.md) |  | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmShipment

> EmptyApiResponse ConfirmShipment(ctx, campaignId, shipmentId).ConfirmShipmentRequest(confirmShipmentRequest).Execute()

Подтверждение отгрузки



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    shipmentId := int64(789) // int64 | Идентификатор отгрузки.
    confirmShipmentRequest := *openapiclient.NewConfirmShipmentRequest([]int64{int64(123)}) // ConfirmShipmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.ConfirmShipment(context.Background(), campaignId, shipmentId).ConfirmShipmentRequest(confirmShipmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.ConfirmShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmShipment`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.ConfirmShipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**shipmentId** | **int64** | Идентификатор отгрузки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **confirmShipmentRequest** | [**ConfirmShipmentRequest**](ConfirmShipmentRequest.md) |  | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaignOffers

> DeleteCampaignOffersResponse DeleteCampaignOffers(ctx, campaignId).DeleteCampaignOffersRequest(deleteCampaignOffersRequest).Execute()

Удаление товаров из ассортимента магазина



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    deleteCampaignOffersRequest := *openapiclient.NewDeleteCampaignOffersRequest([]string{"OfferIds_example"}) // DeleteCampaignOffersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.DeleteCampaignOffers(context.Background(), campaignId).DeleteCampaignOffersRequest(deleteCampaignOffersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.DeleteCampaignOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCampaignOffers`: DeleteCampaignOffersResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.DeleteCampaignOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteCampaignOffersRequest** | [**DeleteCampaignOffersRequest**](DeleteCampaignOffersRequest.md) |  | 

### Return type

[**DeleteCampaignOffersResponse**](DeleteCampaignOffersResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHiddenOffers

> EmptyApiResponse DeleteHiddenOffers(ctx, campaignId).DeleteHiddenOffersRequest(deleteHiddenOffersRequest).Execute()

Возобновление показа товаров



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    deleteHiddenOffersRequest := *openapiclient.NewDeleteHiddenOffersRequest([]openapiclient.HiddenOfferDTO{*openapiclient.NewHiddenOfferDTO()}) // DeleteHiddenOffersRequest | Запрос на возобновление показа оферов.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.DeleteHiddenOffers(context.Background(), campaignId).DeleteHiddenOffersRequest(deleteHiddenOffersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.DeleteHiddenOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHiddenOffers`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.DeleteHiddenOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHiddenOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteHiddenOffersRequest** | [**DeleteHiddenOffersRequest**](DeleteHiddenOffersRequest.md) | Запрос на возобновление показа оферов. | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOffers

> DeleteOffersResponse DeleteOffers(ctx, businessId).DeleteOffersRequest(deleteOffersRequest).Execute()

Удаление товаров из каталога



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    deleteOffersRequest := *openapiclient.NewDeleteOffersRequest([]string{"OfferIds_example"}) // DeleteOffersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.DeleteOffers(context.Background(), businessId).DeleteOffersRequest(deleteOffersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.DeleteOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOffers`: DeleteOffersResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.DeleteOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteOffersRequest** | [**DeleteOffersRequest**](DeleteOffersRequest.md) |  | 

### Return type

[**DeleteOffersResponse**](DeleteOffersResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOffersFromArchive

> DeleteOffersFromArchiveResponse DeleteOffersFromArchive(ctx, businessId).DeleteOffersFromArchiveRequest(deleteOffersFromArchiveRequest).Execute()

Восстановление товаров из архива



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    deleteOffersFromArchiveRequest := *openapiclient.NewDeleteOffersFromArchiveRequest([]string{"OfferIds_example"}) // DeleteOffersFromArchiveRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.DeleteOffersFromArchive(context.Background(), businessId).DeleteOffersFromArchiveRequest(deleteOffersFromArchiveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.DeleteOffersFromArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOffersFromArchive`: DeleteOffersFromArchiveResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.DeleteOffersFromArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOffersFromArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteOffersFromArchiveRequest** | [**DeleteOffersFromArchiveRequest**](DeleteOffersFromArchiveRequest.md) |  | 

### Return type

[**DeleteOffersFromArchiveResponse**](DeleteOffersFromArchiveResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadShipmentAct

> *os.File DownloadShipmentAct(ctx, campaignId, shipmentId).Execute()

Получение акта приема-передачи



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    shipmentId := int64(789) // int64 | Идентификатор отгрузки.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.DownloadShipmentAct(context.Background(), campaignId, shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.DownloadShipmentAct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadShipmentAct`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.DownloadShipmentAct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**shipmentId** | **int64** | Идентификатор отгрузки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadShipmentActRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadShipmentDiscrepancyAct

> *os.File DownloadShipmentDiscrepancyAct(ctx, campaignId, shipmentId).Execute()

Получение акта расхождений



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    shipmentId := int64(789) // int64 | Идентификатор отгрузки.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.DownloadShipmentDiscrepancyAct(context.Background(), campaignId, shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.DownloadShipmentDiscrepancyAct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadShipmentDiscrepancyAct`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.DownloadShipmentDiscrepancyAct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**shipmentId** | **int64** | Идентификатор отгрузки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadShipmentDiscrepancyActRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ms-excel, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadShipmentInboundAct

> *os.File DownloadShipmentInboundAct(ctx, campaignId, shipmentId).Execute()

Получение фактического акта приема-передачи



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    shipmentId := int64(789) // int64 | Идентификатор отгрузки.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.DownloadShipmentInboundAct(context.Background(), campaignId, shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.DownloadShipmentInboundAct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadShipmentInboundAct`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.DownloadShipmentInboundAct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**shipmentId** | **int64** | Идентификатор отгрузки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadShipmentInboundActRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadShipmentPalletLabels

> *os.File DownloadShipmentPalletLabels(ctx, campaignId, shipmentId).Format(format).Execute()

Ярлыки на все упаковки в отгрузке



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    shipmentId := int64(789) // int64 | Идентификатор отгрузки.
    format := openapiclient.ShipmentPalletLabelPageFormatType("A4") // ShipmentPalletLabelPageFormatType | Формат страниц PDF-файла с ярлыками: * `A4` — по 16 ярлыков на странице. * `A8` — по одному ярлыку на странице.  (optional) (default to "A8")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.DownloadShipmentPalletLabels(context.Background(), campaignId, shipmentId).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.DownloadShipmentPalletLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadShipmentPalletLabels`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.DownloadShipmentPalletLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**shipmentId** | **int64** | Идентификатор отгрузки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadShipmentPalletLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | [**ShipmentPalletLabelPageFormatType**](ShipmentPalletLabelPageFormatType.md) | Формат страниц PDF-файла с ярлыками: * &#x60;A4&#x60; — по 16 ярлыков на странице. * &#x60;A8&#x60; — по одному ярлыку на странице.  | [default to &quot;A8&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadShipmentReceptionTransferAct

> *os.File DownloadShipmentReceptionTransferAct(ctx, campaignId).WarehouseId(warehouseId).Execute()

Подтверждение ближайшей отгрузки и получение акта приема-передачи для нее



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    warehouseId := int32(123123) // int32 | Идентификатор склада (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.DownloadShipmentReceptionTransferAct(context.Background(), campaignId).WarehouseId(warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.DownloadShipmentReceptionTransferAct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadShipmentReceptionTransferAct`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.DownloadShipmentReceptionTransferAct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadShipmentReceptionTransferActRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **warehouseId** | **int32** | Идентификатор склада | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadShipmentTransportationWaybill

> *os.File DownloadShipmentTransportationWaybill(ctx, campaignId, shipmentId).Execute()

Получение транспортной накладной



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    shipmentId := int64(789) // int64 | Идентификатор отгрузки.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.DownloadShipmentTransportationWaybill(context.Background(), campaignId, shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.DownloadShipmentTransportationWaybill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadShipmentTransportationWaybill`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.DownloadShipmentTransportationWaybill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**shipmentId** | **int64** | Идентификатор отгрузки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadShipmentTransportationWaybillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ms-excel, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateGoodsRealizationReport

> GenerateReportResponse GenerateGoodsRealizationReport(ctx).GenerateGoodsRealizationReportRequest(generateGoodsRealizationReportRequest).Format(format).Execute()

Отчет по реализации



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    generateGoodsRealizationReportRequest := *openapiclient.NewGenerateGoodsRealizationReportRequest(int64(123), int32(123), int32(123)) // GenerateGoodsRealizationReportRequest | 
    format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. Пока отчеты доступны только в одном формате — они предоставляются в виде электронной таблицы. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GenerateGoodsRealizationReport(context.Background()).GenerateGoodsRealizationReportRequest(generateGoodsRealizationReportRequest).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GenerateGoodsRealizationReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateGoodsRealizationReport`: GenerateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GenerateGoodsRealizationReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateGoodsRealizationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateGoodsRealizationReportRequest** | [**GenerateGoodsRealizationReportRequest**](GenerateGoodsRealizationReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. Пока отчеты доступны только в одном формате — они предоставляются в виде электронной таблицы. | 

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateOrderLabel

> *os.File GenerateOrderLabel(ctx, campaignId, orderId, shipmentId, boxId).Format(format).Execute()

Ярлык‑наклейка на отдельное грузовое место в заказе



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    orderId := int64(789) // int64 | Идентификатор заказа.
    shipmentId := int64(789) // int64 | Идентификатор грузоместа
    boxId := int64(789) // int64 | Идентификатор коробки
    format := openapiclient.PageFormatType("A7") // PageFormatType | Настройка размещения ярлыков на странице. Если параметра нет, возвращается PDF с ярлыками формата A6. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GenerateOrderLabel(context.Background(), campaignId, orderId, shipmentId, boxId).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GenerateOrderLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateOrderLabel`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GenerateOrderLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**orderId** | **int64** | Идентификатор заказа. | 
**shipmentId** | **int64** | Идентификатор грузоместа | 
**boxId** | **int64** | Идентификатор коробки | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateOrderLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **format** | [**PageFormatType**](PageFormatType.md) | Настройка размещения ярлыков на странице. Если параметра нет, возвращается PDF с ярлыками формата A6. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateOrderLabels

> *os.File GenerateOrderLabels(ctx, campaignId, orderId).Format(format).Execute()

Ярлыки‑наклейки на все грузовые места в заказе



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    orderId := int64(789) // int64 | Идентификатор заказа.
    format := openapiclient.PageFormatType("A7") // PageFormatType | Настройка размещения ярлыков на странице. Если параметра нет, возвращается PDF с ярлыками формата A6. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GenerateOrderLabels(context.Background(), campaignId, orderId).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GenerateOrderLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateOrderLabels`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GenerateOrderLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateOrderLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | [**PageFormatType**](PageFormatType.md) | Настройка размещения ярлыков на странице. Если параметра нет, возвращается PDF с ярлыками формата A6. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneratePricesReport

> GenerateReportResponse GeneratePricesReport(ctx).GeneratePricesReportRequest(generatePricesReportRequest).Format(format).Execute()

Отчет «Цены на рынке»



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    generatePricesReportRequest := *openapiclient.NewGeneratePricesReportRequest() // GeneratePricesReportRequest | 
    format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. Пока отчеты доступны только в одном формате — они предоставляются в виде электронной таблицы. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GeneratePricesReport(context.Background()).GeneratePricesReportRequest(generatePricesReportRequest).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GeneratePricesReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneratePricesReport`: GenerateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GeneratePricesReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePricesReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generatePricesReportRequest** | [**GeneratePricesReportRequest**](GeneratePricesReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. Пока отчеты доступны только в одном формате — они предоставляются в виде электронной таблицы. | 

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateShowsSalesReport

> GenerateReportResponse GenerateShowsSalesReport(ctx).GenerateShowsSalesReportRequest(generateShowsSalesReportRequest).Format(format).Execute()

Отчет «Аналитика продаж»



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    generateShowsSalesReportRequest := *openapiclient.NewGenerateShowsSalesReportRequest(time.Now(), time.Now(), openapiclient.ShowsSalesGroupingType("CATEGORIES")) // GenerateShowsSalesReportRequest | 
    format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. Пока отчеты доступны только в одном формате — они предоставляются в виде электронной таблицы. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GenerateShowsSalesReport(context.Background()).GenerateShowsSalesReportRequest(generateShowsSalesReportRequest).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GenerateShowsSalesReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateShowsSalesReport`: GenerateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GenerateShowsSalesReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateShowsSalesReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateShowsSalesReportRequest** | [**GenerateShowsSalesReportRequest**](GenerateShowsSalesReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. Пока отчеты доступны только в одном формате — они предоставляются в виде электронной таблицы. | 

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateStocksOnWarehousesReport

> GenerateReportResponse GenerateStocksOnWarehousesReport(ctx).GenerateStocksOnWarehousesReportRequest(generateStocksOnWarehousesReportRequest).Format(format).Execute()

Отчет по остаткам на складах



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    generateStocksOnWarehousesReportRequest := *openapiclient.NewGenerateStocksOnWarehousesReportRequest(int64(123)) // GenerateStocksOnWarehousesReportRequest | 
    format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. Пока отчеты доступны только в одном формате — они предоставляются в виде электронной таблицы. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GenerateStocksOnWarehousesReport(context.Background()).GenerateStocksOnWarehousesReportRequest(generateStocksOnWarehousesReportRequest).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GenerateStocksOnWarehousesReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateStocksOnWarehousesReport`: GenerateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GenerateStocksOnWarehousesReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateStocksOnWarehousesReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateStocksOnWarehousesReportRequest** | [**GenerateStocksOnWarehousesReportRequest**](GenerateStocksOnWarehousesReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. Пока отчеты доступны только в одном формате — они предоставляются в виде электронной таблицы. | 

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateUnitedMarketplaceServicesReport

> GenerateReportResponse GenerateUnitedMarketplaceServicesReport(ctx).GenerateUnitedMarketplaceServicesReportRequest(generateUnitedMarketplaceServicesReportRequest).Format(format).Execute()

Отчет по стоимости услуг



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    generateUnitedMarketplaceServicesReportRequest := *openapiclient.NewGenerateUnitedMarketplaceServicesReportRequest(int64(123)) // GenerateUnitedMarketplaceServicesReportRequest | 
    format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. Пока отчеты доступны только в одном формате — они предоставляются в виде электронной таблицы. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GenerateUnitedMarketplaceServicesReport(context.Background()).GenerateUnitedMarketplaceServicesReportRequest(generateUnitedMarketplaceServicesReportRequest).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GenerateUnitedMarketplaceServicesReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateUnitedMarketplaceServicesReport`: GenerateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GenerateUnitedMarketplaceServicesReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateUnitedMarketplaceServicesReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateUnitedMarketplaceServicesReportRequest** | [**GenerateUnitedMarketplaceServicesReportRequest**](GenerateUnitedMarketplaceServicesReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. Пока отчеты доступны только в одном формате — они предоставляются в виде электронной таблицы. | 

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateUnitedNettingReport

> GenerateReportResponse GenerateUnitedNettingReport(ctx).GenerateUnitedNettingReportRequest(generateUnitedNettingReportRequest).Format(format).Execute()

Отчет по платежам



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    generateUnitedNettingReportRequest := *openapiclient.NewGenerateUnitedNettingReportRequest(int64(123)) // GenerateUnitedNettingReportRequest | 
    format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. Пока отчеты доступны только в одном формате — они предоставляются в виде электронной таблицы. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GenerateUnitedNettingReport(context.Background()).GenerateUnitedNettingReportRequest(generateUnitedNettingReportRequest).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GenerateUnitedNettingReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateUnitedNettingReport`: GenerateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GenerateUnitedNettingReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateUnitedNettingReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateUnitedNettingReportRequest** | [**GenerateUnitedNettingReportRequest**](GenerateUnitedNettingReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. Пока отчеты доступны только в одном формате — они предоставляются в виде электронной таблицы. | 

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBidsInfoForBusiness

> GetBidsInfoResponse GetBidsInfoForBusiness(ctx, businessId).PageToken(pageToken).Limit(limit).GetBidsInfoRequest(getBidsInfoRequest).Execute()

Информация об установленных ставках



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)
    getBidsInfoRequest := *openapiclient.NewGetBidsInfoRequest() // GetBidsInfoRequest | description (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetBidsInfoForBusiness(context.Background(), businessId).PageToken(pageToken).Limit(limit).GetBidsInfoRequest(getBidsInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetBidsInfoForBusiness``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBidsInfoForBusiness`: GetBidsInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetBidsInfoForBusiness`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBidsInfoForBusinessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 
 **getBidsInfoRequest** | [**GetBidsInfoRequest**](GetBidsInfoRequest.md) | description | 

### Return type

[**GetBidsInfoResponse**](GetBidsInfoResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBidsRecommendations

> GetBidsRecommendationsResponse GetBidsRecommendations(ctx, businessId).GetBidsRecommendationsRequest(getBidsRecommendationsRequest).Execute()

Рекомендованные ставки для заданных товаров



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    getBidsRecommendationsRequest := *openapiclient.NewGetBidsRecommendationsRequest([]string{"Skus_example"}) // GetBidsRecommendationsRequest | description.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetBidsRecommendations(context.Background(), businessId).GetBidsRecommendationsRequest(getBidsRecommendationsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetBidsRecommendations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBidsRecommendations`: GetBidsRecommendationsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetBidsRecommendations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBidsRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getBidsRecommendationsRequest** | [**GetBidsRecommendationsRequest**](GetBidsRecommendationsRequest.md) | description. | 

### Return type

[**GetBidsRecommendationsResponse**](GetBidsRecommendationsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessQuarantineOffers

> GetQuarantineOffersResponse GetBusinessQuarantineOffers(ctx, businessId).GetQuarantineOffersRequest(getQuarantineOffersRequest).PageToken(pageToken).Limit(limit).Execute()

Список товаров, находящихся в карантине (основная цена)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    getQuarantineOffersRequest := *openapiclient.NewGetQuarantineOffersRequest() // GetQuarantineOffersRequest | 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetBusinessQuarantineOffers(context.Background(), businessId).GetQuarantineOffersRequest(getQuarantineOffersRequest).PageToken(pageToken).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetBusinessQuarantineOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinessQuarantineOffers`: GetQuarantineOffersResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetBusinessQuarantineOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessQuarantineOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getQuarantineOffersRequest** | [**GetQuarantineOffersRequest**](GetQuarantineOffersRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 

### Return type

[**GetQuarantineOffersResponse**](GetQuarantineOffersResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> GetCampaignResponse GetCampaign(ctx, campaignId).Execute()

Информация о магазине



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetCampaign(context.Background(), campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaign`: GetCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCampaignResponse**](GetCampaignResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignLogins

> GetCampaignLoginsResponse GetCampaignLogins(ctx, campaignId).Execute()

Логины, связанные с магазином



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetCampaignLogins(context.Background(), campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetCampaignLogins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignLogins`: GetCampaignLoginsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetCampaignLogins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignLoginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCampaignLoginsResponse**](GetCampaignLoginsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignOffers

> GetCampaignOffersResponse GetCampaignOffers(ctx, campaignId).GetCampaignOffersRequest(getCampaignOffersRequest).PageToken(pageToken).Limit(limit).Execute()

Список товаров, размещенных в заданном магазине, с параметрами размещения



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    getCampaignOffersRequest := *openapiclient.NewGetCampaignOffersRequest() // GetCampaignOffersRequest | 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetCampaignOffers(context.Background(), campaignId).GetCampaignOffersRequest(getCampaignOffersRequest).PageToken(pageToken).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetCampaignOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignOffers`: GetCampaignOffersResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetCampaignOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getCampaignOffersRequest** | [**GetCampaignOffersRequest**](GetCampaignOffersRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 

### Return type

[**GetCampaignOffersResponse**](GetCampaignOffersResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignQuarantineOffers

> GetQuarantineOffersResponse GetCampaignQuarantineOffers(ctx, campaignId).GetQuarantineOffersRequest(getQuarantineOffersRequest).PageToken(pageToken).Limit(limit).Execute()

Список товаров, находящихся в карантине (цена в магазине)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    getQuarantineOffersRequest := *openapiclient.NewGetQuarantineOffersRequest() // GetQuarantineOffersRequest | 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetCampaignQuarantineOffers(context.Background(), campaignId).GetQuarantineOffersRequest(getQuarantineOffersRequest).PageToken(pageToken).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetCampaignQuarantineOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignQuarantineOffers`: GetQuarantineOffersResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetCampaignQuarantineOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignQuarantineOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getQuarantineOffersRequest** | [**GetQuarantineOffersRequest**](GetQuarantineOffersRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 

### Return type

[**GetQuarantineOffersResponse**](GetQuarantineOffersResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> GetCampaignsResponse GetCampaigns(ctx).Page(page).PageSize(pageSize).Execute()

Магазины пользователя



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    page := int32(56) // int32 | Номер страницы результатов.  Значение по умолчанию: 1.  Используется вместе с параметром page_size.  `page_number` игнорируется, если задан `page_token`, `limit` или `offset`.  (optional) (default to 1)
    pageSize := int32(56) // int32 | Количество скрытых товаров на странице.  Используется вместе с параметром page_number.  `page_size` игнорируется, если задан `page_token`, `limit` или `offset`.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetCampaigns(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetCampaigns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaigns`: GetCampaignsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Номер страницы результатов.  Значение по умолчанию: 1.  Используется вместе с параметром page_size.  &#x60;page_number&#x60; игнорируется, если задан &#x60;page_token&#x60;, &#x60;limit&#x60; или &#x60;offset&#x60;.  | [default to 1]
 **pageSize** | **int32** | Количество скрытых товаров на странице.  Используется вместе с параметром page_number.  &#x60;page_size&#x60; игнорируется, если задан &#x60;page_token&#x60;, &#x60;limit&#x60; или &#x60;offset&#x60;.  | 

### Return type

[**GetCampaignsResponse**](GetCampaignsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignsByLogin

> GetCampaignsResponse GetCampaignsByLogin(ctx, login).Page(page).PageSize(pageSize).Execute()

Магазины, доступные логину



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    login := "login_example" // string | Логин пользователя.
    page := int32(56) // int32 | Номер страницы результатов.  Значение по умолчанию: 1.  Используется вместе с параметром page_size.  `page_number` игнорируется, если задан `page_token`, `limit` или `offset`.  (optional) (default to 1)
    pageSize := int32(56) // int32 | Количество скрытых товаров на странице.  Используется вместе с параметром page_number.  `page_size` игнорируется, если задан `page_token`, `limit` или `offset`.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetCampaignsByLogin(context.Background(), login).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetCampaignsByLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignsByLogin`: GetCampaignsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetCampaignsByLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Логин пользователя. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsByLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Номер страницы результатов.  Значение по умолчанию: 1.  Используется вместе с параметром page_size.  &#x60;page_number&#x60; игнорируется, если задан &#x60;page_token&#x60;, &#x60;limit&#x60; или &#x60;offset&#x60;.  | [default to 1]
 **pageSize** | **int32** | Количество скрытых товаров на странице.  Используется вместе с параметром page_number.  &#x60;page_size&#x60; игнорируется, если задан &#x60;page_token&#x60;, &#x60;limit&#x60; или &#x60;offset&#x60;.  | 

### Return type

[**GetCampaignsResponse**](GetCampaignsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryContentParameters

> GetCategoryContentParametersResponse GetCategoryContentParameters(ctx, categoryId).Execute()

Списки характеристик товаров по категориям



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    categoryId := int64(789) // int64 | Идентификатор категории на Маркете.  Чтобы узнать идентификатор категории, к которой относится интересующий вас товар, воспользуйтесь запросом [POST /businesses/{businessId}/offer-cards](../../reference/content/getOfferCardsContentStatus.md). 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetCategoryContentParameters(context.Background(), categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetCategoryContentParameters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategoryContentParameters`: GetCategoryContentParametersResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetCategoryContentParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | Идентификатор категории на Маркете.  Чтобы узнать идентификатор категории, к которой относится интересующий вас товар, воспользуйтесь запросом [POST /businesses/{businessId}/offer-cards](../../reference/content/getOfferCardsContentStatus.md).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryContentParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCategoryContentParametersResponse**](GetCategoryContentParametersResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeliveryServices

> GetDeliveryServicesResponse GetDeliveryServices(ctx).Execute()

Справочник служб доставки



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetDeliveryServices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetDeliveryServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeliveryServices`: GetDeliveryServicesResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetDeliveryServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeliveryServicesRequest struct via the builder pattern


### Return type

[**GetDeliveryServicesResponse**](GetDeliveryServicesResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGoodsStats

> GetGoodsStatsResponse GetGoodsStats(ctx, campaignId).GetGoodsStatsRequest(getGoodsStatsRequest).Execute()

Отчет по товарам



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    getGoodsStatsRequest := *openapiclient.NewGetGoodsStatsRequest([]string{"ShopSkus_example"}) // GetGoodsStatsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetGoodsStats(context.Background(), campaignId).GetGoodsStatsRequest(getGoodsStatsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetGoodsStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGoodsStats`: GetGoodsStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetGoodsStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoodsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getGoodsStatsRequest** | [**GetGoodsStatsRequest**](GetGoodsStatsRequest.md) |  | 

### Return type

[**GetGoodsStatsResponse**](GetGoodsStatsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHiddenOffers

> GetHiddenOffersResponse GetHiddenOffers(ctx, campaignId).OfferId(offerId).FeedId(feedId).PageToken(pageToken).Limit(limit).Offset(offset).Page(page).PageSize(pageSize).Execute()

Информация о скрытых вами товарах



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    offerId := []string{"Inner_example"} // []string | Идентификатор скрытого предложения.  (optional)
    feedId := []int64{int64(123)} // []int64 | {% note alert \"Это поле устарело\" %}  Не используйте его — это может привести к ошибкам.  {% endnote %}   Идентификатор прайс-листа.  (optional)
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)
    offset := int32(56) // int32 | Количество скрытых товаров, которые нужно не отображать в выходных данных, начиная с первого.  Скрытые товары выводятся отсортированными в лексикографическом порядке по возрастанию значений marketSku.  Используется вместе с параметром `limit`.  Если задан `offset`, параметры `page_number` и `page_size` игнорируются.  `offset` игнорируется, если задан `page_token`.  (optional)
    page := int32(56) // int32 | Номер страницы результатов.  Значение по умолчанию: 1.  Используется вместе с параметром page_size.  `page_number` игнорируется, если задан `page_token`, `limit` или `offset`.  (optional) (default to 1)
    pageSize := int32(56) // int32 | Количество скрытых товаров на странице.  Используется вместе с параметром page_number.  `page_size` игнорируется, если задан `page_token`, `limit` или `offset`.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetHiddenOffers(context.Background(), campaignId).OfferId(offerId).FeedId(feedId).PageToken(pageToken).Limit(limit).Offset(offset).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetHiddenOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHiddenOffers`: GetHiddenOffersResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetHiddenOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHiddenOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offerId** | **[]string** | Идентификатор скрытого предложения.  | 
 **feedId** | **[]int64** | {% note alert \&quot;Это поле устарело\&quot; %}  Не используйте его — это может привести к ошибкам.  {% endnote %}   Идентификатор прайс-листа.  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 
 **offset** | **int32** | Количество скрытых товаров, которые нужно не отображать в выходных данных, начиная с первого.  Скрытые товары выводятся отсортированными в лексикографическом порядке по возрастанию значений marketSku.  Используется вместе с параметром &#x60;limit&#x60;.  Если задан &#x60;offset&#x60;, параметры &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  &#x60;offset&#x60; игнорируется, если задан &#x60;page_token&#x60;.  | 
 **page** | **int32** | Номер страницы результатов.  Значение по умолчанию: 1.  Используется вместе с параметром page_size.  &#x60;page_number&#x60; игнорируется, если задан &#x60;page_token&#x60;, &#x60;limit&#x60; или &#x60;offset&#x60;.  | [default to 1]
 **pageSize** | **int32** | Количество скрытых товаров на странице.  Используется вместе с параметром page_number.  &#x60;page_size&#x60; игнорируется, если задан &#x60;page_token&#x60;, &#x60;limit&#x60; или &#x60;offset&#x60;.  | 

### Return type

[**GetHiddenOffersResponse**](GetHiddenOffersResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOfferCardsContentStatus

> GetOfferCardsContentStatusResponse GetOfferCardsContentStatus(ctx, businessId).PageToken(pageToken).Limit(limit).GetOfferCardsContentStatusRequest(getOfferCardsContentStatusRequest).Execute()

Получение информации о заполненности карточек



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)
    getOfferCardsContentStatusRequest := *openapiclient.NewGetOfferCardsContentStatusRequest() // GetOfferCardsContentStatusRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetOfferCardsContentStatus(context.Background(), businessId).PageToken(pageToken).Limit(limit).GetOfferCardsContentStatusRequest(getOfferCardsContentStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetOfferCardsContentStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOfferCardsContentStatus`: GetOfferCardsContentStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetOfferCardsContentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferCardsContentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 
 **getOfferCardsContentStatusRequest** | [**GetOfferCardsContentStatusRequest**](GetOfferCardsContentStatusRequest.md) |  | 

### Return type

[**GetOfferCardsContentStatusResponse**](GetOfferCardsContentStatusResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOfferMappingEntries

> GetOfferMappingEntriesResponse GetOfferMappingEntries(ctx, campaignId).OfferId(offerId).ShopSku(shopSku).MappingKind(mappingKind).Status(status).Availability(availability).CategoryId(categoryId).Vendor(vendor).PageToken(pageToken).Limit(limit).Execute()

Список товаров в каталоге



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    offerId := []string{"Inner_example"} // []string | Идентификатор товара в каталоге. (optional)
    shopSku := []string{"Inner_example"} // []string | Ваш SKU товара.  Параметр может быть указан несколько раз, например:  ``` ...shop_sku=123&shop_sku=129&shop_sku=141... ```  В запросе можно указать либо параметр `shopSku`, либо любые параметры для фильтрации товаров. Совместное использование параметра `shopSku` и параметров для фильтрации приведет к ошибке.  (optional)
    mappingKind := openapiclient.OfferMappingKindType("ACTIVE") // OfferMappingKindType | Тип маппинга. (optional)
    status := []openapiclient.OfferProcessingStatusType{openapiclient.OfferProcessingStatusType("UNKNOWN")} // []OfferProcessingStatusType | Фильтрация по статусу публикации товара:  * READY — товар прошел модерацию. * IN_WORK — товар проходит модерацию. * NEED_CONTENT — для товара без SKU на Маркете marketSku нужно найти карточку самостоятельно или создать ее. * NEED_INFO — товар не прошел модерацию из-за ошибок или недостающих сведений в описании товара. * REJECTED — товар не прошел модерацию, так как Маркет не планирует размещать подобные товары. * SUSPENDED — товар не прошел модерацию, так как Маркет пока не размещает подобные товары. * OTHER — товар не прошел модерацию по другой причине.  Можно указать несколько статусов в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  ``` ...status=READY,IN_WORK... ...status=READY&status=IN_WORK... ```  В запросе можно указать либо параметр shopSku, либо любые параметры для фильтрации товаров. Совместное использование параметра shopSku и параметров для фильтрации приведет к ошибке.  (optional)
    availability := []openapiclient.OfferAvailabilityStatusType{openapiclient.OfferAvailabilityStatusType("ACTIVE")} // []OfferAvailabilityStatusType | Фильтрация по планам поставок товара:  * ACTIVE — поставки будут. * INACTIVE — поставок не будет: товар есть на складе, но вы больше не планируете его поставлять. * DELISTED — архив: товар закончился на складе, и его поставок больше не будет.  Можно указать несколько значений в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  ``` ...availability=INACTIVE,DELISTED... ...availability=INACTIVE&availability=DELISTED... ```  В запросе можно указать либо параметр `shopSku`, либо любые параметры для фильтрации товаров. Совместное использование параметра `shopSku` и параметров для фильтрации приведет к ошибке.  (optional)
    categoryId := []int32{int32(123)} // []int32 | Фильтрация по идентификатору категории на Маркете.  Чтобы узнать идентификатор категории, откройте ее страницу и посмотрите на ее URL. Идентификатор — это число после «...?hid=». Например:  ##https://market.yandex.ru/catalog--bezmeny/65590/list?hid=13431995##  Можно указать несколько идентификаторов в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  ``` ...category_id=14727164,14382343... ...category_id=14727164&category_id=14382343... ```  В запросе можно указать либо параметр `shopSku`, либо любые параметры для фильтрации товаров. Совместное использование параметра `shopSku` и параметров для фильтрации приведет к ошибке.  (optional)
    vendor := []string{"Inner_example"} // []string | Фильтрация по бренду товара.  Можно указать несколько брендов в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  ``` ...vendor=Aqua%20Minerale,Borjomi... ...vendor=Aqua%20Minerale&vendor=Borjomi... ```  Чтобы товар попал в результаты фильтрации, его бренд должен точно совпадать с одним из указанных в запросе. Например, если указан бренд Schwarzkopf, то в результатах не будет товаров Schwarzkopf Professional.  Если в названии бренда есть символы, которые не входят в таблицу ASCII (в том числе кириллические символы), используйте для них URL-кодирование. Например, пробел — %20, апостроф «'» — %27 и т. д. Подробнее см. в разделе [Кодирование URL русскоязычной Википедии](https://ru.wikipedia.org/wiki/URL#Кодирование_URL).  В запросе можно указать либо параметр shopSku, либо любые параметры для фильтрации товаров. Совместное использование параметра shopSku и параметров для фильтрации приведет к ошибке.  (optional)
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetOfferMappingEntries(context.Background(), campaignId).OfferId(offerId).ShopSku(shopSku).MappingKind(mappingKind).Status(status).Availability(availability).CategoryId(categoryId).Vendor(vendor).PageToken(pageToken).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetOfferMappingEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOfferMappingEntries`: GetOfferMappingEntriesResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetOfferMappingEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferMappingEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offerId** | **[]string** | Идентификатор товара в каталоге. | 
 **shopSku** | **[]string** | Ваш SKU товара.  Параметр может быть указан несколько раз, например:  &#x60;&#x60;&#x60; ...shop_sku&#x3D;123&amp;shop_sku&#x3D;129&amp;shop_sku&#x3D;141... &#x60;&#x60;&#x60;  В запросе можно указать либо параметр &#x60;shopSku&#x60;, либо любые параметры для фильтрации товаров. Совместное использование параметра &#x60;shopSku&#x60; и параметров для фильтрации приведет к ошибке.  | 
 **mappingKind** | [**OfferMappingKindType**](OfferMappingKindType.md) | Тип маппинга. | 
 **status** | [**[]OfferProcessingStatusType**](OfferProcessingStatusType.md) | Фильтрация по статусу публикации товара:  * READY — товар прошел модерацию. * IN_WORK — товар проходит модерацию. * NEED_CONTENT — для товара без SKU на Маркете marketSku нужно найти карточку самостоятельно или создать ее. * NEED_INFO — товар не прошел модерацию из-за ошибок или недостающих сведений в описании товара. * REJECTED — товар не прошел модерацию, так как Маркет не планирует размещать подобные товары. * SUSPENDED — товар не прошел модерацию, так как Маркет пока не размещает подобные товары. * OTHER — товар не прошел модерацию по другой причине.  Можно указать несколько статусов в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  &#x60;&#x60;&#x60; ...status&#x3D;READY,IN_WORK... ...status&#x3D;READY&amp;status&#x3D;IN_WORK... &#x60;&#x60;&#x60;  В запросе можно указать либо параметр shopSku, либо любые параметры для фильтрации товаров. Совместное использование параметра shopSku и параметров для фильтрации приведет к ошибке.  | 
 **availability** | [**[]OfferAvailabilityStatusType**](OfferAvailabilityStatusType.md) | Фильтрация по планам поставок товара:  * ACTIVE — поставки будут. * INACTIVE — поставок не будет: товар есть на складе, но вы больше не планируете его поставлять. * DELISTED — архив: товар закончился на складе, и его поставок больше не будет.  Можно указать несколько значений в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  &#x60;&#x60;&#x60; ...availability&#x3D;INACTIVE,DELISTED... ...availability&#x3D;INACTIVE&amp;availability&#x3D;DELISTED... &#x60;&#x60;&#x60;  В запросе можно указать либо параметр &#x60;shopSku&#x60;, либо любые параметры для фильтрации товаров. Совместное использование параметра &#x60;shopSku&#x60; и параметров для фильтрации приведет к ошибке.  | 
 **categoryId** | **[]int32** | Фильтрация по идентификатору категории на Маркете.  Чтобы узнать идентификатор категории, откройте ее страницу и посмотрите на ее URL. Идентификатор — это число после «...?hid&#x3D;». Например:  ##https://market.yandex.ru/catalog--bezmeny/65590/list?hid&#x3D;13431995##  Можно указать несколько идентификаторов в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  &#x60;&#x60;&#x60; ...category_id&#x3D;14727164,14382343... ...category_id&#x3D;14727164&amp;category_id&#x3D;14382343... &#x60;&#x60;&#x60;  В запросе можно указать либо параметр &#x60;shopSku&#x60;, либо любые параметры для фильтрации товаров. Совместное использование параметра &#x60;shopSku&#x60; и параметров для фильтрации приведет к ошибке.  | 
 **vendor** | **[]string** | Фильтрация по бренду товара.  Можно указать несколько брендов в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  &#x60;&#x60;&#x60; ...vendor&#x3D;Aqua%20Minerale,Borjomi... ...vendor&#x3D;Aqua%20Minerale&amp;vendor&#x3D;Borjomi... &#x60;&#x60;&#x60;  Чтобы товар попал в результаты фильтрации, его бренд должен точно совпадать с одним из указанных в запросе. Например, если указан бренд Schwarzkopf, то в результатах не будет товаров Schwarzkopf Professional.  Если в названии бренда есть символы, которые не входят в таблицу ASCII (в том числе кириллические символы), используйте для них URL-кодирование. Например, пробел — %20, апостроф «&#39;» — %27 и т. д. Подробнее см. в разделе [Кодирование URL русскоязычной Википедии](https://ru.wikipedia.org/wiki/URL#Кодирование_URL).  В запросе можно указать либо параметр shopSku, либо любые параметры для фильтрации товаров. Совместное использование параметра shopSku и параметров для фильтрации приведет к ошибке.  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 

### Return type

[**GetOfferMappingEntriesResponse**](GetOfferMappingEntriesResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOfferMappings

> GetOfferMappingsResponse GetOfferMappings(ctx, businessId).PageToken(pageToken).Limit(limit).GetOfferMappingsRequest(getOfferMappingsRequest).Execute()

Информация о товарах в каталоге



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)
    getOfferMappingsRequest := *openapiclient.NewGetOfferMappingsRequest() // GetOfferMappingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetOfferMappings(context.Background(), businessId).PageToken(pageToken).Limit(limit).GetOfferMappingsRequest(getOfferMappingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetOfferMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOfferMappings`: GetOfferMappingsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetOfferMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 
 **getOfferMappingsRequest** | [**GetOfferMappingsRequest**](GetOfferMappingsRequest.md) |  | 

### Return type

[**GetOfferMappingsResponse**](GetOfferMappingsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOfferRecommendations

> GetOfferRecommendationsResponse GetOfferRecommendations(ctx, businessId).GetOfferRecommendationsRequest(getOfferRecommendationsRequest).PageToken(pageToken).Limit(limit).Execute()

Рекомендации Маркета, касающиеся цен



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    getOfferRecommendationsRequest := *openapiclient.NewGetOfferRecommendationsRequest() // GetOfferRecommendationsRequest | 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetOfferRecommendations(context.Background(), businessId).GetOfferRecommendationsRequest(getOfferRecommendationsRequest).PageToken(pageToken).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetOfferRecommendations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOfferRecommendations`: GetOfferRecommendationsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetOfferRecommendations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getOfferRecommendationsRequest** | [**GetOfferRecommendationsRequest**](GetOfferRecommendationsRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 

### Return type

[**GetOfferRecommendationsResponse**](GetOfferRecommendationsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> GetOrderResponse GetOrder(ctx, campaignId, orderId).Execute()

Информация о заказе



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    orderId := int64(789) // int64 | Идентификатор заказа.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetOrder(context.Background(), campaignId, orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrder`: GetOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrderResponse**](GetOrderResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderLabelsData

> GetOrderLabelsDataResponse GetOrderLabelsData(ctx, campaignId, orderId).Execute()

Информация на ярлыках‑наклейках



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    orderId := int64(789) // int64 | Идентификатор заказа.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetOrderLabelsData(context.Background(), campaignId, orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetOrderLabelsData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderLabelsData`: GetOrderLabelsDataResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetOrderLabelsData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderLabelsDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrderLabelsDataResponse**](GetOrderLabelsDataResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrders

> GetOrdersResponse GetOrders(ctx, campaignId).Status(status).Substatus(substatus).FromDate(fromDate).ToDate(toDate).SupplierShipmentDateFrom(supplierShipmentDateFrom).SupplierShipmentDateTo(supplierShipmentDateTo).UpdatedAtFrom(updatedAtFrom).UpdatedAtTo(updatedAtTo).DispatchType(dispatchType).Fake(fake).HasCis(hasCis).OnlyWaitingForCancellationApprove(onlyWaitingForCancellationApprove).OnlyEstimatedDelivery(onlyEstimatedDelivery).Page(page).PageSize(pageSize).Execute()

Информация о заказах



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    status := openapiclient.OrderStatusType("PLACING") // OrderStatusType | Статус заказа:  * `CANCELLED` — заказ отменен.  * `DELIVERED` — заказ получен покупателем.  * `DELIVERY` — заказ передан в службу доставки.  * `PICKUP` — заказ доставлен в пункт самовывоза.  * `PROCESSING` — заказ находится в обработке.  * `UNPAID` — заказ оформлен, но еще не оплачен (если выбрана оплата при оформлении).  Также могут возвращаться другие значения. Обрабатывать их не требуется.  (optional)
    substatus := openapiclient.OrderSubstatusType("RESERVATION_EXPIRED") // OrderSubstatusType | Этап обработки заказа (если он имеет статус `PROCESSING`) или причина отмены заказа (если он имеет статус `CANCELLED`).  Возможные значения для заказа в статусе `PROCESSING`:  * `STARTED` — заказ подтвержден, его можно начать обрабатывать. * `READY_TO_SHIP` — заказ собран и готов к отправке. * `SHIPPED` — заказ передан службе доставки.  Возможные значения для заказа в статусе `CANCELLED`:  * `DELIVERY_SERVICE_UNDELIVERED` — служба доставки не смогла доставить заказ.  * `PROCESSING_EXPIRED` — значение более не используется.  * `REPLACING_ORDER` — покупатель решил заменить товар другим по собственной инициативе.  * `RESERVATION_EXPIRED` — покупатель не завершил оформление зарезервированного заказа в течение 10 минут.  * `RESERVATION_FAILED` — Маркет не может продолжить дальнейшую обработку заказа.  * `SHOP_FAILED` — магазин не может выполнить заказ.  * `USER_CHANGED_MIND` — покупатель отменил заказ по личным причинам.  * `USER_NOT_PAID` — покупатель не оплатил заказ (для типа оплаты `PREPAID`) в течение 30 минут.  * `USER_REFUSED_DELIVERY` — покупателя не устроили условия доставки.  * `USER_REFUSED_PRODUCT` — покупателю не подошел товар.  * `USER_REFUSED_QUALITY` — покупателя не устроило качество товара.  * `USER_UNREACHABLE` — не удалось связаться с покупателем. * `USER_WANTS_TO_CHANGE_DELIVERY_DATE` — покупатель хочет получить заказ в другой день. * `CANCELLED_COURIER_NOT_FOUND` — не удалось найти курьера.  Также могут возвращаться другие значения. Обрабатывать их не требуется.  (optional)
    fromDate := time.Now() // string | Начальная дата для фильтрации заказов по дате оформления.  Формат даты: `ДД-ММ-ГГГГ`.  Между начальной и конечной датой (параметр `toDate`) должно быть не больше 30 дней.  Значение по умолчанию: 30 дней назад от текущей даты.  (optional)
    toDate := time.Now() // string | Конечная дата для фильтрации заказов по дате оформления.  Формат даты: `ДД-ММ-ГГГГ`.  Между начальной (параметр `fromDate`) и конечной датой должно быть не больше 30 дней.  Значение по умолчанию: текущая дата.  (optional)
    supplierShipmentDateFrom := time.Now() // string | Начальная дата для фильтрации заказов по дате отгрузки в службу доставки (параметр `shipmentDate`).  Формат даты: `ДД-ММ-ГГГГ`.  Между начальной и конечной датой (параметр `supplierShipmentDateTo`) должно быть не больше 30 дней.  Начальная дата включается в интервал для фильтрации.  (optional)
    supplierShipmentDateTo := time.Now() // string | Конечная дата для фильтрации заказов по дате отгрузки в службу доставки (параметр `shipmentDate`).  Формат даты: `ДД-ММ-ГГГГ`.  Между начальной (параметр `supplierShipmentDateFrom`) и конечной датой должно быть не больше 30 дней.  Конечная дата не включается в интервал для фильтрации.  (optional)
    updatedAtFrom := time.Now() // time.Time | Начальная дата для фильтрации заказов по дате и времени обновления (параметр `updatedAt`).  Формат даты: ISO 8601 со смещением относительно UTC. Например, `2017-11-21T00:42:42+03:00`.  Между начальной и конечной датой (параметр `updatedAtTo`) должно быть не больше 30 дней.  Начальная дата включается в интервал для фильтрации.  (optional)
    updatedAtTo := time.Now() // time.Time | Конечная дата для фильтрации заказов по дате и времени обновления (параметр `updatedAt`).  Формат даты: ISO 8601 со смещением относительно UTC. Например, `2017-11-21T00:42:42+03:00`.  Между начальной (параметр `updatedAtFrom`) и конечной датой должно быть не больше 30 дней.  Конечная дата не включается в интервал для фильтрации.  (optional)
    dispatchType := openapiclient.OrderDeliveryDispatchType("UNKNOWN") // OrderDeliveryDispatchType | Способ отгрузки (optional)
    fake := true // bool | Фильтрация заказов по типам:  * `false` — заказ пользователя.  * `true` — тестовый заказ Маркета.  (optional) (default to false)
    hasCis := true // bool | Нужно ли вернуть только те заказы, в составе которых есть хотя бы один товар с кодом идентификации из системы «Честный ЗНАК»:  * `true` — да;  * `false` — нет.  Такие коды присваиваются товарам, которые подлежат маркировке и относятся к определенным категориям.  (optional) (default to false)
    onlyWaitingForCancellationApprove := true // bool | Фильтрация заказов по наличию запросов покупателей на отмену:  * `true` — возвращаются только заказы, которые находятся в статусе `DELIVERY` или `PICKUP` и которые пользователи решили отменить. Чтобы подтвердить или отклонить отмену, отправьте запрос [PUT campaigns/{campaignId}/orders/{orderId}/cancellation/accept](../../reference/orders/acceptOrderCancellation).  (optional) (default to false)
    onlyEstimatedDelivery := true // bool | Фильтрация заказов с долгой доставкой (31-60 дней) по подтвержденной дате доставки:  * `true` — возвращаются только заказы с неподтвержденной датой доставки. * `false` — фильтрация не применяется.  (optional) (default to false)
    page := int32(56) // int32 | Номер страницы результатов.  Значение по умолчанию: 1.  Используется вместе с параметром page_size.  `page_number` игнорируется, если задан `page_token`, `limit` или `offset`.  (optional) (default to 1)
    pageSize := int32(56) // int32 | Количество скрытых товаров на странице.  Используется вместе с параметром page_number.  `page_size` игнорируется, если задан `page_token`, `limit` или `offset`.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetOrders(context.Background(), campaignId).Status(status).Substatus(substatus).FromDate(fromDate).ToDate(toDate).SupplierShipmentDateFrom(supplierShipmentDateFrom).SupplierShipmentDateTo(supplierShipmentDateTo).UpdatedAtFrom(updatedAtFrom).UpdatedAtTo(updatedAtTo).DispatchType(dispatchType).Fake(fake).HasCis(hasCis).OnlyWaitingForCancellationApprove(onlyWaitingForCancellationApprove).OnlyEstimatedDelivery(onlyEstimatedDelivery).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrders`: GetOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**OrderStatusType**](OrderStatusType.md) | Статус заказа:  * &#x60;CANCELLED&#x60; — заказ отменен.  * &#x60;DELIVERED&#x60; — заказ получен покупателем.  * &#x60;DELIVERY&#x60; — заказ передан в службу доставки.  * &#x60;PICKUP&#x60; — заказ доставлен в пункт самовывоза.  * &#x60;PROCESSING&#x60; — заказ находится в обработке.  * &#x60;UNPAID&#x60; — заказ оформлен, но еще не оплачен (если выбрана оплата при оформлении).  Также могут возвращаться другие значения. Обрабатывать их не требуется.  | 
 **substatus** | [**OrderSubstatusType**](OrderSubstatusType.md) | Этап обработки заказа (если он имеет статус &#x60;PROCESSING&#x60;) или причина отмены заказа (если он имеет статус &#x60;CANCELLED&#x60;).  Возможные значения для заказа в статусе &#x60;PROCESSING&#x60;:  * &#x60;STARTED&#x60; — заказ подтвержден, его можно начать обрабатывать. * &#x60;READY_TO_SHIP&#x60; — заказ собран и готов к отправке. * &#x60;SHIPPED&#x60; — заказ передан службе доставки.  Возможные значения для заказа в статусе &#x60;CANCELLED&#x60;:  * &#x60;DELIVERY_SERVICE_UNDELIVERED&#x60; — служба доставки не смогла доставить заказ.  * &#x60;PROCESSING_EXPIRED&#x60; — значение более не используется.  * &#x60;REPLACING_ORDER&#x60; — покупатель решил заменить товар другим по собственной инициативе.  * &#x60;RESERVATION_EXPIRED&#x60; — покупатель не завершил оформление зарезервированного заказа в течение 10 минут.  * &#x60;RESERVATION_FAILED&#x60; — Маркет не может продолжить дальнейшую обработку заказа.  * &#x60;SHOP_FAILED&#x60; — магазин не может выполнить заказ.  * &#x60;USER_CHANGED_MIND&#x60; — покупатель отменил заказ по личным причинам.  * &#x60;USER_NOT_PAID&#x60; — покупатель не оплатил заказ (для типа оплаты &#x60;PREPAID&#x60;) в течение 30 минут.  * &#x60;USER_REFUSED_DELIVERY&#x60; — покупателя не устроили условия доставки.  * &#x60;USER_REFUSED_PRODUCT&#x60; — покупателю не подошел товар.  * &#x60;USER_REFUSED_QUALITY&#x60; — покупателя не устроило качество товара.  * &#x60;USER_UNREACHABLE&#x60; — не удалось связаться с покупателем. * &#x60;USER_WANTS_TO_CHANGE_DELIVERY_DATE&#x60; — покупатель хочет получить заказ в другой день. * &#x60;CANCELLED_COURIER_NOT_FOUND&#x60; — не удалось найти курьера.  Также могут возвращаться другие значения. Обрабатывать их не требуется.  | 
 **fromDate** | **string** | Начальная дата для фильтрации заказов по дате оформления.  Формат даты: &#x60;ДД-ММ-ГГГГ&#x60;.  Между начальной и конечной датой (параметр &#x60;toDate&#x60;) должно быть не больше 30 дней.  Значение по умолчанию: 30 дней назад от текущей даты.  | 
 **toDate** | **string** | Конечная дата для фильтрации заказов по дате оформления.  Формат даты: &#x60;ДД-ММ-ГГГГ&#x60;.  Между начальной (параметр &#x60;fromDate&#x60;) и конечной датой должно быть не больше 30 дней.  Значение по умолчанию: текущая дата.  | 
 **supplierShipmentDateFrom** | **string** | Начальная дата для фильтрации заказов по дате отгрузки в службу доставки (параметр &#x60;shipmentDate&#x60;).  Формат даты: &#x60;ДД-ММ-ГГГГ&#x60;.  Между начальной и конечной датой (параметр &#x60;supplierShipmentDateTo&#x60;) должно быть не больше 30 дней.  Начальная дата включается в интервал для фильтрации.  | 
 **supplierShipmentDateTo** | **string** | Конечная дата для фильтрации заказов по дате отгрузки в службу доставки (параметр &#x60;shipmentDate&#x60;).  Формат даты: &#x60;ДД-ММ-ГГГГ&#x60;.  Между начальной (параметр &#x60;supplierShipmentDateFrom&#x60;) и конечной датой должно быть не больше 30 дней.  Конечная дата не включается в интервал для фильтрации.  | 
 **updatedAtFrom** | **time.Time** | Начальная дата для фильтрации заказов по дате и времени обновления (параметр &#x60;updatedAt&#x60;).  Формат даты: ISO 8601 со смещением относительно UTC. Например, &#x60;2017-11-21T00:42:42+03:00&#x60;.  Между начальной и конечной датой (параметр &#x60;updatedAtTo&#x60;) должно быть не больше 30 дней.  Начальная дата включается в интервал для фильтрации.  | 
 **updatedAtTo** | **time.Time** | Конечная дата для фильтрации заказов по дате и времени обновления (параметр &#x60;updatedAt&#x60;).  Формат даты: ISO 8601 со смещением относительно UTC. Например, &#x60;2017-11-21T00:42:42+03:00&#x60;.  Между начальной (параметр &#x60;updatedAtFrom&#x60;) и конечной датой должно быть не больше 30 дней.  Конечная дата не включается в интервал для фильтрации.  | 
 **dispatchType** | [**OrderDeliveryDispatchType**](OrderDeliveryDispatchType.md) | Способ отгрузки | 
 **fake** | **bool** | Фильтрация заказов по типам:  * &#x60;false&#x60; — заказ пользователя.  * &#x60;true&#x60; — тестовый заказ Маркета.  | [default to false]
 **hasCis** | **bool** | Нужно ли вернуть только те заказы, в составе которых есть хотя бы один товар с кодом идентификации из системы «Честный ЗНАК»:  * &#x60;true&#x60; — да;  * &#x60;false&#x60; — нет.  Такие коды присваиваются товарам, которые подлежат маркировке и относятся к определенным категориям.  | [default to false]
 **onlyWaitingForCancellationApprove** | **bool** | Фильтрация заказов по наличию запросов покупателей на отмену:  * &#x60;true&#x60; — возвращаются только заказы, которые находятся в статусе &#x60;DELIVERY&#x60; или &#x60;PICKUP&#x60; и которые пользователи решили отменить. Чтобы подтвердить или отклонить отмену, отправьте запрос [PUT campaigns/{campaignId}/orders/{orderId}/cancellation/accept](../../reference/orders/acceptOrderCancellation).  | [default to false]
 **onlyEstimatedDelivery** | **bool** | Фильтрация заказов с долгой доставкой (31-60 дней) по подтвержденной дате доставки:  * &#x60;true&#x60; — возвращаются только заказы с неподтвержденной датой доставки. * &#x60;false&#x60; — фильтрация не применяется.  | [default to false]
 **page** | **int32** | Номер страницы результатов.  Значение по умолчанию: 1.  Используется вместе с параметром page_size.  &#x60;page_number&#x60; игнорируется, если задан &#x60;page_token&#x60;, &#x60;limit&#x60; или &#x60;offset&#x60;.  | [default to 1]
 **pageSize** | **int32** | Количество скрытых товаров на странице.  Используется вместе с параметром page_number.  &#x60;page_size&#x60; игнорируется, если задан &#x60;page_token&#x60;, &#x60;limit&#x60; или &#x60;offset&#x60;.  | 

### Return type

[**GetOrdersResponse**](GetOrdersResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersStats

> GetOrdersStatsResponse GetOrdersStats(ctx, campaignId).PageToken(pageToken).Limit(limit).GetOrdersStatsRequest(getOrdersStatsRequest).Execute()

Отчет по заказам



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)
    getOrdersStatsRequest := *openapiclient.NewGetOrdersStatsRequest() // GetOrdersStatsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetOrdersStats(context.Background(), campaignId).PageToken(pageToken).Limit(limit).GetOrdersStatsRequest(getOrdersStatsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetOrdersStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrdersStats`: GetOrdersStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetOrdersStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 
 **getOrdersStatsRequest** | [**GetOrdersStatsRequest**](GetOrdersStatsRequest.md) |  | 

### Return type

[**GetOrdersStatsResponse**](GetOrdersStatsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrices

> GetPricesResponse GetPrices(ctx, campaignId).PageToken(pageToken).Limit(limit).Archived(archived).Execute()

Список цен



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)
    archived := true // bool | Фильтр по нахождению в архиве (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetPrices(context.Background(), campaignId).PageToken(pageToken).Limit(limit).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrices`: GetPricesResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 
 **archived** | **bool** | Фильтр по нахождению в архиве | [default to false]

### Return type

[**GetPricesResponse**](GetPricesResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricesByOfferIds

> GetPricesByOfferIdsResponse GetPricesByOfferIds(ctx, campaignId).GetPricesByOfferIdsRequest(getPricesByOfferIdsRequest).PageToken(pageToken).Limit(limit).Execute()

Просмотр установленных в магазине цен



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    getPricesByOfferIdsRequest := *openapiclient.NewGetPricesByOfferIdsRequest() // GetPricesByOfferIdsRequest | 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetPricesByOfferIds(context.Background(), campaignId).GetPricesByOfferIdsRequest(getPricesByOfferIdsRequest).PageToken(pageToken).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetPricesByOfferIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPricesByOfferIds`: GetPricesByOfferIdsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetPricesByOfferIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricesByOfferIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getPricesByOfferIdsRequest** | [**GetPricesByOfferIdsRequest**](GetPricesByOfferIdsRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 

### Return type

[**GetPricesByOfferIdsResponse**](GetPricesByOfferIdsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportInfo

> GetReportInfoResponse GetReportInfo(ctx, reportId).Execute()

Статус генерации и скачивание готовых отчетов



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    reportId := "reportId_example" // string | Идентификатор отчета, который вы получили после запуска генерации. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetReportInfo(context.Background(), reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetReportInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportInfo`: GetReportInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetReportInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | Идентификатор отчета, который вы получили после запуска генерации.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReportInfoResponse**](GetReportInfoResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturn

> GetReturnResponse GetReturn(ctx, campaignId, orderId, returnId).Execute()

Информация о возврате или невыкупе



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    orderId := int64(789) // int64 | Идентификатор заказа.
    returnId := int64(789) // int64 | Идентификатор возврата

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetReturn(context.Background(), campaignId, orderId, returnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetReturn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturn`: GetReturnResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetReturn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**orderId** | **int64** | Идентификатор заказа. | 
**returnId** | **int64** | Идентификатор возврата | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetReturnResponse**](GetReturnResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturnApplication

> *os.File GetReturnApplication(ctx, campaignId, orderId, returnId).Execute()

Получение заявления на возврат



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    orderId := int64(789) // int64 | Идентификатор заказа.
    returnId := int64(789) // int64 | Идентификатор возврата

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetReturnApplication(context.Background(), campaignId, orderId, returnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetReturnApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturnApplication`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetReturnApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**orderId** | **int64** | Идентификатор заказа. | 
**returnId** | **int64** | Идентификатор возврата | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturnPhoto

> *os.File GetReturnPhoto(ctx, campaignId, orderId, returnId, itemId, imageHash).Execute()

Получение фотографии возврата



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    orderId := int64(789) // int64 | Идентификатор заказа.
    returnId := int64(789) // int64 | Идентификатор возврата
    itemId := int64(789) // int64 | Идентификатор товара в возврате
    imageHash := "imageHash_example" // string | Хеш ссылки изображения для загрузки

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetReturnPhoto(context.Background(), campaignId, orderId, returnId, itemId, imageHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetReturnPhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturnPhoto`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetReturnPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**orderId** | **int64** | Идентификатор заказа. | 
**returnId** | **int64** | Идентификатор возврата | 
**itemId** | **int64** | Идентификатор товара в возврате | 
**imageHash** | **string** | Хеш ссылки изображения для загрузки | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturns

> GetReturnsResponse GetReturns(ctx, campaignId).PageToken(pageToken).Limit(limit).OrderIds(orderIds).Statuses(statuses).Type_(type_).FromDate(fromDate).ToDate(toDate).FromDate2(fromDate2).ToDate2(toDate2).Execute()

Список возвратов и невыкупов



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)
    orderIds := []int64{int64(123543)} // []int64 | Идентификаторы заказов — для фильтрации результатов.  Несколько идентификаторов перечисляются через запятую без пробела.  (optional)
    statuses := []openapiclient.RefundStatusType{openapiclient.RefundStatusType("STARTED_BY_USER")} // []RefundStatusType | Статусы возвратов или невыкупов — для фильтрации результатов.  Возможные значения:  * CREATED — возврат создан.  * RECEIVED — принят у покупателя.  * IN_TRANSIT — возврат в пути.  * READY_FOR_PICKUP — возврат готов к выдаче магазину.  * PICKED — возврат выдан магазину.  * LOST — возврат утерян при транспортировке.  * CANCELLED — возврат отменен.  * EXPIRED — возврат просрочен, покупатель не передал товар.  * FULFILMENT_RECEIVED — возврат принят на складе Маркета.  * PREPARED_FOR_UTILIZATION — возврат передан в утилизацию.  * UTILIZED — возврат утилизирован.  Несколько статусов перечисляются через запятую.  (optional)
    type_ := openapiclient.ReturnType("UNREDEEMED") // ReturnType | Тип заказа для фильтрации:  * RETURN — возврат.  * UNREDEEMED — невыкуп.  Если не указывать, в ответе будут и возвраты, и невыкупы.  (optional)
    fromDate := time.Now() // string | Начальные дата и время для фильтрации возвратов или невыкупов по дате оформления.  Формат: ГГГГ-ММ-ДД.  (optional)
    toDate := time.Now() // string | Конечные дата и время для фильтрации возвратов или невыкупов по дате оформления.  Формат: ГГГГ-ММ-ДД.  (optional)
    fromDate2 := time.Now() // string | Дата, с которой интересуют возвраты (устаревшее, будет удалено) (optional)
    toDate2 := time.Now() // string | Дата, до которой интересуют возвраты (устаревшее, будет удалено) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetReturns(context.Background(), campaignId).PageToken(pageToken).Limit(limit).OrderIds(orderIds).Statuses(statuses).Type_(type_).FromDate(fromDate).ToDate(toDate).FromDate2(fromDate2).ToDate2(toDate2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetReturns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturns`: GetReturnsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetReturns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 
 **orderIds** | **[]int64** | Идентификаторы заказов — для фильтрации результатов.  Несколько идентификаторов перечисляются через запятую без пробела.  | 
 **statuses** | [**[]RefundStatusType**](RefundStatusType.md) | Статусы возвратов или невыкупов — для фильтрации результатов.  Возможные значения:  * CREATED — возврат создан.  * RECEIVED — принят у покупателя.  * IN_TRANSIT — возврат в пути.  * READY_FOR_PICKUP — возврат готов к выдаче магазину.  * PICKED — возврат выдан магазину.  * LOST — возврат утерян при транспортировке.  * CANCELLED — возврат отменен.  * EXPIRED — возврат просрочен, покупатель не передал товар.  * FULFILMENT_RECEIVED — возврат принят на складе Маркета.  * PREPARED_FOR_UTILIZATION — возврат передан в утилизацию.  * UTILIZED — возврат утилизирован.  Несколько статусов перечисляются через запятую.  | 
 **type_** | [**ReturnType**](ReturnType.md) | Тип заказа для фильтрации:  * RETURN — возврат.  * UNREDEEMED — невыкуп.  Если не указывать, в ответе будут и возвраты, и невыкупы.  | 
 **fromDate** | **string** | Начальные дата и время для фильтрации возвратов или невыкупов по дате оформления.  Формат: ГГГГ-ММ-ДД.  | 
 **toDate** | **string** | Конечные дата и время для фильтрации возвратов или невыкупов по дате оформления.  Формат: ГГГГ-ММ-ДД.  | 
 **fromDate2** | **string** | Дата, с которой интересуют возвраты (устаревшее, будет удалено) | 
 **toDate2** | **string** | Дата, до которой интересуют возвраты (устаревшее, будет удалено) | 

### Return type

[**GetReturnsResponse**](GetReturnsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShipment

> GetShipmentResponse GetShipment(ctx, campaignId, shipmentId).Execute()

Получение информации об отгрузке



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    shipmentId := int64(789) // int64 | Идентификатор отгрузки.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetShipment(context.Background(), campaignId, shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShipment`: GetShipmentResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetShipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**shipmentId** | **int64** | Идентификатор отгрузки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetShipmentResponse**](GetShipmentResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShipmentOrdersInfo

> GetShipmentOrdersInfoResponse GetShipmentOrdersInfo(ctx, campaignId, shipmentId).Execute()

Получение информации о ярлыках



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    shipmentId := int64(789) // int64 | Идентификатор отгрузки.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetShipmentOrdersInfo(context.Background(), campaignId, shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetShipmentOrdersInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShipmentOrdersInfo`: GetShipmentOrdersInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetShipmentOrdersInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**shipmentId** | **int64** | Идентификатор отгрузки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShipmentOrdersInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetShipmentOrdersInfoResponse**](GetShipmentOrdersInfoResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStocks

> GetWarehouseStocksResponse GetStocks(ctx, campaignId).PageToken(pageToken).Limit(limit).GetWarehouseStocksRequest(getWarehouseStocksRequest).Execute()

Информация об остатках и оборачиваемости



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)
    getWarehouseStocksRequest := *openapiclient.NewGetWarehouseStocksRequest() // GetWarehouseStocksRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetStocks(context.Background(), campaignId).PageToken(pageToken).Limit(limit).GetWarehouseStocksRequest(getWarehouseStocksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetStocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStocks`: GetWarehouseStocksResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetStocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 
 **getWarehouseStocksRequest** | [**GetWarehouseStocksRequest**](GetWarehouseStocksRequest.md) |  | 

### Return type

[**GetWarehouseStocksResponse**](GetWarehouseStocksResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuggestedOfferMappingEntries

> GetSuggestedOfferMappingEntriesResponse GetSuggestedOfferMappingEntries(ctx, campaignId).GetSuggestedOfferMappingEntriesRequest(getSuggestedOfferMappingEntriesRequest).Execute()

Рекомендованные карточки для ваших товаров



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    getSuggestedOfferMappingEntriesRequest := *openapiclient.NewGetSuggestedOfferMappingEntriesRequest([]openapiclient.MappingsOfferDTO{*openapiclient.NewMappingsOfferDTO()}) // GetSuggestedOfferMappingEntriesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetSuggestedOfferMappingEntries(context.Background(), campaignId).GetSuggestedOfferMappingEntriesRequest(getSuggestedOfferMappingEntriesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetSuggestedOfferMappingEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSuggestedOfferMappingEntries`: GetSuggestedOfferMappingEntriesResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetSuggestedOfferMappingEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuggestedOfferMappingEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getSuggestedOfferMappingEntriesRequest** | [**GetSuggestedOfferMappingEntriesRequest**](GetSuggestedOfferMappingEntriesRequest.md) |  | 

### Return type

[**GetSuggestedOfferMappingEntriesResponse**](GetSuggestedOfferMappingEntriesResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuggestedOfferMappings

> GetSuggestedOfferMappingsResponse GetSuggestedOfferMappings(ctx, businessId).GetSuggestedOfferMappingsRequest(getSuggestedOfferMappingsRequest).Execute()

Предварительный просмотр карточек на Маркете, соответствующих вашим товарам



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    getSuggestedOfferMappingsRequest := *openapiclient.NewGetSuggestedOfferMappingsRequest() // GetSuggestedOfferMappingsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetSuggestedOfferMappings(context.Background(), businessId).GetSuggestedOfferMappingsRequest(getSuggestedOfferMappingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetSuggestedOfferMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSuggestedOfferMappings`: GetSuggestedOfferMappingsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetSuggestedOfferMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuggestedOfferMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getSuggestedOfferMappingsRequest** | [**GetSuggestedOfferMappingsRequest**](GetSuggestedOfferMappingsRequest.md) |  | 

### Return type

[**GetSuggestedOfferMappingsResponse**](GetSuggestedOfferMappingsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuggestedPrices

> SuggestPricesResponse GetSuggestedPrices(ctx, campaignId).SuggestPricesRequest(suggestPricesRequest).Execute()

Цены для продвижения товаров



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    suggestPricesRequest := *openapiclient.NewSuggestPricesRequest([]openapiclient.SuggestOfferPriceDTO{*openapiclient.NewSuggestOfferPriceDTO()}) // SuggestPricesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetSuggestedPrices(context.Background(), campaignId).SuggestPricesRequest(suggestPricesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetSuggestedPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSuggestedPrices`: SuggestPricesResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetSuggestedPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuggestedPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suggestPricesRequest** | [**SuggestPricesRequest**](SuggestPricesRequest.md) |  | 

### Return type

[**SuggestPricesResponse**](SuggestPricesResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWarehouses

> GetWarehousesResponse GetWarehouses(ctx, businessId).Execute()

Список складов и групп складов



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.GetWarehouses(context.Background(), businessId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.GetWarehouses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWarehouses`: GetWarehousesResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.GetWarehouses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWarehousesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWarehousesResponse**](GetWarehousesResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvideOrderItemIdentifiers

> ProvideOrderItemIdentifiersResponse ProvideOrderItemIdentifiers(ctx, campaignId, orderId).ProvideOrderItemIdentifiersRequest(provideOrderItemIdentifiersRequest).Execute()

Передача уникальных кодов маркировки единиц товара



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    orderId := int64(789) // int64 | Идентификатор заказа.
    provideOrderItemIdentifiersRequest := *openapiclient.NewProvideOrderItemIdentifiersRequest([]openapiclient.OrderItemInstanceModificationDTO{*openapiclient.NewOrderItemInstanceModificationDTO(int64(123), []openapiclient.BriefOrderItemInstanceDTO{*openapiclient.NewBriefOrderItemInstanceDTO()})}) // ProvideOrderItemIdentifiersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.ProvideOrderItemIdentifiers(context.Background(), campaignId, orderId).ProvideOrderItemIdentifiersRequest(provideOrderItemIdentifiersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.ProvideOrderItemIdentifiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvideOrderItemIdentifiers`: ProvideOrderItemIdentifiersResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.ProvideOrderItemIdentifiers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvideOrderItemIdentifiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **provideOrderItemIdentifiersRequest** | [**ProvideOrderItemIdentifiersRequest**](ProvideOrderItemIdentifiersRequest.md) |  | 

### Return type

[**ProvideOrderItemIdentifiersResponse**](ProvideOrderItemIdentifiersResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBidsForBusiness

> EmptyApiResponse PutBidsForBusiness(ctx, businessId).PutSkuBidsRequest(putSkuBidsRequest).Execute()

Включение буста продаж и установка ставок



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    putSkuBidsRequest := *openapiclient.NewPutSkuBidsRequest([]openapiclient.SkuBidItemDTO{*openapiclient.NewSkuBidItemDTO("Sku_example", int32(570))}) // PutSkuBidsRequest | description

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.PutBidsForBusiness(context.Background(), businessId).PutSkuBidsRequest(putSkuBidsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.PutBidsForBusiness``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutBidsForBusiness`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.PutBidsForBusiness`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutBidsForBusinessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putSkuBidsRequest** | [**PutSkuBidsRequest**](PutSkuBidsRequest.md) | description | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRegionChildren

> GetRegionWithChildrenResponse SearchRegionChildren(ctx, regionId).Page(page).PageSize(pageSize).Execute()

Информация о дочерних регионах



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    regionId := int64(789) // int64 | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса `GET /regions`. 
    page := int32(56) // int32 | Номер страницы результатов.  Значение по умолчанию: 1.  Используется вместе с параметром page_size.  `page_number` игнорируется, если задан `page_token`, `limit` или `offset`.  (optional) (default to 1)
    pageSize := int32(56) // int32 | Количество скрытых товаров на странице.  Используется вместе с параметром page_number.  `page_size` игнорируется, если задан `page_token`, `limit` или `offset`.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.SearchRegionChildren(context.Background(), regionId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.SearchRegionChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchRegionChildren`: GetRegionWithChildrenResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.SearchRegionChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **int64** | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса &#x60;GET /regions&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchRegionChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Номер страницы результатов.  Значение по умолчанию: 1.  Используется вместе с параметром page_size.  &#x60;page_number&#x60; игнорируется, если задан &#x60;page_token&#x60;, &#x60;limit&#x60; или &#x60;offset&#x60;.  | [default to 1]
 **pageSize** | **int32** | Количество скрытых товаров на странице.  Используется вместе с параметром page_number.  &#x60;page_size&#x60; игнорируется, если задан &#x60;page_token&#x60;, &#x60;limit&#x60; или &#x60;offset&#x60;.  | 

### Return type

[**GetRegionWithChildrenResponse**](GetRegionWithChildrenResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRegionsById

> GetRegionsResponse SearchRegionsById(ctx, regionId).Execute()

Информация о регионе



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    regionId := int64(789) // int64 | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса `GET /regions`. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.SearchRegionsById(context.Background(), regionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.SearchRegionsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchRegionsById`: GetRegionsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.SearchRegionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **int64** | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса &#x60;GET /regions&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchRegionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRegionsResponse**](GetRegionsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRegionsByName

> GetRegionsResponse SearchRegionsByName(ctx).Name(name).PageToken(pageToken).Limit(limit).Execute()

Метод для поиска регионов по их имени



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    name := "name_example" // string | Название региона.  Важно учитывать регистр: первая буква должна быть заглавной, остальные — строчными. Например, `Москва`. 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.SearchRegionsByName(context.Background()).Name(name).PageToken(pageToken).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.SearchRegionsByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchRegionsByName`: GetRegionsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.SearchRegionsByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRegionsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Название региона.  Важно учитывать регистр: первая буква должна быть заглавной, остальные — строчными. Например, &#x60;Москва&#x60;.  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 

### Return type

[**GetRegionsResponse**](GetRegionsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchShipments

> SearchShipmentsResponse SearchShipments(ctx, campaignId).SearchShipmentsRequest(searchShipmentsRequest).PageToken(pageToken).Limit(limit).Execute()

Получение информации об отгрузках



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    searchShipmentsRequest := *openapiclient.NewSearchShipmentsRequest() // SearchShipmentsRequest | 
    pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token`, параметры `offset`, `page_number` и `page_size` игнорируются.  (optional)
    limit := int32(20) // int32 | Количество товаров на одной странице.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.SearchShipments(context.Background(), campaignId).SearchShipmentsRequest(searchShipmentsRequest).PageToken(pageToken).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.SearchShipments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchShipments`: SearchShipmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.SearchShipments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchShipmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchShipmentsRequest** | [**SearchShipmentsRequest**](SearchShipmentsRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается самая старая страница.  Рекомендуется передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60;, параметры &#x60;offset&#x60;, &#x60;page_number&#x60; и &#x60;page_size&#x60; игнорируются.  | 
 **limit** | **int32** | Количество товаров на одной странице.  | 

### Return type

[**SearchShipmentsResponse**](SearchShipmentsResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOrderShipmentBoxes

> SetOrderShipmentBoxesResponse SetOrderShipmentBoxes(ctx, campaignId, orderId, shipmentId).SetOrderShipmentBoxesRequest(setOrderShipmentBoxesRequest).Execute()

Передача количества грузовых мест в заказе



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    orderId := int64(789) // int64 | Идентификатор заказа.
    shipmentId := int64(789) // int64 | Параметр больше не используется. Вставьте любое число — просто чтобы получился корректный URL. 
    setOrderShipmentBoxesRequest := *openapiclient.NewSetOrderShipmentBoxesRequest([]openapiclient.ParcelBoxDTO{*openapiclient.NewParcelBoxDTO()}) // SetOrderShipmentBoxesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.SetOrderShipmentBoxes(context.Background(), campaignId, orderId, shipmentId).SetOrderShipmentBoxesRequest(setOrderShipmentBoxesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.SetOrderShipmentBoxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetOrderShipmentBoxes`: SetOrderShipmentBoxesResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.SetOrderShipmentBoxes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**orderId** | **int64** | Идентификатор заказа. | 
**shipmentId** | **int64** | Параметр больше не используется. Вставьте любое число — просто чтобы получился корректный URL.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOrderShipmentBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **setOrderShipmentBoxesRequest** | [**SetOrderShipmentBoxesRequest**](SetOrderShipmentBoxesRequest.md) |  | 

### Return type

[**SetOrderShipmentBoxesResponse**](SetOrderShipmentBoxesResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetShipmentPalletsCount

> EmptyApiResponse SetShipmentPalletsCount(ctx, campaignId, shipmentId).SetShipmentPalletsCountRequest(setShipmentPalletsCountRequest).Execute()

Передача количества упаковок в отгрузке



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    shipmentId := int64(789) // int64 | Идентификатор отгрузки.
    setShipmentPalletsCountRequest := *openapiclient.NewSetShipmentPalletsCountRequest(int32(123)) // SetShipmentPalletsCountRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.SetShipmentPalletsCount(context.Background(), campaignId, shipmentId).SetShipmentPalletsCountRequest(setShipmentPalletsCountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.SetShipmentPalletsCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetShipmentPalletsCount`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.SetShipmentPalletsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**shipmentId** | **int64** | Идентификатор отгрузки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetShipmentPalletsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setShipmentPalletsCountRequest** | [**SetShipmentPalletsCountRequest**](SetShipmentPalletsCountRequest.md) |  | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferOrdersFromShipment

> EmptyApiResponse TransferOrdersFromShipment(ctx, campaignId, shipmentId).TransferOrdersFromShipmentRequest(transferOrdersFromShipmentRequest).Execute()

Перенос заказов в следующую отгрузку



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    shipmentId := int64(789) // int64 | Идентификатор отгрузки.
    transferOrdersFromShipmentRequest := *openapiclient.NewTransferOrdersFromShipmentRequest([]int64{int64(123)}) // TransferOrdersFromShipmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.TransferOrdersFromShipment(context.Background(), campaignId, shipmentId).TransferOrdersFromShipmentRequest(transferOrdersFromShipmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.TransferOrdersFromShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferOrdersFromShipment`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.TransferOrdersFromShipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**shipmentId** | **int64** | Идентификатор отгрузки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferOrdersFromShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transferOrdersFromShipmentRequest** | [**TransferOrdersFromShipmentRequest**](TransferOrdersFromShipmentRequest.md) |  | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBusinessPrices

> EmptyApiResponse UpdateBusinessPrices(ctx, businessId).UpdateBusinessPricesRequest(updateBusinessPricesRequest).Execute()

Установка цен



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    updateBusinessPricesRequest := *openapiclient.NewUpdateBusinessPricesRequest([]openapiclient.UpdateBusinessOfferPriceDTO{*openapiclient.NewUpdateBusinessOfferPriceDTO("OfferId_example", *openapiclient.NewUpdatePriceWithDiscountDTO(float32(123), "RUR"))}) // UpdateBusinessPricesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.UpdateBusinessPrices(context.Background(), businessId).UpdateBusinessPricesRequest(updateBusinessPricesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.UpdateBusinessPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBusinessPrices`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.UpdateBusinessPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBusinessPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBusinessPricesRequest** | [**UpdateBusinessPricesRequest**](UpdateBusinessPricesRequest.md) |  | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaignOffers

> EmptyApiResponse UpdateCampaignOffers(ctx, campaignId).UpdateCampaignOffersRequest(updateCampaignOffersRequest).Execute()

Настройка размещения товаров в магазине



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    updateCampaignOffersRequest := *openapiclient.NewUpdateCampaignOffersRequest([]openapiclient.UpdateCampaignOfferDTO{*openapiclient.NewUpdateCampaignOfferDTO("OfferId_example")}) // UpdateCampaignOffersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.UpdateCampaignOffers(context.Background(), campaignId).UpdateCampaignOffersRequest(updateCampaignOffersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.UpdateCampaignOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCampaignOffers`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.UpdateCampaignOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCampaignOffersRequest** | [**UpdateCampaignOffersRequest**](UpdateCampaignOffersRequest.md) |  | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOfferContent

> UpdateOfferContentResponse UpdateOfferContent(ctx, businessId).UpdateOfferContentRequest(updateOfferContentRequest).Execute()

Редактирование категорийных характеристик товара



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    updateOfferContentRequest := *openapiclient.NewUpdateOfferContentRequest([]openapiclient.OfferContentDTO{*openapiclient.NewOfferContentDTO("OfferId_example", int32(123), []openapiclient.ParameterValueDTO{*openapiclient.NewParameterValueDTO(int64(123), "Value_example")})}) // UpdateOfferContentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.UpdateOfferContent(context.Background(), businessId).UpdateOfferContentRequest(updateOfferContentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.UpdateOfferContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOfferContent`: UpdateOfferContentResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.UpdateOfferContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOfferContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOfferContentRequest** | [**UpdateOfferContentRequest**](UpdateOfferContentRequest.md) |  | 

### Return type

[**UpdateOfferContentResponse**](UpdateOfferContentResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOfferMappingEntries

> EmptyApiResponse UpdateOfferMappingEntries(ctx, campaignId).UpdateOfferMappingEntryRequest(updateOfferMappingEntryRequest).Execute()

Добавление и редактирование товаров в каталоге



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    updateOfferMappingEntryRequest := *openapiclient.NewUpdateOfferMappingEntryRequest([]openapiclient.OfferMappingEntryDTO{*openapiclient.NewOfferMappingEntryDTO()}) // UpdateOfferMappingEntryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.UpdateOfferMappingEntries(context.Background(), campaignId).UpdateOfferMappingEntryRequest(updateOfferMappingEntryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.UpdateOfferMappingEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOfferMappingEntries`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.UpdateOfferMappingEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOfferMappingEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOfferMappingEntryRequest** | [**UpdateOfferMappingEntryRequest**](UpdateOfferMappingEntryRequest.md) |  | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOfferMappings

> EmptyApiResponse UpdateOfferMappings(ctx, businessId).UpdateOfferMappingsRequest(updateOfferMappingsRequest).Execute()

Добавление товаров в каталог и редактирование информации о них



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    updateOfferMappingsRequest := *openapiclient.NewUpdateOfferMappingsRequest([]openapiclient.UpdateOfferMappingDTO{*openapiclient.NewUpdateOfferMappingDTO(*openapiclient.NewUpdateOfferDTO("OfferId_example"))}) // UpdateOfferMappingsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.UpdateOfferMappings(context.Background(), businessId).UpdateOfferMappingsRequest(updateOfferMappingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.UpdateOfferMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOfferMappings`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.UpdateOfferMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы узнать идентификатор, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md#businessdto).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOfferMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOfferMappingsRequest** | [**UpdateOfferMappingsRequest**](UpdateOfferMappingsRequest.md) |  | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrderItems

> UpdateOrderItems(ctx, campaignId, orderId).UpdateOrderItemRequest(updateOrderItemRequest).Execute()

Удаление товара из заказа или уменьшение числа единиц



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    orderId := int64(789) // int64 | Идентификатор заказа.
    updateOrderItemRequest := *openapiclient.NewUpdateOrderItemRequest([]openapiclient.OrderItemModificationDTO{*openapiclient.NewOrderItemModificationDTO(int64(123), int32(123))}) // UpdateOrderItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FbsAPI.UpdateOrderItems(context.Background(), campaignId, orderId).UpdateOrderItemRequest(updateOrderItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.UpdateOrderItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrderItemRequest** | [**UpdateOrderItemRequest**](UpdateOrderItemRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrderStatus

> UpdateOrderStatusResponse UpdateOrderStatus(ctx, campaignId, orderId).UpdateOrderStatusRequest(updateOrderStatusRequest).Execute()

Изменение статуса заказа



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    orderId := int64(789) // int64 | Идентификатор заказа.
    updateOrderStatusRequest := *openapiclient.NewUpdateOrderStatusRequest(*openapiclient.NewOrderStatusChangeDTO(openapiclient.OrderStatusType("PLACING"))) // UpdateOrderStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.UpdateOrderStatus(context.Background(), campaignId, orderId).UpdateOrderStatusRequest(updateOrderStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.UpdateOrderStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrderStatus`: UpdateOrderStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.UpdateOrderStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrderStatusRequest** | [**UpdateOrderStatusRequest**](UpdateOrderStatusRequest.md) |  | 

### Return type

[**UpdateOrderStatusResponse**](UpdateOrderStatusResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrderStatuses

> UpdateOrderStatusesResponse UpdateOrderStatuses(ctx, campaignId).UpdateOrderStatusesRequest(updateOrderStatusesRequest).Execute()

Изменение статусов заказа



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    updateOrderStatusesRequest := *openapiclient.NewUpdateOrderStatusesRequest([]openapiclient.OrderStateDTO{*openapiclient.NewOrderStateDTO(int64(123), openapiclient.OrderStatusType("PLACING"))}) // UpdateOrderStatusesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.UpdateOrderStatuses(context.Background(), campaignId).UpdateOrderStatusesRequest(updateOrderStatusesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.UpdateOrderStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrderStatuses`: UpdateOrderStatusesResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.UpdateOrderStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrderStatusesRequest** | [**UpdateOrderStatusesRequest**](UpdateOrderStatusesRequest.md) |  | 

### Return type

[**UpdateOrderStatusesResponse**](UpdateOrderStatusesResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrices

> EmptyApiResponse UpdatePrices(ctx, campaignId).UpdatePricesRequest(updatePricesRequest).Execute()

Установка цен на товары в конкретном магазине



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    updatePricesRequest := *openapiclient.NewUpdatePricesRequest([]openapiclient.OfferPriceDTO{*openapiclient.NewOfferPriceDTO()}) // UpdatePricesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.UpdatePrices(context.Background(), campaignId).UpdatePricesRequest(updatePricesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.UpdatePrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePrices`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.UpdatePrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePricesRequest** | [**UpdatePricesRequest**](UpdatePricesRequest.md) |  | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStocks

> EmptyApiResponse UpdateStocks(ctx, campaignId).UpdateStocksRequest(updateStocksRequest).Execute()

Передача информации об остатках



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/devteamclub/ya-client"
)

func main() {
    campaignId := int64(789) // int64 | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
    updateStocksRequest := *openapiclient.NewUpdateStocksRequest([]openapiclient.StockDTO{*openapiclient.NewStockDTO("Sku_example", int64(123), []openapiclient.StockItemDTO{*openapiclient.NewStockItemDTO(int64(123), openapiclient.StockType("FIT"), time.Now())})}) // UpdateStocksRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FbsAPI.UpdateStocks(context.Background(), campaignId).UpdateStocksRequest(updateStocksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FbsAPI.UpdateStocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStocks`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FbsAPI.UpdateStocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStocksRequest** | [**UpdateStocksRequest**](UpdateStocksRequest.md) |  | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

