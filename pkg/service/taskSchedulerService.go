// Copyright (c) 2025 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package service

import (
	"io"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	ts "extend-custom-task-service/pkg/pb/generic/task_scheduler/v1"
)

type TaskSchedulerServiceImpl struct {
	ts.UnimplementedTaskSchedulerServiceServer
	myService *MyServiceServerImpl
}

func NewTaskSchedulerService(myService *MyServiceServerImpl) *TaskSchedulerServiceImpl {
	return &TaskSchedulerServiceImpl{
		myService: myService,
	}
}

func (s *TaskSchedulerServiceImpl) OnJobTriggered(
	stream grpc.BidiStreamingServer[ts.ExecutionContext, ts.Response],
) error {
	for {
		// Receive ExecutionContext from the stream
		execCtx, err := stream.Recv()
		if err == io.EOF {
			// Client closed the stream
			return nil
		}
		if err != nil {
			logrus.Errorf("Error receiving ExecutionContext: %v", err)
			return err
		}

		// Handle different message types
		switch execCtx.MessageType {
		case ts.MessageType_TASK_START:
			// Increment task execution count when a task starts
			s.myService.IncrementTaskExecutionCount()
			logrus.Infof("Task started with cron expression: %s", execCtx.CronExpression)
		case ts.MessageType_HEART_BEAT:
			// Handle heartbeat - just log it
			logrus.Infof("Received heartbeat with cron expression: %s", execCtx.CronExpression)
		default:
			logrus.Warnf("Unknown message type: %v", execCtx.MessageType)
		}

		// Send a response back
		response := &ts.Response{}
		if err := stream.Send(response); err != nil {
			logrus.Errorf("Error sending response: %v", err)
			return err
		}
	}
}
