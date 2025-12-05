// Copyright (c) 2023 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package service

import (
	"context"
	"sync/atomic"

	pb "extend-custom-task-service/pkg/pb"
)

type MyServiceServerImpl struct {
	pb.UnimplementedServiceServer
	taskExecutionCount int64
}

func NewMyServiceServer() *MyServiceServerImpl {
	return &MyServiceServerImpl{}
}

func (g *MyServiceServerImpl) GetTaskExecutionCount(
	ctx context.Context, req *pb.GetTaskExecutionCountRequest,
) (*pb.GetTaskExecutionCountResponse, error) {
	count := atomic.LoadInt64(&g.taskExecutionCount)

	return &pb.GetTaskExecutionCountResponse{
		Count: count,
	}, nil
}

func (g *MyServiceServerImpl) IncrementTaskExecutionCount() {
	atomic.AddInt64(&g.taskExecutionCount, 1)
}
