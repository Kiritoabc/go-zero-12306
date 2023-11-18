package tOrderItemPassenger

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TOrderItemPassenger0Model = (*customTOrderItemPassenger0Model)(nil)

type (
	// TOrderItemPassenger0Model is an interface to be customized, add more methods here,
	// and implement the added methods in customTOrderItemPassenger0Model.
	TOrderItemPassenger0Model interface {
		tOrderItemPassenger0Model
	}

	customTOrderItemPassenger0Model struct {
		*defaultTOrderItemPassenger0Model
	}
)

// NewTOrderItemPassenger0Model returns a model for the database table.
func NewTOrderItemPassenger0Model(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TOrderItemPassenger0Model {
	return &customTOrderItemPassenger0Model{
		defaultTOrderItemPassenger0Model: newTOrderItemPassenger0Model(conn, c, opts...),
	}
}
