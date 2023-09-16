package persistence

import (
	"reflect"
	"testing"

	thgapi "github.com/rebay1982/thg/thgsink/api"
)

func assertMeasurement(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s measurement, wanted %s measurement", got, want)
	}
}

func assertTags(t *testing.T, got map[string]string, want map[string]string) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s tags, wanted %s tags", got, want)
	}
}

func assertFields(t *testing.T, got map[string]interface{}, want map[string]interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s fields, wanted %s fields", got, want)
	}
}

func Test_InfluxPersister(t *testing.T) {

	t.Run("unpackMeasurement", func(t *testing.T) {
		influxPersister := new(InfluxPersister)
		measurement := thgapi.THGMeasurement{
			Timestamp:   0,
			Sensor:      "Test Sensor",
			Temperature: 25.0,
			Humidity:    50.0,
		}
		expected_fields := map[string]interface{}{
			"temperature": measurement.Temperature,
			"humidity":    measurement.Humidity,
		}
		expected_tags := map[string]string{
			"sensor_id": "Test Sensor",
		}

		got_measurement, got_tags, got_fields := influxPersister.unpackMeasurement(measurement)

		assertMeasurement(t, got_measurement, "thg_measurement")
		assertTags(t, got_tags, expected_tags)
		assertFields(t, got_fields, expected_fields)
	})
}
