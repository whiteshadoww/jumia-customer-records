package repositories

import (
	"go.jumia.org/customers/app/interfaces"
	"go.jumia.org/customers/app/models"
	"go.jumia.org/customers/app/schema"
	"go.jumia.org/customers/pkg/errors"
	"go.jumia.org/customers/pkg/utils"
)

// CustomerRepo customer repository struct
type CustomerRepo struct {
	db interfaces.IDatabase
}

// NewCustomerRepository return new ICustomerRepository interface
func NewCustomerRepository(db interfaces.IDatabase) interfaces.ICustomerRepository {
	return &CustomerRepo{db: db}
}

// List get customer by CustomerQueryParam
func (r *CustomerRepo) List(param *schema.CustomerQueryParam) (*[]models.Customer, error) {
	var query map[string]interface{}
	err := utils.Copy(&query, &param)
	if err != nil {
		return nil, errors.ErrorMarshal.Newm(err.Error())
	}

	//make starts at index 0
	param.Page--

	var customer []models.Customer
	if err := r.db.GetInstance().Where(query).Offset(param.Page * param.PageSize).Limit(param.PageSize).Find(&customer).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return &customer, nil
}
