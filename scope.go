package protocol

import "strings"

// ScopeChannel namespaces an InternalName by the owning application so apps
// cannot collide on a shared Centrifugo channel. It inserts the appID after the
// Centrifugo namespace prefix (which Centrifugo routes on), turning
// "public:news" into "public:app_123.news". The namespace is preserved so
// per-namespace config (presence, cache, etc.) still applies.
//
// Phase 4 subscription/connection tokens MUST authorize these same app-scoped
// names, and every SDK must build channel names the same way.
func ScopeChannel(appID, internalName string) string {
	ns, canonical, found := strings.Cut(internalName, ":")
	if !found {
		// No namespace separator — scope the whole thing defensively.
		return appID + "." + internalName
	}
	return ns + ":" + appID + "." + canonical
}

// AppIDFromScoped is the inverse of ScopeChannel: it extracts the owning app id
// from a scoped Centrifugo channel name. "public:app_A.news" -> ("app_A", true).
// A name with no namespace separator, or no "." after the app segment, is not a
// scoped channel and returns ("", false) — callers fail closed on !ok.
func AppIDFromScoped(scoped string) (string, bool) {
	_, rest, found := strings.Cut(scoped, ":")
	if !found || rest == "" {
		return "", false
	}
	appID, _, found := strings.Cut(rest, ".")
	if !found || appID == "" {
		return "", false
	}
	return appID, true
}
