package publicmember

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-member-api/internal/logic/publicmember"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
)

// swagger:route post /member/reset_password_by_email publicmember ResetPasswordByEmail
//
// Reset password by Email | 通过邮箱重置密码
//
// Reset password by Email | 通过邮箱重置密码
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ResetPasswordByEmailReq
//
// Responses:
//  200: BaseMsgResp

func ResetPasswordByEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ResetPasswordByEmailReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publicmember.NewResetPasswordByEmailLogic(r.Context(), svcCtx)
		resp, err := l.ResetPasswordByEmail(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
