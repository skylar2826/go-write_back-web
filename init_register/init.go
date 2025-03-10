package init_register

import (
	context2 "context"
	"fmt"
	"geektime-go2/cache"
	"geektime-go2/orm/db/dialect"
	"geektime-go2/orm/db/session"
	"geektime-go2/orm/orm_gen/data"
	"geektime-go2/orm/sql/selector"
	"log"
	"strconv"
	"time"
)

func syncDBToCache(ctx context2.Context) {
	s := selector.NewSelector[data.User](DB)
	res, err := s.GetMutli(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	for _, val := range res {
		v := val.(*data.User)
		err = Cache.Set(ctx, strconv.Itoa(v.Id), v, time.Minute)
		if err != nil {
			log.Printf("init error: %s\n", err)
		}
	}
}

// callbackDB 缓存过期时同步删除db
func deleteDBData(key string, val any) {
	query := fmt.Sprintf("Delete From users where `u_id` = ?")
	_, err := DB.DB.ExecContext(context2.Background(), query, key)
	if err != nil {
		log.Println(err)
	}
}

var DB *session.DB
var Cache *cache.BuildInMemoryCache

func init() {
	dataSourceName := fmt.Sprint(Username, ":", Password, "@tcp(", Ip, ":", Port, ")/", DbName)
	var err error
	DB, err = session.Open("mysql", dataSourceName, session.WithDialect(dialect.NewMysqlSQL()))
	if err != nil {
		log.Println(err)
		return
	}

	Cache = cache.NewBuildInMemoryCache(cache.WithOnEvicted(deleteDBData))
	ctx := context2.Background()
	syncDBToCache(ctx)
}
