package mqtt

import (
	"fmt"

	paho "github.com/eclipse/paho.mqtt.golang"
)

func Subscribe(client paho.Client) error {

	token := client.Subscribe("chat/#", 1, func(client paho.Client, msg paho.Message) {
		fmt.Println("[MQTT LOG] Topic:", msg.Topic())
		fmt.Println("[MQTT LOG] Payload:", string(msg.Payload()))
	})

	token.Wait()
	return token.Error()
}