package api

import (
	"backend/model"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/register", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			utils.Fail(c, http.StatusBadRequest, "请求失败: "+err.Error())
			return
		}

		result := db.Create(&user)
		if result.Error != nil {
			utils.Fail(c, http.StatusInternalServerError, "注册失败: "+result.Error.Error())
			return
		}

		if result.RowsAffected == 0 {
			utils.Fail(c, http.StatusConflict, "用户名或邮箱已存在")
			return
		}

		utils.Success(c, nil, "注册成功")
		c.Redirect(http.StatusFound, "/")
	})
}
