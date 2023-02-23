package config

import (
	"github.com/suyuan32/simple-admin-core/pkg/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth         rest.AuthConf
	DatabaseConf config.DatabaseConf
	RedisConf    redis.RedisConf
	CasbinConf   config.CasbinConf
	MmsRpc       zrpc.RpcClientConf
}
