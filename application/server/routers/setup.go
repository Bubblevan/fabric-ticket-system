package routers

import (
	v1 "backend/api/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	configPath := "./config/config.yaml"
	v1.LoginRoutes(r, db)
	v1.RegisterRoutes(r, db)
	v1.BlockchainRoutes(r, configPath)
	v1.TicketsRoutes(r, db)
	v1.OrdersRoutes(r, db)
	v1.UsersRoutes(r, db)
	v1.ProfileRoutes(r, db)
}
