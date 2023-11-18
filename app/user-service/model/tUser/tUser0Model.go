package tUser

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TUser0Model = (*customTUser0Model)(nil)

type (
	// TUser0Model is an interface to be customized, add more methods here,
	// and implement the added methods in customTUser0Model.
	TUser0Model interface {
		tUser0Model
	}

	customTUser0Model struct {
		*defaultTUser0Model
	}
)

// NewTUser0Model returns a model for the database table.
func NewTUser0Model(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TUser0Model {
	return &customTUser0Model{
		defaultTUser0Model: newTUser0Model(conn, c, opts...),
	}
}
