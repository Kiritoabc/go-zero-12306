// Code generated by goctl. DO NOT EDIT.

package tUser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tUser0FieldNames          = builder.RawFieldNames(&TUser0{})
	tUser0Rows                = strings.Join(tUser0FieldNames, ",")
	tUser0RowsExpectAutoSet   = strings.Join(stringx.Remove(tUser0FieldNames, "`id`"), ",")
	tUser0RowsWithPlaceHolder = strings.Join(stringx.Remove(tUser0FieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cache12306User0TUser0IdPrefix                   = "cache:12306User0:tUser0:id:"
	cache12306User0TUser0UsernameDeletionTimePrefix = "cache:12306User0:tUser0:username:deletionTime:"
	cache12306User0TUser0UsernamePrefix             = "cache:12306User0:tUser0:username:"
)

type (
	tUser0Model interface {
		Insert(ctx context.Context, session sqlx.Session, data *TUser0) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TUser0, error)
		FindOneByUsernameDeletionTime(ctx context.Context, username string, deletionTime int64) (*TUser0, error)
		Update(ctx context.Context, session sqlx.Session, data *TUser0) error
		Delete(ctx context.Context, id int64) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		FindOneByUsername(ctx context.Context, session sqlx.Session, username string) (*TUser0, error)
		DeleteWithSession(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultTUser0Model struct {
		sqlc.CachedConn
		table string
	}

	TUser0 struct {
		Id           int64          `db:"id"`            // ID
		Username     string         `db:"username"`      // 用户名
		Password     string         `db:"password"`      // 密码
		RealName     string         `db:"real_name"`     // 真实姓名
		Region       sql.NullString `db:"region"`        // 国家/地区
		IdType       int64          `db:"id_type"`       // 证件类型
		IdCard       string         `db:"id_card"`       // 证件号
		Phone        string         `db:"phone"`         // 手机号
		Telephone    string         `db:"telephone"`     // 固定电话
		Mail         string         `db:"mail"`          // 邮箱
		UserType     int64          `db:"user_type"`     // 旅客类型
		VerifyStatus int64          `db:"verify_status"` // 审核状态
		PostCode     string         `db:"post_code"`     // 邮编
		Address      sql.NullString `db:"address"`       // 地址
		DeletionTime int64          `db:"deletion_time"` // 注销时间戳
		CreateTime   time.Time      `db:"create_time"`   // 创建时间
		UpdateTime   time.Time      `db:"update_time"`   // 修改时间
		DelFlag      int64          `db:"del_flag"`      // 删除标识
	}
)

func newTUser0Model(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTUser0Model {
	return &defaultTUser0Model{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`t_user_0`",
	}
}

// 开启事务
func (m *defaultTUser0Model) Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

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

func (m *defaultTUser0Model) FindOneByUsernameDeletionTime(ctx context.Context, username string, deletionTime int64) (*TUser0, error) {
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

func (m *defaultTUser0Model) Insert(ctx context.Context, session sqlx.Session, data *TUser0) (sql.Result, error) {
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?)", m.table, tUser0RowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Username, data.Password, data.RealName, data.Region, data.IdType, data.IdCard, data.Phone, data.Telephone, data.Mail, data.UserType, data.VerifyStatus, data.PostCode, time.Now(), time.Now(), 0)
		}
		return conn.ExecCtx(ctx, query, data.Username, data.Password, data.RealName, data.Region, data.IdType, data.IdCard, data.Phone, data.Telephone, data.Mail, data.UserType, data.VerifyStatus, data.PostCode, time.Now(), time.Now(), 0)
	})
	return ret, err
}

func (m *defaultTUser0Model) Update(ctx context.Context, session sqlx.Session, data *TUser0) error {
	data, err := m.FindOne(ctx, data.Id)
	if err != nil {
		return err
	}

	_12306User0TUser0IdKey := fmt.Sprintf("%s%v", cache12306User0TUser0IdPrefix, data.Id)
	_12306User0TUser0UsernameDeletionTimeKey := fmt.Sprintf("%s%v:%v", cache12306User0TUser0UsernameDeletionTimePrefix, data.Username, data.DeletionTime)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tUser0RowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Username, data.Password, data.RealName, data.Region, data.IdType, data.IdCard, data.Phone, data.Telephone, data.Mail, data.UserType, data.VerifyStatus, data.PostCode, data.Address, data.DeletionTime, data.DelFlag, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.Username, data.Password, data.RealName, data.Region, data.IdType, data.IdCard, data.Phone, data.Telephone, data.Mail, data.UserType, data.VerifyStatus, data.PostCode, data.Address, data.DeletionTime, data.DelFlag, data.Id)
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

// 自定义
func (m *defaultTUser0Model) FindOneByUsername(ctx context.Context, session sqlx.Session, username string) (*TUser0, error) {
	_12306User0TUser0UsernameKey := fmt.Sprintf("%s%v", cache12306User0TUser0UsernamePrefix, username)
	var resp TUser0
	err := m.QueryRowCtx(ctx, &resp, _12306User0TUser0UsernameKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", tUser0Rows, m.table)
		if session != nil {
			return session.QueryRowsCtx(ctx, &resp, query, username)
		}
		return conn.QueryRowCtx(ctx, &resp, query, username)
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

func (m *defaultTUser0Model) DeleteWithSession(ctx context.Context, session sqlx.Session, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}
	_12306User0TUser0IdKey := fmt.Sprintf("%s%v", cache12306User0TUser0IdPrefix, id)
	_12306User0TUser0UsernameDeletionTimeKey := fmt.Sprintf("%s%v:%v", cache12306User0TUser0UsernameDeletionTimePrefix, data.Username, data.DeletionTime)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `deletion_time` = ? , `del_flag`='1' where `id` = ? and `del_flag` = '0'", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, time.Now(), id)
		}
		return conn.ExecCtx(ctx, query, time.Now(), id)
	}, _12306User0TUser0IdKey, _12306User0TUser0UsernameDeletionTimeKey)
	return err
}
