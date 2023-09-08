package publicmember

import (
	"context"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.BaseMsgResp, err error) {
	if l.svcCtx.Config.ProjectConf.LoginVerify != "captcha" {
		return nil, errorx.NewCodeAbortedError(i18n.PermissionDeny)
	}

	if ok := l.svcCtx.Captcha.Verify("CAPTCHA_"+req.CaptchaId, req.Captcha, true); ok {
		user, err := l.svcCtx.MmsRpc.CreateMember(l.ctx,
			&mms.MemberInfo{
				Username: &req.Username,
				Password: &req.Password,
				Email:    &req.Email,
				Nickname: &req.Username,
				Status:   pointy.GetPointer(uint32(1)),
				RankId:   pointy.GetPointer(l.svcCtx.Config.ProjectConf.DefaultRankId),
			})
		if err != nil {
			return nil, err
		}
		resp = &types.BaseMsgResp{
			Msg: l.svcCtx.Trans.Trans(l.ctx, user.Msg),
		}
		return resp, nil
	} else {
		return nil, errorx.NewCodeError(errorcode.InvalidArgument,
			l.svcCtx.Trans.Trans(l.ctx, "login.wrongCaptcha"))
	}
}
