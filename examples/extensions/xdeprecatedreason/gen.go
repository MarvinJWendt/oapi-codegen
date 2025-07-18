// Package xdeprecatedreason provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package xdeprecatedreason

// Client defines model for Client.
type Client struct {
	Id   *float32 `json:"id,omitempty"`
	Name string   `json:"name"`
}

// ClientWithExtension defines model for ClientWithExtension.
type ClientWithExtension struct {
	// Deprecated: this property has been marked as deprecated upstream, but no `x-deprecated-reason` was set
	DeprecatedWithoutReason *string  `json:"deprecated_without_reason,omitempty"`
	Id                      *float32 `json:"id,omitempty"`
	// Deprecated: Don't use because reasons
	Name string `json:"name"`
}
