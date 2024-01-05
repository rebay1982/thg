package main 

import (
	"fmt"
	"os"
	"flag"
	MQTT "github.com/eclipse/paho.mqtt.golang"

	"math/rand"
	"time"
	"log"
	"context"
	"strconv"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)


// MQTT STUFFZ
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

func testMQTT() {
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

func doInflux() {
	// Initialize a client to InfluxDB
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://thoth.local:8086"
	client := influxdb2.NewClient(url, token)


	org := "THG"
	bucket := "Test"
	writeAPI := client.WriteAPIBlocking(org, bucket)


	writeNewData := false

	if writeNewData {
		fmt.Println("[+] Starting write to DB")
		for sensorId := 0; sensorId < 5; sensorId++ {
			fmt.Printf("  [+] Writing data for sensor %d\n", sensorId)
			for value := 0; value < 5; value++ {
				tags := map[string]string {
					"sensor": strconv.Itoa(sensorId),
				}

				fields := map[string]interface{} {
					"field1": value,
				}

				temperature := write.NewPoint("measurement1", tags, fields, time.Now())
				time.Sleep(1 * time.Second)

				if err := writeAPI.WritePoint(context.Background(), temperature); err != nil {
					log.Fatal(err)
				}
			}
		}
		fmt.Println("[+] Done writing...")
	}

	// FETCH DATA FROM INFLUX
	fmt.Println("[+] Waiting...")
	time.Sleep(5 * time.Second)

	queryAPI := client.QueryAPI(org)
	query := `from(bucket: "Test")
							|> range(start: -10m)
						  |> filter(fn: (r) => r._measurement == "measurement1")`

	fmt.Println("[+] Starting to query the DB (single)")
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("[+] Done querying the DB (single)")



	fmt.Println("[+] Waiting...")
	time.Sleep(5 * time.Second)

	fmt.Println("[+] Starting to query the DB (aggregate)")
	queryAgg := `from(bucket: "Test")
							|> range(start: -20m)
							|> filter(fn: (r) => r._measurement == "measurement1")
							|> mean()`
	results, err = queryAPI.Query(context.Background(), queryAgg)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("[+] Done querying the DB (aggregate)")


}

func previousStuff() {
	// doInflux()
	// Initialize a client to InfluxDB
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://thoth.local:8086"
	client := influxdb2.NewClient(url, token)


	org := "THG"
	bucket := "Test"
	writeAPI := client.WriteAPIBlocking(org, bucket)

	writeNewData := true
	rand.Seed(time.Now().UnixNano())
	for writeNewData {
		for sensorId := 0; sensorId < 5; sensorId++ {
			tags := map[string]string {
				"sensor_id": strconv.Itoa(sensorId),
			}

			// Get random value between 20 and 25
			temp := rand.Intn(5) + 20
			hg := rand.Intn(3) + 45
			fields := map[string]interface{} {
				"temperature": temp,
				"humidity": hg,
			}

			point := write.NewPoint("thg_measurement", tags, fields, time.Now())
			t := time.Now()
			if err := writeAPI.WritePoint(context.Background(), point); err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("%s >>> Wrote measurement (temp: %d, hg: %d) for sensor %d\n", t.Format(time.DateTime), temp, hg, sensorId)
			}
		}

		time.Sleep(30 * time.Second)
	}
}

func produce(client MQTT.Client) {
	fmt.Println("[+] Producing stuff")
	payload := `{"timestamp": %d, "sensor_id": "test_sensor", "temperature": 25.%d, "humidity": 4%d.0}`

	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf(payload, time.Now().Unix(), i, i)
		fmt.Println(msg)
		time.Sleep(1*time.Second)
		token := client.Publish("thg/thg-data", 0, false, msg)
		token.Wait()
	}
	fmt.Println("[+] Done...")
}

func main() {
	initFlags()
	fmt.Println("[+] Setting up connection options")

	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:1883", hostname))
	opts.SetClientID("gosquitto")

	opts.SetUsername(username)
	opts.SetPassword(password)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	produce(c)

	c.Disconnect(250)
}
