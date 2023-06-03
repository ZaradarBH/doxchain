package keeper_test

import (
	"testing"

	"github.com/be-heroes/doxchain/app/apptesting"
	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()

	s.queryClient = types.NewQueryClient(s.QueryHelper)
}

func (s *KeeperTestSuite) GetDid(creator sdk.AccAddress) *types.Did {
	if creator == nil {
		creator = s.CreateRandomAccounts(1)[0]
	}

	did := &types.Did{
		Creator:    creator.String(),
		Url:        "url",
		MethodName: "method",
		MethodId:   "1",
		Path:       "path",
		Fragment:   "fragment",
		Query:      "query",
		Parameters: []*types.DidParameter{
			{
				Name:  "name",
				Value: "value",
			},
		},
	}

	return did
}

func (s *KeeperTestSuite) GetVerificationMethod() *types.VerificationMethod {
	verificationMethod := &types.VerificationMethod{
		Id:         *s.GetDid(nil),
		Type:       "type",
		Controller: *s.GetDid(nil),
		KeyOneof: &types.VerificationMethod_PublicKeyJwk{
			PublicKeyJwk: "publicKeyJwk",
		},
	}

	return verificationMethod
}

func (s *KeeperTestSuite) GetVerificationRelationship() *types.VerificationRelationship {
	verificationRelationship := &types.VerificationRelationship{
		Referenced: []types.Did{
			*s.GetDid(nil),
		},
		Embedded: []types.VerificationMethod{
			*s.GetVerificationMethod(),
		},
	}

	return verificationRelationship
}

func (s *KeeperTestSuite) GetService() *types.Service {
	service := &types.Service{
		Id:              *s.GetDid(nil),
		Type:            "type",
		ServiceEndpoint: "serviceEndpoint",
	}

	return service
}

func (s *KeeperTestSuite) GetDidDocument(creator sdk.AccAddress) *types.DidDocument {
	didDocument := &types.DidDocument{
		Context:    []string{"contex"},
		Id:         *s.GetDid(creator),
		Controller: s.GetDid(nil),
		AlsoKnownAs: []*types.Did{
			s.GetDid(nil),
		},
		VerificationMethod: []*types.VerificationMethod{
			s.GetVerificationMethod(),
		},
		Authentication: []*types.VerificationRelationship{
			s.GetVerificationRelationship(),
		},
		AssertionMethod: []*types.VerificationRelationship{
			s.GetVerificationRelationship(),
		},
		KeyAgreement: []*types.VerificationRelationship{
			s.GetVerificationRelationship(),
		},
		CapabilityInvocation: []*types.VerificationRelationship{
			s.GetVerificationRelationship(),
		},
		CapabilityDelegation: []*types.VerificationRelationship{
			s.GetVerificationRelationship(),
		},
		Service: []*types.Service{
			s.GetService(),
		},
	}

	return didDocument
}
