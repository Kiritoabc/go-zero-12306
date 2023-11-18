package tUserMail

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TUserMail0Model = (*customTUserMail0Model)(nil)

type (
	// TUserMail0Model is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserMail0Model.
	TUserMail0Model interface {
		tUserMail0Model
	}

	customTUserMail0Model struct {
		*defaultTUserMail0Model
	}
)

// NewTUserMail0Model returns a model for the database table.
func NewTUserMail0Model(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TUserMail0Model {
	return &customTUserMail0Model{
		defaultTUserMail0Model: newTUserMail0Model(conn, c, opts...),
	}
}
