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
See the README.md file under `./mosquitto`.

## InfluxDB
See the README.md file under `./influxdb`.

## Grafana
See the README.md file under `./grafana`.

## ESP8266 / Temperature and Humidity probe
See the README.md file under `./arduino`.

## Sink service (Go)
See the README.md file under `./thgsink`
