package protocol

import (
	"encoding/json"
	"os"
	"testing"
)

func TestParseFixtures(t *testing.T) {
	raw, err := os.ReadFile("fixtures/channels.json")
	if err != nil {
		t.Fatal(err)
	}
	var vecs []map[string]any
	if err := json.Unmarshal(raw, &vecs); err != nil {
		t.Fatal(err)
	}
	for _, v := range vecs {
		in := v["in"].(string)
		got, err := Parse(in)
		if err != nil {
			t.Fatalf("Parse(%q) error: %v", in, err)
		}
		if string(got.Type) != v["type"] {
			t.Errorf("%q type=%s want %s", in, got.Type, v["type"])
		}
		if got.Namespace != v["namespace"] {
			t.Errorf("%q ns=%s want %s", in, got.Namespace, v["namespace"])
		}
		if got.InternalName != v["internal"] {
			t.Errorf("%q internal=%s want %s", in, got.InternalName, v["internal"])
		}
		if got.RequiresAuthorization != v["auth"].(bool) {
			t.Errorf("%q auth mismatch", in)
		}
		if got.RequiresEncryption != v["encryption"].(bool) {
			t.Errorf("%q enc mismatch", in)
		}
		if got.UsesCache != v["cache"].(bool) {
			t.Errorf("%q cache mismatch", in)
		}
	}
}

func TestParseRejectsInvalid(t *testing.T) {
	for _, bad := range []string{"", "private-", "has space", "bad$char", string(make([]byte, 201))} {
		if _, err := Parse(bad); err == nil {
			t.Errorf("Parse(%q) should error", bad)
		}
	}
}
