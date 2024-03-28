package bootstrap

import (
	"myBulebell/models"
	"myBulebell/pkg/auth"
	"myBulebell/pkg/cache"
	"myBulebell/pkg/conf"
	"myBulebell/pkg/email"
	"myBulebell/pkg/logger"
	"myBulebell/pkg/snowflake"
)

func Init() {
	conf.Init()
	logger.Init()
	models.Init()
	cache.Init()
	auth.Init()
	email.Init()
	snowflake.Init()
}

func Shutdown() error {
	return cache.Store.Close()
}
