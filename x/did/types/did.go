package types

import (
	"fmt"
)

func (did *Did) GetFullyQualifiedDidIdentifier() string {
	return fmt.Sprintf("did:%s:%s", did.MethodName, did.MethodId)
}
