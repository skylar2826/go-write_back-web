package filter

import (
	"fmt"
	"geektime-go2/web/context"
	"time"
)

func TimeFilterBuilder(next Filter) Filter {
	return func(c *context.Context) {
		start := time.Now().UnixNano()
		next(c)
		now := time.Now().UnixNano()
		fmt.Printf("执行花费时间：%d\n", now-start)
	}
}
