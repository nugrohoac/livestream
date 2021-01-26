package _interface

import (
	"context"

	"github.com/nugrohoac/livestream/entity"
)

// LivestreamService is interface for live stream service at service/livestream
type LivestreamService interface {
	Create(ctx context.Context, livestream entity.LiveStream) (entity.LiveStream, error)
}

// LivestreamRepository is interface for data source of live stream at repository for all data source, mysql, http,
type LivestreamRepository interface {
	Create(ctx context.Context, livestream entity.LiveStream) (entity.LiveStream, error)
	Fetch(ctx context.Context, filter entity.LivestreamFilter) ([]entity.LiveStream, []string, string, error)
}
