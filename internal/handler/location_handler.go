package handler

import (
	"github.com/gin-gonic/gin"
	"fleet-management/internal/service"
)

// func InsertLocation(c *gin.Context) {
// 	var req struct {
// 		VehicleID string  `json:"vehicle_id" binding:"required"`
// 		Latitude  float64 `json:"latitude" binding:"required"`
// 		Longitude float64 `json:"longitude" binding:"required"`
// 		Timestamp int64   `json:"timestamp" binding:"required"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err := service.InsertLocation(req.VehicleID, req.Latitude, req.Longitude, req.Timestamp)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, gin.H{"message": "Location inserted successfully"})
// }
// func InsertLocation(c *gin.Context) {
// 	err := service.InsertLocation("B1234XYZ", -6.2, 106.4194)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, gin.H{"message": "Location inserted successfully"})
// }

func GetLatestLocation(c *gin.Context) {
	vehicleID := c.Param("vehicle_id")

	location, err := service.GetLatestLocation(vehicleID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, location)
}

func GetLocationHistory(c *gin.Context) {
	vehicleID := c.Param("vehicle_id")
	start := c.Query("start")
	end := c.Query("end")

	data, err := service.GetLocationHistory(vehicleID, start, end)
	// data, err := service.GetLocationHistory(vehicleID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}