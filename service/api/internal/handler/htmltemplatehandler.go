package handler

import (
	"html/template"
	"net/http"

	"cleaningservice/service/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HtmlTemplateHandler(globalTemplate *template.Template, templateName string, svcCtx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		t := globalTemplate
		err := t.ExecuteTemplate(w, templateName, r.URL.Query())
		if err != nil {
			httpx.Error(w, err)
		}
	}
}
