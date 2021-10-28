package admin

import (
	"net/http"

	"Pandax/panda-gateway/internal/logic/admin"
	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func AdminAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminPostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := admin.NewAdminAddLogic(r.Context(), ctx)
		resp, err := l.AdminAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
