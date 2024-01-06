# Grafana
Grafana will be used to display metrics and visualize data that has been captured by the different sensors.

## Installation
Installation is as simple as pulling the OSS image from Docker hub using
```
docker run -d --name=grafana -p 3000:3000 grafana/grafana-oss
```

There are other options to take into account like providing credentials, setting up default dashboards, etc.
These will be reviewed under the configuration section below.

## Configuration
Configuration options abound. We can refer to the
[official documentation](https://grafana.com/docs/grafana/latest/setup-grafana/configure-docker/) for a Docker
installation.

### Providing Credentials
Credentials can be provided to Grafana by using (Docker secrets)
[https://grafana.com/docs/grafana/latest/setup-grafana/configure-docker/#configure-grafana-with-docker-secrets]. In this
project, none were provided and the admin password was setup once though the UI once the container was lauched.

### Configuring Dashboards
The dashboard configuration will depend on how the data is pushed into InfluxDB. In the case of this project, data is 
published as follows.

```
"time": <...>
"measurement": "thg_measurement",
"tags": {
    "sensor_id": <...>
},
"fields": {
    "temperature": <...>,
    "humidity": <...>
}

```

Grafana uses the Flux query language (for an InfluxDB datasource) to retrieve the temperature and humidity values 
for each sensor. Below are the queries used for this project.
```
# Temperature
from(bucket: "Test")
  |> range(start: v.timeRangeStart, stop:v.timeRangeStop)
  |> filter(fn: (r) => r._measurement == "thg_measurement" and r._field == "temperature")
  |> keep(columns: ["_time", "_value", "sensor_id"])

# Humidity
from(bucket: "Test")
  |> range(start: v.timeRangeStart, stop:v.timeRangeStop)
  |> filter(fn: (r) => r._measurement == "thg_measurement" and r._field == "humidity")
  |> keep(columns: ["_time", "_value", "sensor_id"])

```
