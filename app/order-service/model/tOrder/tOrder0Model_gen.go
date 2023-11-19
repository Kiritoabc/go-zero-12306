// Code generated by goctl. DO NOT EDIT.

package tOrder

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
	tOrder0FieldNames          = builder.RawFieldNames(&TOrder0{})
	tOrder0Rows                = strings.Join(tOrder0FieldNames, ",")
	tOrder0RowsExpectAutoSet   = strings.Join(stringx.Remove(tOrder0FieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tOrder0RowsWithPlaceHolder = strings.Join(stringx.Remove(tOrder0FieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cache12306Order0TOrder0IdPrefix = "cache:12306Order0:tOrder0:id:"
)

type (
	tOrder0Model interface {
		Insert(ctx context.Context, data *TOrder0) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TOrder0, error)
		Update(ctx context.Context, data *TOrder0) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTOrder0Model struct {
		sqlc.CachedConn
		table string
	}

	TOrder0 struct {
		Id            int64          `db:"id"`             // ID
		OrderSn       sql.NullString `db:"order_sn"`       // 订单号
		UserId        sql.NullInt64  `db:"user_id"`        // 用户ID
		Username      sql.NullString `db:"username"`       // 用户名
		TrainId       sql.NullInt64  `db:"train_id"`       // 列车ID
		TrainNumber   sql.NullString `db:"train_number"`   // 列车车次
		RidingDate    sql.NullTime   `db:"riding_date"`    // 乘车日期
		Departure     sql.NullString `db:"departure"`      // 出发站点
		Arrival       sql.NullString `db:"arrival"`        // 到达站点
		DepartureTime sql.NullTime   `db:"departure_time"` // 出发时间
		ArrivalTime   sql.NullTime   `db:"arrival_time"`   // 到达时间
		Source        sql.NullInt64  `db:"source"`         // 订单来源
		Status        sql.NullInt64  `db:"status"`         // 订单状态
		OrderTime     sql.NullTime   `db:"order_time"`     // 下单时间
		PayType       sql.NullInt64  `db:"pay_type"`       // 支付方式
		PayTime       sql.NullTime   `db:"pay_time"`       // 支付时间
		CreateTime    sql.NullTime   `db:"create_time"`    // 创建时间
		UpdateTime    sql.NullTime   `db:"update_time"`    // 修改时间
		DelFlag       sql.NullInt64  `db:"del_flag"`       // 删除标识
	}
)

// todo: 添加后缀
func newTOrder0Model(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTOrder0Model {
	return &defaultTOrder0Model{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`t_order_0`",
	}
}

func (m *defaultTOrder0Model) Delete(ctx context.Context, id int64) error {
	_12306Order0TOrder0IdKey := fmt.Sprintf("%s%v", cache12306Order0TOrder0IdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, _12306Order0TOrder0IdKey)
	return err
}

func (m *defaultTOrder0Model) FindOne(ctx context.Context, id int64) (*TOrder0, error) {
	_12306Order0TOrder0IdKey := fmt.Sprintf("%s%v", cache12306Order0TOrder0IdPrefix, id)
	var resp TOrder0
	err := m.QueryRowCtx(ctx, &resp, _12306Order0TOrder0IdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tOrder0Rows, m.table)
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

func (m *defaultTOrder0Model) Insert(ctx context.Context, data *TOrder0) (sql.Result, error) {
	_12306Order0TOrder0IdKey := fmt.Sprintf("%s%v", cache12306Order0TOrder0IdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tOrder0RowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.OrderSn, data.UserId, data.Username, data.TrainId, data.TrainNumber, data.RidingDate, data.Departure, data.Arrival, data.DepartureTime, data.ArrivalTime, data.Source, data.Status, data.OrderTime, data.PayType, data.PayTime, data.DelFlag)
	}, _12306Order0TOrder0IdKey)
	return ret, err
}

func (m *defaultTOrder0Model) Update(ctx context.Context, data *TOrder0) error {
	_12306Order0TOrder0IdKey := fmt.Sprintf("%s%v", cache12306Order0TOrder0IdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tOrder0RowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.OrderSn, data.UserId, data.Username, data.TrainId, data.TrainNumber, data.RidingDate, data.Departure, data.Arrival, data.DepartureTime, data.ArrivalTime, data.Source, data.Status, data.OrderTime, data.PayType, data.PayTime, data.DelFlag, data.Id)
	}, _12306Order0TOrder0IdKey)
	return err
}

func (m *defaultTOrder0Model) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cache12306Order0TOrder0IdPrefix, primary)
}

func (m *defaultTOrder0Model) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tOrder0Rows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTOrder0Model) tableName() string {
	return m.table
}