package publicmember

import (
	"context"
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"time"

	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/suyuan32/simple-admin-common/utils/jwt"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

var MemberRankData = make(map[uint64]string)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if l.svcCtx.Config.ProjectConf.LoginVerify != "captcha" && l.svcCtx.Config.ProjectConf.LoginVerify != "all" {
		return nil, errorx.NewCodeAbortedError("login.loginTypeForbidden")
	}

	var isPass bool

	if !l.svcCtx.Config.ProjectConf.UseCaptcha {
		isPass = true
	} else {
		isPass = l.svcCtx.Captcha.Verify(config.RedisCaptchaPrefix+req.CaptchaId, req.Captcha, true)
	}

	if isPass {
		user, err := l.svcCtx.MmsRpc.GetMemberByUsername(l.ctx,
			&mms.UsernameReq{
				Username: req.Username,
			})
		if err != nil {
			return nil, err
		}

		if !encrypt.BcryptCheck(req.Password, *user.Password) {
			return nil, errorx.NewCodeInvalidArgumentError("login.wrongUsernameOrPassword")
		}

		// check whether is expired
		if (time.Now().UnixMilli() - *user.ExpiredAt) >= 0 {
			return nil, errorx.NewCodeAbortedError("login.expiredAccount")
		}

		if user.Status != nil && *user.Status != uint32(common.StatusNormal) {
			return nil, errorx.NewCodeInvalidArgumentError("login.userBanned")
		}

		// Check the remaining available time
		expire := l.svcCtx.Config.Auth.AccessExpire
		if (*user.ExpiredAt - time.Now().UnixMilli()) < expire*1000 {
			expire = (*user.ExpiredAt - time.Now().UnixMilli()) / 1000
		}

		token, err := jwt.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(),
			expire, jwt.WithOption("userId", user.Id), jwt.WithOption("rankId",
				user.RankCode), jwt.WithOption("roleId", "invalid"))

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
		expiredAt := time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)).UnixMilli()
		_, err = l.svcCtx.MmsRpc.CreateToken(l.ctx, &mms.TokenInfo{
			Uuid:      user.Id,
			Token:     pointy.GetPointer(token),
			Source:    pointy.GetPointer("mms_member"),
			Status:    pointy.GetPointer(uint32(common.StatusNormal)),
			Username:  user.Username,
			ExpiredAt: pointy.GetPointer(expiredAt),
		})

		if err != nil {
			return nil, err
		}

		resp = &types.LoginResp{
			BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, "login.loginSuccessTitle")},
			Data: types.LoginInfo{
				UserId:   *user.Id,
				Token:    token,
				RankId:   *user.RankCode,
				Nickname: *user.Nickname,
				RankName: l.svcCtx.Trans.Trans(l.ctx, MemberRankData[*user.RankId]),
				Avatar:   *user.Avatar,
				Expire:   uint64(expiredAt),
			},
		}
		return resp, nil
	} else {
		return nil, errorx.NewCodeError(errorcode.InvalidArgument, "login.wrongCaptcha")
	}
}

// genRankCache used to generate cache for member rank to improve performance
func (l *LoginLogic) genRankCache() error {
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
