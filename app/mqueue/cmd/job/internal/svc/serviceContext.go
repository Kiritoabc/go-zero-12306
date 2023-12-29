package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-12306/app/mqueue/cmd/job/internal/config"
	"go-zero-12306/app/ticket-service/cmd/rpc/ticket"
)
import "github.com/hibiken/asynq"

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	TicketRpc   ticket.Ticket
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
		TicketRpc:   ticket.NewTicket(zrpc.MustNewClient(c.TicketRpcConf)),
	}
}
