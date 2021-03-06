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
	"time"
)
// Key Properties that describe a key.
type Key struct {
	// Specifies the MIME type that represents the key resource. Currently, only the default is supported.
	Type string `json:"type"`
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	Id string `json:"id,omitempty"`
	// A unique, human-readable alias to assign to your key.    To protect your privacy, do not use personal data, such as your name or location, as an alias for your key.
	Name string `json:"name"`
	// A text field used to provide a more detailed description of the key. The maximum length is 240 characters.    To protect your privacy, do not use personal data, such as your name or location, as a description for your  key.
	Description string `json:"description,omitempty"`
	// Up to 30 tags can be created. Tags can be between 2-30 characters, including spaces. Special characters not permitted include the angled bracket, comma, colon, ampersand, and vertical pipe character (|).    To protect your privacy, do not use personal data, such as your name or location, as a tag for your key. 
	Tags []string `json:"tags,omitempty"`
	// The key state based on NIST SP 800-57. States are integers and correspond to the Pre-activation = 0, Active = 1, Deactivated = 3, and Destroyed = 5 values.
	State int32 `json:"state,omitempty"`
	// The date the key material expires. The date format follows RFC 3339. You can set an expiration date on any  key on its creation. If you create a key without specifying an expiration date, the key does not expire.
	ExpirationDate time.Time `json:"expirationDate,omitempty"`
	// A boolean value that determines whether the key material can leave the service.       If set to `false`, Key Protect designates the key as a nonextractable root key used for `wrap` and `unwrap` actions. If set to `true`, Key Protect designates the key as a standard key that you can store in your apps and services. Once set to `false` it cannot be changed to `true`.
	Extractable bool `json:"extractable,omitempty"`
	// The Cloud Resource Name (CRN) that uniquely identifies your cloud network resources.
	Crn string `json:"crn,omitempty"`
	// A boolean value that shows whether your key was originally imported or generated in Key Protect. The value is set by Key Protect based on how the key material is initially added to the service.    A value of `true` indicates that you must provide new key material when it's time to rotate the key. A value  of `false` indicates that Key Protect will generate the new key material on a `rotate` operation, as it did in key creation.
	Imported bool `json:"imported,omitempty"`
}
