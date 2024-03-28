package email

import (
	"myBulebell/pkg/conf"
	"myBulebell/pkg/logger"
)

var Client Driver

func Init() {
	if conf.EmailConf.User == "" {
		return
	}

	logger.L().Debug("Initializing email sending queue...")

	Client = NewSMTPClient()
}
