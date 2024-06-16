package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/dymensionxyz/dymension/v3/x/incentives/keeper"
	"github.com/osmosis-labs/osmosis/v15/app/apptesting"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	querier keeper.Querier
}

// SetupTest sets incentives parameters from the suite's context
func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.querier = keeper.NewQuerier(*suite.App.IncentivesKeeper)
	lockableDurations := suite.App.IncentivesKeeper.GetLockableDurations(suite.Ctx)
	lockableDurations = append(lockableDurations, 2*time.Second)
	suite.App.IncentivesKeeper.SetLockableDurations(suite.Ctx, lockableDurations)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
