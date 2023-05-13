package utils

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JwtTokenFactory struct {
	Context       *sdk.Context
	SigningMethod jwt.SigningMethod
}

type JwtTokenFactoryOption func(jtf *JwtTokenFactory)

func WithContext(ctx *sdk.Context) JwtTokenFactoryOption {
	return func(jtf *JwtTokenFactory) {
		jtf.Context = ctx
	}
}

func WithSigningMethod(method jwt.SigningMethod) JwtTokenFactoryOption {
	return func(jtf *JwtTokenFactory) {
		jtf.SigningMethod = method
	}
}

// NewJwtTokenFactory initializes a new token factory.
func NewJwtTokenFactory(opts ...JwtTokenFactoryOption) *JwtTokenFactory {
	jtf := &JwtTokenFactory{}

	for _, opt := range opts {
		opt(jtf)
	}

	if jtf.SigningMethod == nil {
		jtf.SigningMethod = jwt.SigningMethodHS256
	}

	return jtf
}

// Create returns a new jwt token with the configured signing method. Defaults to HS256
func (jtf JwtTokenFactory) Create(tenant string, creator string, clientId string, expireOffSet time.Duration) *jwt.Token {
	jwtToken := jwt.New(jtf.SigningMethod)
	claims := jwtToken.Claims.(jwt.MapClaims)
	issuedAt := jtf.Context.BlockTime()

	claims["jti"] = strings.Replace(uuid.New().String(), "-", "", -1)
	claims["iss"] = tenant
	claims["sub"] = creator
	//TODO: Decide on best strategy for infering the audience based on the available state
	claims["aud"] = clientId	
	claims["iat"] = issuedAt.Unix()
	claims["exp"] = issuedAt.Add(expireOffSet).Unix()

	return jwtToken
}
