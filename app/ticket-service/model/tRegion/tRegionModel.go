package tRegion

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TRegionModel = (*customTRegionModel)(nil)

type (
	// TRegionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTRegionModel.
	TRegionModel interface {
		tRegionModel
	}

	customTRegionModel struct {
		*defaultTRegionModel
	}
)

// NewTRegionModel returns a model for the database table.
func NewTRegionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TRegionModel {
	return &customTRegionModel{
		defaultTRegionModel: newTRegionModel(conn, c, opts...),
	}
}
