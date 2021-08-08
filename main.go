package main

import (
	"flag"
	"log"
	"mqtt-client-Go/protocols"
)

var Host     = flag.String("host", "192.168.2.112", "server hostname or IP")
var Port     = flag.Int("port", 8882, "server port")
var Action   = flag.String("action", "pubsub", "pub/sub/pubsub action")
var Protocol = flag.String("protocol", "ws", "mqtt/mqtts/ws/wss")
var Topic    = flag.String("topic", "cameras/go", "publish/subscribe topic")
var Username = flag.String("username", "admin2", "mqtt broker username")
var Password = flag.String("password", "instar2", "mqtt broker password")
var Qos      = flag.Int("qos", 1, "MQTT QOS")
var Tls      = flag.Bool("tls", false, "Enable TLS/SSL")
var CaCert   = flag.String("cacert", "./broker.ca.crt", "tls certificate")

func main() {
	flag.Parse()
	config := protocols.Config{Host: *Host, Port: *Port, Action: *Action, Topic: *Topic, Username: *Username, Password: *Password, Qos: *Qos, Tls: *Tls, CaCert: *CaCert}
	protocol := *Protocol
	switch protocol {
	case "mqtt":
		protocols.MQTTConnection(config)
	case "mqtts":
		protocols.MQTTSConnection(config)
	case "ws":
		protocols.WSConnection(config)
	case "wss":
		protocols.WSSConnection(config)
	default:
		log.Fatalf("Unsupported protocol: %s", protocol)
	}
}
