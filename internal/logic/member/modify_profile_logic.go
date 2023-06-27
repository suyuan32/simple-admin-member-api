package member

import (
	"context"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	userId string
}

func NewModifyProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyProfileLogic {
	return &ModifyProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		userId: ctx.Value("userId").(string),
	}
}

func (l *ModifyProfileLogic) ModifyProfile(req *types.ModifyProfileReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.MmsRpc.UpdateMember(l.ctx, &mms.MemberInfo{
		Id:       &l.userId,
		Nickname: req.Nickname,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Avatar:   req.Avatar,
	})

	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{
		Code: 0,
		Msg:  result.Msg,
	}, nil
}
