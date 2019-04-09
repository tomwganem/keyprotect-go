# \LockersBetaApi

All URIs are relative to *https://us-south.kms.cloud.ibm.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocker**](LockersBetaApi.md#GetLocker) | **Get** /lockers/{id} | Retrieve a transport key by ID
[**GetLockerMetadata**](LockersBetaApi.md#GetLockerMetadata) | **Get** /lockers | Retrieve transport key metadata
[**PostLocker**](LockersBetaApi.md#PostLocker) | **Post** /lockers | Create a new transport key


# **GetLocker**
> LockerKeyResponse GetLocker(ctx, authorization, bluemixInstance, id, optional)
Retrieve a transport key by ID

Retrieves the transport encryption key with a timestamp, remaining retrieval count and an import token. The payload is the public key that you can use to encrypt your key material.      **Note:** This method will fail if `maxAllowedRetrievals` has been reached. Use `GET /lockers/{id}` to check `remainingRetrievals`. The default value is 1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Your IBM Cloud access token. | 
  **bluemixInstance** | [**string**](.md)| The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace &#x60;Bluemix-Instance&#x60; with the &#x60;Bluemix-Org&#x60; and &#x60;Bluemix-Space&#x60; parameters. | 
  **id** | [**string**](.md)| The v4 UUID that uniquely identifies the key. | 
 **optional** | ***GetLockerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetLockerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bluemixSpace** | [**optional.Interface of string**](.md)| The IBM Cloud space v4 UUID. | 
 **bluemixOrg** | [**optional.Interface of string**](.md)| The IBM Cloud organization v4 UUID. | 
 **correlationId** | [**optional.Interface of string**](.md)| The v4 UUID used to correlate and track transactions. | 
 **prefer** | **optional.String**| Alters server behavior for POST or DELETE operations. A header with &#x60;return&#x3D;minimal&#x60; causes the service to  return only the key identifier, or metadata. A header containing &#x60;return&#x3D;representation&#x60; returns both the key  material and metadata in the response entity-body. If the key has been designated as a root key, the  system cannot return the key material.      **Note:** During POST operations, Key Protect may not immediately return the key material due to key generation time. To retrieve the key material, you can perform a subsequent &#x60;GET /keys/{id}&#x60; request. | 

### Return type

[**LockerKeyResponse**](LockerKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.ibm.kms.locker+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLockerMetadata**
> LockerMetadata GetLockerMetadata(ctx, authorization, bluemixInstance, optional)
Retrieve transport key metadata

Retrieves the metadata about the transport encryption key that is associated with your service instance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Your IBM Cloud access token. | 
  **bluemixInstance** | [**string**](.md)| The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace &#x60;Bluemix-Instance&#x60; with the &#x60;Bluemix-Org&#x60; and &#x60;Bluemix-Space&#x60; parameters. | 
 **optional** | ***GetLockerMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetLockerMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bluemixSpace** | [**optional.Interface of string**](.md)| The IBM Cloud space v4 UUID. | 
 **bluemixOrg** | [**optional.Interface of string**](.md)| The IBM Cloud organization v4 UUID. | 
 **correlationId** | [**optional.Interface of string**](.md)| The v4 UUID used to correlate and track transactions. | 

### Return type

[**LockerMetadata**](LockerMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.ibm.kms.locker_metadata+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostLocker**
> LockerMetadata PostLocker(ctx, authorization, bluemixInstance, createLockerRequest, optional)
Create a new transport key

Creates a transport encryption key that you can use to encrypt and import root keys into the service.     When you call `POST /lockers`, Key Protect creates an RSA key-pair from its HSMs. The service stores the encrypted private key in your service instance, and then returns the public key in the response entity-body. You can create only one transport key per service instance. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Your IBM Cloud access token. | 
  **bluemixInstance** | [**string**](.md)| The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace &#x60;Bluemix-Instance&#x60; with the &#x60;Bluemix-Org&#x60; and &#x60;Bluemix-Space&#x60; parameters. | 
  **createLockerRequest** | [**CreateLockerRequest**](CreateLockerRequest.md)| The base request to create a transport encryption key. | 
 **optional** | ***PostLockerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostLockerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bluemixSpace** | [**optional.Interface of string**](.md)| The IBM Cloud space v4 UUID. | 
 **bluemixOrg** | [**optional.Interface of string**](.md)| The IBM Cloud organization v4 UUID. | 
 **correlationId** | [**optional.Interface of string**](.md)| The v4 UUID used to correlate and track transactions. | 

### Return type

[**LockerMetadata**](LockerMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.ibm.kms.locker_metadata+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

