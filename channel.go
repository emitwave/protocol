package protocol

import (
	"fmt"
	"strings"
)

type ChannelType string

const (
	Public           ChannelType = "public"
	Private          ChannelType = "private"
	PrivateEncrypted ChannelType = "private_encrypted"
	Presence         ChannelType = "presence"
	Cache            ChannelType = "cache"
	PrivateCache     ChannelType = "private_cache"
	PresenceCache    ChannelType = "presence_cache"
)

type ParsedChannel struct {
	Type                  ChannelType
	OriginalName          string
	CanonicalName         string
	Namespace             string
	InternalName          string
	RequiresAuthorization bool
	RequiresPresence      bool
	RequiresEncryption    bool
	UsesCache             bool
}

const MaxChannelNameLength = 200

type rule struct {
	prefix     string
	typ        ChannelType
	namespace  string
	auth       bool
	presence   bool
	encryption bool
	cache      bool
}

// longest-prefix-first
var rules = []rule{
	{"private-encrypted-", PrivateEncrypted, "encrypted", true, false, true, false},
	{"presence-cache-", PresenceCache, "presence_cache", true, true, false, true},
	{"private-cache-", PrivateCache, "private_cache", true, false, false, true},
	{"presence-", Presence, "presence", true, true, false, false},
	{"private-", Private, "private", true, false, false, false},
	{"cache-", Cache, "cache", false, false, false, true},
}

func validName(s string) bool {
	if s == "" || len(s) > MaxChannelNameLength {
		return false
	}
	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z', r >= 'a' && r <= 'z', r >= '0' && r <= '9':
		case r == '_' || r == '-' || r == '.' || r == ':':
		default:
			return false
		}
	}
	return true
}

// Parse maps an external channel name to its canonical type + Centrifugo namespace.
func Parse(name string) (ParsedChannel, error) {
	if !validName(name) {
		return ParsedChannel{}, fmt.Errorf("invalid channel name")
	}
	for _, rl := range rules {
		if strings.HasPrefix(name, rl.prefix) {
			canon := name[len(rl.prefix):]
			if canon == "" {
				return ParsedChannel{}, fmt.Errorf("empty canonical name")
			}
			return ParsedChannel{
				Type:                  rl.typ,
				OriginalName:          name,
				CanonicalName:         canon,
				Namespace:             rl.namespace,
				InternalName:          rl.namespace + ":" + canon,
				RequiresAuthorization: rl.auth,
				RequiresPresence:      rl.presence,
				RequiresEncryption:    rl.encryption,
				UsesCache:             rl.cache,
			}, nil
		}
	}
	// No prefix → public.
	return ParsedChannel{
		Type:          Public,
		OriginalName:  name,
		CanonicalName: name,
		Namespace:     "public",
		InternalName:  "public:" + name,
	}, nil
}
