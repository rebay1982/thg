# thg
Temperature and Hygrometer IoT project.


# Setup, installation, and configuration
The initial intention was to use docker for both the MQTT broker and InfluxDB but the only Raspberry Pi that was laying
around was an original B+. Resources were quite limited. On top of having limited resources, the majority of the
services needed for the project have long moved on from the old `armhf` 32bit architecture.

This made everything an uphill battle so after some thought, the decision to replace the Raspberry Pi with something
else was made. The rational is that the project isn't about cross platform compilation and making things work on an
obsolete architecture. The Raspberry Pi B+ was replaced with a Trigkey S5 5500U mini PC sporting a widely supported
`x86_64` architecture along with a hefty upgrade in resources and potential at a relatively low price.


## OS Installation
The original idea was running the Rapsberry Pi as a headless server. This hasn't changed after moving on to the new 
mini PC. Archlinux was chosen as the prime candidate for a barebones console-only installation.

1. Install Archlinux by following the Arch Wiki documentation.
2. Install SSH, configure it to only work public key authentication, and
[harden](https://www.ssh-audit.com/hardening_guides.html) it.
3. Install Avahi with the `nss-mdns` plugin to make the server discoverable on the local network.
4. Install Docker engine and docker-compose (might come in handy later).

This is a living section of the document. As more services are installed, steps will be added to it. A good candidate
for addition would be Kubernetes if the need for an orchestrator comes up (or simply to play around it with it). These
steps aren't specific to thet project, but lay down the foundations for the infrastructure on which the project with 
run.


## Mosquitto
Mosquitto will be the message broker, receiving data from publishers (ESP8266/DTH11) and making them available to
subscribers to consume. It's pretty light weight and well suited for IoT integrations.

The Eclipse Foundation maintains a Docker image for Mosquitto, making it super convenient to run the broker as a
container. See the details [here](https://hub.docker.com/_/eclipse-mosquitto/).


### Installing
Fetch the latest image with
```
docker pull eclipse-mosquitto
```


### Running
Run the container with the `-v <conf_file>:/mosquitto/config/mosquitto.conf` parametter to provide a customized 
configuration file.


### Authenticated access to the broker
**WIP**
Enabling authenticated access to the broker might be a bit more complicated than running it natively. I will have to
look into this.

To configure remote access with authentication, run the following replacing USERNAME with a user name of your choosing.
```
sudo mosquitto_passwd -c /etc/mosquitto/passwd <USERNAME>
```


### Configuring
As discussed above, configuration is done by providing the mosquitto.conf file as a parameter to the run command. The
basic configuration is described in this section.

This should be added to the absolute top of the file:
```
per_listener_settings true
```

and the following block at the end of the file:
```
allow_anonymous false
listener 1883
password_file /etc/mosquitto/passwd
```
It prevents anonymous access to the broker, sets the MTTQ listener to port `1883`, and specifies the authentication
file to use. See the [Authenticated accesss to the broker](#authenticated-access-to-the-broker) section for more
details on setting this up.


## InfluxDB
**WIP**
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
