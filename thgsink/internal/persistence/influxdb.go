package persistence

import (
	"time"

	thgapi "github.com/rebay1982/thg/thgsink/api"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

// InfluxConfig struct containing necessary configurations to instanciate an InfluxDB2 client and an async write API.
type InfluxConfig struct {
	Token  string
	Url    string
	Org    string
	Bucket string
}

type InfluxPersister struct {
	writeAPI api.WriteAPI
}

// NewInfluxPersister Method to create a new InfluxPersister, used to write data points to InfluxDB.
func NewInfluxPersister(config InfluxConfig) *InfluxPersister {
	client := influxdb2.NewClient(config.Url, config.Token)

	influxPersister := new(InfluxPersister)
	influxPersister.writeAPI = client.WriteAPI(config.Org, config.Bucket)

	return influxPersister
}

// WriteTHGData Method receives a THGMeasurement data structure and persists it to InfluxDB.
func (p *InfluxPersister) WriteTHGData(data thgapi.THGMeasurement) {
	measurement, tags, fields := p.unpackMeasurement(data)

	measurementTime := time.Unix(data.Timestamp, 0).UTC()
	point := write.NewPoint(measurement, tags, fields, measurementTime)
	p.writeAPI.WritePoint(point)
}

func (p InfluxPersister) unpackMeasurement(data thgapi.THGMeasurement) (
	measurement string, tags map[string]string, fields map[string]interface{}) {

	measurement = "thg_measurement"
	tags = map[string]string{
		"sensor_id": data.Sensor,
	}

	fields = map[string]interface{}{
		"temperature": data.Temperature,
		"humidity":    data.Humidity,
	}
	return measurement, tags, fields
}
