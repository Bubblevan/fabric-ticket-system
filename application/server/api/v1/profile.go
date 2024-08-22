package api

import (
	"backend/middleware"
	"backend/model"
	"backend/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProfileRoutes(r *gin.Engine, db *gorm.DB) {
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/profile", func(c *gin.Context) {
		log.Println("Entering ProfileRoutes") // 调试信息1：进入ProfileRoutes

		loggedInUser, exists := c.Get("currentUser")
		if !exists {
			log.Println("currentUser does not exist in context") // 调试信息2：currentUser不存在
			utils.Fail(c, http.StatusUnauthorized, "profile显示用户未登录")
			log.Println("User not logged in") // 调试信息3：用户未登录
			return
		}

		user, ok := loggedInUser.(model.User)
		if !ok {
			log.Println("Failed to assert loggedInUser as models.User") // 调试信息4：断言失败
			utils.Fail(c, http.StatusInternalServerError, "服务器内部错误")
			return
		}

		log.Printf("Current User: %+v\n", user) // 调试信息5：当前用户信息

		utils.Success(c, user, "获取个人信息成功")
	})
}
