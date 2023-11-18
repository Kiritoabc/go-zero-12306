package tUserReuse

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TUserReuseModel = (*customTUserReuseModel)(nil)

type (
	// TUserReuseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserReuseModel.
	TUserReuseModel interface {
		tUserReuseModel
	}

	customTUserReuseModel struct {
		*defaultTUserReuseModel
	}
)

// NewTUserReuseModel returns a model for the database table.
func NewTUserReuseModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TUserReuseModel {
	return &customTUserReuseModel{
		defaultTUserReuseModel: newTUserReuseModel(conn, c, opts...),
	}
}
