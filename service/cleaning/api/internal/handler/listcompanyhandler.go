package handler

import (
	"net/http"

	"cleaningservice/service/cleaning/api/internal/logic"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListCompanyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListCompanyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewListCompanyLogic(r.Context(), svcCtx)
		resp, err := l.ListCompany(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
