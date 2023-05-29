package keeper_test

import (
	"testing"

	"github.com/be-heroes/doxchain/app"
	"github.com/be-heroes/doxchain/app/apptesting"
	"github.com/be-heroes/doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
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

// go test -v -run ^TestKeeperTestSuite/TestSendAddedToWatchList$ github.com/be-heroes/doxchain/x/abs/keeper
func (s *KeeperTestSuite) TestSendAddedToWatchList() {
	s.SetupTest()

	sender := s.TestAccs[0]
	s.FundAcc(sender, sdk.NewCoins(sdk.NewInt64Coin(app.DefaultDenomUnit, 1000000000)))
	toAddress := s.TestAccs[1]

	_, err := bankkeeper.NewMsgServerImpl(s.App.BankKeeper).Send(s.Ctx, banktypes.NewMsgSend(sender, toAddress, sdk.NewCoins(sdk.NewInt64Coin("udox", 1000000))))
	s.Require().NoError(err)
	// sender should be added to watch list
	res := s.App.AbsKeeper.GetAddressWatchlist(s.Ctx, sender)
	s.Require().Equal(sender.String(), res.Address)
}
