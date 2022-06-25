package handler

import (
	"net/http"

	"cleaningservice/service/cleaning/api/internal/logic"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateAddressRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateAddressLogic(r.Context(), svcCtx)
		resp, err := l.CreateAddress(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
