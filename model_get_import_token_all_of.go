/*
 * IBM Key Protect API
 *
 * IBM Key Protect helps you provision encrypted keys for apps across IBM Cloud. As you manage the lifecycle of your keys, you can benefit from knowing that your keys are secured by cloud-based FIPS 140-2 Level 3 hardware security modules (HSMs) that protect against theft of information. You can use the Key Protect API to store, generate, and retrieve your key material. Keys within the service can protect any type of data in your symmetric key based encryption solution.
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package keyprotect
import (
	"os"
)
// GetImportTokenAllOf struct for GetImportTokenAllOf
type GetImportTokenAllOf struct {
	// The public encryption key that you can use to encrypt key material before you import it into the service.     This value is a PEM-encoded public key in PKIX format. Because PEM encoding is a binary format, the value is base64 encoded.
	Payload *os.File `json:"payload,omitempty"`
	// The nonce value that is used to verify a key import request. Encrypt and provide the encrypted nonce value when you use `POST /keys` to securely import a key to the service.
	Nonce *os.File `json:"nonce,omitempty"`
}
