package tOrderItem

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TOrderItem0Model = (*customTOrderItem0Model)(nil)

type (
	// TOrderItem0Model is an interface to be customized, add more methods here,
	// and implement the added methods in customTOrderItem0Model.
	TOrderItem0Model interface {
		tOrderItem0Model
	}

	customTOrderItem0Model struct {
		*defaultTOrderItem0Model
	}
)

// NewTOrderItem0Model returns a model for the database table.
func NewTOrderItem0Model(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TOrderItem0Model {
	return &customTOrderItem0Model{
		defaultTOrderItem0Model: newTOrderItem0Model(conn, c, opts...),
	}
}
