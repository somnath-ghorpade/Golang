package middleware

import (
	"net/http"

	"GinWebServer/common"

	"github.com/gin-gonic/gin"
)

func ValidateUser(c *gin.Context) {
	tokenData := c.GetHeader("Authorization")

	if tokenData == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	err := common.DecodeJWTToken(tokenData)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.Next()
}
