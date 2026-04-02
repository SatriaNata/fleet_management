package repository
import (
	"context"
	"fleet-management/internal/db"
)

type VehicleLocation struct {
	VehicleID *string  `json:"vehicle_id"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Timestamp *int64   `json:"timestamp"`
}

func InsertLocation(vehicle_id string, latitude, longitude float64, timestamp int64) error {
	query := `INSERT INTO vehicle_locations (vehicle_id, latitude, longitude, timestamp) VALUES ($1, $2, $3, $4)`
	_, err := db.DB.Exec(context.Background(), query, vehicle_id, latitude, longitude, timestamp)
	return err
}

func GetLatestLocation(vehicle_id string) (*VehicleLocation, error) {
	query := `SELECT vehicle_id, latitude, longitude, timestamp FROM vehicle_locations WHERE vehicle_id = $1 ORDER BY timestamp DESC LIMIT 1`
	row := db.DB.QueryRow(context.Background(), query, vehicle_id)

	var loc VehicleLocation
	err := row.Scan(&loc.VehicleID, &loc.Latitude, &loc.Longitude, &loc.Timestamp)
	if err != nil {
		return nil, err
	}

	return &loc, nil
}

func GetLocationHistory(vehicle_id string, start string, end string) ([]VehicleLocation, error) {
	query := `SELECT vehicle_id, latitude, longitude, timestamp FROM vehicle_locations WHERE vehicle_id = $1 AND timestamp >= $2 AND timestamp <= $3 ORDER BY timestamp DESC`
	rows, err := db.DB.Query(context.Background(), query, vehicle_id, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []VehicleLocation
	for rows.Next() {
		var loc VehicleLocation
		err := rows.Scan(&loc.VehicleID, &loc.Latitude, &loc.Longitude, &loc.Timestamp)
		if err != nil {
			return nil, err
		}
		history = append(history, loc)
	}

	return history, nil
}