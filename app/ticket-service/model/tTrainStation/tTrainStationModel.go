package tTrainStation

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TTrainStationModel = (*customTTrainStationModel)(nil)

type (
	// TTrainStationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTrainStationModel.
	TTrainStationModel interface {
		tTrainStationModel
	}

	customTTrainStationModel struct {
		*defaultTTrainStationModel
	}
)

// NewTTrainStationModel returns a model for the database table.
func NewTTrainStationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TTrainStationModel {
	return &customTTrainStationModel{
		defaultTTrainStationModel: newTTrainStationModel(conn, c, opts...),
	}
}
