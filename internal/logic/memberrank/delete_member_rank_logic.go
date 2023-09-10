package memberrank

import (
	"context"
	"github.com/suyuan32/simple-admin-member-api/internal/logic/publicmember"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMemberRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMemberRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberRankLogic {
	return &DeleteMemberRankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMemberRankLogic) DeleteMemberRank(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.MmsRpc.DeleteMemberRank(l.ctx, &mms.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	publicmember.MemberRankData = make(map[uint64]string)

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
