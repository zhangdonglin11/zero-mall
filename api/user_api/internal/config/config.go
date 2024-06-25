package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	SMS struct {
		AccessKeyID     string
		AccessKeySecret string
	}
	Redis redis.RedisConf

	UserRpc zrpc.RpcClientConf
}
