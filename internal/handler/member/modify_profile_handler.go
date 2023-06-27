package member

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-member-api/internal/logic/member"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
)

// swagger:route post /member/profile member ModifyProfile
//
// Modify users' own profile | 用户修改信息接口
//
// Modify users' own profile | 用户修改信息接口
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ModifyProfileReq
//
// Responses:
//  200: BaseMsgResp

func ModifyProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyProfileReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := member.NewModifyProfileLogic(r.Context(), svcCtx)
		resp, err := l.ModifyProfile(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
