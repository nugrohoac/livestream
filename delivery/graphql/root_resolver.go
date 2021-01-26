package graphql

import (
	"github.com/nugrohoac/livestream/delivery/graphql/mutation"
	"github.com/nugrohoac/livestream/delivery/graphql/query"
)

type rootResolver struct {
	mutation.LivestreamMutation
	query.LivestreamQuery
}

// NewRootResolver is main resolver
func NewRootResolver(livestreamMutation mutation.LivestreamMutation, livestreamQuery query.LivestreamQuery) *rootResolver {
	return &rootResolver{
		LivestreamMutation: livestreamMutation,
		LivestreamQuery:    livestreamQuery,
	}
}
