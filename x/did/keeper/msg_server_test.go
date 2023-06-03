package keeper_test

import (
	"github.com/be-heroes/doxchain/x/did/keeper"
	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// go test -v -run ^TestKeeperTestSuite/TestCreateDIDDocument$ github.com/be-heroes/doxchain/x/did/keeper
func (s *KeeperTestSuite) TestCreateDIDDocument() {
	s.SetupTest()

	creator := s.TestAccs[0]

	msgServer := keeper.NewMsgServerImpl(s.App.DidKeeper)
	msg := types.NewMsgCreateDidDocumentRequest(creator.String(), *s.GetDidDocument(creator))
	err := msg.ValidateBasic()
	s.Require().NoError(err)
	res, err := msgServer.CreateDidDocument(sdk.WrapSDKContext(s.Ctx), msg)
	s.Require().NoError(err)
	s.Require().NotNil(res)

	// check if has stored
	didDocument, found := s.App.DidKeeper.GetDidDocument(s.Ctx, msg.DidDocument.Id.GetW3CIdentifier())
	s.Require().True(found)
	s.Require().NotNil(didDocument)
}

// go test -v -run ^TestKeeperTestSuite/TestCreateDID$ github.com/be-heroes/doxchain/x/did/keeper
func (s *KeeperTestSuite) TestCreateDID() {
	s.SetupTest()

	creator := s.TestAccs[0]

	msgServer := keeper.NewMsgServerImpl(s.App.DidKeeper)
	msg := types.NewMsgCreateDidRequest(creator.String(), *s.GetDid(creator))
	err := msg.ValidateBasic()
	s.Require().NoError(err)
	res, err := msgServer.CreateDid(sdk.WrapSDKContext(s.Ctx), msg)
	s.Require().NoError(err)
	s.Require().NotNil(res)

	// check if has stored
	did, found := s.App.DidKeeper.GetDid(s.Ctx, msg.Did.GetW3CIdentifier())
	s.Require().True(found)
	s.Require().NotNil(did)
}
