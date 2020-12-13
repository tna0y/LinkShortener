package actions

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"service/pkg/entities"
)

func TestCreateLinkTestSuite(t *testing.T) {
	suite.Run(t, new(CreateLinkTestSuite))
}

type CreateLinkTestSuite struct {
	ActionsTestSuite
}

func (suite *CreateLinkTestSuite) TestSuccess() {
	_, err := suite.actions.CreateLink(suite.ctx, CreateLinkArgs{
		ShortID: "hello",
		Target:  "https://google.com",
		TTL:     0,
	}, entities.NewRequester("12345"))

	suite.Require().NoError(err)
}

func (suite *CreateLinkTestSuite) TestInvalidID() {
	_, err := suite.actions.CreateLink(suite.ctx, CreateLinkArgs{
		ShortID: "../../../",
		Target:  "https://google.com",
		TTL:     0,
	}, entities.NewRequester("12345"))

	suite.Require().Equal(entities.ErrInvalidShortID, err)
}

func (suite *CreateLinkTestSuite) TestInvalidTarget() {
	_, err := suite.actions.CreateLink(suite.ctx, CreateLinkArgs{
		ShortID: "hello",
		Target:  "google.com",
		TTL:     0,
	}, entities.NewRequester("12345"))

	suite.Require().Equal(entities.ErrInvalidTargetURL, err)
}

func (suite *CreateLinkTestSuite) TestAlreadyExists() {
	_, err := suite.actions.CreateLink(suite.ctx, CreateLinkArgs{
		ShortID: "hello",
		Target:  "https://google.com",
		TTL:     0,
	}, entities.NewRequester("12345"))

	suite.Require().NoError(err)

	_, err = suite.actions.CreateLink(suite.ctx, CreateLinkArgs{
		ShortID: "hello",
		Target:  "https://yandex.ru",
		TTL:     0,
	}, entities.NewRequester("123"))

	suite.Require().Error(entities.ErrExists, err)
}


func (suite *CreateLinkTestSuite) TestTimeout() {
	_, err := suite.actions.CreateLink(suite.ctx, CreateLinkArgs{
		ShortID: "hello",
		Target:  "https://google.com",
		TTL:     1,
	}, entities.NewRequester("12345"))

	suite.Require().NoError(err)

	time.Sleep(time.Second * 2)

	_, err = suite.actions.CreateLink(suite.ctx, CreateLinkArgs{
		ShortID: "hello",
		Target:  "https://yandex.ru",
		TTL:     0,
	}, entities.NewRequester("123"))

	suite.Require().NoError(err)
}