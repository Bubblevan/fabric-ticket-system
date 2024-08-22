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

func TicketsRoutes(r *gin.Engine, db *gorm.DB) {
	// 受保护的路由
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.POST("/createTicket", func(c *gin.Context) {
		var ticket model.Ticket
		if err := c.ShouldBindJSON(&ticket); err != nil {
			utils.Fail(c, http.StatusBadRequest, "无效请求")
			log.Printf("Binding error: %v", err)
			return
		}

		// 获取当前登录用户
		loggedInUser, exists := c.Get("currentUser")
		if !exists {
			utils.Fail(c, http.StatusUnauthorized, "用户未登录")
			log.Println("User not logged in")
			return
		}
		user := loggedInUser.(model.User)

		// 检查角色
		if user.Role == nil || !*user.Role {
			utils.Fail(c, http.StatusUnauthorized, "无权限")
			log.Println("User does not have permission")
			return
		}

		log.Printf("Creating ticket: %+v", ticket)

		// 将门票信息存入数据库
		result := db.Create(&ticket)
		if result.Error != nil {
			utils.Fail(c, http.StatusInternalServerError, "创建门票失败")
			log.Printf("Database error: %v", result.Error)
			return
		}

		utils.Success(c, ticket, "创建门票成功")
	})

	r.GET("/tickets/:id", func(c *gin.Context) {
		ticketID := c.Param("id")
		var ticket model.Ticket
		result := db.First(&ticket, ticketID)
		if result.Error != nil {
			utils.Fail(c, http.StatusNotFound, "门票未找到")
			log.Printf("Database error: %v", result.Error)
			return
		}
		utils.Success(c, ticket, "获取门票成功")
	})

	r.GET("/tickets", func(c *gin.Context) {
		var tickets []model.Ticket
		result := db.Find(&tickets)
		if result.Error != nil {
			utils.Fail(c, http.StatusInternalServerError, "获取门票列表失败")
			return
		}
		utils.Success(c, tickets, "获取门票列表成功")
	})
}
