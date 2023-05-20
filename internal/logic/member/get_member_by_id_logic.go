package member

import (
	"context"

	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMemberByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberByIdLogic {
	return &GetMemberByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMemberByIdLogic) GetMemberById(req *types.UUIDReq) (resp *types.MemberInfoResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetMemberById(l.ctx, &mms.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.MemberInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.MemberInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:   data.Status,
			Username: data.Username,
			Password: data.Password,
			Nickname: data.Nickname,
			RankId:   data.RankId,
			Mobile:   data.Mobile,
			Email:    data.Email,
			Avatar:   data.Avatar,
		},
	}, nil
}
