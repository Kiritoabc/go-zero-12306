// Code generated by goctl. DO NOT EDIT.

package tTrainStation

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tTrainStationFieldNames          = builder.RawFieldNames(&TTrainStation{})
	tTrainStationRows                = strings.Join(tTrainStationFieldNames, ",")
	tTrainStationRowsExpectAutoSet   = strings.Join(stringx.Remove(tTrainStationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tTrainStationRowsWithPlaceHolder = strings.Join(stringx.Remove(tTrainStationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cache12306TicketTTrainStationIdPrefix = "cache:12306Ticket:tTrainStation:id:"
)

type (
	tTrainStationModel interface {
		// 自定义
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		SlectListByTrainId(ctx context.Context, builder squirrel.SelectBuilder, id int64) ([]*TTrainStation, error)
	}

	defaultTTrainStationModel struct {
		sqlc.CachedConn
		table string
	}

	TTrainStation struct {
		Id            int64          `db:"id"`             // ID
		TrainId       int64          `db:"train_id"`       // 车次ID
		StationId     int64          `db:"station_id"`     // 车站ID
		Sequence      string         `db:"sequence"`       // 站点顺序
		Departure     string         `db:"departure"`      // 出发站点
		Arrival       sql.NullString `db:"arrival"`        // 到达站点
		StartRegion   string         `db:"start_region"`   // 起始城市
		EndRegion     sql.NullString `db:"end_region"`     // 终点城市
		ArrivalTime   time.Time      `db:"arrival_time"`   // 到站时间
		DepartureTime time.Time      `db:"departure_time"` // 出站时间
		StopoverTime  sql.NullInt64  `db:"stopover_time"`  // 停留时间，单位分
		CreateTime    time.Time      `db:"create_time"`    // 创建时间
		UpdateTime    time.Time      `db:"update_time"`    // 修改时间
		DelFlag       int64          `db:"del_flag"`       // 删除标识
	}
)

func newTTrainStationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTTrainStationModel {
	return &defaultTTrainStationModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`t_train_station`",
	}
}

func (m *defaultTTrainStationModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cache12306TicketTTrainStationIdPrefix, primary)
}

func (m *defaultTTrainStationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tTrainStationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTTrainStationModel) tableName() string {
	return m.table
}

// 自定义
// 事务
func (m *defaultTTrainStationModel) Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultTTrainStationModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultTTrainStationModel) SlectListByTrainId(ctx context.Context, builder squirrel.SelectBuilder, id int64) ([]*TTrainStation, error) {
	builder = builder.Columns(tTrainStationRows)
	query, values, err := builder.Where("`train_id` = ?", id).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*TTrainStation
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
