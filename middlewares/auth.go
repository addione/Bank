package middlewares

import (
	"net/http"

	"github.com/addione/New/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No TokenFound"})
		return
	}

	user, err := helpers.NewHelpersDIContainer().GetJwtTokenHelper().VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token"})
	}
	context.Set("userInfo", user)
	context.Next()
}
