package publicoauth

import (
	"context"
	"github.com/suyuan32/simple-admin-common/utils/jwt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-member-api/internal/logic/publicmember"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"
	"github.com/zeromicro/go-zero/core/errorx"
	"time"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WechatMiniProgramLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWechatMiniProgramLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WechatMiniProgramLoginLogic {
	return &WechatMiniProgramLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *WechatMiniProgramLoginLogic) WechatMiniProgramLogin(req *types.WechatMiniProgramLoginReq) (resp *types.CallbackResp, err error) {
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

	if data.Total == 0 {
		return nil, errorx.NewCodeInvalidArgumentError("login.bindWechatToAccount")
	}

	token, err := jwt.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire, jwt.WithOption("userId", data.Data[0].Id), jwt.WithOption("rankId",
			data.Data[0].RankCode), jwt.WithOption("roleId", "invalid"))

	// add token into database
	expiredAt := time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)).UnixMilli()
	_, err = l.svcCtx.MmsRpc.CreateToken(l.ctx, &mms.TokenInfo{
		Uuid:      data.Data[0].Id,
		Token:     pointy.GetPointer(token),
		Source:    &l.svcCtx.Config.ProjectConf.WechatMiniOauthProvider,
		Status:    pointy.GetPointer(uint32(1)),
		ExpiredAt: pointy.GetPointer(expiredAt),
	})

	if err != nil {
		return nil, err
	}

	return &types.CallbackResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, "login.loginSuccessTitle")},
		Data: types.CallbackInfo{
			UserId:   *data.Data[0].Id,
			Token:    token,
			Expire:   expiredAt,
			Avatar:   *data.Data[0].Avatar,
			RankId:   *data.Data[0].RankCode,
			RankName: publicmember.MemberRankData[*data.Data[0].RankId],
			Nickname: *data.Data[0].Nickname,
		},
	}, nil
}
