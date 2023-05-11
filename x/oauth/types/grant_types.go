package types

type GrantType int

const (
	ClientCredentials GrantType = iota
	DeviceCode
)

func (gt GrantType) String() string {
	switch gt {
	case ClientCredentials:
		return "client_credentials"
	case DeviceCode:
		return "device_code"
	}

	return "unknown"
}
