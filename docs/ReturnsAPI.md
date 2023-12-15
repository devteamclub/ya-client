# \ReturnsAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReturn**](ReturnsAPI.md#GetReturn) | **Get** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId} | Информация о возврате или невыкупе
[**GetReturnApplication**](ReturnsAPI.md#GetReturnApplication) | **Get** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/application | Получение заявления на возврат
[**GetReturnPhoto**](ReturnsAPI.md#GetReturnPhoto) | **Get** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/decision/{itemId}/image/{imageHash} | Получение фотографии возврата
[**GetReturns**](ReturnsAPI.md#GetReturns) | **Get** /campaigns/{campaignId}/returns | Список возвратов и невыкупов
[**SetReturnDecision**](ReturnsAPI.md#SetReturnDecision) | **Post** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/decision | Принятие/изменение решения по позиции в возврате
[**SubmitReturnDecision**](ReturnsAPI.md#SubmitReturnDecision) | **Post** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/decision/submit | Подтверждение решения по возвратным позициям



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
    resp, r, err := apiClient.ReturnsAPI.GetReturn(context.Background(), campaignId, orderId, returnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.GetReturn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturn`: GetReturnResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.GetReturn`: %v\n", resp)
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
    resp, r, err := apiClient.ReturnsAPI.GetReturnApplication(context.Background(), campaignId, orderId, returnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.GetReturnApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturnApplication`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.GetReturnApplication`: %v\n", resp)
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
    resp, r, err := apiClient.ReturnsAPI.GetReturnPhoto(context.Background(), campaignId, orderId, returnId, itemId, imageHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.GetReturnPhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturnPhoto`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.GetReturnPhoto`: %v\n", resp)
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
    resp, r, err := apiClient.ReturnsAPI.GetReturns(context.Background(), campaignId).PageToken(pageToken).Limit(limit).OrderIds(orderIds).Statuses(statuses).Type_(type_).FromDate(fromDate).ToDate(toDate).FromDate2(fromDate2).ToDate2(toDate2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.GetReturns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturns`: GetReturnsResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.GetReturns`: %v\n", resp)
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


## SetReturnDecision

> EmptyApiResponse SetReturnDecision(ctx, campaignId, orderId, returnId).SetReturnDecisionRequest(setReturnDecisionRequest).Execute()

Принятие/изменение решения по позиции в возврате



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
    setReturnDecisionRequest := *openapiclient.NewSetReturnDecisionRequest(int64(123), openapiclient.ReturnRequestDecisionType("REFUND_MONEY")) // SetReturnDecisionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnsAPI.SetReturnDecision(context.Background(), campaignId, orderId, returnId).SetReturnDecisionRequest(setReturnDecisionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.SetReturnDecision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetReturnDecision`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.SetReturnDecision`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiSetReturnDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **setReturnDecisionRequest** | [**SetReturnDecisionRequest**](SetReturnDecisionRequest.md) |  | 

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


## SubmitReturnDecision

> EmptyApiResponse SubmitReturnDecision(ctx, campaignId, orderId, returnId).Execute()

Подтверждение решения по возвратным позициям



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
    resp, r, err := apiClient.ReturnsAPI.SubmitReturnDecision(context.Background(), campaignId, orderId, returnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.SubmitReturnDecision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitReturnDecision`: EmptyApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.SubmitReturnDecision`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiSubmitReturnDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

