package main 
import (
	"log"
	"fleet-management/internal/api"
	"github.com/gin-gonic/gin"
	"fleet-management/internal/db"
	"fleet-management/internal/mqtt"
	"fleet-management/internal/rabbitmq"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.ConnectDB()
	db.Migration()
	db.Seed()
	rabbitmq.ConnectRabbitMQ()
	go mqtt.StartSubscriber()
	go rabbitmq.StartConsumer()

	router := gin.Default()
	api.RegisterRoutes(router)

	router.Run(":8080")
}