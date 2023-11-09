package publicmember

import (
	"context"
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/jwt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"
	"github.com/zeromicro/go-zero/core/errorx"
	"time"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByEmailLogic {
	return &LoginByEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *LoginByEmailLogic) LoginByEmail(req *types.LoginByEmailReq) (resp *types.LoginResp, err error) {
	if l.svcCtx.Config.ProjectConf.LoginVerify != "email" && l.svcCtx.Config.ProjectConf.LoginVerify != "sms_or_email" &&
		l.svcCtx.Config.ProjectConf.LoginVerify != "all" {
		return nil, errorx.NewCodeAbortedError("login.loginTypeForbidden")
	}

	captchaData, err := l.svcCtx.Redis.Get("CAPTCHA_" + req.Email)
	if err != nil {
		logx.Errorw("failed to get captcha data in redis for email validation", logx.Field("detail", err),
			logx.Field("data", req))
		return nil, errorx.NewCodeInvalidArgumentError(i18n.Failed)
	}

	if captchaData == req.Captcha {
		memberData, err := l.svcCtx.MmsRpc.GetMemberList(l.ctx, &mms.MemberListReq{
			Page:     1,
			PageSize: 1,
			Email:    &req.Email,
		})
		if err != nil {
			return nil, err
		}

		if memberData.Total == 0 {
			return nil, errorx.NewCodeInvalidArgumentError("login.userNotExist")
		}

		_, err = l.svcCtx.Redis.Del("CAPTCHA_" + req.Email)
		if err != nil {
			logx.Errorw("failed to delete captcha in redis", logx.Field("detail", err))
		}

		token, err := jwt.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(),
			l.svcCtx.Config.Auth.AccessExpire, jwt.WithOption("userId", memberData.Data[0].Id), jwt.WithOption("rankId",
				memberData.Data[0].RankCode), jwt.WithOption("roleId", "invalid"))
		if err != nil {
			return nil, err
		}

		// get rank data
		if len(MemberRankData) == 0 {
			err = l.genRankCache()
			if err != nil {
				return nil, err
			}
		}

		// add token into database
		expiredAt := time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)).Unix()
		_, err = l.svcCtx.MmsRpc.CreateToken(l.ctx, &mms.TokenInfo{
			Uuid:      memberData.Data[0].Id,
			Token:     pointy.GetPointer(token),
			Source:    pointy.GetPointer("mms_member"),
			Status:    pointy.GetPointer(uint32(common.StatusNormal)),
			ExpiredAt: pointy.GetPointer(expiredAt),
		})

		if err != nil {
			return nil, err
		}

		return &types.LoginResp{
			BaseDataInfo: types.BaseDataInfo{
				Code: 0,
				Msg:  l.svcCtx.Trans.Trans(l.ctx, "login.loginSuccessTitle"),
				Data: "",
			},
			Data: types.LoginInfo{
				UserId:   *memberData.Data[0].Id,
				RankId:   *memberData.Data[0].RankCode,
				Token:    token,
				Expire:   uint64(expiredAt),
				Avatar:   *memberData.Data[0].Avatar,
				Nickname: *memberData.Data[0].Nickname,
				RankName: l.svcCtx.Trans.Trans(l.ctx, MemberRankData[*memberData.Data[0].RankId]),
			},
		}, nil
	}

	return nil, errorx.NewInvalidArgumentError("login.wrongCaptcha")
}

// genRankCache used to generate cache for member rank to improve performance
func (l *LoginByEmailLogic) genRankCache() error {
	list, err := l.svcCtx.MmsRpc.GetMemberRankList(l.ctx, &mms.MemberRankListReq{
		Page:     1,
		PageSize: 1000,
	})
	if err != nil {
		return err
	}

	for _, v := range list.Data {
		MemberRankData[*v.Id] = *v.Name
	}

	return err
}
