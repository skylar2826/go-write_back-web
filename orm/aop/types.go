package aop

import (
	"context"
	"geektime-go2/orm"
)

type Handler func(ctx context.Context, qc *orm.QueryContext) *orm.QueryResult

type Middleware func(next Handler) Handler
