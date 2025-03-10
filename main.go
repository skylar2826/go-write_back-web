package main

import (
	context2 "context"
	"fmt"
	"geektime-go2/init_register"
	"geektime-go2/orm/orm_gen/data"
	"geektime-go2/orm/predicate"
	"geektime-go2/orm/sql/insertor"
	"geektime-go2/orm/sql/selector"
	"geektime-go2/web/context"
	__shutdown "geektime-go2/web/graceful_shutdown"
	"geektime-go2/web/handler/handle_based_on_tree"
	"geektime-go2/web/handler_func"
	"geektime-go2/web/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
	"time"
)

func routeHandler2(c *context.Context) {
	u := &data.User{}
	u.Username = c.R.FormValue("Username")
	u.Email = c.R.FormValue("Email")
	var err error
	u.Id, err = strconv.Atoi(c.R.FormValue("Id"))
	if err != nil {
		_ = c.SystemErrorJson(err)
	}

	// Insert into users (`u_username`,`u_id`,`email`) Values (?,?,?),(?,?,?);
	//i := insertor.NewInserter[data.User](init_register.DB).Columns(predicate.C("Username"), predicate.C("Email")).Values(u)

	//res := i.Execute(c)
	//if res.Err != nil {
	//	_ = c.SystemErrorJson(res.Err)
	//}

	err = init_register.Cache.Set(*c, strconv.Itoa(u.Id), u, time.Minute)
	if err != nil {
		_ = c.SystemErrorJson(err)
	}

	_ = c.OkJson(fmt.Sprintf("200 注册成功\n"))
}

func syncCacheToDB(ctx context2.Context) error {
	for key, value := range init_register.Cache.Data {
		id, err := strconv.Atoi(key)
		if err != nil {
			log.Println(err)
		}
		var res any
		res, err = selector.NewSelector[data.User](init_register.DB).Where(data.UserIdEq(id)).Get(ctx)
		if err != nil {
			log.Println(err)
		}
		if res == nil {
			result := insertor.NewInserter[data.User](init_register.DB).Columns(predicate.C("Id"), predicate.C("Username"), predicate.C("Email")).Values(value.Val.(*data.User)).Execute(ctx)
			if result.Err != nil {
				log.Println(result.Err)
			}
		}
	}
	log.Printf("同步数据库完成\n")
	return nil
}

func main() {
	hdl := handle_based_on_tree.NewHandleBasedOnTree()

	s := server.NewSdkHttpServer("web-server", hdl, server.WithMiddlewares(__shutdown.G.GracefulShutdownFilterBuilder))

	s.Route(http.MethodPost, "/login", handler_func.SignUp)
	s.Route(http.MethodPost, "/register", routeHandler2)

	go func() {
		http.Handle("/metics", promhttp.Handler())
		_ = http.ListenAndServe("127.0.0.1:8082", nil)
	}()

	// 拒绝请求，关闭服务，释放资源
	go func() {
		__shutdown.WaitForShutdown(__shutdown.RejectRequestHook, syncCacheToDB, __shutdown.BuildServerHook(s))
	}()

	_ = s.Start("127.0.0.1:8081")
}
