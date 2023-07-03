package base

import (
	"context"
	"strings"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	if l.svcCtx.Config.CoreRpc.Enabled {
		err = l.insertApiData()
		if err != nil {
			if strings.Contains(err.Error(), "common.createFailed") {
				return nil, errorx.NewInvalidArgumentError(i18n.AlreadyInit)
			}
			return nil, errorx.NewCodeInternalError(err.Error())
		}

		err = l.insertMenuData()
		if err != nil {
			return nil, errorx.NewCodeInternalError(err.Error())
		}

	}

	data, err := l.svcCtx.MmsRpc.InitDatabase(l.ctx, &mms.Empty{})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{
		Code: 0,
		Msg:  l.svcCtx.Trans.Trans(l.ctx, data.Msg),
	}, nil
}
