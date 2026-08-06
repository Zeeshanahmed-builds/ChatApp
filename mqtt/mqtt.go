package mqtt

import (
	"os"

	paho "github.com/eclipse/paho.mqtt.golang"
)

func Connect() (paho.Client, error) {

	// Get MQTT broker from environment variable, default to localhost
	mqttBroker := os.Getenv("MQTT_BROKER")
	if mqttBroker == "" {
		mqttBroker = "tcp://emqx:1883"
	}
	opts := paho.NewClientOptions()
	opts.AddBroker(mqttBroker)
	opts.SetClientID("go-server")

	client := paho.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return client, nil
}
