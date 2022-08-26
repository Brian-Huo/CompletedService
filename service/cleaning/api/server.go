package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"cleaningservice/common/errorx"
	"cleaningservice/common/orderqueue"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/config"
	"cleaningservice/service/cleaning/api/internal/handler"
	"cleaningservice/service/cleaning/api/internal/logic"
	"cleaningservice/service/cleaning/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/server.yaml", "the config file")

func notAllowedFn(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // 允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Authorization,true") // header的类型
	w.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
}

func orderQueueInit(svcCtx *svc.ServiceContext) {
	for {
		time.Sleep(time.Second * time.Duration(variables.Check_time_unit))

		// Signal invoice queue
		l := logic.NewSendInvoiceRequestLogic(context.TODO(), svcCtx)
		err := l.SendInvoiceRequest()
		if err != nil {
			logx.Alert("Send Invoice emails failed")
		}

		// Signal reminder queue
		if orderqueue.GetIteration() == variables.Check_time_clock {
			l := logic.NewSendReminderRequestLogic(context.TODO(), svcCtx)
			err := l.SendReminderRequest()
			if err != nil {
				logx.Alert("Send reminder emails failed")
			}
		}
	}
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
	go orderQueueInit(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
