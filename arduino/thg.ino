#include <NTPClient.h>
#include <WiFiUdp.h>
#include <ESP8266WiFi.h>
#include <Ticker.h>
#include "DHT.h"
#include <AsyncMqttClient.h>

#define WIFI_SSID ""
#define WIFI_PASS ""
#define SENSOR_ID ""
#define DHT_TYPE DHT11

#define MQTT_HOST ".local"
#define MQTT_PORT 1883
#define MQTT_PUB_TOPIC ""
#define MQTT_USER ""
#define MQTT_PASS ""

const int DHTPin = 5;
DHT dht(DHTPin, DHT_TYPE);

// Define NTP Client to get time
WiFiUDP ntpUDP;
NTPClient timeClient(ntpUDP, "pool.ntp.org");

AsyncMqttClient mqttClient;
Ticker mqttReconnectTimer;

WiFiEventHandler wifiConnectHandler;
WiFiEventHandler wifiDisconnectHandler;
Ticker wifiReconnectTimer;

void connectToWifi() {
  Serial.println("Connecting to Wi-Fi...");
  WiFi.begin(WIFI_SSID, WIFI_PASS);
}

void onWifiConnect(const WiFiEventStationModeGotIP& event) {
  Serial.println("Connected to Wi-Fi.");
  connectToMqtt();
}

void connectToMqtt() {
  Serial.println("Connecting to MQTT...");
  mqttClient.connect();
}

void onWifiDisconnect(const WiFiEventStationModeDisconnected& event) {
  Serial.println("Disconnected from Wi-Fi.");
  mqttReconnectTimer.detach(); // ensure we don't reconnect to MQTT while reconnecting to Wi-Fi
  wifiReconnectTimer.once(2, connectToWifi);
}

void onMqttConnect(bool sessionPresent) {
  Serial.println("Connected to MQTT.");
  Serial.print("Session present: ");
  Serial.println(sessionPresent);
}

void onMqttDisconnect(AsyncMqttClientDisconnectReason reason) {
  Serial.println("Disconnected from MQTT.");

  if (WiFi.isConnected()) {
    mqttReconnectTimer.once(2, connectToMqtt);
  }
}

void setup(){
  Serial.begin(115200);
  delay(10);

  wifiConnectHandler = WiFi.onStationModeGotIP(onWifiConnect);
  wifiDisconnectHandler = WiFi.onStationModeDisconnected(onWifiDisconnect);

  mqttClient.onConnect(onMqttConnect);
  mqttClient.onDisconnect(onMqttDisconnect);
  //mqttClient.onPublish(onMqttPublish);
  mqttClient.setServer(MQTT_HOST, MQTT_PORT);
  mqttClient.setCredentials(MQTT_USER, MQTT_PASS);
  
  connectToWifi();
  while ( WiFi.status() != WL_CONNECTED ) {
    delay ( 500 );
    Serial.print ( "." );
  }

  dht.begin();
  timeClient.begin();
}

void loop() {
  timeClient.update();

  long timestamp = timeClient.getEpochTime();
  float h = dht.readHumidity();
  float t = dht.readTemperature();

  char buffer[128];
  sprintf(buffer, "{\"timestamp\": %d, \"sensor_id\": \"%s\", \"temperature\": %.2f, \"humidity\": %.2f}", timestamp, SENSOR_ID, t, h);
  Serial.println(buffer);
  
  mqttClient.publish(MQTT_PUB_TOPIC, 0, false, buffer);
  delay(60000);
}
