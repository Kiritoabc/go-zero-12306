package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"go-zero-12306/app/mqueue/cmd/job/internal/svc"
)

type RegionTrainStationJobHandler struct {
	svcCtx *svc.ServiceContext
}

func NewRegionTrainStationHandler(svcCtx *svc.ServiceContext) *RegionTrainStationJobHandler {
	return &RegionTrainStationJobHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask 地区站点查询定时任务
func (l *RegionTrainStationJobHandler) ProcessTask(ctx context.Context, task *asynq.Task) error {

	// 解析任务参数
	// ...

	// 执行任务逻辑
	// ...

	return nil
}
