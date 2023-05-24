package did

import (
	"fmt"
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	REGEX_DID_SEPERATOR_CHAR                 = ":"
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

func IsValidDid(didUrl string) bool {
	return REGEX_DID_URL.MatchString(didUrl)
}

func CreateModuleDidUrl(moduleName string, moduleTypeName string, creator string) (result string, err error) {
	//Clean up moduleTypeName to ensure it adheres to the did standard
	cleanTypeName := strings.Replace(strings.Replace(moduleTypeName, "*", "", -1), ".", "_", -1)
	result = fmt.Sprintf("%s_%s%s%s", moduleName, cleanTypeName, REGEX_DID_SEPERATOR_CHAR, creator)

	if !IsValidDid(result) {
		return result, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Could not generate valid Did based on input params")
	}

	return result, nil
}
