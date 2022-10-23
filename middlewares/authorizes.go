package middlewares

import (
	"net/http"
	"strings"

	"github.com/Thanaporn4226/Project-sa-65/services"
	"github.com/gin-gonic/gin"
)

func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "No Authorization header provided",
			})
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Incorrect Format of Authorization Token"})
			return
		}

		jwtWrapper := services.JwtWrapper{
			SecretKey: "Secret",
			Issuer:    "AuthService",
		}

		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("user_id", claims.User_id)
		c.Set("role_name", claims.Role_name)

		c.Next()
	}
}
