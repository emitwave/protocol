# emitwave/protocol (Go)

Shared request-signing + channel-name protocol for the EmitWave Go SDKs.
Imported by [`github.com/emitwave/go-sdk`](https://github.com/emitwave/go-sdk).

```go
import "github.com/emitwave/protocol"
```

## Realtime authentication

Browsers authenticate through their host application's same-origin endpoints:

1. `POST /emitwave/connect-token` authenticates the host user. The backend
   signs `POST /apps/{appId}/connection_tokens` with its server credential.
2. Private and presence subscriptions call `POST /emitwave/channel-auth`.
   The backend evaluates its channel policy, then signs
   `POST /apps/{appId}/subscription_tokens`.
3. The returned subscription token is bound to the app-scoped channel.
   Encrypted channels also receive a derived per-channel key.

Public channels skip channel authorization, but every browser connection still
uses a short-lived connection token. Publishing credentials remain on trusted
servers and never grant browser subscription access.

MIT licensed.
