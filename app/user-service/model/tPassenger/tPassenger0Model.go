package tPassenger

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TPassenger0Model = (*customTPassenger0Model)(nil)

type (
	// TPassenger0Model is an interface to be customized, add more methods here,
	// and implement the added methods in customTPassenger0Model.
	TPassenger0Model interface {
		tPassenger0Model
	}

	customTPassenger0Model struct {
		*defaultTPassenger0Model
	}
)

// NewTPassenger0Model returns a model for the database table.
func NewTPassenger0Model(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TPassenger0Model {
	return &customTPassenger0Model{
		defaultTPassenger0Model: newTPassenger0Model(conn, c, opts...),
	}
}
