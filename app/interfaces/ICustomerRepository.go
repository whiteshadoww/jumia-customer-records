package interfaces

import (
	"go.jumia.org/customers/app/models"
	"go.jumia.org/customers/app/schema"
)

// ICustomerRepository interface
type ICustomerRepository interface {
	List(queryParam *schema.CustomerQueryParam) (*[]models.Customer, error)
}
