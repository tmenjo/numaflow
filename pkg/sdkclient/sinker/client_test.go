/*
Copyright 2022 The Numaproj Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sinker

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	sinkpb "github.com/numaproj/numaflow-go/pkg/apis/proto/sink/v1"
	"github.com/numaproj/numaflow-go/pkg/apis/proto/sink/v1/sinkmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestClient_IsReady(t *testing.T) {
	var ctx = context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := sinkmock.NewMockSinkClient(ctrl)
	mockStream := sinkmock.NewMockSink_SinkFnClient(ctrl)
	mockClient.EXPECT().SinkFn(gomock.Any(), gomock.Any()).Return(mockStream, nil)
	mockClient.EXPECT().IsReady(gomock.Any(), gomock.Any()).Return(&sinkpb.ReadyResponse{Ready: true}, nil)
	mockClient.EXPECT().IsReady(gomock.Any(), gomock.Any()).Return(&sinkpb.ReadyResponse{Ready: false}, fmt.Errorf("mock connection refused"))

	testClient, err := NewFromClient(ctx, mockClient)
	assert.NoError(t, err)
	reflect.DeepEqual(testClient, &client{
		grpcClt: mockClient,
	})

	ready, err := testClient.IsReady(ctx, &emptypb.Empty{})
	assert.True(t, ready)
	assert.NoError(t, err)

	ready, err = testClient.IsReady(ctx, &emptypb.Empty{})
	assert.False(t, ready)
	assert.EqualError(t, err, "mock connection refused")
}

func TestClient_SinkFn(t *testing.T) {
	var ctx = context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSinkClient := sinkmock.NewMockSink_SinkFnClient(ctrl)
	mockSinkClient.EXPECT().Send(gomock.Any()).Return(nil).AnyTimes()
	mockSinkClient.EXPECT().Recv().Return(&sinkpb.SinkResponse{
		Results: []*sinkpb.SinkResponse_Result{
			{
				Id:     "temp-id",
				Status: sinkpb.Status_SUCCESS,
			},
		},
	}, nil)
	mockSinkClient.EXPECT().Recv().Return(&sinkpb.SinkResponse{
		Status: &sinkpb.TransmissionStatus{
			Eot: true,
		},
	}, nil)

	mockClient := sinkmock.NewMockSinkClient(ctrl)
	mockClient.EXPECT().SinkFn(gomock.Any(), gomock.Any()).Return(mockSinkClient, nil)

	testClient, err := NewFromClient(ctx, mockClient)
	assert.NoError(t, err)
	reflect.DeepEqual(testClient, &client{
		grpcClt: mockClient,
	})

	response, err := testClient.SinkFn(ctx, []*sinkpb.SinkRequest{
		{
			Request: &sinkpb.SinkRequest_Request{
				Id: "temp-id",
			},
		},
	})
	assert.Equal(t, []*sinkpb.SinkResponse{
		{
			Results: []*sinkpb.SinkResponse_Result{
				{
					Id:     "temp-id",
					Status: sinkpb.Status_SUCCESS,
				},
			},
		},
	}, response)
	assert.NoError(t, err)

}
