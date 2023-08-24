package types

import (
	"fmt"

	regexpUtils "github.com/be-heroes/doxchain/utils/regexp"
)

func (did *Did) SetW3CIdentifier() string {
	did.W3CIdentifier = fmt.Sprintf("%s%s%s%s%s", regexpUtils.REGEX_DID_PREFIX, regexpUtils.REGEX_DID_SEPERATOR_CHAR, did.MethodName, regexpUtils.REGEX_DID_SEPERATOR_CHAR, did.MethodId)
	return did.W3CIdentifier
}

func (did *Did) IsUserIdentifier() bool {
	matchIndexes := regexpUtils.REGEX_DOXC_ADDRESS.FindAllStringIndex(did.GetW3CIdentifier(), -1)

	return len(matchIndexes) == 2 && matchIndexes[0][1]+len(regexpUtils.REGEX_DID_SEPERATOR_CHAR) == matchIndexes[1][0]
}

func (did *Did) IsModuleIdentifier() bool {
	return len(regexpUtils.REGEX_DOXC_MODULE.FindAllStringSubmatch(did.GetW3CIdentifier(), -1)) == 1
}
