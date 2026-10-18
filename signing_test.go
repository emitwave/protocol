package protocol

import "testing"

func TestSignVerifyRoundTrip(t *testing.T) {
	secret := "ew_sk_test_FIXTURESECRET1234567890"
	body := []byte(`{"name":"x","channel":"news","data":{}}`)
	q := CanonicalQuery(map[string]string{
		"auth_key":       "ew_pk_test_abc",
		"auth_timestamp": "1700000000",
		"auth_version":   "1.0",
		"body_sha256":    BodySHA256(body),
		"auth_signature": "SHOULD_BE_EXCLUDED",
	})
	canon := CanonicalString("POST", "/apps/app_1/events", q, BodySHA256(body))
	sig := Sign(canon, secret)
	if !Verify(canon, secret, sig) {
		t.Fatal("verify should accept a valid signature")
	}
	if Verify(canon, secret, sig+"00") {
		t.Fatal("verify must reject a tampered signature")
	}
	if Verify(canon, "wrong-secret", sig) {
		t.Fatal("verify must reject a wrong secret")
	}
}

func TestCanonicalQueryExcludesSignatureAndSorts(t *testing.T) {
	q := CanonicalQuery(map[string]string{"b": "2", "a": "1", "auth_signature": "zzz"})
	if q != "a=1&b=2" {
		t.Fatalf("got %q want a=1&b=2", q)
	}
}

func TestBodySHA256Known(t *testing.T) {
	if BodySHA256([]byte("")) != "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" {
		t.Fatal("empty-body sha256 mismatch")
	}
}
