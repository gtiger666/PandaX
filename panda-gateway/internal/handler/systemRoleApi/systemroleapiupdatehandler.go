package systemRoleApi

import (
	"net/http"

	"Pandax/panda-gateway/internal/logic/systemRoleApi"
	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func SystemRoleApiUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SystemRoleApiPostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := systemRoleApi.NewSystemRoleApiUpdateLogic(r.Context(), ctx)
		resp, err := l.SystemRoleApiUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
