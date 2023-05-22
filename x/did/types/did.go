package types

import (
	"fmt"
	"regexp"
)

var (
	//TODO: This regex is most likely wrong. Verify that it will match an actual bech32 address
	REGEX_BECH32 = regexp.MustCompile("[13][a-km-zA-HJ-NP-Z1-9]{25,34}")
)

func (did *Did) GetW3CIdentifier() string {
	return fmt.Sprintf("did:%s:%s", did.MethodName, did.MethodId)
}

func (did *Did) IsUserIdentifier() bool {	
	w3cIdentifier := did.GetW3CIdentifier()

	return len(REGEX_BECH32.FindStringSubmatch(w3cIdentifier)) == 2
}

func (did *Did) IsModuleIdentifier() bool {
	//TODO: Implement this logic. Name syntax for module DIDs is {MODULE_NAME}_{MODULETYPENAME}{REGEX_DID_SEPERATOR_CHAR}{CREATOR}
	return false
}