# \ReportsAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateGoodsMovementReport**](ReportsAPI.md#GenerateGoodsMovementReport) | **Post** /reports/goods-movement/generate | Отчет по движению товаров
[**GenerateGoodsRealizationReport**](ReportsAPI.md#GenerateGoodsRealizationReport) | **Post** /reports/goods-realization/generate | Отчет по реализации
[**GeneratePricesReport**](ReportsAPI.md#GeneratePricesReport) | **Post** /reports/prices/generate | Отчет «Цены на рынке»
[**GenerateShowsSalesReport**](ReportsAPI.md#GenerateShowsSalesReport) | **Post** /reports/shows-sales/generate | Отчет «Аналитика продаж»
[**GenerateStocksOnWarehousesReport**](ReportsAPI.md#GenerateStocksOnWarehousesReport) | **Post** /reports/stocks-on-warehouses/generate | Отчет по остаткам на складах
[**GenerateUnitedMarketplaceServicesReport**](ReportsAPI.md#GenerateUnitedMarketplaceServicesReport) | **Post** /reports/united-marketplace-services/generate | Отчет по стоимости услуг
[**GenerateUnitedNettingReport**](ReportsAPI.md#GenerateUnitedNettingReport) | **Post** /reports/united-netting/generate | Отчет по платежам
[**GetReportInfo**](ReportsAPI.md#GetReportInfo) | **Get** /reports/info/{reportId} | Статус генерации и скачивание готовых отчетов



## GenerateGoodsMovementReport

> GenerateReportResponse GenerateGoodsMovementReport(ctx).GenerateGoodsMovementReportRequest(generateGoodsMovementReportRequest).Format(format).Execute()

Отчет по движению товаров



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
    generateGoodsMovementReportRequest := *openapiclient.NewGenerateGoodsMovementReportRequest(int64(123), time.Now(), time.Now()) // GenerateGoodsMovementReportRequest | 
    format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. Пока отчеты доступны только в одном формате — они предоставляются в виде электронной таблицы. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.GenerateGoodsMovementReport(context.Background()).GenerateGoodsMovementReportRequest(generateGoodsMovementReportRequest).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateGoodsMovementReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateGoodsMovementReport`: GenerateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateGoodsMovementReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateGoodsMovementReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateGoodsMovementReportRequest** | [**GenerateGoodsMovementReportRequest**](GenerateGoodsMovementReportRequest.md) |  | 
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
    resp, r, err := apiClient.ReportsAPI.GenerateGoodsRealizationReport(context.Background()).GenerateGoodsRealizationReportRequest(generateGoodsRealizationReportRequest).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateGoodsRealizationReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateGoodsRealizationReport`: GenerateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateGoodsRealizationReport`: %v\n", resp)
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
    resp, r, err := apiClient.ReportsAPI.GeneratePricesReport(context.Background()).GeneratePricesReportRequest(generatePricesReportRequest).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GeneratePricesReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneratePricesReport`: GenerateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GeneratePricesReport`: %v\n", resp)
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
    resp, r, err := apiClient.ReportsAPI.GenerateShowsSalesReport(context.Background()).GenerateShowsSalesReportRequest(generateShowsSalesReportRequest).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateShowsSalesReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateShowsSalesReport`: GenerateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateShowsSalesReport`: %v\n", resp)
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
    resp, r, err := apiClient.ReportsAPI.GenerateStocksOnWarehousesReport(context.Background()).GenerateStocksOnWarehousesReportRequest(generateStocksOnWarehousesReportRequest).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateStocksOnWarehousesReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateStocksOnWarehousesReport`: GenerateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateStocksOnWarehousesReport`: %v\n", resp)
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
    resp, r, err := apiClient.ReportsAPI.GenerateUnitedMarketplaceServicesReport(context.Background()).GenerateUnitedMarketplaceServicesReportRequest(generateUnitedMarketplaceServicesReportRequest).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateUnitedMarketplaceServicesReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateUnitedMarketplaceServicesReport`: GenerateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateUnitedMarketplaceServicesReport`: %v\n", resp)
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
    resp, r, err := apiClient.ReportsAPI.GenerateUnitedNettingReport(context.Background()).GenerateUnitedNettingReportRequest(generateUnitedNettingReportRequest).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateUnitedNettingReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateUnitedNettingReport`: GenerateReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateUnitedNettingReport`: %v\n", resp)
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
    resp, r, err := apiClient.ReportsAPI.GetReportInfo(context.Background(), reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportInfo`: GetReportInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportInfo`: %v\n", resp)
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

