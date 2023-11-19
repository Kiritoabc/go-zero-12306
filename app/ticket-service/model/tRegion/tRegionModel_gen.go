// Code generated by goctl. DO NOT EDIT.

package tRegion

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tRegionFieldNames          = builder.RawFieldNames(&TRegion{})
	tRegionRows                = strings.Join(tRegionFieldNames, ",")
	tRegionRowsExpectAutoSet   = strings.Join(stringx.Remove(tRegionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tRegionRowsWithPlaceHolder = strings.Join(stringx.Remove(tRegionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cache12306TicketTRegionIdPrefix = "cache:12306Ticket:tRegion:id:"
)

type (
	tRegionModel interface {
		Insert(ctx context.Context, data *TRegion) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TRegion, error)
		Update(ctx context.Context, data *TRegion) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTRegionModel struct {
		sqlc.CachedConn
		table string
	}

	TRegion struct {
		Id          int64          `db:"id"`           // ID
		Name        sql.NullString `db:"name"`         // 地区名称
		FullName    sql.NullString `db:"full_name"`    // 地区全名
		Code        sql.NullString `db:"code"`         // 地区编码
		Initial     sql.NullString `db:"initial"`      // 地区首字母
		Spell       sql.NullString `db:"spell"`        // 拼音
		PopularFlag sql.NullInt64  `db:"popular_flag"` // 热门标识
		CreateTime  sql.NullTime   `db:"create_time"`  // 创建时间
		UpdateTime  sql.NullTime   `db:"update_time"`  // 修改时间
		DelFlag     sql.NullInt64  `db:"del_flag"`     // 删除标识
	}
)

func newTRegionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTRegionModel {
	return &defaultTRegionModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`t_region`",
	}
}

func (m *defaultTRegionModel) Delete(ctx context.Context, id int64) error {
	_12306TicketTRegionIdKey := fmt.Sprintf("%s%v", cache12306TicketTRegionIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, _12306TicketTRegionIdKey)
	return err
}

func (m *defaultTRegionModel) FindOne(ctx context.Context, id int64) (*TRegion, error) {
	_12306TicketTRegionIdKey := fmt.Sprintf("%s%v", cache12306TicketTRegionIdPrefix, id)
	var resp TRegion
	err := m.QueryRowCtx(ctx, &resp, _12306TicketTRegionIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tRegionRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTRegionModel) Insert(ctx context.Context, data *TRegion) (sql.Result, error) {
	_12306TicketTRegionIdKey := fmt.Sprintf("%s%v", cache12306TicketTRegionIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, tRegionRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.FullName, data.Code, data.Initial, data.Spell, data.PopularFlag, data.DelFlag)
	}, _12306TicketTRegionIdKey)
	return ret, err
}

func (m *defaultTRegionModel) Update(ctx context.Context, data *TRegion) error {
	_12306TicketTRegionIdKey := fmt.Sprintf("%s%v", cache12306TicketTRegionIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tRegionRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.FullName, data.Code, data.Initial, data.Spell, data.PopularFlag, data.DelFlag, data.Id)
	}, _12306TicketTRegionIdKey)
	return err
}

func (m *defaultTRegionModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cache12306TicketTRegionIdPrefix, primary)
}

func (m *defaultTRegionModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tRegionRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTRegionModel) tableName() string {
	return m.table
}