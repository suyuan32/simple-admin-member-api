package base

import (
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

func (l *InitDatabaseLogic) insertApiData() (err error) {
	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/member/create"),
		Description: pointy.GetPointer("apiDesc.createMember"),
		ApiGroup:    pointy.GetPointer("member"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/member/update"),
		Description: pointy.GetPointer("apiDesc.updateMember"),
		ApiGroup:    pointy.GetPointer("member"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/member/delete"),
		Description: pointy.GetPointer("apiDesc.deleteMember"),
		ApiGroup:    pointy.GetPointer("member"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/member/list"),
		Description: pointy.GetPointer("apiDesc.getMemberList"),
		ApiGroup:    pointy.GetPointer("member"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/member"),
		Description: pointy.GetPointer("apiDesc.getMemberById"),
		ApiGroup:    pointy.GetPointer("member"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	// MEMBER RANK

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/member_rank/create"),
		Description: pointy.GetPointer("apiDesc.createMemberRank"),
		ApiGroup:    pointy.GetPointer("member_rank"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/member_rank/update"),
		Description: pointy.GetPointer("apiDesc.updateMemberRank"),
		ApiGroup:    pointy.GetPointer("member_rank"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/member_rank/delete"),
		Description: pointy.GetPointer("apiDesc.deleteMemberRank"),
		ApiGroup:    pointy.GetPointer("member_rank"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/member_rank/list"),
		Description: pointy.GetPointer("apiDesc.getMemberRankList"),
		ApiGroup:    pointy.GetPointer("member_rank"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/member_rank"),
		Description: pointy.GetPointer("apiDesc.getMemberRankById"),
		ApiGroup:    pointy.GetPointer("member_rank"),
		Method:      pointy.GetPointer("Post"),
	})

	if err != nil {
		return err
	}

	return nil
}

func (l *InitDatabaseLogic) insertMenuData() error {
	menuData, err := l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:     pointy.GetPointer(uint32(2)),
		ParentId:  pointy.GetPointer(common.DefaultParentId),
		Name:      pointy.GetPointer("MemberManagementDirectory"),
		Component: pointy.GetPointer("LAYOUT"),
		Sort:      pointy.GetPointer(uint32(1)),
		Disabled:  pointy.GetPointer(false),
		Meta: &core.Meta{
			Title:              pointy.GetPointer("route.memberManagement"),
			Icon:               pointy.GetPointer("ic:round-person-outline"),
			HideMenu:           pointy.GetPointer(false),
			HideBreadcrumb:     pointy.GetPointer(false),
			IgnoreKeepAlive:    pointy.GetPointer(false),
			HideTab:            pointy.GetPointer(false),
			CarryParam:         pointy.GetPointer(false),
			HideChildrenInMenu: pointy.GetPointer(false),
			Affix:              pointy.GetPointer(false),
		},
		MenuType: pointy.GetPointer(uint32(0)),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:     pointy.GetPointer(uint32(2)),
		ParentId:  pointy.GetPointer(menuData.Id),
		Path:      pointy.GetPointer("/member"),
		Name:      pointy.GetPointer("MemberManagement"),
		Component: pointy.GetPointer("/mms/member/index"),
		Sort:      pointy.GetPointer(uint32(1)),
		Disabled:  pointy.GetPointer(false),
		Meta: &core.Meta{
			Title:              pointy.GetPointer("route.memberManagement"),
			Icon:               pointy.GetPointer("ic:round-person-outline"),
			HideMenu:           pointy.GetPointer(false),
			HideBreadcrumb:     pointy.GetPointer(false),
			IgnoreKeepAlive:    pointy.GetPointer(false),
			HideTab:            pointy.GetPointer(false),
			CarryParam:         pointy.GetPointer(false),
			HideChildrenInMenu: pointy.GetPointer(false),
			Affix:              pointy.GetPointer(false),
		},
		MenuType: pointy.GetPointer(uint32(1)),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:     pointy.GetPointer(uint32(2)),
		ParentId:  pointy.GetPointer(menuData.Id),
		Path:      pointy.GetPointer("/member_rank"),
		Name:      pointy.GetPointer("MemberRankManagement"),
		Component: pointy.GetPointer("/mms/memberRank/index"),
		Sort:      pointy.GetPointer(uint32(2)),
		Disabled:  pointy.GetPointer(false),
		Meta: &core.Meta{
			Title:              pointy.GetPointer("route.memberRankManagement"),
			Icon:               pointy.GetPointer("ic:round-person-outline"),
			HideMenu:           pointy.GetPointer(false),
			HideBreadcrumb:     pointy.GetPointer(false),
			IgnoreKeepAlive:    pointy.GetPointer(false),
			HideTab:            pointy.GetPointer(false),
			CarryParam:         pointy.GetPointer(false),
			HideChildrenInMenu: pointy.GetPointer(false),
			Affix:              pointy.GetPointer(false),
		},
		MenuType: pointy.GetPointer(uint32(1)),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:     pointy.GetPointer(uint32(2)),
		ParentId:  pointy.GetPointer(menuData.Id),
		Path:      pointy.GetPointer("/member_token"),
		Name:      pointy.GetPointer("MemberTokenManagement"),
		Component: pointy.GetPointer("/mms/token/index"),
		Sort:      pointy.GetPointer(uint32(3)),
		Disabled:  pointy.GetPointer(false),
		Meta: &core.Meta{
			Title:              pointy.GetPointer("route.tokenManagement"),
			Icon:               pointy.GetPointer("ant-design:lock-outlined"),
			HideMenu:           pointy.GetPointer(false),
			HideBreadcrumb:     pointy.GetPointer(false),
			IgnoreKeepAlive:    pointy.GetPointer(false),
			HideTab:            pointy.GetPointer(false),
			CarryParam:         pointy.GetPointer(false),
			HideChildrenInMenu: pointy.GetPointer(false),
			Affix:              pointy.GetPointer(false),
		},
		MenuType: pointy.GetPointer(uint32(1)),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:     pointy.GetPointer(uint32(2)),
		ParentId:  pointy.GetPointer(menuData.Id),
		Path:      pointy.GetPointer("/member_oauth"),
		Name:      pointy.GetPointer("MemberOauthManagement"),
		Component: pointy.GetPointer("/mms/oauth/index"),
		Sort:      pointy.GetPointer(uint32(3)),
		Disabled:  pointy.GetPointer(false),
		Meta: &core.Meta{
			Title:              pointy.GetPointer("route.oauthManagement"),
			Icon:               pointy.GetPointer("ant-design:unlock-filled"),
			HideMenu:           pointy.GetPointer(false),
			HideBreadcrumb:     pointy.GetPointer(false),
			IgnoreKeepAlive:    pointy.GetPointer(false),
			HideTab:            pointy.GetPointer(false),
			CarryParam:         pointy.GetPointer(false),
			HideChildrenInMenu: pointy.GetPointer(false),
			Affix:              pointy.GetPointer(false),
		},
		MenuType: pointy.GetPointer(uint32(1)),
	})

	if err != nil {
		return err
	}

	return err
}
