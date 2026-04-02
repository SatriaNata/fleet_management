package geofence
import (
	"math"
	"os"
	"log"
	"strconv"
)


/* Calculate radius */
func Haversine(lat1, lng1, lat2, lng2 float64) float64 {
	
	const R = 6371000 // Earth radius in meter
	dLat := (lat2 - lat1) * math.Pi / 180
	dLng := (lng2 - lng1) * math.Pi / 180
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*math.Sin(dLng/2)*math.Sin(dLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func IsInsideGeofence(lat, lng float64) bool {

	strgeofenceLat := os.Getenv("GEOFENCE_LAT")
	strgeofenceLng := os.Getenv("GEOFENCE_LNG")
	strradius := os.Getenv("GEOFENCE_RADIUS")

	geofenceLat, _ := strconv.ParseFloat(strgeofenceLat, 64)
	geofenceLng, _ := strconv.ParseFloat(strgeofenceLng, 64)
	radius, _ := strconv.ParseFloat(strradius, 64)

	distance := Haversine(geofenceLat, geofenceLng, lat, lng)
	log.Printf("Geofence Center: (%f, %f)", geofenceLat, geofenceLng)
	log.Printf("Vehicle Location: (%f, %f)", lat, lng)
	log.Printf("Calculated distance from geofence center: %.2f meters", distance)
	return distance <= radius
}