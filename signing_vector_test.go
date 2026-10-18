package protocol

import (
	"encoding/json"
	"os"
	"testing"
)

func TestSharedSigningVector(t *testing.T) {
	// The shared cross-language vector lives in the monorepo at ../fixtures, and
	// is also vendored into this module's own fixtures/ so the test passes when
	// the module is published to its standalone repo.
	raw, err := os.ReadFile("fixtures/gateway-signing.json")
	if err != nil {
		raw, err = os.ReadFile("../fixtures/gateway-signing.json")
	}
	if err != nil {
		t.Fatal(err)
	}
	var v struct {
		Method, Path, Secret, Body, Signature string
		Query                                 map[string]string
	}
	if err := json.Unmarshal(raw, &v); err != nil {
		t.Fatal(err)
	}
	if got := BodySHA256([]byte(v.Body)); got != v.Query["body_sha256"] {
		t.Fatalf("body_sha256 %s != %s", got, v.Query["body_sha256"])
	}
	cs := CanonicalString(v.Method, v.Path, CanonicalQuery(v.Query), v.Query["body_sha256"])
	if got := Sign(cs, v.Secret); got != v.Signature {
		t.Fatalf("signature %s != %s", got, v.Signature)
	}
}
