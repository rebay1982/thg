package main 

import (
	"fmt"
	"os"
	"flag"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var h MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Println("*** Received Message ***")
	fmt.Printf("  TOPIC: %s\n", msg.Topic())
	fmt.Printf("  MSG: %s\n", msg.Payload())
}

var hostname string
var username string
var password string

func initFlags() {
	flag.StringVar(&hostname, "host", "localhost", "Hostname of the running Mosquitto broker")
	flag.StringVar(&username, "username", "", "Username for authentication against the broker")
	flag.StringVar(&password, "password", "", "Password for authentication against the broker")

	flag.Parse()
}

func main() {
	fmt.Println("[+] Parsing command line flags.")
	initFlags()

	fmt.Println("[+] Setting up connection options")

	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:1883", hostname))
	opts.SetClientID("gosquitto")
	opts.SetDefaultPublishHandler(h)

	if username != "" {
		opts.SetUsername(username)
		opts.SetPassword(password)
	}

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("gosquitto/test", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	produce(c)

	if token := c.Unsubscribe("gosquitto/test"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}

func produce(client MQTT.Client) {
	fmt.Println("[+] Producing stuff")
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("this is msg #%d", i)
		token := client.Publish("gosquitto/test", 0, false, msg)
		token.Wait()
	}
	fmt.Println("[+] Done...")
}
