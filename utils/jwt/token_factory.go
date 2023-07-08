package jwt

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/golang-jwt/jwt"
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

func (jtf JwtTokenFactory) Create(tenant string, creator string, clientId string, expireOffSet time.Duration) *jwt.Token {
	jwtToken := jwt.New(jtf.SigningMethod)
	claims := jwtToken.Claims.(jwt.MapClaims)
	issuedAt := jtf.Context.BlockTime()

	//TODO: We will need some indexer logic to handle scenarios where people are requesting multiple tokens in the same block.
	claims["jti"] = creator + issuedAt.String()
	claims["iss"] = tenant
	claims["sub"] = creator
	//TODO: We need to figure out how to wire in the audience
	claims["aud"] = ""
	claims["iat"] = issuedAt.Unix()
	claims["exp"] = issuedAt.Add(expireOffSet).Unix()

	return jwtToken
}
