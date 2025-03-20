package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// GenerateNonce creates a secure random nonce for CSP
func GenerateNonce() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic("failed to generate nonce: " + err.Error())
	}
	return base64.StdEncoding.EncodeToString(b)
}

// GetHtmxNonce returns a CSP nonce specifically for HTMX scripts
func GetHtmxNonce() string {
	return GenerateNonce()
}

// GetResponseTargetsNonce returns a CSP nonce for response-targets scripts
func GetResponseTargetsNonce() string {
	return GenerateNonce()
}

// GetCSSCSP returns the exact CSP entry for your hashed CSS file
func GetCSSCSP() string {
	// Assuming your hashed CSS path from manifest
	cssPath := GetCSSHashFile()
	return fmt.Sprintf("'self' %s", cssPath)
}
