package memberrank

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-member-rpc/mms"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewUpdateMemberRankLogic(r *http.Request, svcCtx *svc.ServiceContext) *UpdateMemberRankLogic {
	return &UpdateMemberRankLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *UpdateMemberRankLogic) UpdateMemberRank(req *types.MemberRankInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.MmsRpc.UpdateMemberRank(l.ctx,
		&mms.MemberRankInfo{
			Id:          req.Id,
			Name:        req.Name,
			Description: req.Description,
			Remark:      req.Remark,
			Code:        req.Code,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, data.Msg)}, nil
}
