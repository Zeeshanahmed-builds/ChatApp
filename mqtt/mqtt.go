package mqtt

import (
	paho "github.com/eclipse/paho.mqtt.golang"
)

func Connect() (paho.Client, error) {
	opts := paho.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("go-server")

	client := paho.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return client, nil
}