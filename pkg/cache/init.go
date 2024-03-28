package cache

import (
	"fmt"
	"myBulebell/pkg/conf"
	"myBulebell/pkg/logger"
)

var Store Driver = NewLocalStore()

func Init() {
	if conf.RedisConf.Host != "" {
		Store = NewRedisStore(
			conf.RedisConf.Network,
			fmt.Sprintf("%s:%d", conf.RedisConf.Host, conf.RedisConf.Port),
			conf.RedisConf.User,
			conf.RedisConf.Password,
			conf.RedisConf.DB,
			conf.RedisConf.PoolSize,
		)
	}
	err := Store.Ping()
	if err != nil {
		logger.L().Panic(err)
	}
	if err = Store.Restore(DefaultCacheFile); err != nil {
		logger.L().Warn(err)
	}
}
