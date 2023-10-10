# THGSink

This project serves as the glue between MQTT and the InfluxDB backend. The image can be found on docker hub under 
`rebay1982/thgsink`. 


## Configuration

The project requires configuration of eight environment variables:
`THG_INFLUXDB_WRITE_TOKEN` A token configured in Influx to allow the service to write to the time series database.
`THG_INFLUXDB_URL` The URL to the Influx database.
`THG_INFLUXDB_ORGANIZATION` The organization configured in Influx.
`THG_INFLUXDB_BUCKET` The Influx bucket to store the data into.

`THG_MQTT_HOSTNAME` The MQTT broker's hostname.
`THG_MQTT_TOPIC` The MQTT topic to listen to.
`THG_MQTT_CLIENT_NAME` The MQTT client name.
`THG_MQTT_USER` The configured MQTT user name.
`THG_MQTT_PASS` The confgured MQTT password.

These need to be configured in a configuration file and used as a volume to the `/thgsink/thgsink.conf` file.


## Building

If you want to build your own image or version, it is possible to do so by using the provided Makefile and use
the `docker-build` target.


## Running

Run the `thgsink` container using the following command and also providing the necessary configuration file:
`docker run -it -d --name thgsink -v ./config/thgsink.conf:/thgsink/thgsink.conf rebay1982/thgsink:latest`
