package steward

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/scoir/canis/pkg/datastore/mocks"
)

var target *Steward

type AdminTestSuite struct {
	suite.Suite
	Store *mocks.Store
}

func (suite *AdminTestSuite) SetupTest() {
	suite.Store = &mocks.Store{}

	target = &Steward{store: suite.Store}
}

func TestAdminTestSuite(t *testing.T) {
	suite.Run(t, new(AdminTestSuite))
}
