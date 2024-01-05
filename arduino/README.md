# Arduino
This part of the project contains the source file to program the ESP8266 microcontroller to read temperature and 
humidity readings from the DHT11 sensor, connect to the configured WiFi network, and publish the data to the MQTT
broker.

The data published follows the following JSON format:

```
{
    "timestamp": <epoc time>,
    "sensor_id": <string containing the sensor's id>,
    "temperature": <float containing the temperature in celcius, with two decimal points.>
    "humidity": <float containing the humidity percentage, with two decimal points.>,
}
```

## Building
Building the project requires installing Arduino IDE, and the required libraries for MQTT, NTPClient, and the DHT
sensor.

[Random Nerd Tutorials](https://randomnerdtutorials.com/getting-started-with-esp8266-wifi-transceiver-review/) contains
all the detailed information you will need to get familiar with the ESP8266 and install Arduino IDE.

For the library dependencies:
 - The MQTT library and its instructions can be found [here](https://github.com/marvinroger/async-mqtt-client).
 - The NTPClient library can be directly installed from the Arduino IDE library manager.
 - The DHT library and its instructions can be found [here](https://github.com/adafruit/DHT-sensor-library).

