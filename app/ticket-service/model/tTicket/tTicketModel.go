package tTicket

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TTicketModel = (*customTTicketModel)(nil)

type (
	// TTicketModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTicketModel.
	TTicketModel interface {
		tTicketModel
	}

	customTTicketModel struct {
		*defaultTTicketModel
	}
)

// NewTTicketModel returns a model for the database table.
func NewTTicketModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TTicketModel {
	return &customTTicketModel{
		defaultTTicketModel: newTTicketModel(conn, c, opts...),
	}
}
