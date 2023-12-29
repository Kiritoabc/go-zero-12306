package svc

import (
	"fmt"
	"github.com/hibiken/asynq"
	"go-zero-12306/app/mqueue/cmd/scheduler/internal/config"
	"time"
)

func newScheduler(c config.Config) *asynq.Scheduler {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(asynq.RedisClientOpt{
		Addr: c.Redis.Host,
	}, &asynq.SchedulerOpts{
		Location: location,
		EnqueueErrorHandler: func(task *asynq.Task, opts []asynq.Option, err error) {
			fmt.Printf("Scheduler EnqueueErrorHandler <<<<<<<===>>>>> err : %+v , task : %+v", err, task)
		},
	})
}
