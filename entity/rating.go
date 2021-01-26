package entity

// Rating ...
type Rating struct {
	ID           string          `json:"id"`
	ObjectType   string          `json:"object_type"`
	ObjectID     string          `json:"object_id"`
	LiveStreamID string          `json:"live_stream_id"`
	UserID       int32           `json:"user_id"`
	Score        int32           `json:"score"`
	Comments     []RatingComment `json:"comments"`
	BaseEntity
}

// RatingComment ...
type RatingComment struct {
	ID       string `json:"id"`
	RatingID string `json:"rating_id"`
	Comment  string `json:"comment"`
	BaseEntity
}
