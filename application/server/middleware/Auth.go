package middleware

import (
	models "backend/model"
	"backend/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		log.Println("Middleware Token String:", tokenString)

		if tokenString == "" {
			utils.Fail(c, http.StatusUnauthorized, "中间件显示用户未登录")
			c.Abort()
			return
		}

		// 去掉 "Bearer " 前缀
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		log.Println("Trimmed Token String:", tokenString)

		// 解析和验证 JWT
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			log.Println("Token Parsing Error:", err)
			utils.Fail(c, http.StatusUnauthorized, "无效的token")
			c.Abort()
			return
		}

		log.Println("Parsed Claims:", claims)

		// 在上下文中存储用户信息
		var user models.User
		if err := utils.DB.First(&user, claims.UserID).Error; err != nil {
			log.Println("User Query Error:", err)
			utils.Fail(c, http.StatusUnauthorized, "用户不存在")
			c.Abort()
			return
		}

		log.Println("User Found:", user)

		c.Set("currentUser", user)
		c.Next()
	}
}
