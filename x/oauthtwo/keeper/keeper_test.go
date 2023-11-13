package keeper_test

import (
	"crypto/rand"
	"encoding/base64"
	"testing"

	"github.com/be-heroes/doxchain/app/apptesting"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
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

func (s *KeeperTestSuite) mockRequestToken() *types.MsgTokenRequest {
	creator := s.TestAccs[0]
	clientSecret := s.generateRandomString(32)
	grantType := types.GrantType_GRANT_TYPE_CLIENT_CREDENTIALS_GRANT

	tokenRequest := &types.MsgTokenRequest{
		Creator:      creator.String(),
		ClientSecret: clientSecret,
		GrantType:    grantType,
	}

	return tokenRequest
}

func (s *KeeperTestSuite) generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	s.Require().NoError(err)

	return base64.URLEncoding.EncodeToString(bytes)
}
