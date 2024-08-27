package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// maxResourceNameLength represents the maximum allowed length for resource names
const maxResourceNameLength = 64

// shortenResourceName shortens the resource name if it exceeds maxResourceNameLength.
// It replaces the last 5 characters with the first 5 characters of a hash (base64 encoded).
func ResourceNameShortener(name string) string {
	if len(name) <= maxResourceNameLength {
		return name
	}

	// Create a SHA-256 hash of the name
	hash := sha256.Sum256([]byte(name))

	// Convert the hash to a base64 string
	hashBase64 := base64.RawURLEncoding.EncodeToString(hash[:])

	// Truncate the name and replace the last 5 characters with the first 5 characters of the hash
	truncatedName := name[:maxResourceNameLength-6] + "-" + hashBase64[:5]

	return truncatedName
}
func main() {
	name := "ephemeral-persistent-volume-claim-for-application-persistent-volume-claim"
	fmt.Println(ResourceNameShortener(name))
}
