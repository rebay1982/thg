# THG Sink
This part of the project is the sink service that subscribes to the MQTT broker and persists all data messages received
from the temp/humidity sensors into InfluxDB.

## Building
Building the project is as simple as installing the latest Golang release along with `make`.

The `build` targer in the included `Makefile` is to build the Go binary. This binary can then be deployed to the
destination machine as long as it is running the same operating system and architecture as the machine used to build the
binary. Take note that the `build` target can be modified to crossbuild to different target OSes and architectures. Feel
free to add whatever you need here.

## Configuration and deployment
Deployment is done by copying the executable built in the (previous section)[#building] to 
the destination machine. The `sink.sh` file under `./config` needs to be configured and copied alongside the binary.
`sink.sh` contains some variable environments that need to be populated with hostnames and credentials:

 - THG_INFLUXDN_WRITE_TOKEN: Needs to contain an API token that was created in InfluxDB to serve as write credentials for
the service.
 - THG_MQTT_USER and THG_MQTT_PASS: Username and password to the MQTT broker (Mosquitto in this case).

Finally, copy the `thgsink.service` file to your distribution's `systemd` folder. Install, enable, and start the service
as you would with any `systemd` unit.

## Docker
Why isn't this in Docker? There is an issue with the reverse name lookup function when the service is being run in
Docker and I haven't found a solution to the problem, yet. In the interest of project completion, this is why a
`systemd` unit file has been included to run the binary natively on the destiantion server.

When the Docker deployment gets sorted out it will be possible to use secrets instead of environment variables to
provide credentials to the service.
