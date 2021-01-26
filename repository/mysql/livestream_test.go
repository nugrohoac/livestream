package mysql_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nugrohoac/livestream/testdata"

	"github.com/stretchr/testify/suite"

	"github.com/nugrohoac/livestream/entity"
	"github.com/nugrohoac/livestream/repository/mysql"
)

type livestreamTestSuite struct {
	mysql.TestSuite
}

func TestNewLiveStreamMysql(t *testing.T) {
	if testing.Short() {
		t.Skip(`Skip livestream integration test repository`)
	}

	suite.Run(t, new(livestreamTestSuite))
}

func (l *livestreamTestSuite) TestLivestreamRepository_Create() {
	ctx := context.Background()
	var livestream entity.LiveStream
	testdata.GoldenJSONUnmarshal(l.T(), "livestream", &livestream)

	livestream.ID = ""

	livestreamRepository := mysql.NewLiveStreamMysql(l.DBConn)
	storedLivestream, err := livestreamRepository.Create(ctx, livestream)

	assert.NoError(l.T(), err)
	assert.NotEmpty(l.T(), storedLivestream.ID)

	var count int
	row := l.DBConn.QueryRowContext(ctx, "SELECT COUNT(id) FROM live_stream")
	err = row.Scan(&count)
	assert.NoError(l.T(), err)
	assert.Equal(l.T(), 1, count)
}
