package keeper_test

import (
	"cosmossdk.io/math"
	"github.com/be-heroes/doxchain/app"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (s *KeeperTestSuite) send(sender authtypes.AccountI, receiver sdk.AccAddress, priv cryptotypes.PrivKey, amt sdk.Coin, hasError bool) {
	msg := banktypes.NewMsgSend(sender.GetAddress(), receiver, sdk.NewCoins(amt))

	handler := s.App.MsgServiceRouter().Handler(msg)
	_, err := handler(s.Ctx, msg)
	if hasError {
		s.Require().Error(err)
	} else {
		s.Require().NoError(err)
		sender.SetSequence(sender.GetSequence() + 1)
	}
}

// go test -v -run ^TestKeeperTestSuite/TestSendWatchList$ github.com/be-heroes/doxchain/x/abs/keeper
// send and amount equal to the throttled rolling average which should pass
func (s *KeeperTestSuite) TestSendWatchList() {

	testCases := []struct {
		name     string
		amount   math.Int
		hasError bool
	}{
		{
			name:     "send an amount equal to throttled rolling average",
			amount:   s.App.AbsKeeper.GetThrottledRollingAverage(s.Ctx),
			hasError: false,
		},
		{
			name:     "send an amount greater than throttled rolling average",
			amount:   s.App.AbsKeeper.GetThrottledRollingAverage(s.Ctx).Add(sdk.NewInt(1)),
			hasError: true,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()

			sender, priv := s.CreateSigningAccount()
			sendAcc := s.App.AccountKeeper.NewAccountWithAddress(s.Ctx, sender)
			s.App.AccountKeeper.SetAccount(s.Ctx, sendAcc)
			s.FundAcc(sender, sdk.NewCoins(sdk.NewInt64Coin(app.DefaultDenomUnit, 1000000000)))
			toAddress := s.TestAccs[1]

			// entry has to be empty
			entryBefore := s.App.AbsKeeper.HasAddressWatchlist(s.Ctx, sender)
			s.False(entryBefore)

			s.send(sendAcc, toAddress, priv, sdk.NewInt64Coin(app.DefaultDenomUnit, tc.amount.Int64()), tc.hasError)

			// sender should be added to watch list
			if !tc.hasError {
				entry := s.App.AbsKeeper.HasAddressWatchlist(s.Ctx, sender)
				s.True(entry)
			}
		})
	}
}

// go test -v -run ^TestKeeperTestSuite/TestResetAfterBlockExpire$ github.com/be-heroes/doxchain/x/abs/keeper
func (s *KeeperTestSuite) TestResetAfterBlockExpire() {
	s.SetupTest()

	sender, priv := s.CreateSigningAccount()
	sendAcc := s.App.AccountKeeper.NewAccountWithAddress(s.Ctx, sender)
	s.App.AccountKeeper.SetAccount(s.Ctx, sendAcc)
	s.FundAcc(sender, sdk.NewCoins(sdk.NewInt64Coin(app.DefaultDenomUnit, 1000000000)))
	toAddress := s.TestAccs[1]

	// entry should be added to watch list
	s.send(sendAcc, toAddress, priv, sdk.NewInt64Coin(app.DefaultDenomUnit, 100), false)
	entry := s.App.AbsKeeper.GetAddressWatchlist(s.Ctx, sender)

	// will reset at expired height
	blockExpireOffset := s.App.AbsKeeper.GetBlockExpireOffset(s.Ctx)
	expiredHeight := blockExpireOffset.Int64() + int64(entry.BlockHeight)
	s.Ctx = s.Ctx.WithBlockHeight(expiredHeight)

	// entry should be removed from watch list
	s.send(sendAcc, toAddress, priv, sdk.NewInt64Coin(app.DefaultDenomUnit, 100), false)
	addressExisted := s.App.AbsKeeper.HasAddressWatchlist(s.Ctx, sender)
	s.False(addressExisted)
}
