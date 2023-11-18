package tTrainStationPrice

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TTrainStationPriceModel = (*customTTrainStationPriceModel)(nil)

type (
	// TTrainStationPriceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTrainStationPriceModel.
	TTrainStationPriceModel interface {
		tTrainStationPriceModel
	}

	customTTrainStationPriceModel struct {
		*defaultTTrainStationPriceModel
	}
)

// NewTTrainStationPriceModel returns a model for the database table.
func NewTTrainStationPriceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TTrainStationPriceModel {
	return &customTTrainStationPriceModel{
		defaultTTrainStationPriceModel: newTTrainStationPriceModel(conn, c, opts...),
	}
}
