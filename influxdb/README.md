# InfluxDB
InfluxDB was chosen as the time series data base. The sensors in this IoT application will just be pushing data at
regular time intervals, with temperature and humidity readings. This is essentially time series data. On top of it,
this seems to integrate well and easily in Grafana.

## Installing and configuration
At the time of this writing, the latest InfluxDB version was InfluxDB v2.7. InfluxData has a pretty straight forward 
[guide](https://docs.influxdata.com/influxdb/v2.7/install/?t=Docker) on how to set this up so there is no reason to
repeat it here.

We'll be going for an installation that allows us to persist the InfluxDB data outside of the container if we ever need
to move it for some reason.

To have the container persist onto the host file system, we first create the destination directory and run the container
by providing the `--volume` parameter.
```
mkdir <influxdb data path> && \
docker run \
-it -d \
--name influxdb \
-p 8086:8086 \
--volume <influxdb data path>:/var/lib/influxdb2 \
influxdb:2.7.0
```

With the container running, we can now go through the onboarding and initial setup at `<host>:8086/`. We'll create a
'THG' bucket to hold the data for our project and that is about it.
