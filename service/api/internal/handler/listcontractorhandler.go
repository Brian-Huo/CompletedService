package handler

import (
	"net/http"

	"cleaningservice/service/api/internal/logic"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListContractorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListContractorRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewListContractorLogic(r.Context(), svcCtx)
		resp, err := l.ListContractor(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
