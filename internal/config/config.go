package config

import (
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/suyuan32/simple-admin-common/utils/captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth         rest.AuthConf
	DatabaseConf config.DatabaseConf
	RedisConf    redis.RedisConf
	CasbinConf   casbin.CasbinConf
	MmsRpc       zrpc.RpcClientConf
	Captcha      captcha.Conf
	CoreRpc      zrpc.RpcClientConf
	ProjectConf  ProjectConf
	CROSConf     config.CROSConf
}

type ProjectConf struct {
	UseCaptcha    bool
	DefaultRankId uint64
}
