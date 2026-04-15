package ex7

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrRecipientRequired = errors.New("recipient is required")
	ErrBodyRequired      = errors.New("body is required")
)

type SendFailedErr struct {
	Err error
}

func (se *SendFailedErr) Error() string {
	return fmt.Sprintf("send failed: %v", se.Err)
}

func (se *SendFailedErr) Is(err error) bool {
	if e, ok := errors.AsType[*SendFailedErr](err); ok {
		return errors.Is(se.Err, e.Err)
	}
	return false
}

func (se *SendFailedErr) Unwrap() error {
	return se.Err
}

// MessageSender is an abstraction over how messages are delivered.
type MessageSender interface {
	Send(recipient string, body string) error
}

// NotificationService uses a MessageSender to deliver notifications.
type NotificationService struct {
	Sender MessageSender
}

// Notify validates the inputs and sends a message through the Sender.
// It returns an error if the recipient is empty, the body is empty,
// or the Sender itself returns an error.
func (ns NotificationService) Notify(recipient string, body string) error {
	recipient = strings.TrimSpace(recipient)
	if recipient == "" {
		return ErrRecipientRequired
	}
	body = strings.TrimSpace(body)
	if body == "" {
		return ErrBodyRequired
	}
	err := ns.Sender.Send(recipient, body)
	if err != nil {
		return &SendFailedErr{Err: err}
	}
	return nil
}
