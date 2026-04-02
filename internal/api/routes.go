package api
import (
	"github.com/gin-gonic/gin"
	"fleet-management/internal/handler"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "System successfully running...",
		})
	})

	router.GET("/vehicles/:vehicle_id/location", handler.GetLatestLocation)

	router.GET("/vehicles/:vehicle_id/history", handler.GetLocationHistory)
}