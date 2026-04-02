package service

import (
	"errors"
	"time"

	"fleet-management/internal/repository"
)

func InsertLocation(vehicleID string, lat, long float64, timestamp int64) error {
	if vehicleID == "" {
		return errors.New("vehicle_id is required")
	}
	saveTime := timestamp
	if timestamp == 0 {
		saveTime = time.Now().Unix()
	}
	return repository.InsertLocation(vehicleID, lat, long, saveTime)
}

func GetLatestLocation(vehicleID string) (interface{}, error) {
	return repository.GetLatestLocation(vehicleID)
}

func GetLocationHistory(vehicleID, start, end string) (interface{}, error) {
// func GetLocationHistory(vehicleID string) (interface{}, error) {
	// bisa tambah validasi waktu di sini
	return repository.GetLocationHistory(vehicleID, start, end)
	// return repository.GetLocationHistory(vehicleID)
}