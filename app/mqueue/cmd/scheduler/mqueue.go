package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-12306/app/mqueue/cmd/scheduler/internal/config"
	"go-zero-12306/app/mqueue/cmd/scheduler/internal/logic"
	"go-zero-12306/app/mqueue/cmd/scheduler/internal/svc"
	"os"
)

var configFile = flag.String("f", "etc/mqueue.yml", "config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c, conf.UseEnv())
	logx.DisableStat()
	// log、prometheus、trace、metricsUrl.
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	mqueueScheduler := logic.NewCronScheduler(ctx, svcContext)
	mqueueScheduler.Register()

	if err := svcContext.Scheduler.Run(); err != nil {
		logx.Errorf("!!!MqueueSchedulerErr!!!  run err:%+v", err)
		os.Exit(1)
	}
}
