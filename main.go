package main 
import (
	"log"
	"fleet-management/internal/api"
)

func main() {
	router := gin.Default()
	api.RegisterRoutes(router)

	router.Run(":8080")
}