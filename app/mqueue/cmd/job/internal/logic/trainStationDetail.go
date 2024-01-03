package logic

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"go-zero-12306/app/mqueue/cmd/job/internal/svc"
)

type TrainStationDetailJobHandler struct {
	svcCtx *svc.ServiceContext
}

func NewTrainStationDetailJobHandler(svcCtx *svc.ServiceContext) *TrainStationDetailJobHandler {
	return &TrainStationDetailJobHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask 站点详细信息定时任务
func (l *TrainStationDetailJobHandler) ProcessTask(ctx context.Context, task *asynq.Task) error {
	fmt.Println("TrainStationDetailJobHandler------start work")

	return nil
}
