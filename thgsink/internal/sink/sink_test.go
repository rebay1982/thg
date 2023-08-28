package sink

import (
	"testing"
)

func Test_MQTTSink(t *testing.T) {
	t.Run("getClientOptions", func(t *testing.T) {
		mqttSink := new(MQTTSink)
		config := MQTTConfig {
			Hostname: "www.test.com",
			Topic: "test/topic"
			ClientName: "MQTT Test Client",
			Username: "user",
			Password: "password",
		}

		opts := mqttSink.getClientOptions(config)
		
		if opts.ClientID != "MQTT Test Client" {
			t.Errorf("Failed client name validation. Got %s, want %s", opts.ClientID, "MQTT Test Client")
		}
		if len(opts.Servers) != 1 {
			t.Errorf("Failed server validation. Got %d, want %d", len(opts.Servers), 1)
		} else {
			if opts.Servers[0].String() != "tcp://www.test.com:1883" {
				t.Errorf("Failed server url validation. Got %s, want %s", opts.Servers[0].String(), "tcp://www.test.com:1883")
			}
		}
		if opts.Username != "user" {
			t.Errorf("Failed user name validation. Got %s, want %s", opts.Username, "user")
		}
		if opts.Password != "password" {
			t.Errorf("Failed password validation. Got %s, want %s", opts.Password, "password")
		}
	})
}
