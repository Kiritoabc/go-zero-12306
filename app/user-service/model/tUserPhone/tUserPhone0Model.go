package tUserPhone

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TUserPhone0Model = (*customTUserPhone0Model)(nil)

type (
	// TUserPhone0Model is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserPhone0Model.
	TUserPhone0Model interface {
		tUserPhone0Model
	}

	customTUserPhone0Model struct {
		*defaultTUserPhone0Model
	}
)

// NewTUserPhone0Model returns a model for the database table.
func NewTUserPhone0Model(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TUserPhone0Model {
	return &customTUserPhone0Model{
		defaultTUserPhone0Model: newTUserPhone0Model(conn, c, opts...),
	}
}
