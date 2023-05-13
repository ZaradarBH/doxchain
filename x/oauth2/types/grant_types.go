package types

type GrantType int

const (
	ClientCredentialsGrant GrantType = iota
	DeviceCodeGrant
)

func (gt GrantType) String() string {
	switch gt {
	case ClientCredentialsGrant:
		return "client_credentials"
	case DeviceCodeGrant:
		return "urn:ietf:params:oauth:grant-type:device_code"
	}

	return "unknown"
}
