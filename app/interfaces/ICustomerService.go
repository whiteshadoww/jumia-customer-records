package interfaces

import (
	"context"

	"go.jumia.org/customers/app/models"
	"go.jumia.org/customers/app/schema"
)

// ICustomerService interface
type ICustomerService interface {
	List(ctx context.Context, param *schema.CustomerQueryParam) (*[]models.Customer, error)
}
