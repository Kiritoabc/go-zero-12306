package tTrainStationRelation

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TTrainStationRelationModel = (*customTTrainStationRelationModel)(nil)

type (
	// TTrainStationRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTrainStationRelationModel.
	TTrainStationRelationModel interface {
		tTrainStationRelationModel
	}

	customTTrainStationRelationModel struct {
		*defaultTTrainStationRelationModel
	}
)

// NewTTrainStationRelationModel returns a model for the database table.
func NewTTrainStationRelationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TTrainStationRelationModel {
	return &customTTrainStationRelationModel{
		defaultTTrainStationRelationModel: newTTrainStationRelationModel(conn, c, opts...),
	}
}
