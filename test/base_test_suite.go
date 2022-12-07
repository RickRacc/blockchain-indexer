package test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"go-bonotans/di/infra"
	"os"
)

//type BaseTestSuiteInternal interface {
//	SetupSuiteInternal()
//	TearDownSuiteInternal()
//}

type BaseTestSuite struct {
	suite.Suite
	Pool *sql.DB
}

func (suite *BaseTestSuite) SetupSuite() {
	var err error

	diInfra := infra.DiInfra{}
	suite.Pool, err = diInfra.ProvideDbPool()
	if err != nil {
		panic(err)
	}
}

func (suite *BaseTestSuite) TearDownSuite() {
	err := suite.Pool.Close()
	if err != nil {
		os.Exit(1)
	}
}
