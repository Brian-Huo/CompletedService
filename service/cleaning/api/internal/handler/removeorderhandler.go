package handler

import (
	"net/http"

	"cleaningservice/service/cleaning/api/internal/logic"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveOrderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRemoveOrderLogic(r.Context(), svcCtx)
		resp, err := l.RemoveOrder(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
