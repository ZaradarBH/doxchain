package types

import (
	"fmt"
)

func (did *Did) GetW3CIdentifier() string {
	return fmt.Sprintf("did:%s:%s", did.MethodName, did.MethodId)
}
