package rabbitmq
import (
	"encoding/json"
	"log"
)

type GeofenceAlert struct {
	VehicleID string `json:"vehicle_id"`
	Event     string `json:"event"`
	Location  struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	Timestamp int64 `json:"timestamp"`
}

func StartConsumer() {
	messages, err := ch.Consume(
		"geofence_alerts", // queue
		"",                // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)

	if err != nil {
		log.Fatal("❌ Failed to register consumer:", err)
	}

	log.Println("Worker started, waiting for messages...")

	forever := make(chan bool)
	go func() {
		for d := range messages {
			log.Println("Received data from rabbit event:", string(d.Body))
			var alert GeofenceAlert
			err := json.Unmarshal(d.Body, &alert)
			if err != nil {
				log.Printf("❌ Error parsing message: %v\n", err)
				d.Nack(false, false)
				continue
			}
			log.Printf("✅ [Geofence Alert radius in 50 meters] - Vehicle: %s, Location: (%f, %f), Timestamp: %d\n",
				alert.VehicleID, 
				alert.Location.Latitude, 
				alert.Location.Longitude, 
				alert.Timestamp,
			)
			d.Ack(false)
		}
	}()
	<-forever
}