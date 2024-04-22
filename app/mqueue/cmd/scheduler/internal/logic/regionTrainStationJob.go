package logic

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-12306/app/mqueue/cmd/job/jobtype"
)

func (l *MqueueScheduler) regionTrainStationScheduler() {
	task := asynq.NewTask(jobtype.RegionTrainStationJob, nil)
	// every ten minute exec
	entryID, err := l.svcCtx.Scheduler.Register("*/1 * * * *", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【settleRecordScheduler】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【regionTrainStationScheduler】 registered an  entry: %q \n", entryID)
}
