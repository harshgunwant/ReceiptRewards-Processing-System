package routes

import (
	"FetchRewardsAssessment/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all API endpoints
func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/receipts/process", handlers.ProcessReceipt)
	router.GET("/receipts/:id/points", handlers.GetPoints)

	return router
}
