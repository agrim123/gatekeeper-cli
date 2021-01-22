package notifier

import (
	"errors"
)

type CustomNotifier struct {
}

func (c CustomNotifier) Notify(message string) error {
	// will force to fallback to default notifier
	return errors.New("something")
}

func (c CustomNotifier) Name() string {
	return "custom_notifier"
}
