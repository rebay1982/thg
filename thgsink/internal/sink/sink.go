package sink

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTConfig struct {
	Hostname string
	Topic    string
	// Nice to have: Add port.
	ClientName string
	Username   string
	Password   string
}

type MQTTSink struct {
	client mqtt.Client
}

func NewMQTTSink(config MQTTConfig) *MQTTSink {
	sink := new(MQTTSink)
	opts := sink.getClientOptions(config)
	client := mqtt.NewClient(opts)

	sink.client = client
	return sink
}

// getClientOPtions Method to take a MQTTConfig struct and return mqqt client options.
func (s MQTTSink) getClientOptions(config MQTTConfig) *mqtt.ClientOptions {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:1883", config.Hostname))

	opts.SetClientID(config.ClientName)
	if config.Username != "" {
		opts.SetUsername(config.Username)
		opts.SetPassword(config.Password)
	}

	return opts
}

// ConnectClient Method to connect the MQTT client to the MQTT broker.
func (s *MQTTSink) ConnectClient() error {
	if token := s.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// DisconnectClient Method to disconnect the MQTT client from the MQTT broker.
func (s *MQTTSink) DisconnectClient() {
	// Wait 1s for work to be completed.
	s.client.Disconnect(1000)
}

// Subscribe Method to subsrcibe to an MQTT topic, providing the message handler to be used.
func (s *MQTTSink) Subscribe(topic string, handler mqtt.MessageHandler) error {

	// QOS 0: Fire and forget to minimize complexity for IoT application.
	if token := s.client.Subscribe(topic, 0, handler); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

// Unsubscribe Method to unsubscribe from an MQTT topic.
func (s *MQTTSink) Unsubscribe(topic string) error {
	if token := s.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}
