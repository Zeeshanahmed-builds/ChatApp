package message_service_imp
import (
	"chat-app/repo/message_repo"
	"chat-app/service/message_service"
)

type MessageServiceImp struct {
	messageRepo message_repo.MessageRepo
}

func NewMessageService(input NewMessageServiceImp) message_service.MessageService {
	return &MessageServiceImp{
		messageRepo: input.MessageRepo,
	}
}

type NewMessageServiceImp struct {
	MessageRepo message_repo.MessageRepo
}