package api

import (
	"backend/model"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User = model.User

func UsersRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/users", func(c *gin.Context) { GetUsers(c, db) })
	r.GET("/users/:id", func(c *gin.Context) { GetUser(c, db) })
	r.POST("/users", func(c *gin.Context) { CreateUser(c, db) })
	r.PUT("/users/:id", func(c *gin.Context) { UpdateUser(c, db) })
	r.DELETE("/users/:id", func(c *gin.Context) { DeleteUser(c, db) })
}

// 获取所有用户
func GetUsers(c *gin.Context, db *gorm.DB) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, users, "Users retrieved successfully")
}

// 根据ID获取用户
func GetUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		utils.Fail(c, http.StatusNotFound, "User not found")
		return
	}
	utils.Success(c, user, "User retrieved successfully")
}

// 创建新用户
func CreateUser(c *gin.Context, db *gorm.DB) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := db.Create(&user).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, user, "User created successfully")
}

// 更新用户信息
func UpdateUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		utils.Fail(c, http.StatusNotFound, "User not found")
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := db.Save(&user).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, user, "User updated successfully")
}

// 删除用户
func DeleteUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := db.Delete(&User{}, id).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, nil, "User deleted successfully")
}
