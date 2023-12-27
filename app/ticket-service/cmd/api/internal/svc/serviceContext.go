package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-12306/app/ticket-service/cmd/api/internal/config"
	"go-zero-12306/app/ticket-service/cmd/rpc/ticket"
)

type ServiceContext struct {
	Config    config.Config
	TicketRpc ticket.Ticket
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		TicketRpc: ticket.NewTicket(zrpc.MustNewClient(c.TicketRpcConf)),
	}
}
