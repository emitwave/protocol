package protocol

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
)

// BodySHA256 returns the lowercase hex SHA-256 of the raw request body.
func BodySHA256(body []byte) string {
	sum := sha256.Sum256(body)
	return hex.EncodeToString(sum[:])
}

// CanonicalQuery sorts the auth params lexically and joins k=v with &,
// excluding auth_signature itself.
func CanonicalQuery(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "auth_signature" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, k+"="+params[k])
	}
	return strings.Join(parts, "&")
}

// CanonicalString assembles METHOD\nPATH\nQUERY\nBODY_SHA256.
func CanonicalString(method, path, canonicalQuery, bodySHA256 string) string {
	return method + "\n" + path + "\n" + canonicalQuery + "\n" + bodySHA256
}

// Sign returns hex(HMAC-SHA256(canonical, secret)).
func Sign(canonical, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(canonical))
	return hex.EncodeToString(mac.Sum(nil))
}

// Verify compares a provided signature constant-time.
func Verify(canonical, secret, sig string) bool {
	expected := Sign(canonical, secret)
	return hmac.Equal([]byte(expected), []byte(sig))
}
