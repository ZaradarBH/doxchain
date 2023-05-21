package did

import (
	didTypes "github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DidTokenFactory struct {
	Context *sdk.Context
}

type DidTokenFactoryOption func(didf *DidTokenFactory)

func NewDidTokenFactory(opts ...DidTokenFactoryOption) *DidTokenFactory {
	didf := &DidTokenFactory{}

	for _, opt := range opts {
		opt(didf)
	}

	return didf
}

func (didf DidTokenFactory) Create(creator string, url string) (did *didTypes.Did) {
	if IsValidDid(url) {
		did.Creator = creator
		did.Url = url
		did.MethodName = REGEX_DID_METHOD_NAME.FindStringSubmatch(url)[0]
		did.MethodId = REGEX_DID_METHOD_ID.FindStringSubmatch(url)[0]
		did.Path = REGEX_DID_PATH.FindStringSubmatch(url)[0]
		did.Query = REGEX_DID_QUERY.FindStringSubmatch(url)[0]
		did.Fragment = REGEX_DID_FRAGMENT.FindStringSubmatch(url)[0]
		params := REGEX_DID_PARAMS.FindStringSubmatch(url)[0]

		if len(params) > 0 {
			did.Parameters = []*didTypes.DidParameter{}
			match := REGEX_DID_PARAM.FindStringSubmatch(params)

			for i, name := range REGEX_DID_PARAM.SubexpNames() {
				if i > 0 && i <= len(match) {
					didParam := didTypes.DidParameter{
						Name:  name,
						Value: match[i],
					}

					did.Parameters = append(did.Parameters, &didParam)
				}
			}
		}
	}

	return did
}
