package routes

import (
	"github.com/gin-gonic/gin"
	"myBulebell/middleware"
	"myBulebell/pkg/auth"
	"myBulebell/pkg/conf"
	"myBulebell/pkg/hashid"
	"myBulebell/routes/controller"
)

func Init() *gin.Engine {

	if conf.AppConf.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	base := r.Group(conf.ServerConf.Prefix)
	{
		base.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		user := base.Group("/user")
		{
			user.POST("/register", controller.UserRegister)
			user.GET("/activate/:id",
				middleware.SignRequired(auth.General),
				middleware.HashID(hashid.UserID),
				controller.UserActive,
			)
		}
	}

	return r
}
