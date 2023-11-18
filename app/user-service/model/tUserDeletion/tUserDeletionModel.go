package tUserDeletion

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TUserDeletionModel = (*customTUserDeletionModel)(nil)

type (
	// TUserDeletionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserDeletionModel.
	TUserDeletionModel interface {
		tUserDeletionModel
	}

	customTUserDeletionModel struct {
		*defaultTUserDeletionModel
	}
)

// NewTUserDeletionModel returns a model for the database table.
func NewTUserDeletionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TUserDeletionModel {
	return &customTUserDeletionModel{
		defaultTUserDeletionModel: newTUserDeletionModel(conn, c, opts...),
	}
}
