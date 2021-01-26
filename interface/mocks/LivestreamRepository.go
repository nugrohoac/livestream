// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/nugrohoac/livestream/entity"
	mock "github.com/stretchr/testify/mock"
)

// LivestreamRepository is an autogenerated mock type for the LivestreamRepository type
type LivestreamRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, livestream
func (_m *LivestreamRepository) Create(ctx context.Context, livestream entity.LiveStream) (entity.LiveStream, error) {
	ret := _m.Called(ctx, livestream)

	var r0 entity.LiveStream
	if rf, ok := ret.Get(0).(func(context.Context, entity.LiveStream) entity.LiveStream); ok {
		r0 = rf(ctx, livestream)
	} else {
		r0 = ret.Get(0).(entity.LiveStream)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.LiveStream) error); ok {
		r1 = rf(ctx, livestream)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
