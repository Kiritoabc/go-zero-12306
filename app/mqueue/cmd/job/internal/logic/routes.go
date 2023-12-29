package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"go-zero-12306/app/mqueue/cmd/job/internal/svc"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register job
func (l *CronJob) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	// scheduler job
	//mux.Handle()

	return mux
}
