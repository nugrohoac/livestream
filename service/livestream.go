package service

import (
	"context"

	"github.com/nugrohoac/livestream/entity"
	_interface "github.com/nugrohoac/livestream/interface"
)

type livestreamService struct {
	livestreamRepo _interface.LivestreamRepository
}

// Create is used to handle store live stream
func (l livestreamService) Create(ctx context.Context, livestream entity.LiveStream) (entity.LiveStream, error) {
	return l.livestreamRepo.Create(ctx, livestream)
}

// Fetch ...
func (l livestreamService) Fetch(ctx context.Context, filter entity.LivestreamFilter) ([]*entity.LiveStream, *string, error) {
	livestreams, _, cursor, err := l.livestreamRepo.Fetch(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	return livestreams, cursor, nil
}

// NewLivestreamService is used to initiate new instance of live stream service
func NewLivestreamService(livestreamRepo _interface.LivestreamRepository) _interface.LivestreamService {
	return livestreamService{
		livestreamRepo: livestreamRepo,
	}
}
