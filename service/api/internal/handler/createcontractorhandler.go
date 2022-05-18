package handler

import (
	"net/http"

	"cleaningservice/service/api/internal/logic"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateContractorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateContractorRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateContractorLogic(r.Context(), svcCtx)
		resp, err := l.CreateContractor(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
