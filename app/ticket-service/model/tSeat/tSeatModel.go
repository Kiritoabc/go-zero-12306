package tSeat

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TSeatModel = (*customTSeatModel)(nil)

type (
	// TSeatModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSeatModel.
	TSeatModel interface {
		tSeatModel
	}

	customTSeatModel struct {
		*defaultTSeatModel
	}
)

// NewTSeatModel returns a model for the database table.
func NewTSeatModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TSeatModel {
	return &customTSeatModel{
		defaultTSeatModel: newTSeatModel(conn, c, opts...),
	}
}
