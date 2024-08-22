package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var (
	SDK           *fabsdk.FabricSDK
	channelClient *channel.Client
	channelName   = "mychannel"
	chaincodeName = "fabasg"
	orgName       = "Org1"
	orgAdmin      = "Admin"
	org1Peer0     = "peer0.org1.example.com"
	org2Peer0     = "peer0.org2.example.com"
)

type Order struct {
	ID         string    `json:"id"`
	UserID     uint      `json:"userID"`
	TicketID   uint      `json:"ticketID"`
	Num        int       `json:"num"`
	TotalPrice float64   `json:"totalPrice"`
	OrderDate  time.Time `json:"orderDate"`
}

func ChannelExecute(funcName string, args [][]byte) (channel.Response, error) {
	var err error
	configPath := "./config.yaml"
	configProvider := config.FromFile(configPath)
	SDK, err := fabsdk.New(configProvider)
	if err != nil {
		log.Fatalf("Failed to create new SDK: %s", err)
	}
	ctx := SDK.ChannelContext(channelName, fabsdk.WithOrg(orgName), fabsdk.WithUser(orgAdmin))
	channelClient, err = channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}

	response, err := channelClient.Execute(channel.Request{
		ChaincodeID: chaincodeName,
		Fcn:         funcName,
		Args:        args,
	})
	if err != nil {
		return response, err
	}
	SDK.Close()
	return response, nil
}

func main() {
	r := gin.Default()
	r.GET("/queryAllOrders", func(c *gin.Context) {
		var result channel.Response
		result, err := ChannelExecute("queryAllOrders", [][]byte{})
		fmt.Println(result)
		if err != nil {
			log.Fatalf("Failed to evaluate transaction: %s\n", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"message": "Query All Success",
			"result":  string(result.Payload),
		})
	})
	r.POST("/createOrder", func(c *gin.Context) {
		var order Order
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    "400",
				"message": "Invalid request payload",
				"error":   err.Error(),
			})
			return
		}
		var result channel.Response
		result, err := ChannelExecute("CreateOrder", [][]byte{
			[]byte(order.ID),
			[]byte(strconv.FormatUint(uint64(order.UserID), 10)),
			[]byte(strconv.FormatUint(uint64(order.TicketID), 10)),
			[]byte(strconv.Itoa(order.Num)),
			[]byte(strconv.FormatFloat(order.TotalPrice, 'f', -1, 64)),
			[]byte(order.OrderDate.Format(time.RFC3339)),
		})
		fmt.Println(result)
		if err != nil {
			log.Printf("Failed to evaluate transaction: %s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    "400",
				"message": "Create Failed",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"message": "Create Success",
			"result":  string(result.Payload),
		})
	})
	r.Run(":9099")
}

func ChannelQuery(funcName string, args [][]byte) (channel.Response, error) {
	configPath := "./config.yaml"
	configProvider := config.FromFile(configPath)
	SDK, err := fabsdk.New(configProvider)
	if err != nil {
		log.Fatalf("Failed to create new SDK: %s", err)
	}
	defer SDK.Close()

	ctx := SDK.ChannelContext(channelName, fabsdk.WithOrg(orgName), fabsdk.WithUser(orgAdmin))
	channelClient, err := channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}

	response, err := channelClient.Query(channel.Request{
		ChaincodeID: chaincodeName,
		Fcn:         funcName,
		Args:        args,
	})
	if err != nil {
		return response, err
	}

	return response, nil
}

// // CreateOrderHandler 创建订单的处理器
// func CreateOrderHandler(c *gin.Context) {
// 	var order struct {
// 		ID        string `json:"id"`
// 		UserID    uint   `json:"userID"`   // 直接接收uint类型
// 		TicketID  uint   `json:"ticketID"` // 直接接收uint类型
// 		Num       int    `json:"num"`
// 		OrderDate string `json:"orderDate"`
// 	}

