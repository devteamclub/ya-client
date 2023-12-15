# \StocksAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActualStocks**](StocksAPI.md#GetActualStocks) | **Get** /campaigns/{campaignId}/warehouses/{warehouseId}/stocks/actual | Запрос информации об остатках
[**GetStocks**](StocksAPI.md#GetStocks) | **Post** /campaigns/{campaignId}/offers/stocks | Информация об остатках и оборачиваемости
[**UpdateStocks**](StocksAPI.md#UpdateStocks) | **Put** /campaigns/{campaignId}/offers/stocks | Передача информации об остатках



## GetActualStocks

> GetActualStocksResponse GetActualStocks(ctx, campaignId, warehouseId).Sku(sku).Execute()

Запрос информации об остатках



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
    warehouseId := int64(789) // int64 | Идентификатор склада.
    sku := []string{"Inner_example"} // []string | Фильтр по SKU (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StocksAPI.GetActualStocks(context.Background(), campaignId, warehouseId).Sku(sku).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StocksAPI.GetActualStocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActualStocks`: GetActualStocksResponse
    fmt.Fprintf(os.Stdout, "Response from `StocksAPI.GetActualStocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании в API и магазина в кабинете. Каждая кампания в API соответствует магазину в кабинете.  Чтобы узнать идентификаторы своих магазинов, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 
**warehouseId** | **int64** | Идентификатор склада. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActualStocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sku** | **[]string** | Фильтр по SKU | 

### Return type

[**GetActualStocksResponse**](GetActualStocksResponse.md)

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
    resp, r, err := apiClient.StocksAPI.GetStocks(context.Background(), campaignId).PageToken(pageToken).Limit(limit).GetWarehouseStocksRequest(getWarehouseStocksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StocksAPI.GetStocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStocks`: GetWarehouseStocksResponse
    fmt.Fprintf(os.Stdout, "Response from `StocksAPI.GetStocks`: %v\n", resp)
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
    resp, r, err := apiClient.StocksAPI.UpdateStocks(context.Background(), campaignId).UpdateStocksRequest(updateStocksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StocksAPI.UpdateStocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStocks`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `StocksAPI.UpdateStocks`: %v\n", resp)
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

