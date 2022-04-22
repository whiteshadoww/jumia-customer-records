package services

import (
	"go.uber.org/dig"
)

// Inject services
func Inject(container *dig.Container) error {
	_ = container.Provide(NewCustomerService)
	return nil
}
