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

These need to be configured in an environment file and passed along to the `--env-file` when starting the container.


## Building

If you want to build your own image or version, it is possible to do so by using the provided `Makefile` and use
the `docker-build` target.

If you want to build from source without building the docker image, you can use the `build` target. You can also run
local unit tests using the `test` target.

## Running

Although this project was meant to run in a container, there is nothing preventing it from being run natively.
Environment variables will still need to be set prior to running the binary.

To run using docker, the docker image can either be built locally or pulled from the `rebay1982` repo as
`thgsink:latest`.

Run the `thgsink` container using the following command while also providing the necessary environment file.
`docker run -it -d --name thgsink --env-file <path_to_env_file> rebay1982/thgsink:latest`
