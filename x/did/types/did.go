package types

import (
	"fmt"
	"regexp"
)

var (
	REGEX_BECH32 = regexp.MustCompile("^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$")
)

func (did *Did) GetW3CIdentifier() string {
	return fmt.Sprintf("did:%s:%s", did.MethodName, did.MethodId)
}

func (did *Did) IsUserIdentifier() bool {	
	w3cIdentifier := did.GetW3CIdentifier()

	return len(REGEX_BECH32.FindStringSubmatch(w3cIdentifier)) > 0 && len(REGEX_BECH32.FindStringSubmatch(w3cIdentifier)) > 0
}
