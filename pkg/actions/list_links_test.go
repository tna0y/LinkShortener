package actions

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"

	"service/pkg/entities"
)

func TestListLinkTestSuite(t *testing.T) {
	suite.Run(t, new(ListLinkTestSuite))
}

type ListLinkTestSuite struct {
	ActionsTestSuite
}

func (suite *ListLinkTestSuite) TestSuccess() {

	var argsList []CreateLinkArgs
	for i := 0; i < 10; i++ {
		args := CreateLinkArgs{
			ShortID: "hello" + strconv.Itoa(i),
			Target:  "https://google.com" + strconv.Itoa(i),
			TTL:     0,
		}
		argsList = append(argsList, args)
		_, err := suite.actions.CreateLink(suite.ctx, args, entities.NewRequester("12345"))
		suite.Require().NoError(err)
	}

	res, err := suite.actions.ListLinks(suite.ctx, entities.NewRequester("12345"))

	suite.Require().NoError(err)
	suite.Require().Equal(len(argsList), len(res))
	for i := 0; i < 10; i++ {
		suite.Require().Equal(argsList[i].Target, res[i].Target)
		suite.Require().Equal(argsList[i].ShortID, res[i].ShortID)
	}

}
