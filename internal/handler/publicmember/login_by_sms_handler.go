package publicmember

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-member-api/internal/logic/publicmember"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
)

// swagger:route post /member/login_by_sms publicmember LoginBySms
//
// Log in by SMS | 短信登录
//
// Log in by SMS | 短信登录
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: LoginBySmsReq
//
// Responses:
//  200: LoginResp

func LoginBySmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginBySmsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publicmember.NewLoginBySmsLogic(r.Context(), svcCtx)
		resp, err := l.LoginBySms(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
