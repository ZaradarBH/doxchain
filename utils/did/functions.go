package did

import (
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	regexpUtils "github.com/be-heroes/doxchain/utils/regexp"
)

func CreateModuleDidUrl(moduleName string, moduleTypeName string, creator string) (result string, err error) {
	cleanTypeName := strings.Replace(strings.Replace(moduleTypeName, "*", "", -1), ".", "_", -1)
	result = fmt.Sprintf("%s_%s%s%s", moduleName, cleanTypeName, regexpUtils.REGEX_DID_SEPERATOR_CHAR, creator)

	if !regexpUtils.REGEX_DID_URL.MatchString(result) {
		return result, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Could not generate valid Did based on input params")
	}

	return result, nil
}
