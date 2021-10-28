package systemRoleApi

import (
	"net/http"

	"Pandax/panda-gateway/internal/logic/systemRoleApi"
	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func SystemRoleApiListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SystemRoleApiListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := systemRoleApi.NewSystemRoleApiListLogic(r.Context(), ctx)
		resp, err := l.SystemRoleApiList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
