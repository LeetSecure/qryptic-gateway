package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leetsecure/qryptic-gateway/internal/utils/auth"
)

func ControllerAuthCheckMiddleware(c *gin.Context) {
	// log := logger.Default()
	Bearer_Schema := "Bearer "
	authorisation := c.GetHeader("Authorization")

	if len(authorisation) <= len(Bearer_Schema)+1 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorised"})
		c.Abort()
		return
	}

	token := authorisation[len(Bearer_Schema):]

	_, err := auth.VerifyControllerAuthToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorised"})
		c.Abort()
		return
	}
	c.Next()
}
