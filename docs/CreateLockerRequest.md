# CreateLockerRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | **float32** | The time in seconds from the creation of a transport key that determines how long the key remains valid.     The minimum value is &#x60;300&#x60; seconds (5 minutes), and the maximum value is &#x60;86400&#x60; (24 hours). The default value is &#x60;600&#x60; (10 minutes). | [optional] [default to 600]
**MaxAllowedRetrievals** | **float32** | The number of times that a transport key can be retrieved within its expiration time before it is no longer accessible.  | [optional] [default to 1]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


