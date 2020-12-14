package actions

import (
	"context"

	"github.com/stretchr/testify/suite"

	"service/pkg/adapters/sql"
)

type ActionsTestSuite struct {
	suite.Suite
	actions *Actions
	ctx     context.Context
}

func (suite *ActionsTestSuite) SetupTest() {
	storage, err := sql.NewSQLiteStorage(":memory:")
	suite.Require().NoError(err)

	suite.actions = NewActions(storage)

	suite.ctx = context.Background()
}
