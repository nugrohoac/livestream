package mutation

import (
	"context"
	"time"

	"github.com/nugrohoac/livestream/delivery/graphql/resolver"

	"github.com/nugrohoac/livestream/entity"

	_interface "github.com/nugrohoac/livestream/interface"
)

// LivestreamMutation ...
type LivestreamMutation struct {
	livestreamService _interface.LivestreamService
}

type pushNotifInput struct {
	Title       *string
	Description *string
}

type livestreamInput struct {
	Title                   string
	Type                    string
	RoomURL                 string
	Description             *string
	PosterURL               *string
	IsPublish               bool
	StartAt                 string
	EndAt                   string
	PostVideoURL            *string
	TeacherID               *int32
	LiveStreamBackupURL     *string
	ThumbnailVideo          *string
	IsOnAir                 *bool
	MaxQuestionLobby        *int32
	PushNotifKelasOnline    *pushNotifInput
	TriggerPushNotification *bool
}

// CreateLiveStream is implement of mutation in file mutation.graphql
func (l *LivestreamMutation) CreateLiveStream(ctx context.Context, param livestreamInput) (*resolver.LivestreamResolver, error) {
	var livestream entity.LiveStream

	livestream.Title = param.Title
	livestream.Type = param.Type
	livestream.RoomUrl = param.RoomURL
	livestream.IsPublish = param.IsPublish
	livestream.PushNotif = entity.PushNotif{}

	if param.Description != nil {
		livestream.Description = *param.Description
	}

	if param.PosterURL != nil {
		livestream.PosterUrl = *param.PosterURL
	}

	startTime, err := time.Parse(time.RFC3339, param.StartAt)
	if err != nil {
		return nil, err
	}
	livestream.StartAt = startTime

	endTime, err := time.Parse(time.RFC3339, param.EndAt)
	if err != nil {
		return nil, err
	}
	livestream.EndAt = endTime

	if param.PostVideoURL != nil {
		livestream.PostVideoUrl = *param.PostVideoURL
	}

	if param.TeacherID != nil {
		livestream.TeacherID = *param.TeacherID
	}

	if param.LiveStreamBackupURL != nil {
		livestream.LiveStreamBackupUrl = *param.LiveStreamBackupURL
	}

	if param.ThumbnailVideo != nil {
		livestream.ThumbnailVideo = *param.ThumbnailVideo
	}

	if param.IsOnAir != nil {
		livestream.IsOnAir = *param.IsOnAir
	}

	if param.MaxQuestionLobby != nil {
		livestream.MaxQuestionLobby = *param.MaxQuestionLobby
	}

	if param.PushNotifKelasOnline != nil {
		if param.PushNotifKelasOnline.Title != nil {
			livestream.PushNotif.Title = *param.PushNotifKelasOnline.Title
		}

		if param.PushNotifKelasOnline.Title != nil {
			livestream.PushNotif.Description = *param.PushNotifKelasOnline.Description
		}
	}

	if param.TriggerPushNotification != nil {
		livestream.TriggerPushNotification = *param.TriggerPushNotification
	}

	storedLivestream, err := l.livestreamService.Create(ctx, livestream)
	if err != nil {
		return nil, err
	}

	return &resolver.LivestreamResolver{Data: &storedLivestream}, nil
}

// NewLivestreamMutation is used to create new instance of livestream mutation
func NewLivestreamMutation(livestreamService _interface.LivestreamService) LivestreamMutation {
	return LivestreamMutation{
		livestreamService: livestreamService,
	}
}
