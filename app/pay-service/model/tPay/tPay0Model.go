package tPay

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TPay0Model = (*customTPay0Model)(nil)

type (
	// TPay0Model is an interface to be customized, add more methods here,
	// and implement the added methods in customTPay0Model.
	TPay0Model interface {
		tPay0Model
	}

	customTPay0Model struct {
		*defaultTPay0Model
	}
)

// NewTPay0Model returns a model for the database table.
func NewTPay0Model(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TPay0Model {
	return &customTPay0Model{
		defaultTPay0Model: newTPay0Model(conn, c, opts...),
	}
}
