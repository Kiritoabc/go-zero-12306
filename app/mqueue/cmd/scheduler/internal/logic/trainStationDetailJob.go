package logic

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-12306/app/mqueue/cmd/job/jobtype"
)

func (l *MqueueScheduler) trainStationDetailScheduler() {

	task := asynq.NewTask(jobtype.TrainStationDetailJob, nil)
	// every ten minute exec
	entryID, err := l.svcCtx.Scheduler.Register("*/10 * * * *", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【settleRecordScheduler】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【trainStationDetailScheduler】 registered an  entry: %q \n", entryID)
}
