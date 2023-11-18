package tCarriage

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TCarriageModel = (*customTCarriageModel)(nil)

type (
	// TCarriageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTCarriageModel.
	TCarriageModel interface {
		tCarriageModel
	}

	customTCarriageModel struct {
		*defaultTCarriageModel
	}
)

// NewTCarriageModel returns a model for the database table.
func NewTCarriageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TCarriageModel {
	return &customTCarriageModel{
		defaultTCarriageModel: newTCarriageModel(conn, c, opts...),
	}
}
