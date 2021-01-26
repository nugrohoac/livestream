package mysql

import (
	"database/sql"

	"github.com/sirupsen/logrus"

	// enabling file source for migration
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// TestSuite is struct for test suite
type TestSuite struct {
	suite.Suite
	DSN    string
	DBConn *sql.DB
	M      *Migration
}

var (
	mysqlDriver      = `mysql`
	mysqlHost        = `localhost`
	mysqlExposedPort = `3307`
	mysqlUser        = `root1`
	mysqlPassword    = `password123`
	mysqlDatabase    = `livestream`
)

// SetupSuite is method for setup the test suite
func (s *TestSuite) SetupSuite() {
	var err error
	s.DSN = mysqlUser + `:` + mysqlPassword + `@tcp(` + mysqlHost + `:` + mysqlExposedPort + `)/` + mysqlDatabase
	s.DBConn, err = sql.Open(mysqlDriver, s.DSN+`?parseTime=true&loc=Asia%2FJakarta`)
	require.NoError(s.T(), err)

	err = s.DBConn.Ping()
	require.NoError(s.T(), err)

	s.M, err = RunMigration(s.DSN + `?parseTime=true`)
	require.NoError(s.T(), err)

	logrus.Info("Starting to Migrate Up Data")
	errUp, okUp := s.M.Up()
	for _, element := range errUp {
		require.NoError(s.T(), element)
	}

	require.True(s.T(), okUp)
	require.Len(s.T(), errUp, 0)
}

// TearDownSuite is method which will be run when the test suite is done
func (s *TestSuite) TearDownSuite() {
	logrus.Info("Starting to Migrate Down Data")
	err, ok := s.M.Down()
	require.True(s.T(), ok)
	require.Len(s.T(), err, 0)

	errClose := s.DBConn.Close()
	require.NoError(s.T(), errClose)
}

// TearDownTest is called when starting migrate down
func (s *TestSuite) TearDownTest() {
	query := `SELECT TABLE_NAME FROM information_schema.tables WHERE table_schema='` + mysqlDatabase + `'`
	rows, err := s.DBConn.Query(query)
	require.NoError(s.T(), err)

	for rows.Next() {
		var tableName string

		if errScan := rows.Scan(&tableName); errScan != nil {
			logrus.Error("error scan : ", errScan)
		}

		if tableName == "schema_migrations" {
			continue
		}

		queryTruncate := "TRUNCATE TABLE " + tableName
		_, err = s.DBConn.Exec(queryTruncate)
		require.NoError(s.T(), err)
	}

	err = rows.Close()
	require.NoError(s.T(), err)
}
