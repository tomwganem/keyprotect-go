/*
 * IBM Key Protect API
 *
 * IBM Key Protect helps you provision encrypted keys for apps across IBM Cloud. As you manage the lifecycle of your keys, you can benefit from knowing that your keys are secured by cloud-based FIPS 140-2 Level 2 hardware security modules (HSMs) that protect against theft of information. You can use the Key Protect API to store, generate, and retrieve your key material. Keys within the service can protect any type of data in your symmetric key based encryption solution.
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package keyprotect

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type LockersBetaApiService service

/*
LockersBetaApiService Retrieve a transport key by ID
Retrieves the transport encryption key with a timestamp, remaining retrieval count and an import token. The payload is the public key that you can use to encrypt your key material.      **Note:** This method will fail if &#x60;maxAllowedRetrievals&#x60; has been reached. Use &#x60;GET /lockers/{id}&#x60; to check &#x60;remainingRetrievals&#x60;. The default value is 1.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization Your IBM Cloud access token.
 * @param bluemixInstance The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace `Bluemix-Instance` with the `Bluemix-Org` and `Bluemix-Space` parameters.
 * @param id The v4 UUID that uniquely identifies the key.
 * @param optional nil or *GetLockerOpts - Optional Parameters:
 * @param "BluemixSpace" (optional.Interface of string) -  The IBM Cloud space v4 UUID.
 * @param "BluemixOrg" (optional.Interface of string) -  The IBM Cloud organization v4 UUID.
 * @param "CorrelationId" (optional.Interface of string) -  The v4 UUID used to correlate and track transactions.
 * @param "Prefer" (optional.String) -  Alters server behavior for POST or DELETE operations. A header with `return=minimal` causes the service to  return only the key identifier, or metadata. A header containing `return=representation` returns both the key  material and metadata in the response entity-body. If the key has been designated as a root key, the  system cannot return the key material.      **Note:** During POST operations, Key Protect may not immediately return the key material due to key generation time. To retrieve the key material, you can perform a subsequent `GET /keys/{id}` request.
@return LockerKeyResponse
*/

type GetLockerOpts struct {
	BluemixSpace  optional.Interface
	BluemixOrg    optional.Interface
	CorrelationId optional.Interface
	Prefer        optional.String
}

func (a *LockersBetaApiService) GetLocker(ctx context.Context, authorization string, bluemixInstance string, id string, localVarOptionals *GetLockerOpts) (LockerKeyResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LockerKeyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/lockers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/vnd.ibm.kms.locker+json", "application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	localVarHeaderParams["Bluemix-Instance"] = parameterToString(bluemixInstance, "")
	if localVarOptionals != nil && localVarOptionals.BluemixSpace.IsSet() {
		localVarHeaderParams["Bluemix-Space"] = parameterToString(localVarOptionals.BluemixSpace.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.BluemixOrg.IsSet() {
		localVarHeaderParams["Bluemix-Org"] = parameterToString(localVarOptionals.BluemixOrg.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.CorrelationId.IsSet() {
		localVarHeaderParams["Correlation-Id"] = parameterToString(localVarOptionals.CorrelationId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.Prefer.IsSet() {
		localVarHeaderParams["Prefer"] = parameterToString(localVarOptionals.Prefer.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v LockerKeyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
LockersBetaApiService Retrieve transport key metadata
Retrieves the metadata about the transport encryption key that is associated with your service instance.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization Your IBM Cloud access token.
 * @param bluemixInstance The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace `Bluemix-Instance` with the `Bluemix-Org` and `Bluemix-Space` parameters.
 * @param optional nil or *GetLockerMetadataOpts - Optional Parameters:
 * @param "BluemixSpace" (optional.Interface of string) -  The IBM Cloud space v4 UUID.
 * @param "BluemixOrg" (optional.Interface of string) -  The IBM Cloud organization v4 UUID.
 * @param "CorrelationId" (optional.Interface of string) -  The v4 UUID used to correlate and track transactions.
@return LockerMetadata
*/

type GetLockerMetadataOpts struct {
	BluemixSpace  optional.Interface
	BluemixOrg    optional.Interface
	CorrelationId optional.Interface
}

func (a *LockersBetaApiService) GetLockerMetadata(ctx context.Context, authorization string, bluemixInstance string, localVarOptionals *GetLockerMetadataOpts) (LockerMetadata, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LockerMetadata
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/lockers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/vnd.ibm.kms.locker_metadata+json", "application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	localVarHeaderParams["Bluemix-Instance"] = parameterToString(bluemixInstance, "")
	if localVarOptionals != nil && localVarOptionals.BluemixSpace.IsSet() {
		localVarHeaderParams["Bluemix-Space"] = parameterToString(localVarOptionals.BluemixSpace.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.BluemixOrg.IsSet() {
		localVarHeaderParams["Bluemix-Org"] = parameterToString(localVarOptionals.BluemixOrg.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.CorrelationId.IsSet() {
		localVarHeaderParams["Correlation-Id"] = parameterToString(localVarOptionals.CorrelationId.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v LockerMetadata
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
LockersBetaApiService Create a new transport key
Creates a transport encryption key that you can use to encrypt and import root keys into the service.     When you call &#x60;POST /lockers&#x60;, Key Protect creates an RSA key-pair from its HSMs. The service stores the encrypted private key in your service instance, and then returns the public key in the response entity-body. You can create only one transport key per service instance.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization Your IBM Cloud access token.
 * @param bluemixInstance The IBM Cloud instance ID that identifies your Key Protect service instance.       To manage resources within a specified Cloud Foundry org and space, replace `Bluemix-Instance` with the `Bluemix-Org` and `Bluemix-Space` parameters.
 * @param createLockerRequest The base request to create a transport encryption key.
 * @param optional nil or *PostLockerOpts - Optional Parameters:
 * @param "BluemixSpace" (optional.Interface of string) -  The IBM Cloud space v4 UUID.
 * @param "BluemixOrg" (optional.Interface of string) -  The IBM Cloud organization v4 UUID.
 * @param "CorrelationId" (optional.Interface of string) -  The v4 UUID used to correlate and track transactions.
@return LockerMetadata
*/

type PostLockerOpts struct {
	BluemixSpace  optional.Interface
	BluemixOrg    optional.Interface
	CorrelationId optional.Interface
}

func (a *LockersBetaApiService) PostLocker(ctx context.Context, authorization string, bluemixInstance string, createLockerRequest CreateLockerRequest, localVarOptionals *PostLockerOpts) (LockerMetadata, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LockerMetadata
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/lockers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/vnd.ibm.kms.locker_metadata+json", "application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	localVarHeaderParams["Bluemix-Instance"] = parameterToString(bluemixInstance, "")
	if localVarOptionals != nil && localVarOptionals.BluemixSpace.IsSet() {
		localVarHeaderParams["Bluemix-Space"] = parameterToString(localVarOptionals.BluemixSpace.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.BluemixOrg.IsSet() {
		localVarHeaderParams["Bluemix-Org"] = parameterToString(localVarOptionals.BluemixOrg.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.CorrelationId.IsSet() {
		localVarHeaderParams["Correlation-Id"] = parameterToString(localVarOptionals.CorrelationId.Value(), "")
	}
	// body params
	localVarPostBody = &createLockerRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v LockerMetadata
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ErrorCollection
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
