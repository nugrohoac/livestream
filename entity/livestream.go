package entity

import "time"

// LiveStream as entity of live stream
type LiveStream struct {
	ID                      string    `json:"id"`
	Title                   string    `json:"title"`
	Type                    string    `json:"type"`
	RoomUrl                 string    `json:"room_url"`
	Description             string    `json:"description"`
	PosterUrl               string    `json:"poster_url"`
	PostVideoUrl            string    `json:"post_video_url"`
	IsPublish               bool      `json:"is_publish"`
	StartAt                 time.Time `json:"start_at"`
	EndAt                   time.Time `json:"end_at"`
	ClassIDs                []string  `json:"class_ids"`
	CourseIDs               []string  `json:"course_ids"`
	TeacherID               int32     `json:"teacher_id"`
	LiveStreamBackupUrl     string    `json:"live_stream_backup_url"`
	ThumbnailVideo          string    `json:"thumbnail_video"`
	IsOnAir                 bool      `json:"is_on_air"`
	MaxQuestionLobby        int32     `json:"max_question_lobby"`
	RoomID                  string    `json:"room_id"`
	PushNotif               PushNotif `json:"push_notif"`
	TriggerPushNotification bool      `json:"trigger_push_notification"`
	Ratings                 []Rating  `json:"ratings"`
}

// PushNotif is use to hold information
// status 0 is not sent
// status 1 is sent
type PushNotif struct {
	Status      int8   `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// LivestreamFilter ....
type LivestreamFilter struct {
	Num    int32
	Cursor string
}
