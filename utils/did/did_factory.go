package did

import (
	"fmt"
	"regexp"

	didTypes "github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	REGEX_DID_SEPERATOR_CHAR				 = "@"
	REGEX_DID_ID_CHAR                        = "[a-zA-Z0-9_.-]"
	REGEX_DID_PARAM_CHAR                     = "[a-zA-Z0-9_.:%-]"
	REGEX_DID_METHOD_NAME_CAPTURE_IDENTIFIER = "MethodName"
	REGEX_DID_METHOD_ID_CAPTURE_IDENTIFIER   = "MethodId"
	REGEX_DID_PATH_CAPTURE_IDENTIFIER        = "Path"
	REGEX_DID_QUERY_CAPTURE_IDENTIFIER       = "Query"
	REGEX_DID_FRAGMENT_CAPTURE_IDENTIFIER    = "Fragment"
	REGEX_DID_PARAMS_CAPTURE_IDENTIFIER      = "Params"
	REGEX_DID_PARAM_NAME_CAPTURE_IDENTIFIER  = "ParamName"
	REGEX_DID_PARAM_VALUE_CAPTURE_IDENTIFIER = "ParamValue"
)

var (
	REGEX_DID_METHOD_NAME = regexp.MustCompile(fmt.Sprintf("(?P<%s>[a-zA-Z0-9_]+)", REGEX_DID_METHOD_NAME_CAPTURE_IDENTIFIER))
	REGEX_DID_METHOD_ID   = regexp.MustCompile(fmt.Sprintf("(?P<%s>%s+(:%s+)*)", REGEX_DID_METHOD_ID_CAPTURE_IDENTIFIER, REGEX_DID_ID_CHAR, REGEX_DID_ID_CHAR))
	REGEX_DID_PATH        = regexp.MustCompile(fmt.Sprintf("(?P<%s>\\/[^#?]*)?", REGEX_DID_PATH_CAPTURE_IDENTIFIER))
	REGEX_DID_QUERY       = regexp.MustCompile(fmt.Sprintf("(?P<%s>[?][^#]*)?", REGEX_DID_QUERY_CAPTURE_IDENTIFIER))
	REGEX_DID_FRAGMENT    = regexp.MustCompile(fmt.Sprintf("(?P<%s>\\#.*)?", REGEX_DID_FRAGMENT_CAPTURE_IDENTIFIER))
	REGEX_DID_PARAM       = regexp.MustCompile(fmt.Sprintf(";(?P<%s>%s+)=(?P<%s>%s*)", REGEX_DID_PARAM_NAME_CAPTURE_IDENTIFIER, REGEX_DID_PARAM_CHAR, REGEX_DID_PARAM_VALUE_CAPTURE_IDENTIFIER, REGEX_DID_PARAM_CHAR))
	REGEX_DID_PARAMS      = regexp.MustCompile(fmt.Sprintf("(?P<%s>(%s)*)", REGEX_DID_PARAMS_CAPTURE_IDENTIFIER, REGEX_DID_PARAM))
	REGEX_DID_URL         = regexp.MustCompile(fmt.Sprintf("^did:%s:%s%s%s%s%s$", REGEX_DID_METHOD_NAME, REGEX_DID_METHOD_ID, REGEX_DID_PARAMS, REGEX_DID_PATH, REGEX_DID_QUERY, REGEX_DID_FRAGMENT))
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

func (didf DidTokenFactory) Create(creator string, url string) *didTypes.Did {
	did := &didTypes.Did{
		Creator: creator,
	}

	if REGEX_DID_URL.MatchString(url) {
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
	} else {
		did.MethodName = creator
		did.MethodId = creator
		did.Path = creator
		did.Query = creator
		did.Fragment = creator		
		did.Url = fmt.Sprintf("did:%s:%s/%s?%s#%s", did.MethodName, did.MethodId, did.Path, did.Query, did.Fragment)
	}

	return did
}
