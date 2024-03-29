package test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"go-bonotans/di/infra"
	"os"
)

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

func (suite *BaseTestSuite) BeforeTest(suiteName, testName string) {
	suite.deleteAllData()
}

func (suite *BaseTestSuite) AfterTest(suiteName, testName string) {
}

func (suite *BaseTestSuite) deleteAllData() {
	_, err := suite.Pool.Exec("TRUNCATE indexer_position CASCADE")
	if err != nil {
		panic(err)
	}
	_, err = suite.Pool.Exec("TRUNCATE sequencer_position CASCADE")
	if err != nil {
		panic(err)
	}
	_, err = suite.Pool.Exec("TRUNCATE block CASCADE")
	if err != nil {
		panic(err)
	}
	_, err = suite.Pool.Exec("TRUNCATE eth_transaction CASCADE")
	if err != nil {
		panic(err)
	}
	_, err = suite.Pool.Exec("TRUNCATE transaction_payment CASCADE")
	if err != nil {
		panic(err)
	}
}
