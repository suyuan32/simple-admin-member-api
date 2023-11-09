package vipdemo

import (
	"context"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestVipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestVipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestVipLogic {
	return &TestVipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *TestVipLogic) TestVip() (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
