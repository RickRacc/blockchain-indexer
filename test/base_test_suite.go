package test

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"go-bonotans/di/infra"
	"os"
)

type BaseTestSuiteInternal interface {
	SetupAllSuiteInternal()
	TearDownAllSuiteInternal()
}

type BaseTestSuite struct {
	suite.Suite
	Pool *sql.DB
}

func (suite *BaseTestSuite) SetupAllSuite() {
	var err error

	diInfra := infra.DiInfra{}
	suite.Pool, err = diInfra.ProvideDbPool()
	if err != nil {
		os.Exit(1)
	}
	suite.SetupAllSuiteInternal()
}

func (suite *BaseTestSuite) TearDownAllSuite() {
	suite.TearDownAllSuiteInternal()
	err := suite.Pool.Close()
	if err != nil {
		os.Exit(1)
	}
}

func (suite *BaseTestSuite) SetupAllSuiteInternal() {

}

func (suite *BaseTestSuite) TearDownAllSuiteInternal() {

}
