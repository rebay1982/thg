package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	thgapi "github.com/rebay1982/thg/thgsink/api"
	"github.com/rebay1982/thg/thgsink/internal/sink"
	"github.com/rebay1982/thg/thgsink/internal/persistence"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// getenv Tool to allow fallback values when retrieving environment variables.
func getenv(key, fallback string) (string) {
	value := os.Getenv(key)
	if len(value) > 0 {
		return value
	}
	return fallback
}

func getConfigurations() (persistence.InfluxConfig, sink.MQTTConfig) {
	persistConfig := persistence.InfluxConfig{
		Token: getenv("THG_INFLUXDB_WRITE_TOKEN", ""),
		Url: getenv("THG_INFLUXDB_URL", ""),
		Org: getenv("THG_INFLUXDB_ORGANIZATION", ""),
		Bucket: getenv("THG_INFLUXDB_BUCKET", ""),
	}
	sinkConfig := sink.MQTTConfig{
		Hostname: getenv("THG_MQTT_HOSTNAME", ""),
		Topic: getenv("THG_MQTT_TOPIC", ""),
		ClientName: getenv("THG_MQTT_CLIENT_NAME", ""),
		Username: getenv("THG_MQTT_USER", ""),
		Password: getenv("THG_MQTT_PASS", ""),
	}

	return persistConfig, sinkConfig
}

var influxPersister *persistence.InfluxPersister
func mqttMessageHandler(client mqtt.Client, msg mqtt.Message) {
	jsonPayload := msg.Payload()
	var measurement thgapi.THGMeasurement

	if err := json.Unmarshal([]byte(jsonPayload), &measurement); err != nil {
		log.Printf("Failed to unmarshall json payload into measurement: %s", jsonPayload)
	}

	influxPersister.WriteTHGData(measurement)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	// Initialize configuration, persister, and sink
	persistConfig, sinkConfig := getConfigurations()
	influxPersister = persistence.NewInfluxPersister(persistConfig)
	mqttSink := sink.NewMQTTSink(sinkConfig)

	// Connect sink client
	if err :=	mqttSink.ConnectClient(); err != nil {
		log.Printf("Unable to connect MQTT client:\n%v", err)
		os.Exit(1)
	}

	// Subscribe sink to topic
	if err := mqttSink.Subscribe(sinkConfig.Topic, mqttMessageHandler); err != nil {
		log.Printf("Unable to subscribe to MQTT topic %s:\n%v", sinkConfig.Topic, err)
		os.Exit(1)
	}

	log.Println("Listening for messages...")
	sig := <-sigs

	log.Printf("Gracefully exiting after receiving %v\n", sig)
	if err := mqttSink.Unsubscribe(sinkConfig.Topic); err != nil {
		log.Printf("Unable to ubsubrcrube from MQTT topic %s:\n%v", sinkConfig.Topic, err)
	}
	mqttSink.DisconnectClient()

	log.Println("Exited.")
}
