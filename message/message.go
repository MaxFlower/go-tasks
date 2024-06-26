package message

import "fmt"

type Message struct {
	uuid        string
	text        string
	messageType string
}

type MessageService interface {
	NewMessage(i, tx, tp string) Message
	Add(m Message)
	// Remove(id string)
	Print()
	// Test()
}

type SevenNewsService struct {
	Messages []Message
}

func (s *SevenNewsService) NewMessage(i, tx, tp string) Message {
	return Message{
		uuid:        i,
		text:        tx,
		messageType: tp,
	}
}

func (s *SevenNewsService) Add(m Message) {
	s.Messages = append(s.Messages, m)
}

func (s *SevenNewsService) Print() {
	fmt.Println(s.Messages)
}
