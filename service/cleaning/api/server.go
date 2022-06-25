package main

import (
	"flag"
	"fmt"
	"net/http"

	"cleaningservice/common/errorx"
	"cleaningservice/common/orderqueue"
	"cleaningservice/service/cleaning/api/internal/config"
	"cleaningservice/service/cleaning/api/internal/handler"
	"cleaningservice/service/cleaning/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/server.yaml", "the config file")

func notAllowedFn(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // 允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Authorization,true") // header的类型
	w.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, notAllowedFn))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// Sel-defined error
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Order queue initialized.\n")
	go orderqueue.OrderQueueStart()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
