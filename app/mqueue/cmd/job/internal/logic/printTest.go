package logic

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"go-zero-12306/app/mqueue/cmd/job/internal/svc"
)

type PrintJobHandler struct {
	svcCtx *svc.ServiceContext
}

func NewPrintJobHandler(svcCtx *svc.ServiceContext) *PrintJobHandler {
	return &PrintJobHandler{
		svcCtx: svcCtx,
	}
}

func (l *PrintJobHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {

	fmt.Println("shcedule job demo ----> hello world")

	return nil
}
