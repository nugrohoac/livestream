package mysql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/nugrohoac/livestream/pkg"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"github.com/nugrohoac/livestream/entity"
	_interface "github.com/nugrohoac/livestream/interface"
)

type livestreamRepository struct {
	db *sql.DB
}

// Create ...
func (r livestreamRepository) Create(ctx context.Context, livestream entity.LiveStream) (entity.LiveStream, error) {
	livestream.ID = uuid.NewV4().String()

	livestreamType := 1
	switch livestream.Type {
	case pkg.WEBINAR:
		livestreamType = 2
		break
	case pkg.OBRAL_BAJU:
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
			"created_at",
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
			time.Now(),
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

// Fetch return array live stream, ids livestream, and error
func (r livestreamRepository) Fetch(ctx context.Context, filter entity.LivestreamFilter) ([]*entity.LiveStream, []*string, *string, error) {
	qSelect := sq.Select(
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
	).From("live_stream").
		OrderBy("created_at DESC")

	if filter.Num != 0 {
		qSelect = qSelect.Limit(uint64(filter.Num))
	}

	if filter.Cursor != "" {
		timeCursor, err := pkg.DecodeCursor(filter.Cursor)
		if err != nil {
			return nil, nil, nil, err
		}

		qSelect = qSelect.Where("created_at < ?", timeCursor)
	}

	query, args, err := qSelect.ToSql()
	if err != nil {
		if err == sql.ErrNoRows {
			return make([]*entity.LiveStream, 0), make([]*string, 0), nil, err
		}

		return nil, nil, nil, errors.Wrap(err, "error fetch live stream")
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "error execute query")
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			fmt.Println("error close connection : ", errClose)
		}
	}()

	var (
		livestreams   = make([]*entity.LiveStream, 0)
		livestreamIDs = make([]*string, 0)
		timeCursor    time.Time
		classIDs      string
		courseIDs     string
	)

	for rows.Next() {
		var (
			livestream    entity.LiveStream
			bytePushNotif []byte
			pushNotif     entity.PushNotif
		)

		if err = rows.Scan(
			&livestream.ID,
			&livestream.Title,
			&livestream.Type,
			&livestream.RoomUrl,
			&livestream.Description,
			&livestream.PosterUrl,
			&livestream.PostVideoUrl,
			&livestream.IsPublish,
			&livestream.StartAt,
			&livestream.EndAt,
			&classIDs,
			&courseIDs,
			&livestream.TeacherID,
			&livestream.LiveStreamBackupUrl,
			&livestream.ThumbnailVideo,
			&livestream.IsOnAir,
			&livestream.MaxQuestionLobby,
			&livestream.RoomID,
			&bytePushNotif,
			&timeCursor,
		); err != nil {
			return nil, nil, nil, errors.Wrap(err, "error scan row")
		}

		if err = json.Unmarshal(bytePushNotif, &pushNotif); err != nil {
			return nil, nil, nil, err
		}

		livestream.ClassIDs = strings.Split(classIDs, ",")
		livestream.CourseIDs = strings.Split(courseIDs, ",")
		livestream.PushNotif = pushNotif
		livestreams = append(livestreams, &livestream)
		livestreamIDs = append(livestreamIDs, &livestream.ID)
	}

	cursor, err := pkg.EncodeCursor(timeCursor)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "error encode cursor")
	}

	return livestreams, livestreamIDs, &cursor, nil
}

// NewLiveStreamMysql ...
func NewLiveStreamMysql(db *sql.DB) _interface.LivestreamRepository {
	return livestreamRepository{
		db: db,
	}
}
