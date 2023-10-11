package publicoauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-member-api/internal/logic/publicoauth"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
)

// swagger:route post /oauth/login/wechat/mini_program publicoauth WechatMiniProgramLogin
//
// Wechat Mini program login | 微信小程序登录
//
// Wechat Mini program login | 微信小程序登录
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: WechatMiniProgramLoginReq
//
// Responses:
//  200: CallbackResp

func WechatMiniProgramLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WechatMiniProgramLoginReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publicoauth.NewWechatMiniProgramLoginLogic(r.Context(), svcCtx)
		resp, err := l.WechatMiniProgramLogin(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
