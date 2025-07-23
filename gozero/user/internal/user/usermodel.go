package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		withSession(session sqlx.Session) UserModel
		FindByUsername(ctx context.Context, username string) (*User, error)
		FindByUsernameAndPwd(ctx context.Context, username, password string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *customUserModel) withSession(session sqlx.Session) UserModel {
	return NewUserModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customUserModel) FindByUsername(ctx context.Context, username string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sql.ErrNoRows, sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customUserModel) FindByUsernameAndPwd(ctx context.Context, username, password string) (*User, error) {
	// 构造 SQL 查询语句，增加了密码字段
	query := fmt.Sprintf("select %s from %s where `username` = ? and `password` = ? limit 1", userRows, m.table)
	var resp User
	// 执行查询
	err := m.conn.QueryRowCtx(ctx, &resp, query, username, password)
	switch err {
	case nil:
		// 如果没有错误，返回查询结果
		return &resp, nil
	case sql.ErrNoRows, sqlx.ErrNotFound:
		// 如果没有找到记录，返回 nil
		return nil, nil
	default:
		// 其他错误，返回错误
		return nil, err
	}
}
