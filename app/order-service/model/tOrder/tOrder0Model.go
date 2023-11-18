package tOrder

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TOrder0Model = (*customTOrder0Model)(nil)

type (
	// TOrder0Model is an interface to be customized, add more methods here,
	// and implement the added methods in customTOrder0Model.
	TOrder0Model interface {
		tOrder0Model
	}

	customTOrder0Model struct {
		*defaultTOrder0Model
	}
)

// NewTOrder0Model returns a model for the database table.
func NewTOrder0Model(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TOrder0Model {
	return &customTOrder0Model{
		defaultTOrder0Model: newTOrder0Model(conn, c, opts...),
	}
}
