package utils

import (
	"fmt"
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	didTypes "github.com/be-heroes/doxchain/x/did/types"
)

const (
	REGEX_DID_ID_CHAR = "[a-zA-Z0-9_.-]";
	REGEX_DID_PARAM_CHAR = "[a-zA-Z0-9_.:%-]"
	REGEX_DID_METHOD_NAME_CAPTURE_IDENTIFIER = "MethodName"
	REGEX_DID_METHOD_ID_CAPTURE_IDENTIFIER = "MethodId"
	REGEX_DID_PATH_CAPTURE_IDENTIFIER = "Path";
	REGEX_DID_QUERY_CAPTURE_IDENTIFIER = "Query"
	REGEX_DID_FRAGMENT_CAPTURE_IDENTIFIER = "Fragment"
	REGEX_DID_PARAMS_CAPTURE_IDENTIFIER = "Params"
	REGEX_DID_PARAM_NAME_CAPTURE_IDENTIFIER = "ParamName"
	REGEX_DID_PARAM_VALUE_CAPTURE_IDENTIFIER = "ParamValue"
)

var (
	REGEX_DID_METHOD_NAME = regexp.MustCompile(fmt.Sprintf("(?<%s>[a-zA-Z0-9_]+)", REGEX_DID_METHOD_NAME_CAPTURE_IDENTIFIER))
	REGEX_DID_METHOD_ID = regexp.MustCompile(fmt.Sprintf("(?<%s>%s+(:%s+)*)", REGEX_DID_METHOD_ID_CAPTURE_IDENTIFIER, REGEX_DID_ID_CHAR, REGEX_DID_ID_CHAR))
	REGEX_DID_PATH = regexp.MustCompile(fmt.Sprintf("(?<%s>\\/[^#?]*)?", REGEX_DID_PATH_CAPTURE_IDENTIFIER))
	REGEX_DID_QUERY = regexp.MustCompile(fmt.Sprintf("(?<%s>[?][^#]*)?", REGEX_DID_QUERY_CAPTURE_IDENTIFIER))
	REGEX_DID_FRAGMENT = regexp.MustCompile(fmt.Sprintf("(?<%s>\\#.*)?", REGEX_DID_FRAGMENT_CAPTURE_IDENTIFIER))
	REGEX_DID_PARAM = regexp.MustCompile(fmt.Sprintf(";(?<%s>%s+)=(?<%s>%s*)", REGEX_DID_PARAM_NAME_CAPTURE_IDENTIFIER, REGEX_DID_PARAM_CHAR, REGEX_DID_PARAM_VALUE_CAPTURE_IDENTIFIER, REGEX_DID_PARAM_CHAR))
	REGEX_DID_PARAMS = regexp.MustCompile(fmt.Sprintf("(?<%s>(%s)*)", REGEX_DID_PARAMS_CAPTURE_IDENTIFIER, REGEX_DID_PARAM))
	REGEX_DID_URL = regexp.MustCompile(fmt.Sprintf("^did:%s:%s%s%s%s%s$", REGEX_DID_METHOD_NAME, REGEX_DID_METHOD_ID, REGEX_DID_PARAMS, REGEX_DID_PATH, REGEX_DID_QUERY, REGEX_DID_FRAGMENT))
)

type DidTokenFactory struct {
	Context       *sdk.Context
}

type DidTokenFactoryOption func(didf *DidTokenFactory)

// NewDidTokenFactory initializes a new did factory.
func NewDidTokenFactory(opts ...DidTokenFactoryOption) *DidTokenFactory {
	didf := &DidTokenFactory{}

	for _, opt := range opts {
		opt(didf)
	}

	return didf
}

// Create returns a new did
func (didf DidTokenFactory) Create(creator string, url string) *didTypes.Did {
	//TODO: Master regexp package in golang and parse url string to populate remaning did fields based on regex matches
	return &didTypes.Did{
		Creator: creator,
		Url: url,
	}
}
