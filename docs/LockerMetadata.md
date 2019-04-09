# LockerMetadata

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The v4 UUID used to uniquely identify the resource, as specified by RFC 4122. | [optional] 
**CreationDate** | [**time.Time**](time.Time.md) | The date the transport encryption key was created. The date format follows RFC 3339. | [optional] 
**ExpirationDate** | [**time.Time**](time.Time.md) | The date the transport encryption key expires. The date format follows RFC 3339. | [optional] 
**MaxAllowedRetrievals** | **float32** | The number of times that a transport key can be retrieved within its expiration time before it is no longer  accessible.  | [optional] [default to 1]
**RemainingRetrievals** | **float32** | The remaining number of times the transport encryption key can be retrieved before becoming inaccessible. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


