// Code generated by goctl. DO NOT EDIT.

package model

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
	followsFieldNames          = builder.RawFieldNames(&Follows{})
	followsRows                = strings.Join(followsFieldNames, ",")
	followsRowsExpectAutoSet   = strings.Join(stringx.Remove(followsFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	followsRowsWithPlaceHolder = strings.Join(stringx.Remove(followsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLiujunUserFollowsIdPrefix = "cache:liujunUser:follows:id:"
)

type (
	followsModel interface {
		Insert(ctx context.Context, data *Follows) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Follows, error)
		Update(ctx context.Context, data *Follows) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFollowsModel struct {
		sqlc.CachedConn
		table string
	}

	Follows struct {
		Id         int64          `db:"id"`          // 主键
		UserId     sql.NullString `db:"user_id"`     // 关注的人
		FollowId   sql.NullString `db:"follow_id"`   // 被关注的人
		CreateTime int64          `db:"create_time"` // 该条记录创建时间
		UpdateTime int64          `db:"update_time"` // 该条最后一次信息修改时间
		IsDelete   int64          `db:"is_delete"`   // 逻辑删除
	}
)

func newFollowsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultFollowsModel {
	return &defaultFollowsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`follows`",
	}
}

func (m *defaultFollowsModel) withSession(session sqlx.Session) *defaultFollowsModel {
	return &defaultFollowsModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`follows`",
	}
}

func (m *defaultFollowsModel) Delete(ctx context.Context, id int64) error {
	liujunUserFollowsIdKey := fmt.Sprintf("%s%v", cacheLiujunUserFollowsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, liujunUserFollowsIdKey)
	return err
}

func (m *defaultFollowsModel) FindOne(ctx context.Context, id int64) (*Follows, error) {
	liujunUserFollowsIdKey := fmt.Sprintf("%s%v", cacheLiujunUserFollowsIdPrefix, id)
	var resp Follows
	err := m.QueryRowCtx(ctx, &resp, liujunUserFollowsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", followsRows, m.table)
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

func (m *defaultFollowsModel) Insert(ctx context.Context, data *Follows) (sql.Result, error) {
	liujunUserFollowsIdKey := fmt.Sprintf("%s%v", cacheLiujunUserFollowsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, followsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.UserId, data.FollowId, data.IsDelete)
	}, liujunUserFollowsIdKey)
	return ret, err
}

func (m *defaultFollowsModel) Update(ctx context.Context, data *Follows) error {
	liujunUserFollowsIdKey := fmt.Sprintf("%s%v", cacheLiujunUserFollowsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, followsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.FollowId, data.IsDelete, data.Id)
	}, liujunUserFollowsIdKey)
	return err
}

func (m *defaultFollowsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLiujunUserFollowsIdPrefix, primary)
}

func (m *defaultFollowsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", followsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFollowsModel) tableName() string {
	return m.table
}
