package protocol

// Stable machine-readable error codes (roadmap §3.8).
const (
	CodeInvalidAppKey         = "INVALID_APP_KEY"
	CodeApplicationDisabled   = "APPLICATION_DISABLED"
	CodeCredentialRevoked     = "CREDENTIAL_REVOKED"
	CodeInvalidSignature      = "INVALID_SIGNATURE"
	CodeExpiredRequest        = "EXPIRED_REQUEST"
	CodeRateLimited           = "RATE_LIMITED"
	CodePayloadTooLarge       = "PAYLOAD_TOO_LARGE"
	CodeInvalidChannel        = "INVALID_CHANNEL"
	CodeCentrifugoUnavailable = "CENTRIFUGO_UNAVAILABLE"
)
