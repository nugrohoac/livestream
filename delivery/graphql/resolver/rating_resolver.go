package resolver

import "github.com/nugrohoac/livestream/entity"

type RatingResolver struct {
	Data *entity.Rating
}

func (r *RatingResolver) ID() *string {
	return &r.Data.ID
}

func (r *RatingResolver) ObjectType() *string {
	return &r.Data.ObjectType
}

func (r *RatingResolver) ObjectID() *string {
	return &r.Data.ObjectID
}

func (r *RatingResolver) LiveStreamID() *string {
	return &r.Data.LiveStreamID
}

func (r *RatingResolver) UserID() *int32 {
	return &r.Data.UserID
}

func (r *RatingResolver) Score() *int32 {
	return &r.Data.Score
}
