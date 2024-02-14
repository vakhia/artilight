package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/pkg/token"
	"net/http"
)

func Authenticate(jwtService token.IJwtService) gin.HandlerFunc {
	return func(context *gin.Context) {
		authToken := context.Request.Header.Get("Authorization")

		if authToken == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
			return
		}

		userId, err := jwtService.VerifyToken(authToken)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
			return
		}

		context.Set("userId", userId)
		context.Next()
	}
}
