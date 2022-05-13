package handler

import (
	"net/http"
	"text/template"

	"cleaningservice/service/api/internal/svc"
)

func IndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles("view/index.html"))
		tmp.Execute(w, nil)
	}
}
