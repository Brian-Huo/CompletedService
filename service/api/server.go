package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"cleaningservice/service/api/internal/config"
	"cleaningservice/service/api/internal/handler"
	"cleaningservice/service/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/server.yaml", "the config file")
var funcMap = template.FuncMap{}

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

	var pageRouters []rest.Route
	globalTemplate, _ := template.New("").Funcs(funcMap).ParseGlob("./internal/view/www/*")
	for _, tpl := range globalTemplate.Templates() {
		pattern := tpl.Name()
		if !strings.HasPrefix(pattern, "/") {
			pattern = "/" + pattern
		}
		templateName := tpl.Name()
		if 0 != len(templateName) {
			pageRouters = append(pageRouters, rest.Route{
				Method:  http.MethodGet,
				Path:    pattern,
				Handler: handler.HtmlTemplateHandler(globalTemplate, templateName, ctx),
			})
			logx.Infof("register page %s %s", pattern, templateName)
			server.AddRoutes(pageRouters)
		}
		templateName = "/"
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
