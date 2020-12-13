package actions

import (
	"context"

	"github.com/stretchr/testify/suite"

	"service/pkg/adapters/sqlite"
)

type ActionsTestSuite struct {
	suite.Suite
	actions *Actions
	ctx     context.Context
}

func (suite *ActionsTestSuite) SetupTest() {
	storage, err := sqlite.NewSQLiteStorage(":memory:")
	suite.Require().NoError(err)

	suite.actions = NewActions(storage)

	suite.ctx = context.Background()
}