// 	if err := c.ShouldBindJSON(&order); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// 查询对应的票
// 	var ticket model.Ticket
// 	if err := utils.DB.First(&ticket, order.TicketID).Error; err != nil {
// 		log.Printf("Error finding ticket with ID %d: %v", order.TicketID, err)
// 		c.JSON(404, gin.H{"error": "Ticket not found"})
// 		return
// 	}
// 	log.Printf("Found ticket: %+v", ticket)

// 	// 判断剩余票数是否足够
// 	if ticket.Num < uint(order.Num) {
// 		c.JSON(400, gin.H{"error": "Not enough tickets available"})
// 		return
// 	}

// 	// 更新票的数量
// 	ticket.Num -= uint(order.Num)
// 	if err := utils.DB.Save(&ticket).Error; err != nil {
// 		c.JSON(500, gin.H{"error": "Failed to update ticket quantity"})
// 		return
// 	}

// 	// 将字符串形式的 OrderDate 转换为 time.Time 类型
// 	orderDateTime, err := time.Parse(time.RFC3339, order.OrderDate)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "Invalid order date format"})
// 		return
// 	}

// 	// 将订单信息存入数据库
// 	newOrder := model.Order{
// 		ID:         order.ID,
// 		UserID:     order.UserID,
// 		TicketID:   order.TicketID,
// 		Num:        order.Num,
// 		TotalPrice: float64(order.Num) * ticket.Price,
// 		OrderDate:  orderDateTime,
// 	}
// 	result := utils.DB.Create(&newOrder)
// 	if result.Error != nil {
// 		utils.Fail(c, http.StatusInternalServerError, "创建订单失败")
// 		log.Printf("Database error: %v", result.Error)
// 		return
// 	}

// 	// 调用智能合约创建订单
// 	orderJSON, err := json.Marshal(newOrder)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Failed to marshal order"})
// 		return
// 	}

// 	args := [][]byte{
// 		[]byte(orderJSON),
// 	}
// 	response, err := ChannelExecute("CreateOrder", args)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, gin.H{"message": "Order created", "result": string(response.Payload)})
// }

// // QueryOrderHandler 查询订单的处理器
// func QueryOrderHandler(c *gin.Context) {
// 	orderID := c.Param("id")
// 	response, err := ChannelQuery("QueryOrder", [][]byte{[]byte(orderID)})
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"result": string(response.Payload)})
// }

// // QueryOrdersByUserIDHandler 按UserID查询订单的处理器
// func QueryOrdersByUserIDHandler(c *gin.Context) {
// 	userID := c.Param("userID")
// 	response, err := ChannelQuery("QueryOrdersByUserID", [][]byte{[]byte(userID)})
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"result": string(response.Payload)})
// }

// // QueryOrdersByTicketIDHandler 按TicketID查询订单的处理器
// func QueryOrdersByTicketIDHandler(c *gin.Context) {
// 	ticketID := c.Param("ticketID")
// 	response, err := ChannelQuery("QueryOrdersByTicketID", [][]byte{[]byte(ticketID)})
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"result": string(response.Payload)})
// }

// // QueryOrdersByOrderDateHandler 按OrderDate查询订单的处理器
// func QueryOrdersByOrderDateHandler(c *gin.Context) {
// 	orderDate := c.Param("orderDate")
// 	response, err := ChannelQuery("QueryOrdersByOrderDate", [][]byte{[]byte(orderDate)})
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"result": string(response.Payload)})
// }

// // QueryOrderByIdInUserHandler 在特定的UserID下按id查询订单的处理器
// func QueryOrderByIdInUserHandler(c *gin.Context) {
// 	userID := c.Param("userID")
// 	orderID := c.Param("orderID")
// 	response, err := ChannelQuery("QueryOrderByIdInUser", [][]byte{[]byte(userID), []byte(orderID)})
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"result": string(response.Payload)})
// }
