package did

import (
	didTypes "github.com/be-heroes/doxchain/x/did/types"
	regexpUtils "github.com/be-heroes/doxchain/utils/regexp"
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
	did.Creator = creator
	
	if regexpUtils.REGEX_DID_URL.MatchString(url) {
		did.Url = url
		did.MethodName = regexpUtils.REGEX_DID_METHOD_NAME.FindStringSubmatch(url)[0]
		did.MethodId = regexpUtils.REGEX_DID_METHOD_ID.FindStringSubmatch(url)[0]
		did.Path = regexpUtils.REGEX_DID_PATH.FindStringSubmatch(url)[0]
		did.Query = regexpUtils.REGEX_DID_QUERY.FindStringSubmatch(url)[0]
		did.Fragment = regexpUtils.REGEX_DID_FRAGMENT.FindStringSubmatch(url)[0]
		params := regexpUtils.REGEX_DID_PARAMS.FindStringSubmatch(url)[0]

		if len(params) > 0 {
			did.Parameters = []*didTypes.DidParameter{}
			match := regexpUtils.REGEX_DID_PARAM.FindStringSubmatch(params)

			for i, name := range regexpUtils.REGEX_DID_PARAM.SubexpNames() {
				if i > 0 && i <= len(match) {
					didParam := didTypes.DidParameter{
						Name:  name,
						Value: match[i],
					}

					did.Parameters = append(did.Parameters, &didParam)
				}
			}
		} else {
			did.MethodName = creator
			did.MethodId = creator
		}
	}

	return did
}
