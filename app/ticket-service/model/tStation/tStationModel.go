package tStation

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TStationModel = (*customTStationModel)(nil)

type (
	// TStationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTStationModel.
	TStationModel interface {
		tStationModel
	}

	customTStationModel struct {
		*defaultTStationModel
	}
)

// NewTStationModel returns a model for the database table.
func NewTStationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TStationModel {
	return &customTStationModel{
		defaultTStationModel: newTStationModel(conn, c, opts...),
	}
}
