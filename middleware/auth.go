package middleware

import (
	"github.com/gin-gonic/gin"
	"myBulebell/pkg/auth"
	"myBulebell/pkg/serializer"
	"net/http"
)

func SignRequired(authInstance auth.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		switch c.Request.Method {
		case http.MethodPut, http.MethodPost, http.MethodPatch:
		//TODO check request
		default:
			err = auth.CheckURL(authInstance, c.Request.URL)
		}

		if err != nil {
			c.JSON(http.StatusOK, serializer.Err(serializer.CodeCredentialInvalid, err.Error(), err))
			c.Abort()
			return
		}
		c.Next()
	}
}
