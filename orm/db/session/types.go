package session

import (
	"context"
	"database/sql"
	sql2 "geektime-go2/orm"
	"geektime-go2/orm/db/dialect"
	"geektime-go2/orm/db/register"
	"geektime-go2/orm/db/valuer"
)

type Core struct {
	R            *register.Register
	Model        *register.Model
	ValueCreator valuer.ValueCreator
	Dialect      dialect.Dialect
}

// Session 是 db、tx的顶层抽象
type Session interface {
	GetCore() Core
	ExecContext(ctx context.Context, query string, args ...any) *sql2.QueryResult
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

var _ Session = &DB{}
var _ Session = &Tx{}
