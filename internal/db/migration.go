package db
import (
	"context"
	"log"
)

func Migration() {
	query := `
	CREATE TABLE IF NOT EXISTS vehicle_locations (
		id SERIAL PRIMARY KEY,
		vehicle_id VARCHAR(20),
		latitude DOUBLE PRECISION,
		longitude DOUBLE PRECISION,
		timestamp BIGINT
	);
	`
	
	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to run migration: ", err)
	}

	log.Println("Database migration completed successfully ✅")
}