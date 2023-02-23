package member

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-member-rpc/mms"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateMemberLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateMemberLogic {
	return &CreateMemberLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateMemberLogic) CreateMember(req *types.MemberInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.MmsRpc.CreateMember(l.ctx,
		&mms.MemberInfo{
			Id:       req.Id,
			Status:   req.Status,
			Username: req.Username,
			Password: req.Password,
			Nickname: req.Nickname,
			RankId:   req.RankId,
			Mobile:   req.Mobile,
			Email:    req.Email,
			Avatar:   req.Avatar,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, data.Msg)}, nil
}
