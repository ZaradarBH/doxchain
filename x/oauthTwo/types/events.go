package types

// Treasury module event types
const (
	EventTypeTokenExpired             = "token_expired"
	EventTypeDeviceCodeExpired        = "device_code_expired"
	EventTypeAuthorizationCodeExpired = "authorization_code_expired"

	AttributeKeyIdentifier        = "identifier"
	AttributeKeyDeviceCode        = "device_code"
	AttributeKeyAuthorizationCode = "authorization_code"
)
