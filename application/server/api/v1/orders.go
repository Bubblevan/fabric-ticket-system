package api

import (
	"backend/middleware"
	"backend/model"
	"backend/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrdersRoutes(r *gin.Engine, db *gorm.DB) {
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.POST("/createOrder", func(c *gin.Context) {
		var order Order

		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// 查询对应的票
		var ticket model.Ticket
		if err := db.First(&ticket, order.TicketID).Error; err != nil {
			log.Printf("Error finding ticket with ID %d: %v", order.TicketID, err)
			c.JSON(404, gin.H{"error": "Ticket not found"})
			return
		}
		log.Printf("Found ticket: %+v", ticket)

		// 判断剩余票数是否足够
		if ticket.Num < uint(order.Num) {
			c.JSON(400, gin.H{"error": "Not enough tickets available"})
			return
		}

		// 更新票的数量
		ticket.Num -= uint(order.Num)
		if err := db.Save(&ticket).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update ticket quantity"})
			return
		}

		// 将订单信息存入数据库
		newOrder := model.Order{
			ID:         order.ID,
			UserID:     order.UserID,
			TicketID:   order.TicketID,
			Num:        order.Num,
			TotalPrice: float64(order.Num) * ticket.Price,
			OrderDate:  order.OrderDate,
		}
		result := db.Create(&newOrder)
		if result.Error != nil {
			utils.Fail(c, http.StatusInternalServerError, "创建订单失败")
			log.Printf("Database error: %v", result.Error)
			return
		}

		// 调用智能合约创建订单
		args := [][]byte{
			[]byte(newOrder.ID),
			[]byte(strconv.FormatUint(uint64(newOrder.UserID), 10)),
			[]byte(strconv.FormatUint(uint64(newOrder.TicketID), 10)),
			[]byte(strconv.Itoa(newOrder.Num)),
			[]byte(strconv.FormatFloat(newOrder.TotalPrice, 'f', -1, 64)),
			[]byte(newOrder.OrderDate.Format(time.RFC3339)),
		}
		response, err := ChannelExecute("CreateOrder", args)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Order created", "result": string(response.Payload)})
	})
	r.GET("/queryOrder/:id", func(c *gin.Context) {
		orderID := c.Param("id")
		response, err := ChannelQuery("QueryOrder", [][]byte{[]byte(orderID)})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"result": string(response.Payload)})
	})
	r.GET("queryAllOrders", func(c *gin.Context) {
		response, err := ChannelQuery("QueryAllOrders", [][]byte{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"result": string(response.Payload)})
	})
	r.POST("/changeTicketsOwner", func(c *gin.Context) {
		var request struct {
			OrderID   string `json:"orderID"`
			NewUserID uint   `json:"newUserID"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		args := [][]byte{
			[]byte(request.OrderID),
			[]byte(strconv.FormatUint(uint64(request.NewUserID), 10)),
		}

		_, err := ChannelExecute("ChangeTicketsOwner", args)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Tickets owner changed successfully"})
	})
	r.GET("/queryOrdersByUserID", func(c *gin.Context) {
		userIDStr := c.Query("userID")
		userID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid userID"})
			return
		}

		args := [][]byte{
			[]byte(strconv.FormatUint(userID, 10)),
		}

		response, err := ChannelQuery("QueryOrdersByUserID", args)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"result": string(response.Payload)})
	})
	r.GET("/queryOrdersByTicketID", func(c *gin.Context) {
		ticketIDStr := c.Query("ticketID")
		ticketID, err := strconv.ParseUint(ticketIDStr, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ticketID"})
			return
		}

		args := [][]byte{
			[]byte(strconv.FormatUint(ticketID, 10)),
		}

		response, err := ChannelQuery("QueryOrdersByTicketID", args)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"result": string(response.Payload)})
	})
}
