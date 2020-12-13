package actions

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"service/pkg/entities"
)

func TestDeleteLinkTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteLinkTestSuite))
}

type DeleteLinkTestSuite struct {
	ActionsTestSuite
}

func (suite *DeleteLinkTestSuite) TestSuccess() {
	_, err := suite.actions.CreateLink(suite.ctx, CreateLinkArgs{
		ShortID: "hello",
		Target:  "https://google.com",
		TTL:     0,
	}, entities.NewRequester("12345"))

	suite.Require().NoError(err)

	err = suite.actions.DeleteLink(suite.ctx,"hello", entities.NewRequester("12345"))

	suite.Require().NoError(err)
}


func (suite *DeleteLinkTestSuite) TestNotFound() {
	err := suite.actions.DeleteLink(suite.ctx,"hello", entities.NewRequester("12345"))

	suite.Require().Error(entities.ErrNotFound, err)
}


func (suite *DeleteLinkTestSuite) TestPermissionDenied() {
	_, err := suite.actions.CreateLink(suite.ctx, CreateLinkArgs{
		ShortID: "hello",
		Target:  "https://google.com",
		TTL:     0,
	}, entities.NewRequester("12345"))

	suite.Require().NoError(err)

	err = suite.actions.DeleteLink(suite.ctx,"hello", entities.NewRequester("1"))

	suite.Require().Error(entities.ErrPermissionDenied, err)
}