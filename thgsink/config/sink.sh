#!/bin/bash
export THG_INFLUXDB_WRITE_TOKEN=""
export THG_INFLUXDB_URL="http://:8086"
export THG_INFLUXDB_ORGANIZATION="THG"
export THG_INFLUXDB_BUCKET="THG"

export THG_MQTT_HOSTNAME=""
export THG_MQTT_TOPIC="thg/thg-data"
export THG_MQTT_CLIENT_NAME="thg-sink-client"
export THG_MQTT_USER=""
export THG_MQTT_PASS=""

/home/rbaydoun/sink/thgsink 1>/dev/null 2>>/var/log/thgsink.log
