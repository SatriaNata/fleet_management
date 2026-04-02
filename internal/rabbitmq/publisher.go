package rabbitmq
import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"time"
)

var conn *amqp.Connection
var ch *amqp.Channel

func ConnectRabbitMQ() {
	var err error
	for i := 0; i < 20; i++ {
		conn, err = amqp.Dial(os.Getenv("RABBITMQ_URL"))
		if err == nil {
			break
		}
		log.Printf("Attempt %d: ❌ Failed to connect to RabbitMQ: %v\n", i+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatal("❌ Failed to connect to RabbitMQ:", err)
	}

	ch, err = conn.Channel()
	if err != nil {
		log.Fatal("❌ Failed to open a channel:", err)
	}

	err = ch.ExchangeDeclare(
		"fleet.events", // name
		"fanout",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // app
		false,          // no-wait
		nil,            // arguments
	)

	if err != nil {
		log.Fatal("❌ Failed to declare exchange:", err)
	}
	
	//queues
	q, err := ch.QueueDeclare(
		"geofence_alerts", // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)

	if err != nil {
		log.Fatal("❌ Failed to declare queue:", err)
	}

	//binding
	err = ch.QueueBind(
		q.Name,          // queue name
		"",              // routing key
		"fleet.events",  // exchange
		false,
		nil,
	)

	if err != nil {
		log.Fatal("❌ Failed to bind queue:", err)
	}

	log.Println("✅ Connected to RabbitMQ")
}

func PublishGeofenceAlert(vehicleID string, lat, long float64, timestamp int64) {
	log.Println("Publishing to RabbitMQ...")
	alert := map[string]interface{}{
		"vehicle_id": vehicleID,
		"event": "geofence_alert",
		"location": map[string]float64{
			"latitude": lat,
			"longitude": long,
		},
		"timestamp": timestamp,
	}

	body, _ := json.Marshal(alert)
	err := ch.Publish(
		"fleet.events", // exchange
		"", 			// routing key
		false, 			// mandatory
		false, 			// immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body: body,
		},
	)
	
	if err != nil {
		log.Println("❌ Publish error:", err)
	} else {
		log.Println("✅ Publish success")
	}
}
