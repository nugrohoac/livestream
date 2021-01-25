package mysql

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"github.com/nugrohoac/livestream/entity"
	_interface "github.com/nugrohoac/livestream/interface"
	_const "github.com/nugrohoac/livestream/pkg/const"
)

type livestreamRepository struct {
	db *sql.DB
}

// Create ...
func (r livestreamRepository) Create(ctx context.Context, livestream entity.LiveStream) (entity.LiveStream, error) {
	livestream.ID = uuid.NewV4().String()

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
	if err != nil {
		return entity.LiveStream{}, errors.Wrap(err, "[Create Live Stream]error marshal push notif")
	}

	queryInsertLivestream, args, err := sq.Insert("live_stream").
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
		).
		Values(
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
		).
		ToSql()
	if err != nil {
		return entity.LiveStream{}, errors.Wrap(err, "[Create Live Stream] error generate sql insert live stream")
	}

	if _, err = r.db.ExecContext(ctx, queryInsertLivestream, args...); err != nil {
		return entity.LiveStream{}, errors.Wrap(err, "[Create Live Stream] error exec context store live stream")
	}

	return livestream, nil
}

func NewLiveStreamMysql(db *sql.DB) _interface.LivestreamRepository {
	return livestreamRepository{
		db: db,
	}
}
