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
	McmsRpc      zrpc.RpcClientConf
	CoreRpc      zrpc.RpcClientConf
	Captcha      captcha.Conf
	ProjectConf  ProjectConf
	CROSConf     config.CROSConf
}

type ProjectConf struct {
	UseCaptcha              bool
	DefaultRankId           uint64
	EmailCaptchaExpiredTime int    `json:",default=600"`
	SmsTemplateId           string `json:",optional"`
	SmsAppId                string `json:",optional"`
	SmsSignName             string `json:",optional"`
	RegisterVerify          string `json:",default=captcha,options=[captcha,email,sms,sms_or_email]"`
	LoginVerify             string `json:",default=captcha,options=[captcha,email,sms,sms_or_email,all]"`
	ResetVerify             string `json:",default=email,options=[email,sms,sms_or_email]"`
	WechatMiniOauthProvider string `json:",default=wechat_mini"`
	AllowInit               bool   `json:",default=true"`
}
