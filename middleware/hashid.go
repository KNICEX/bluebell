package middleware

import (
	"github.com/gin-gonic/gin"
	"myBulebell/pkg/hashid"
	"myBulebell/pkg/serializer"
	"net/http"
)

func HashID(IDType int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Param("id") != "" {
			id, err := hashid.DecodeId(c.Param("id"), IDType)
			if err == nil {
				c.Set("object_id", id)
				c.Next()
				return
			}
			c.JSON(http.StatusOK, serializer.ParamErr("Failed to parse object ID", nil))
			c.Abort()
			return
		}
		c.Next()
	}
}
