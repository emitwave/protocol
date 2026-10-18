package protocol

import "testing"

func TestScopeChannelInsertsAppID(t *testing.T) {
	pc, _ := Parse("news")
	got := ScopeChannel("app_123", pc.InternalName)
	if got != "public:app_123.news" {
		t.Fatalf("got %q want public:app_123.news", got)
	}
}

func TestScopeChannelIsolatesApps(t *testing.T) {
	pc, _ := Parse("news")
	a := ScopeChannel("app_A", pc.InternalName)
	b := ScopeChannel("app_B", pc.InternalName)
	if a == b {
		t.Fatalf("two apps must get distinct channels, both %q", a)
	}
}

func TestScopeChannelPreservesNamespace(t *testing.T) {
	pc, _ := Parse("private-encrypted-chat.9") // namespace "encrypted"
	got := ScopeChannel("app_1", pc.InternalName)
	if got != "encrypted:app_1.chat.9" {
		t.Fatalf("got %q want encrypted:app_1.chat.9", got)
	}
}

func TestAppIDFromScoped(t *testing.T) {
	cases := []struct {
		in    string
		appID string
		ok    bool
	}{
		{"public:app_A.news", "app_A", true},
		{"cache:app_1.stock.AAPL", "app_1", true},
		{"presence:app_2.room.1", "app_2", true},
		{"public:app_A.", "app_A", true}, // empty canonical still has an app segment
		{"nocolon", "", false},
		{"public:app_A", "", false}, // no dot after app
		{"", "", false},
		{"public:", "", false},
	}
	for _, c := range cases {
		gotID, gotOK := AppIDFromScoped(c.in)
		if gotID != c.appID || gotOK != c.ok {
			t.Errorf("AppIDFromScoped(%q) = (%q,%v), want (%q,%v)", c.in, gotID, gotOK, c.appID, c.ok)
		}
	}
}
