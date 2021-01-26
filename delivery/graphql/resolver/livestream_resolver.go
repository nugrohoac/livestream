package resolver

import (
	"time"

	"github.com/nugrohoac/livestream/entity"
)

// LivestreamResolver ...
type LivestreamResolver struct {
	Data *entity.LiveStream
}

func (l *LivestreamResolver) ID() *string {
	return &l.Data.ID
}

func (l *LivestreamResolver) Title() *string {
	return &l.Data.Title
}

func (l *LivestreamResolver) Type() *string {
	return &l.Data.Type
}

func (l *LivestreamResolver) RoomUrl() *string {
	return &l.Data.RoomUrl
}

func (l *LivestreamResolver) Description() *string {
	return &l.Data.Description
}

func (l *LivestreamResolver) PosterUrl() *string {
	return &l.Data.PosterUrl
}

func (l *LivestreamResolver) PostVideoUrl() *string {
	return &l.Data.PostVideoUrl
}

func (l *LivestreamResolver) IsPublish() *bool {
	return &l.Data.IsPublish
}

func (l *LivestreamResolver) StartAt() *string {
	startAt := l.Data.StartAt.Format(time.RFC3339)
	return &startAt
}

func (l *LivestreamResolver) EndAt() *string {
	EndAt := l.Data.EndAt.Format(time.RFC3339)
	return &EndAt
}

func (l *LivestreamResolver) TeacherID() *int32 {
	return &l.Data.TeacherID
}

func (l *LivestreamResolver) LiveStreamBackupURL() *string {
	return &l.Data.LiveStreamBackupUrl
}

func (l *LivestreamResolver) ThumbnailVideo() *string {
	return &l.Data.ThumbnailVideo
}

func (l *LivestreamResolver) IsOnAir() *bool {
	return &l.Data.IsOnAir
}

func (l *LivestreamResolver) MaxQuestionLobby() *int32 {
	return &l.Data.MaxQuestionLobby
}

func (l *LivestreamResolver) PushNotifKelasOnline() *PushNotifResolver {
	if l.Data.PushNotif.Title == "" {
		return nil
	}

	return &PushNotifResolver{
		Data: &entity.PushNotif{
			Title:       l.Data.PushNotif.Title,
			Description: l.Data.PushNotif.Description,
		},
	}
}

func (l *LivestreamResolver) TriggerPushNotification() *bool {
	status := false
	if l.Data.PushNotif.Status == 0 {
		return &status
	}

	status = true
	return &status
}

func (l *LivestreamResolver) Ratings() *[]*RatingResolver {
	if len(l.Data.Ratings) == 0 {
		return nil
	}

	ratings := make([]*RatingResolver, 0)
	for _, rating := range l.Data.Ratings {
		ratings = append(ratings, &RatingResolver{
			Data: &entity.Rating{
				ID:           rating.ID,
				ObjectType:   rating.ObjectType,
				ObjectID:     rating.ObjectID,
				LiveStreamID: rating.LiveStreamID,
				UserID:       rating.UserID,
				Score:        rating.Score,
			},
		})
	}

	return &ratings
}
