package user

import (
	"github.com/gin-gonic/gin"
	"myBulebell/models"
	"myBulebell/pkg/auth"
	"myBulebell/pkg/conf"
	"myBulebell/pkg/email"
	"myBulebell/pkg/hashid"
	"myBulebell/pkg/logger"
	"myBulebell/pkg/serializer"
	"myBulebell/pkg/snowflake"
	"myBulebell/pkg/utils"
	"net/url"
	"strings"
	"time"
)

type RegisterService struct {
	Password string `json:"password" binding:"required,len=64"`
	Email    string `json:"email" binding:"required,email"`
}

func (service *RegisterService) Register() serializer.Response {

	user := models.User{
		UserId:   snowflake.GenID(),
		Username: strings.Split(service.Email, "@")[0] + "_" + utils.RandomString(5),
		Email:    service.Email,
		Status:   models.NotActivated,
	}
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(serializer.CodeEncryptError, "encrypt password error", err)
	}
	if err := models.DB.Create(&user).Error; err != nil {
		expectedUser, err := models.GetUserByEmail(service.Email)
		if expectedUser.Status == models.NotActivated {
			user = *expectedUser
			user.SetPassword(service.Password)
		} else {
			return serializer.Err(serializer.CodeEmailExist, "Email already in use", err)
		}
	}
	models.DB.Update(&user)

	base := models.GetSiteURL()
	userId := hashid.HashId(uint64(user.UserId), hashid.UserID)
	controller, _ := url.Parse(conf.ServerConf.Prefix + "/user/activate/" + userId)

	activeURL := auth.SignURL(auth.General, base.ResolveReference(controller), time.Minute*10)

	err := email.Send(user.Email, "Activate your account", activeURL.String())
	if err != nil {
		logger.L().Error("send email error", err)
		return serializer.Err(serializer.CodeEmailSendErr, "send email error", err)
	}

	return serializer.Response{}
}

type ActiveService struct {
	UserId string `uri:"id" binding:"required"`
}

func (service *ActiveService) Active(c *gin.Context) serializer.Response {
	uid, _ := c.Get("object_id")
	user, err := models.GetUserByUserId(int64(uid.(uint64)))
	if err != nil {
		return serializer.Err(serializer.CodeParamErr, "user not found", err)
	}
	if user.Status != models.NotActivated {
		return serializer.Err(serializer.CodeParamErr, "user already activated", err)
	}
	user.SetStatus(models.Active)
	return serializer.Response{}
}
