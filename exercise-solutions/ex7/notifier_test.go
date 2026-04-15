package ex7

import (
	"errors"
	"testing"
)

// SenderStub implements MessageSender by delegating to a function field.
type SenderStub struct {
	sendFunc func(recipient string, body string) error
}

func (s SenderStub) Send(recipient string, body string) error {
	return s.sendFunc(recipient, body)
}

var errTestConnectionRefused = errors.New("connection refused")

func TestNotify(t *testing.T) {
	workingSender := func(recipient string, body string) error {
		return nil
	}
	data := []struct {
		name   string
		sender func(recipient string, body string) error
		to     string
		body   string
		err    error
	}{
		{
			name:   "successful send",
			sender: workingSender,
			to:     "alice@example.com",
			body:   "Hello Alice",
			err:    nil,
		},
		{
			name:   "empty recipient",
			sender: workingSender,
			to:     "",
			body:   "Hello",
			err:    ErrRecipientRequired,
		},
		{
			name:   "empty body",
			sender: workingSender,
			to:     "alice@example.com",
			body:   "   ",
			err:    ErrBodyRequired,
		},
		{
			name: "sender returns error",
			sender: func(recipient string, body string) error {
				return errTestConnectionRefused
			},
			to:   "carol@example.com",
			body: "Important update",
			err:  &SendFailedErr{Err: errTestConnectionRefused},
		},
	}

	var ns NotificationService
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			ns.Sender = SenderStub{d.sender}
			err := ns.Notify(d.to, d.body)
			if !errors.Is(err, d.err) {
				t.Errorf("expected %v, got %v", d.err, err)
			}
		})
	}
}
