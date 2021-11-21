package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: handle custom headers to expose. if there is any .

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, PATCH, POST, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, accept, Content-Length")
		// c.Writer.Header().Set("Access-Control-Expose-Headers", "")
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
