package repositories

import (
	"go.uber.org/dig"
)

// Inject repositories
func Inject(container *dig.Container) error {
	_ = container.Provide(NewCustomerRepository)
	return nil
}
