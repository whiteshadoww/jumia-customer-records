package api

import (
	"go.uber.org/dig"
)

// Inject apis
func Inject(container *dig.Container) error {
	_ = container.Provide(NewCustomerAPI)
	return nil
}
