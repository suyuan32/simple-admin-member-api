package vipdemo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-member-api/internal/logic/vipdemo"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"
)

// swagger:route get /member/vvip vipdemo TestVvip
//
// Test VVIP authorization | 测试 VVIP 授权
//
// Test VVIP authorization | 测试 VVIP 授权
//
// Responses:
//  200: BaseMsgResp

func TestVvipHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := vipdemo.NewTestVvipLogic(r.Context(), svcCtx)
		resp, err := l.TestVvip()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
