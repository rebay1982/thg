package api

// Structure representing JSON data received from the MQTT broker for a temperature and himidity measurement.
type THGMeasurement struct {
	Timestamp   int64   `json:"timestamp"`
	Sensor      string  `json:"sensor_id"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}
