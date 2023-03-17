package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CustomLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("", params.ClientIP, params.BodySize, params.Request, params.Method, params.Path)
	})
}
