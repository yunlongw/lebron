package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/zhoushuguang/lebron/apps/app/api/internal/logic/user"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"
	"github.com/zhoushuguang/lebron/pkg/result"
)

func AddReceiveAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReceiveAddressAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewAddReceiveAddressLogic(r.Context(), svcCtx)
		resp, err := l.AddReceiveAddress(&req)
		result.HttpResult(r, w, resp, err)
	}
}
