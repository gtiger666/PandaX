package systemDepartment

import (
	"net/http"

	"Pandax/panda-gateway/internal/logic/systemDepartment"
	"Pandax/panda-gateway/internal/svc"
	"Pandax/panda-gateway/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func SystemDepartmentParentListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SystemDepartmentListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := systemDepartment.NewSystemDepartmentParentListLogic(r.Context(), ctx)
		resp, err := l.SystemDepartmentParentList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
