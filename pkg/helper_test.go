package pkg_test

import (
	"testing"
	"time"

	"github.com/nugrohoac/livestream/pkg"

	"github.com/stretchr/testify/require"
)

func TestEncodeCursor(t *testing.T) {
	timeNow := time.Now()
	encodedCursor, err := pkg.EncodeCursor(timeNow)
	require.NoError(t, err)
	require.NotEmpty(t, encodedCursor)
}

func TestDecodeCursor(t *testing.T) {
	timeNow := time.Now()
	encodedCursor, err := pkg.EncodeCursor(timeNow)
	require.NoError(t, err)
	require.NotEmpty(t, encodedCursor)

	timeDecode, err := pkg.DecodeCursor(encodedCursor)
	require.NoError(t, err)
	require.Equal(t, timeNow.Year(), timeDecode.Year())
	require.Equal(t, timeNow.Month(), timeDecode.Month())
	require.Equal(t, timeNow.Day(), timeDecode.Day())
	require.Equal(t, timeNow.Hour(), timeDecode.Hour())
	require.Equal(t, timeNow.Minute(), timeDecode.Minute())
	require.Equal(t, timeNow.Second(), timeDecode.Second())
}
