package api

import (
	"backend/model"
	"backend/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/login", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			utils.Fail(c, http.StatusBadRequest, "无效请求")
			return
		}

		var foundUser model.User
		result := db.Where("username = ? AND role = ?", user.Username, user.Role).First(&foundUser)
		if result.Error != nil {
			utils.Fail(c, http.StatusUnauthorized, "用户名或密码错误")
			return
		}

		if foundUser.Password != user.Password {
			utils.Fail(c, http.StatusUnauthorized, "用户名或密码错误")
			return
		}

		// 生成JWT token
		token, err := utils.GenerateToken(foundUser.ID)
		if err != nil {
			utils.Fail(c, http.StatusInternalServerError, "生成token失败")
			return
		}

		// 记录token到日志
		log.Printf("Generated token: %s\n", token)

		utils.Success(c, gin.H{"token": token}, "登录成功")
	})
}
