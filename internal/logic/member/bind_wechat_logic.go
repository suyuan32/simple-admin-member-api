package member

import (
	"context"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindWechatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBindWechatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindWechatLogic {
	return &BindWechatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *BindWechatLogic) BindWechat(req *types.BindWechatReq) (resp *types.BaseMsgResp, err error) {
	wechatIdData, err := l.svcCtx.MmsRpc.WechatMiniProgramLogin(l.ctx, &mms.OauthLoginReq{
		State:    req.Code,
		Provider: l.svcCtx.Config.ProjectConf.WechatMiniOauthProvider,
	})
	if err != nil {
		return nil, err
	}

	data, err := l.svcCtx.MmsRpc.GetMemberList(l.ctx, &mms.MemberListReq{Page: 1, PageSize: 1, WechatId: &wechatIdData.Msg})
	if err != nil {
		return nil, err
	}

	if data.Total != 0 {
		return nil, errorx.NewCodeInvalidArgumentError("login.alreadyBindWechatToAccount")
	}

	userId := l.ctx.Value("userId").(string)

	result, err := l.svcCtx.MmsRpc.UpdateMember(l.ctx, &mms.MemberInfo{
		Id:       &userId,
		WechatId: &wechatIdData.Msg,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: result.Msg}, nil
}
