package query

import (
	"context"

	"github.com/nugrohoac/livestream/delivery/graphql/resolver"
	"github.com/nugrohoac/livestream/entity"
	_interface "github.com/nugrohoac/livestream/interface"
)

// LivestreamQuery ...
type LivestreamQuery struct {
	livestreamService _interface.LivestreamService
}

// GetLivestreams ..
func (l *LivestreamQuery) GetLivestreams(ctx context.Context, params *entity.ParamLivestreamQuery) (*resolver.LiveStreamFeedResolver, error) {
	filter := entity.LivestreamFilter{
		Num:    20,
		Cursor: "",
	}

	if params != nil {
		if params.Filter != nil {
			if params.Filter.Num != nil {
				filter.Num = *params.Filter.Num
			}

			if params.Filter.Cursor != nil {
				filter.Cursor = *params.Filter.Cursor
			}
		}
	}

	livestreams, cursor, err := l.livestreamService.Fetch(ctx, filter)
	if err != nil {
		return nil, err
	}

	livestramResolvers := make([]*resolver.LivestreamResolver, 0)

	for _, livestream := range livestreams {
		livestramResolvers = append(livestramResolvers, &resolver.LivestreamResolver{Data: livestream})
	}

	return &resolver.LiveStreamFeedResolver{
		Data:         &livestramResolvers,
		CursorHolder: cursor,
	}, err
}

// NewLivestreamQuery ...
func NewLivestreamQuery(livestreamService _interface.LivestreamService) LivestreamQuery {
	return LivestreamQuery{
		livestreamService,
	}
}
