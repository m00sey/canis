package messagepickup

import (
	"github.com/hyperledger/aries-framework-go/pkg/controller/command"
)

type provider interface{}

type Command struct{}

func New(ctx provider, registrar command.MessageHandler, notifier command.Notifier) (*Command, error) {
	return &Command{}, nil
}
