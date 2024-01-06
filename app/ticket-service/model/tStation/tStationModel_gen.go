// Code generated by goctl. DO NOT EDIT.

package tStation

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"go-zero-12306/common/globalkey"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tStationFieldNames          = builder.RawFieldNames(&TStation{})
	tStationRows                = strings.Join(tStationFieldNames, ",")
	tStationRowsExpectAutoSet   = strings.Join(stringx.Remove(tStationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tStationRowsWithPlaceHolder = strings.Join(stringx.Remove(tStationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cache12306TicketTStationIdPrefix = "cache:12306Ticket:tStation:id:"
)

type (
	tStationModel interface {
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		List(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*TStation, error)
		SelectListByName(ctx context.Context, builder squirrel.SelectBuilder, name string) ([]*TStation, error)
	}

	defaultTStationModel struct {
		sqlc.CachedConn
		table string
	}

	TStation struct {
		Id         int64          `db:"id"`          // ID
		Code       sql.NullString `db:"code"`        // 车站编号
		Name       sql.NullString `db:"name"`        // 车站名称
		Spell      sql.NullString `db:"spell"`       // 拼音
		Region     sql.NullString `db:"region"`      // 车站地区
		RegionName sql.NullString `db:"region_name"` // 车站地区名称
		CreateTime sql.NullTime   `db:"create_time"` // 创建时间
		UpdateTime sql.NullTime   `db:"update_time"` // 修改时间
		DelFlag    sql.NullInt64  `db:"del_flag"`    // 删除标识
	}
)

func newTStationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTStationModel {
	return &defaultTStationModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`t_station`",
	}
}

func (m *defaultTStationModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cache12306TicketTStationIdPrefix, primary)
}

func (m *defaultTStationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tStationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTStationModel) tableName() string {
	return m.table
}

// 自定义

func (m *defaultTStationModel) Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultTStationModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultTStationModel) List(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*TStation, error) {

	builder = builder.Columns(tStationRows)
	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_flag = ?", globalkey.DelFlagNo).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*TStation

	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTStationModel) SelectListByName(ctx context.Context, builder squirrel.SelectBuilder, name string) ([]*TStation, error) {
	builder = builder.Columns(tStationRows)

	query, value, err := builder.Where("del_flag = ?", globalkey.DelFlagNo).
		Where(squirrel.Or{
			squirrel.Like{"name": name},
			squirrel.Like{"spell": name},
		}).ToSql()
	var resp []*TStation
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, value...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
