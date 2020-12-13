package actions

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"service/pkg/entities"
)

func TestGetLinkTestSuite(t *testing.T) {
	suite.Run(t, new(GetLinkTestSuite))
}

type GetLinkTestSuite struct {
	ActionsTestSuite
}

func (suite *GetLinkTestSuite) TestSuccess() {

	args := CreateLinkArgs{
		ShortID: "hello",
		Target:  "https://google.com",
		TTL:     0,
	}

	_, err := suite.actions.CreateLink(suite.ctx, args, entities.NewRequester("12345"))

	suite.Require().NoError(err)

	res, err := suite.actions.GetLink(suite.ctx, "hello", entities.NewRequester("12345"))
	suite.Require().NoError(err)

	suite.Require().Equal(args.Target, res.Target)
}

func (suite *GetLinkTestSuite) TestNotFound() {
	_, err := suite.actions.GetLink(suite.ctx, "hello", entities.NewRequester("12345"))

	suite.Require().Error(entities.ErrNotFound, err)
}

func (suite *GetLinkTestSuite) TestPermissionDenied() {
	_, err := suite.actions.CreateLink(suite.ctx, CreateLinkArgs{
		ShortID: "hello",
		Target:  "https://google.com",
		TTL:     0,
	}, entities.NewRequester("12345"))

	suite.Require().NoError(err)

	_, err = suite.actions.GetLink(suite.ctx, "hello", entities.NewRequester("1"))

	suite.Require().Error(entities.ErrPermissionDenied, err)
}
