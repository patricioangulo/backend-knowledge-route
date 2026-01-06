package main

type MessageService struct{}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (s *MessageService) HelloMessage() string {
	return "Hello!"
}

func (s *MessageService) GoodbyeMessage() string {
	return "Goodbye!"
}
