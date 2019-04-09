# \KeysApi

All URIs are relative to *https://us-south.kms.cloud.ibm.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionOnKey**](KeysApi.md#ActionOnKey) | **Post** /keys/{id} | Invoke an action on a key
[**CreateKey**](KeysApi.md#CreateKey) | **Post** /keys | Create a new key
[**DeleteKey**](KeysApi.md#DeleteKey) | **Delete** /keys/{id} | Delete a key by ID
[**GetKey**](KeysApi.md#GetKey) | **Get** /keys/{id} | Retrieve a key by ID
[**GetKeyCollectionMetadata**](KeysApi.md#GetKeyCollectionMetadata) | **Head** /keys | Retrieve the number of keys
[**GetKeys**](KeysApi.md#GetKeys) | **Get** /keys | Retrieve a list of keys


# **ActionOnKey**
> KeyAction ActionOnKey(ctx, authorization, bluemixInstance, id, action, keyAction, optional)
Invoke an action on a key

Invokes an action, such as a `wrap`, `unwrap`, or `rotate` operation, on a specified root key.       **Note:** When you unwrap a wrapped data encryption key (WDEK) by using a rotated root key, the service returns a new ciphertext in the response entity-body. Each ciphertext remains available for `unwrap` actions. If you unwrap a DEK with a previous ciphertext, the service also returns the latest ciphertext in the response. Use the latest  ciphertext for future unwrap operations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Your IBM Cloud access token. | 
  **bluemixInstance** | [**string**](.md)| The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace &#x60;Bluemix-Instance&#x60; with the &#x60;Bluemix-Org&#x60; and &#x60;Bluemix-Space&#x60; parameters. | 
  **id** | [**string**](.md)| The root key that is used as the wrapping key. It must be a v4 UUID for an active key. | 
  **action** | **string**| The action to perform on the specified key. | 
  **keyAction** | [**KeyAction**](KeyAction.md)| The request to perform a key wrap operation. | 
 **optional** | ***ActionOnKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionOnKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **bluemixSpace** | [**optional.Interface of string**](.md)| The IBM Cloud space v4 UUID. | 
 **bluemixOrg** | [**optional.Interface of string**](.md)| The IBM Cloud organization v4 UUID. | 
 **correlationId** | [**optional.Interface of string**](.md)| The v4 UUID used to correlate and track transactions. | 
 **prefer** | **optional.String**| Alters server behavior for POST or DELETE operations. A header with &#x60;return&#x3D;minimal&#x60; causes the service to  return only the key identifier, or metadata. A header containing &#x60;return&#x3D;representation&#x60; returns both the key  material and metadata in the response entity-body. If the key has been designated as a root key, the  system cannot return the key material.      **Note:** During POST operations, Key Protect may not immediately return the key material due to key generation time. To retrieve the key material, you can perform a subsequent &#x60;GET /keys/{id}&#x60; request. | 

### Return type

[**KeyAction**](KeyAction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateKey**
> CreateKeyCollection CreateKey(ctx, authorization, bluemixInstance, createKeyCollection, optional)
Create a new key

Creates a new key with specified key material.      Key Protect designates the resource as either a root key or a standard key based on the `extractable` value that you specify. A successful `POST /keys` operation adds the key to the service and  returns the details of the request in the response entity-body, if the Prefer header is set to  `return=representation`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Your IBM Cloud access token. | 
  **bluemixInstance** | [**string**](.md)| The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace &#x60;Bluemix-Instance&#x60; with the &#x60;Bluemix-Org&#x60; and &#x60;Bluemix-Space&#x60; parameters. | 
  **createKeyCollection** | [**CreateKeyCollection**](CreateKeyCollection.md)| The base request for creating a new key. | 
 **optional** | ***CreateKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bluemixSpace** | [**optional.Interface of string**](.md)| The IBM Cloud space v4 UUID. | 
 **bluemixOrg** | [**optional.Interface of string**](.md)| The IBM Cloud organization v4 UUID. | 
 **correlationId** | [**optional.Interface of string**](.md)| The v4 UUID used to correlate and track transactions. | 
 **prefer** | **optional.String**| Alters server behavior for POST or DELETE operations. A header with &#x60;return&#x3D;minimal&#x60; causes the service to  return only the key identifier, or metadata. A header containing &#x60;return&#x3D;representation&#x60; returns both the key  material and metadata in the response entity-body. If the key has been designated as a root key, the  system cannot return the key material.      **Note:** During POST operations, Key Protect may not immediately return the key material due to key generation time. To retrieve the key material, you can perform a subsequent &#x60;GET /keys/{id}&#x60; request. | 

### Return type

[**CreateKeyCollection**](CreateKeyCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.ibm.kms.key+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteKey**
> DeleteKeyCollection DeleteKey(ctx, authorization, bluemixInstance, id, optional)
Delete a key by ID

Deletes a key by specifying the ID of the key.    **Important:** When you delete a key, you permanently shred its contents and associated data. The action cannot be reversed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Your IBM Cloud access token. | 
  **bluemixInstance** | [**string**](.md)| The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace &#x60;Bluemix-Instance&#x60; with the &#x60;Bluemix-Org&#x60; and &#x60;Bluemix-Space&#x60; parameters. | 
  **id** | [**string**](.md)| The v4 UUID that uniquely identifies the key. | 
 **optional** | ***DeleteKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bluemixSpace** | [**optional.Interface of string**](.md)| The IBM Cloud space v4 UUID. | 
 **bluemixOrg** | [**optional.Interface of string**](.md)| The IBM Cloud organization v4 UUID. | 
 **correlationId** | [**optional.Interface of string**](.md)| The v4 UUID used to correlate and track transactions. | 
 **prefer** | **optional.String**| Alters server behavior for POST or DELETE operations. A header with &#x60;return&#x3D;minimal&#x60; causes the service to  return only the key identifier, or metadata. A header containing &#x60;return&#x3D;representation&#x60; returns both the key  material and metadata in the response entity-body. If the key has been designated as a root key, the  system cannot return the key material.      **Note:** During POST operations, Key Protect may not immediately return the key material due to key generation time. To retrieve the key material, you can perform a subsequent &#x60;GET /keys/{id}&#x60; request. | 

### Return type

[**DeleteKeyCollection**](DeleteKeyCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKey**
> GetKeyCollection GetKey(ctx, authorization, bluemixInstance, id, optional)
Retrieve a key by ID

Retrieves the details of a key by specifying the ID of the key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Your IBM Cloud access token. | 
  **bluemixInstance** | [**string**](.md)| The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace &#x60;Bluemix-Instance&#x60; with the &#x60;Bluemix-Org&#x60; and &#x60;Bluemix-Space&#x60; parameters. | 
  **id** | [**string**](.md)| The v4 UUID that uniquely identifies the key. | 
 **optional** | ***GetKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bluemixSpace** | [**optional.Interface of string**](.md)| The IBM Cloud space v4 UUID. | 
 **bluemixOrg** | [**optional.Interface of string**](.md)| The IBM Cloud organization v4 UUID. | 
 **correlationId** | [**optional.Interface of string**](.md)| The v4 UUID used to correlate and track transactions. | 

### Return type

[**GetKeyCollection**](GetKeyCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeyCollectionMetadata**
> GetKeyCollectionMetadata(ctx, authorization, bluemixInstance, optional)
Retrieve the number of keys

Returns the same HTTP headers as a GET request without returning the entity-body. This operation returns the number of keys in your instance in a header called `Key-Total`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Your IBM Cloud access token. | 
  **bluemixInstance** | [**string**](.md)| The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace &#x60;Bluemix-Instance&#x60; with the &#x60;Bluemix-Org&#x60; and &#x60;Bluemix-Space&#x60; parameters. | 
 **optional** | ***GetKeyCollectionMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetKeyCollectionMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bluemixSpace** | [**optional.Interface of string**](.md)| The IBM Cloud space v4 UUID. | 
 **bluemixOrg** | [**optional.Interface of string**](.md)| The IBM Cloud organization v4 UUID. | 
 **correlationId** | [**optional.Interface of string**](.md)| The v4 UUID used to correlate and track transactions. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeys**
> ListKeyCollection GetKeys(ctx, authorization, bluemixInstance, optional)
Retrieve a list of keys

Retrieves a list of keys that are stored in your Key Protect service instance.    **Note:** `GET /keys` will not return the key material in the response body. You can retrieve the key material for a standard key with a subsequent `GET /keys/{id}` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Your IBM Cloud access token. | 
  **bluemixInstance** | [**string**](.md)| The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace &#x60;Bluemix-Instance&#x60; with the &#x60;Bluemix-Org&#x60; and &#x60;Bluemix-Space&#x60; parameters. | 
 **optional** | ***GetKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bluemixSpace** | [**optional.Interface of string**](.md)| The IBM Cloud space v4 UUID. | 
 **bluemixOrg** | [**optional.Interface of string**](.md)| The IBM Cloud organization v4 UUID. | 
 **correlationId** | [**optional.Interface of string**](.md)| The v4 UUID used to correlate and track transactions. | 
 **limit** | **optional.Int32**| The number of keys to retrieve. By default, &#x60;GET /keys&#x60; returns the first 2000 keys. To retrieve a different set of keys, use &#x60;limit&#x60; with &#x60;offset&#x60; to page through your available resources. The maximum value for &#x60;limit&#x60; is 5000.    **Usage:** If you have 20 keys in your instance, and you want to retrieve only the first 5 keys, use  &#x60;../keys?limit&#x3D;5&#x60;. | [default to 2000]
 **offset** | **optional.Int32**| The number of keys to skip. By specifying &#x60;offset&#x60;, you retrieve a subset of keys that starts with the &#x60;offset&#x60; value. Use &#x60;offset&#x60; with &#x60;limit&#x60; to page through your available resources.    **Usage:** If you have 100 keys in your instance, and you want to retrieve keys 26 through 50, use  &#x60;../keys?offset&#x3D;25&amp;limit&#x3D;50&#x60;.     | [default to 0]

### Return type

[**ListKeyCollection**](ListKeyCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

