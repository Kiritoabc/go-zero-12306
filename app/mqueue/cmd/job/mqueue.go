package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-12306/app/mqueue/cmd/job/internal/config"
	"go-zero-12306/app/mqueue/cmd/job/internal/logic"
	"go-zero-12306/app/mqueue/cmd/job/internal/svc"
	"os"
)

var configFile = flag.String("f", "etc/mqueue.yml", "the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c, conf.UseEnv())
	// log
	if err := c.SetUp(); err != nil {
		panic(err)
	}
	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	cronjob := logic.NewCronJob(ctx, svcContext)
	mux := cronjob.Register()

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
