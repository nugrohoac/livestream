package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/nugrohoac/livestream/entity"
	"github.com/nugrohoac/livestream/interface/mocks"
	"github.com/nugrohoac/livestream/service"
	"github.com/nugrohoac/livestream/testdata"
)

func TestLivestreamService_Create(t *testing.T) {
	var livestreamResponse entity.LiveStream
	testdata.GoldenJSONUnmarshal(t, "livestream", &livestreamResponse)

	livestreamParameter := livestreamResponse
	livestreamParameter.ID = ""

	tests := map[string]struct {
		livestreamParameter entity.LiveStream
		createLivestream    testdata.FuncCaller
		expectedResponse    entity.LiveStream
		expectedErr         error
	}{
		"success": {
			livestreamParameter: livestreamParameter,
			createLivestream: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, livestreamParameter},
				Output:   []interface{}{livestreamResponse, nil},
			},
			expectedResponse: livestreamResponse,
			expectedErr:      nil,
		},
		"internal server error": {
			livestreamParameter: livestreamParameter,
			createLivestream: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, livestreamParameter},
				Output:   []interface{}{entity.LiveStream{}, errors.New("internal server error")},
			},
			expectedResponse: entity.LiveStream{},
			expectedErr:      errors.New("internal server error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			mockLivestreamRepo := new(mocks.LivestreamRepository)
			if test.createLivestream.IsCalled {
				mockLivestreamRepo.On("Create", test.createLivestream.Input...).
					Return(test.createLivestream.Output...).
					Once()
			}

			livestreamService := service.NewLivestreamService(mockLivestreamRepo)
			response, err := livestreamService.Create(context.Background(), test.livestreamParameter)
			mockLivestreamRepo.AssertExpectations(t)
			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.Equal(t, test.expectedResponse, response)
		})
	}
}
