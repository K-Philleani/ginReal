package middleware

import (
	"ginReal/common"
	"ginReal/model"
	"ginReal/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{"tokenString": "err"}, "权限不足")
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{"Valid": "err"}, "权限不足")
			ctx.Abort()
			return
		}
		userId := claims.UserID
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{"err": "token认证失败"}, "权限不足")
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Set("currentPhonne", user.Phone)
		ctx.Next()
	}
}
