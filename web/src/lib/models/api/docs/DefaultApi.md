# DefaultApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getGreeting**](DefaultApi.md#getgreeting) | **POST** /#get_greeting |  |



## getGreeting

> ResponseBodyOfTheGetGreetingMethod getGreeting(requestBodyOfTheGetGreetingMethod)



### Example

```ts
import {
  Configuration,
  DefaultApi,
} from '';
import type { GetGreetingRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new DefaultApi();

  const body = {
    // RequestBodyOfTheGetGreetingMethod
    requestBodyOfTheGetGreetingMethod: ...,
  } satisfies GetGreetingRequest;

  try {
    const data = await api.getGreeting(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **requestBodyOfTheGetGreetingMethod** | [RequestBodyOfTheGetGreetingMethod](RequestBodyOfTheGetGreetingMethod.md) |  | |

### Return type

[**ResponseBodyOfTheGetGreetingMethod**](ResponseBodyOfTheGetGreetingMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | JSON-RPC response body |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

