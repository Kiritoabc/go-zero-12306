package tTrain

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TTrainModel = (*customTTrainModel)(nil)

type (
	// TTrainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTrainModel.
	TTrainModel interface {
		tTrainModel
	}

	customTTrainModel struct {
		*defaultTTrainModel
	}
)

// NewTTrainModel returns a model for the database table.
func NewTTrainModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TTrainModel {
	return &customTTrainModel{
		defaultTTrainModel: newTTrainModel(conn, c, opts...),
	}
}
