# thg
Temperature and Hygrometer IoT project.



# Setup, installation, and configuration

The initial intention was to use docker for both the MQTT broker and InfluxDB but the only Raspberry Pi I had on hand
was an original B+ so resources are quite limited.

## Docker
The ideal situation would be to run everything as docker containers so that a rebuild/reconfiguration can be done 
easily and quickly.


### Installation
Installation on the Raspberry Pi is relatively straight forward, though there are doubts a first gen B+ Raspberry Pi 
has the horse power to run multiple things under docker.

Installation through the basic installation script. This will figure out the details about the platform and install
all the necessary components for you.
```
curl -sSL https://get.docker.com | sh
```

## Mosquitto
Mosquitto will be the message broker, receiving data from pusblishers (ESP8266/DTH11) and writting that information
into a data store (InfluxDB for now).

To install on a RaspberryPi, follow this
(guide)[https://randomnerdtutorials.com/how-to-install-mosquitto-broker-on-raspberry-pi/].

In essence, run the following to install and configure with remote privileges.
```
sudo apt update && sudo apt upgrade                 # Update apt.
sudo apt install -y mosquitto mosquitto-clients     # Install packages.
sudo systemctl enable mosquitto.service             # Enable Mosquitto broker service.
```

To configure remote access with authentication, run the following replacing USERNAME with a user name of your choosing.
```
sudo mosquitto_passwd -c /etc/mosquitto/passwd <USERNAME>
```

Then edit the `/etc/mosquitto/mosquitto.conf` configuration file to prevent anonymous remote access. Add the following
to the absolute top of the file:
```
per_listener_settings true
```

and the following block at the end of the file:
```
allow_anonymous false
listener 1883
password_file /etc/mosquitto/passwd
```

Restart the service and validate the status:
```
sudo systemctl restart mosquitto.service
sudo systemctl status mosquitto.service
```

## InfluxDB
The initial idea was to run InfluxDB in a docker container. Unfortunately, there is no image based for the 
`linux/arm/v6` architecture. It will not be possible to go down this route (easily).

The fall back option will be to run InfluxDB on the Raspberry Pi natively, like Mosquitto.

### Installation
Basic repositories don't have InfluxDB > 1.6.7-rc0. It was necessary to go to https://portal.influxdata.com/downloads/
and use the instructions to install InfluxDB 1.8.10. For the `armhf` architecture, `influxdb2` isn't available so
`influxdb` (`1.8.10`) will have to do.

```
wget -q https://repos.influxdata.com/influxdata-archive_compat.key
echo '393e8779c89ac8d958f81f942f9ad7fb82a25e133faddaf92e15b16e6ac9ce4c influxdata-archive_compat.key' | sha256sum -c && cat influxdata-archive_compat.key | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/influxdata-archive_compat.gpg > /dev/null
echo 'deb [signed-by=/etc/apt/trusted.gpg.d/influxdata-archive_compat.gpg] https://repos.influxdata.com/debian stable main' | sudo tee /etc/apt/sources.list.d/influxdata.list

sudo apt-get update && sudo apt-get install influxdb
```
