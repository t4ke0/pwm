package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, PATCH, POST, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, accept, Content-Length, token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "token")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// BasePATH
const BasePATH string = "/api"

func GetBasePATHroute(engine *gin.Engine) *gin.RouterGroup {
	return engine.Group(BasePATH)
}
