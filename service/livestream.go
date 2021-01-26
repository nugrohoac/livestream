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

// NewLivestreamService is used to initiate new instance of live stream service
func NewLivestreamService(livestreamRepo _interface.LivestreamRepository) _interface.LivestreamService {
	return livestreamService{
		livestreamRepo: livestreamRepo,
	}
}
