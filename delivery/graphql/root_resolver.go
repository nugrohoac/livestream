package graphql

import (
	"github.com/nugrohoac/livestream/delivery/graphql/mutation"
)

type rootResolver struct {
	mutation.LivestreamMutation
}

// NewRootResolver is main resolver
func NewRootResolver(livestreamMutation mutation.LivestreamMutation) *rootResolver {
	return &rootResolver{
		livestreamMutation,
	}
}
