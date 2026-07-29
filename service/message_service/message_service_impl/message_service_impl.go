package message_service_imp
import (
	"chat-app/repo/message_repo"
	"chat-app/service/message_service"
	paho "github.com/eclipse/paho.mqtt.golang"
)

type MessageServiceImp struct {
	messageRepo message_repo.MessageRepo
	mqttClient  paho.Client
}

func NewMessageService(input NewMessageServiceImp) message_service.MessageService {
	return &MessageServiceImp{
		messageRepo: input.MessageRepo,
		mqttClient: input.MQTTClient,
	}
}

type NewMessageServiceImp struct {
	MessageRepo message_repo.MessageRepo
	MQTTClient  paho.Client
}