package svc

import (
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"github.com/suyuan32/simple-admin-common/utils/captcha"
	"github.com/suyuan32/simple-admin-core/rpc/coreclient"
	"github.com/suyuan32/simple-admin-message-center/mcmsclient"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/suyuan32/simple-admin-member-rpc/mmsclient"

	"github.com/suyuan32/simple-admin-member-api/internal/config"
	i18n2 "github.com/suyuan32/simple-admin-member-api/internal/i18n"
	"github.com/suyuan32/simple-admin-member-api/internal/middleware"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	Casbin    *casbin.Enforcer
	Authority rest.Middleware
	Vip       rest.Middleware
	Trans     *i18n.Translator
	MmsRpc    mmsclient.Mms
	McmsRpc   mcmsclient.Mcms
	Captcha   *base64Captcha.Captcha
	CoreRpc   coreclient.Core
	Redis     redis.UniversalClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := c.RedisConf.MustNewUniversalRedis()

	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(c.I18nConf, i18n2.LocaleFS)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Vip:       middleware.NewVipMiddleware().Handle,
		Trans:     trans,
		Redis:     rds,
		Casbin:    cbn,
		MmsRpc:    mmsclient.NewMms(zrpc.NewClientIfEnable(c.MmsRpc)),
		Captcha:   captcha.MustNewOriginalRedisCaptcha(c.Captcha, rds),
		CoreRpc:   coreclient.NewCore(zrpc.NewClientIfEnable(c.CoreRpc)),
		McmsRpc:   mcmsclient.NewMcms(zrpc.NewClientIfEnable(c.McmsRpc)),
	}
}
