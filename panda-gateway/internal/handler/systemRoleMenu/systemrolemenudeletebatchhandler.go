package systemRoleMenu

import (
	"net/http"

	"Pandax/panda-gateway/internal/logic/systemRoleMenu"
	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func SystemRoleMenuDeleteBatchHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SystemRoleMenuDelBatchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := systemRoleMenu.NewSystemRoleMenuDeleteBatchLogic(r.Context(), ctx)
		resp, err := l.SystemRoleMenuDeleteBatch(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
