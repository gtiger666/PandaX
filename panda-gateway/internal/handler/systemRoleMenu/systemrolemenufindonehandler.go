package systemRoleMenu

import (
	"net/http"

	"Pandax/panda-gateway/internal/logic/systemRoleMenu"
	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func SystemRoleMenuFindOneHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SystemRoleMenuDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := systemRoleMenu.NewSystemRoleMenuFindOneLogic(r.Context(), ctx)
		resp, err := l.SystemRoleMenuFindOne(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
