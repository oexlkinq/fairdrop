# DefaultApi

All URIs are relative to *https://plck.ru/fairdrop*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**foldersCreatePost**](#folderscreatepost) | **POST** /folders/create | создать папку|
|[**foldersFolderDelete**](#foldersfolderdelete) | **DELETE** /folders/{folder} | удалить папку|
|[**foldersFolderFilenameGet**](#foldersfolderfilenameget) | **GET** /folders/{folder}/{filename} | скачать файл|
|[**foldersFolderGet**](#foldersfolderget) | **GET** /folders/{folder} | получить список файлов папки|
|[**foldersFolderPost**](#foldersfolderpost) | **POST** /folders/{folder} | загрузить файлы в папку|
|[**foldersPushPost**](#folderspushpost) | **POST** /folders/push | аплоад файлов с созданием папки|

# **foldersCreatePost**
> Folder foldersCreatePost()

создаёт новую папку и возвращает информацию о ней

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

const { status, data } = await apiInstance.foldersCreatePost();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**Folder**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**5XX** | внутренняя ошибка сервера |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **foldersFolderDelete**
> foldersFolderDelete()

удаляет папку

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let folder: string; //пароль-идентификатор папки (default to undefined)

const { status, data } = await apiInstance.foldersFolderDelete(
    folder
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **folder** | [**string**] | пароль-идентификатор папки | defaults to undefined|


### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**404** | запрошенная папка не существует |  -  |
|**5XX** | внутренняя ошибка сервера |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **foldersFolderFilenameGet**
> File foldersFolderFilenameGet()


### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let folder: string; //пароль-идентификатор папки (default to undefined)
let filename: string; //имя файла (default to undefined)

const { status, data } = await apiInstance.foldersFolderFilenameGet(
    folder,
    filename
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **folder** | [**string**] | пароль-идентификатор папки | defaults to undefined|
| **filename** | [**string**] | имя файла | defaults to undefined|


### Return type

**File**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**404** | запрошенные файл или папка не найдены |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **foldersFolderGet**
> Array<string> foldersFolderGet()

возвращает список из имен файлов содержащихся в папке

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let folder: string; //пароль-идентификатор папки (default to undefined)

const { status, data } = await apiInstance.foldersFolderGet(
    folder
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **folder** | [**string**] | пароль-идентификатор папки | defaults to undefined|


### Return type

**Array<string>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**404** | папка не найдена |  -  |
|**5XX** | внутренняя ошибка сервера |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **foldersFolderPost**
> foldersFolderPost()

записывает файлы переданные в форме в указанную папку

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let folder: string; //пароль-идентификатор папки (default to undefined)
let files: Array<File>; // (default to undefined)

const { status, data } = await apiInstance.foldersFolderPost(
    folder,
    files
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **folder** | [**string**] | пароль-идентификатор папки | defaults to undefined|
| **files** | **Array&lt;File&gt;** |  | defaults to undefined|


### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**404** | запрошенная папка не существует |  -  |
|**413** | объем данных в форме превышает лимит |  -  |
|**5XX** | внутренняя ошибка сервера |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **foldersPushPost**
> Folder foldersPushPost()

создаёт новую папку с полученными файлами и возвращает информацию о ней

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let files: Array<File>; // (default to undefined)

const { status, data } = await apiInstance.foldersPushPost(
    files
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **files** | **Array&lt;File&gt;** |  | defaults to undefined|


### Return type

**Folder**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**413** | объем данных в форме превышает лимит |  -  |
|**5XX** | внутренняя ошибка сервера |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

