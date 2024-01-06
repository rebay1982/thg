# Mosquitto
Mosquitto will be the message broker, receiving data from publishers (ESP8266/DTH11) and making it available for 
subscribers to consume. Its light weight make is well suited for IoT integrations.

The Eclipse Foundation maintains a Docker image for Mosquitto, making it super convenient to run the broker as a
container. See the details [here](https://hub.docker.com/_/eclipse-mosquitto/).


## Installing
Fetch the latest image with
```
docker pull eclipse-mosquitto
```

## Configuring
The configuration is done by providing the `mosquitto.conf` file as a parameter to the `docker run` command. The
basic configuration provided here is described in this section.

This should be added to the absolute top of the file:
```
per_listener_settings true
```
This configuration specifies if the the current configuration is for the current listener only. False would make this
apply to all listeners.

The following block at the end of the file:
```
allow_anonymous false
listener 1883
password_file /mosquitto/config/passwd
```
This prevents anonymous access to the broker, sets the MTTQ listener to port `1883`, and specifies the authentication
file to use. See the [Authenticated accesss to the Broker](#authenticated-access-to-the-broker) section for more
details on setting up the password file.

## Authenticated access to the Broker
By default, the configuration file expects to receive a `/mosquitto/config/passwd` file. There are none provided for
obvious security reasons but it can be created using the `mosquitto_passwd` command line tool.

To start the `mosquitto` container, we must first start it with a blank `passwd` file. To start the container with the
blank `passwd` file, see the [Running](#running) section below. Once the container is running, we can use the following
command to generate authentication content for the `passwd` file:

```
docker exec -it mosquitto mosquitto_passwd -c /mosquitto/config/passwd <USER_NAME>
```

Once done, we can restart the container and have authentication setup correctly.


## Running
To run the container with docker run, we must specify two volumes. The configuration file, and the password file.
Run the container with the `-v <path_to_conf_file>:/mosquitto/config/mosquitto.conf` parametter to provide the
customized configuration file and `-v <path_to_passwd_file>:/mosquitto/config/passwd` to provide the password file.
Other parameters include specifying the ports and other details. For those in a hurry, the full command is provided
below. More details can be found on the image's official documentation on docker hub 
[here](https://hub.docker.com/_/eclipse-mosquitto).

The configuration provided here has authentication enabled. Instructions on generating a password file is detailed under
(Authenticated Access to the Broker)[#authenticated-access-to-the-broker].

Once the password file is generated and configuration is setup, run the image with the following command:
```
docker run -it -d --name mosquitto -p 1883:1883 -p 9001:9001 -v ./mosquitto.conf:/mosquitto/config/mosquitto.conf -v ./passwd:/mosquitto/config/passwd eclipse-mosquitto:latest
```

## Docker Compose
WIP: Provide a docker-compose file to fire everything up.
