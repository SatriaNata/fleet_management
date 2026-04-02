package db

import (
	"context"
	"log"
)

func Seed() {
	query := `
	INSERT INTO vehicle_locations (vehicle_id, latitude, longitude, timestamp)
	SELECT 'B1234XYZ', -6.2088, 106.8456, EXTRACT(EPOCH FROM NOW())::BIGINT
	WHERE NOT EXISTS (
		SELECT 1 FROM vehicle_locations WHERE vehicle_id = 'B1234XYZ'
	);
	`
	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to run seed: %v\n", err)
	}
	log.Println("Seeder executed ✅")
}