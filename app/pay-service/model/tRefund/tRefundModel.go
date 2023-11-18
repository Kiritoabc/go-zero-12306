package tRefund

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TRefundModel = (*customTRefundModel)(nil)

type (
	// TRefundModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTRefundModel.
	TRefundModel interface {
		tRefundModel
	}

	customTRefundModel struct {
		*defaultTRefundModel
	}
)

// NewTRefundModel returns a model for the database table.
func NewTRefundModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TRefundModel {
	return &customTRefundModel{
		defaultTRefundModel: newTRefundModel(conn, c, opts...),
	}
}
