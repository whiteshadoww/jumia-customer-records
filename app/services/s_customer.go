package services

import (
	"context"
	"go.jumia.org/customers/app/interfaces"
	"go.jumia.org/customers/app/models"
	"go.jumia.org/customers/app/schema"
	"strings"
)

var (
	maxPageSize = 30
	defaultPage = 0
)

// CustomerService customer service
type CustomerService struct {
	customerRepo interfaces.ICustomerRepository
}

// NewCustomerService return new ICustomerService interface
func NewCustomerService(customer interfaces.ICustomerRepository) interfaces.ICustomerService {
	return &CustomerService{
		customerRepo: customer,
	}
}

// List customers by query
func (u *CustomerService) List(ctx context.Context, param *schema.CustomerQueryParam) (*[]models.Customer, error) {

	if param.PageSize > maxPageSize {
		param.PageSize = maxPageSize
	} else if param.PageSize <= 0 {
		param.PageSize = defaultPage
	}

	param.Country = strings.Title(strings.ToLower(param.Country))

	customer, err := u.customerRepo.List(param)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
