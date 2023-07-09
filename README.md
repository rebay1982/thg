# thg
Temperature and Hygrometer IoT project.



# Setup, installation, and configuration

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

