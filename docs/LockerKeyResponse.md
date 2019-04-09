# LockerKeyResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The v4 UUID used to uniquely identify the resource, as specified by RFC 4122. | [optional] 
**CreationDate** | [**time.Time**](time.Time.md) | The date the transport encryption key was created. The date format follows RFC 3339. | [optional] 
**ExpirationDate** | [**time.Time**](time.Time.md) | The date the transport encryption key expires. The date format follows RFC 3339. | [optional] 
**Payload** | [***os.File**](*os.File.md) | The transport encryption key that you can use to encrypt key material before you import it into the service.      This value is a DER-encoded public key in PKIX format. Because DER encoding is a binary format, the value is base64 encoded. | [optional] 
**ImportToken** | [***os.File**](*os.File.md) | The import token that is used to verify the liveliness and integrity of a transport key. If you use a transport key to encrypt key material before you upload it to the service, you must supply the retrieved import token as a parameter in the &#x60;POST /keys&#x60; request. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


