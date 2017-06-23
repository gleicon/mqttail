## MQTTail

Console client for [mqtt](http://mqtt.org/) written in golang. Tails a topic, integrates with scripting.

### Build

$ make all

### Usage

	$ ./mqttail -h
	Usage: mqttail [-s ip:port] [-f] <topic filter>
	mqttail consumes and print messages from a mqtt topic.
	-s broker(s) separated by comma. Defaults to localhost:1883
	-q QoS level (0, 1, 2 - default QoS 0)

