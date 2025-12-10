// Copyright (c) 2025 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package service

import (
	"context"

	ts "extend-custom-task-service/pkg/pb/task_scheduler"

	"github.com/sirupsen/logrus"
)

type TaskSchedulerServiceImpl struct {
	ts.UnimplementedScheduledTaskHandlerServer
	myService *MyServiceServerImpl
}

func NewTaskSchedulerService(myService *MyServiceServerImpl) *TaskSchedulerServiceImpl {
	return &TaskSchedulerServiceImpl{
		myService: myService,
	}
}

func (t *TaskSchedulerServiceImpl) RunScheduledTask(ctx context.Context, req *ts.ScheduledTaskRequest) (*ts.ScheduledTaskResponse, error) {
	logrus.Infof("Task started")

	// Increment task execution count in the main service
	t.myService.IncrementTaskExecutionCount()

	return &ts.ScheduledTaskResponse{
		Success:        true,
		Message:        "Task executed successfully",
		HttpStatusCode: 200,
	}, nil
}
