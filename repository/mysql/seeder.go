package mysql

import (
	"database/sql"
	"encoding/json"
	"strings"
	"testing"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/require"

	"github.com/nugrohoac/livestream/entity"
	_const "github.com/nugrohoac/livestream/pkg/const"
)

// SeedLivestreams ...
func SeedLivestreams(db *sql.DB, t *testing.T, livestreams []entity.LiveStream) {
	qInsert := sq.Insert("live_stream").
		Columns(
			"id",
			"title",
			"type",
			"room_url",
			"description",
			"poster_url",
			"post_video_url",
			"is_publish",
			"start_at",
			"end_at",
			"class_ids",
			"course_ids",
			"teacher_id",
			"live_stream_backup_url",
			"thumbnail_video",
			"is_on_air",
			"max_question_lobby",
			"room_id",
			"push_notif",
			"created_at",
		)

	timeNow := time.Now()
	for _, livestream := range livestreams {
		timeNow = timeNow.Add(1 * time.Second)

		livestreamType := 1
		switch livestream.Type {
		case _const.WEBINAR:
			livestreamType = 2
			break
		case _const.OBRAL_BAJU:
			livestreamType = 3
			break
		}

		bytePushNotif, err := json.Marshal(livestream.PushNotif)
		require.NoError(t, err)

		qInsert = qInsert.Values(
			livestream.ID,
			livestream.Title,
			livestreamType,
			livestream.RoomUrl,
			livestream.Description,
			livestream.PosterUrl,
			livestream.PostVideoUrl,
			livestream.IsPublish,
			livestream.StartAt,
			livestream.EndAt,
			strings.Join(livestream.ClassIDs, ","),
			strings.Join(livestream.CourseIDs, ","),
			livestream.TeacherID,
			livestream.LiveStreamBackupUrl,
			livestream.ThumbnailVideo,
			livestream.IsOnAir,
			livestream.MaxQuestionLobby,
			livestream.RoomID,
			bytePushNotif,
			timeNow,
		)
	}

	query, args, err := qInsert.ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}
