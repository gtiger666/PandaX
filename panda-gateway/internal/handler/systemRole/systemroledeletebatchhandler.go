package systemRole

import (
	"net/http"

	"Pandax/panda-gateway/internal/logic/systemRole"
	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func SystemRoleDeleteBatchHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SystemRoleDelBatchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := systemRole.NewSystemRoleDeleteBatchLogic(r.Context(), ctx)
		resp, err := l.SystemRoleDeleteBatch(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
