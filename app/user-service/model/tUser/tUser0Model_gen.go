// Code generated by goctl. DO NOT EDIT.

package tUser

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
	tUser0FieldNames          = builder.RawFieldNames(&TUser0{})
	tUser0Rows                = strings.Join(tUser0FieldNames, ",")
	tUser0RowsExpectAutoSet   = strings.Join(stringx.Remove(tUser0FieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tUser0RowsWithPlaceHolder = strings.Join(stringx.Remove(tUser0FieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cache12306User0TUser0IdPrefix                   = "cache:12306User0:tUser0:id:"
	cache12306User0TUser0UsernameDeletionTimePrefix = "cache:12306User0:tUser0:username:deletionTime:"
)

type (
	tUser0Model interface {
		Insert(ctx context.Context, data *TUser0) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TUser0, error)
		FindOneByUsernameDeletionTime(ctx context.Context, username sql.NullString, deletionTime int64) (*TUser0, error)
		Update(ctx context.Context, data *TUser0) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTUser0Model struct {
		sqlc.CachedConn
		table string
	}

	TUser0 struct {
		Id           int64          `db:"id"`            // ID
		Username     sql.NullString `db:"username"`      // 用户名
		Password     sql.NullString `db:"password"`      // 密码
		RealName     sql.NullString `db:"real_name"`     // 真实姓名
		Region       string         `db:"region"`        // 国家/地区
		IdType       sql.NullInt64  `db:"id_type"`       // 证件类型
		IdCard       sql.NullString `db:"id_card"`       // 证件号
		Phone        sql.NullString `db:"phone"`         // 手机号
		Telephone    sql.NullString `db:"telephone"`     // 固定电话
		Mail         sql.NullString `db:"mail"`          // 邮箱
		UserType     sql.NullInt64  `db:"user_type"`     // 旅客类型
		VerifyStatus sql.NullInt64  `db:"verify_status"` // 审核状态
		PostCode     sql.NullString `db:"post_code"`     // 邮编
		Address      sql.NullString `db:"address"`       // 地址
		DeletionTime int64          `db:"deletion_time"` // 注销时间戳
		CreateTime   sql.NullTime   `db:"create_time"`   // 创建时间
		UpdateTime   sql.NullTime   `db:"update_time"`   // 修改时间
		DelFlag      sql.NullInt64  `db:"del_flag"`      // 删除标识
	}
)

func newTUser0Model(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTUser0Model {
	return &defaultTUser0Model{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`t_user_0`",
	}
}

func (m *defaultTUser0Model) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	_12306User0TUser0IdKey := fmt.Sprintf("%s%v", cache12306User0TUser0IdPrefix, id)
	_12306User0TUser0UsernameDeletionTimeKey := fmt.Sprintf("%s%v:%v", cache12306User0TUser0UsernameDeletionTimePrefix, data.Username, data.DeletionTime)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, _12306User0TUser0IdKey, _12306User0TUser0UsernameDeletionTimeKey)
	return err
}

func (m *defaultTUser0Model) FindOne(ctx context.Context, id int64) (*TUser0, error) {
	_12306User0TUser0IdKey := fmt.Sprintf("%s%v", cache12306User0TUser0IdPrefix, id)
	var resp TUser0
	err := m.QueryRowCtx(ctx, &resp, _12306User0TUser0IdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tUser0Rows, m.table)
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

func (m *defaultTUser0Model) FindOneByUsernameDeletionTime(ctx context.Context, username sql.NullString, deletionTime int64) (*TUser0, error) {
	_12306User0TUser0UsernameDeletionTimeKey := fmt.Sprintf("%s%v:%v", cache12306User0TUser0UsernameDeletionTimePrefix, username, deletionTime)
	var resp TUser0
	err := m.QueryRowIndexCtx(ctx, &resp, _12306User0TUser0UsernameDeletionTimeKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? and `deletion_time` = ? limit 1", tUser0Rows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, username, deletionTime); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTUser0Model) Insert(ctx context.Context, data *TUser0) (sql.Result, error) {
	_12306User0TUser0IdKey := fmt.Sprintf("%s%v", cache12306User0TUser0IdPrefix, data.Id)
	_12306User0TUser0UsernameDeletionTimeKey := fmt.Sprintf("%s%v:%v", cache12306User0TUser0UsernameDeletionTimePrefix, data.Username, data.DeletionTime)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tUser0RowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Username, data.Password, data.RealName, data.Region, data.IdType, data.IdCard, data.Phone, data.Telephone, data.Mail, data.UserType, data.VerifyStatus, data.PostCode, data.Address, data.DeletionTime, data.DelFlag)
	}, _12306User0TUser0IdKey, _12306User0TUser0UsernameDeletionTimeKey)
	return ret, err
}

func (m *defaultTUser0Model) Update(ctx context.Context, newData *TUser0) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	_12306User0TUser0IdKey := fmt.Sprintf("%s%v", cache12306User0TUser0IdPrefix, data.Id)
	_12306User0TUser0UsernameDeletionTimeKey := fmt.Sprintf("%s%v:%v", cache12306User0TUser0UsernameDeletionTimePrefix, data.Username, data.DeletionTime)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tUser0RowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Username, newData.Password, newData.RealName, newData.Region, newData.IdType, newData.IdCard, newData.Phone, newData.Telephone, newData.Mail, newData.UserType, newData.VerifyStatus, newData.PostCode, newData.Address, newData.DeletionTime, newData.DelFlag, newData.Id)
	}, _12306User0TUser0IdKey, _12306User0TUser0UsernameDeletionTimeKey)
	return err
}

func (m *defaultTUser0Model) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cache12306User0TUser0IdPrefix, primary)
}

func (m *defaultTUser0Model) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tUser0Rows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTUser0Model) tableName() string {
	return m.table
}