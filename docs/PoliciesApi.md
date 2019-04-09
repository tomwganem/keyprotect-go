# \PoliciesApi

All URIs are relative to *https://us-south.kms.cloud.ibm.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPolicy**](PoliciesApi.md#GetPolicy) | **Get** /keys/{id}/policies | Retrieve a list of policies
[**PutPolicy**](PoliciesApi.md#PutPolicy) | **Put** /keys/{id}/policies | Replace an existing policy


# **GetPolicy**
> CreatePolicyCollection GetPolicy(ctx, authorization, bluemixInstance, id, optional)
Retrieve a list of policies

Retrieves a list of policies that are associated with a specified key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Your IBM Cloud access token. | 
  **bluemixInstance** | [**string**](.md)| The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace &#x60;Bluemix-Instance&#x60; with the &#x60;Bluemix-Org&#x60; and &#x60;Bluemix-Space&#x60; parameters. | 
  **id** | [**string**](.md)| The v4 UUID that uniquely identifies the key. | 
 **optional** | ***GetPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **correlationId** | [**optional.Interface of string**](.md)| The v4 UUID used to correlate and track transactions. | 

### Return type

[**CreatePolicyCollection**](CreatePolicyCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPolicy**
> CreatePolicyCollection PutPolicy(ctx, authorization, bluemixInstance, id, createPolicyCollection, optional)
Replace an existing policy

Replaces the policy that is associated with a specified key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Your IBM Cloud access token. | 
  **bluemixInstance** | [**string**](.md)| The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace &#x60;Bluemix-Instance&#x60; with the &#x60;Bluemix-Org&#x60; and &#x60;Bluemix-Space&#x60; parameters. | 
  **id** | [**string**](.md)| The v4 UUID that uniquely identifies the key. | 
  **createPolicyCollection** | [**CreatePolicyCollection**](CreatePolicyCollection.md)| The base request for creating a new policies resource. | 
 **optional** | ***PutPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationId** | [**optional.Interface of string**](.md)| The v4 UUID used to correlate and track transactions. | 
 **prefer** | **optional.String**| Alters server behavior for POST or DELETE operations. A header with &#x60;return&#x3D;minimal&#x60; causes the service to  return only the key identifier, or metadata. A header containing &#x60;return&#x3D;representation&#x60; returns both the key  material and metadata in the response entity-body. If the key has been designated as a root key, the  system cannot return the key material.      **Note:** During POST operations, Key Protect may not immediately return the key material due to key generation time. To retrieve the key material, you can perform a subsequent &#x60;GET /keys/{id}&#x60; request. | 

### Return type

[**CreatePolicyCollection**](CreatePolicyCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.ibm.kms.key+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

