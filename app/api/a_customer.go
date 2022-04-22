package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/quangdangfit/gocommon/logger"
	"go.jumia.org/customers/app/interfaces"
	"go.jumia.org/customers/app/schema"
	"go.jumia.org/customers/pkg/errors"
	gohttp "go.jumia.org/customers/pkg/http"
)

// CustomerAPI handle customer api
type CustomerAPI struct {
	service interfaces.ICustomerService
}

// NewCustomerAPI return new CustomerAPI pointer
func NewCustomerAPI(service interfaces.ICustomerService) *CustomerAPI {
	return &CustomerAPI{service: service}
}

// List customer by query
func (u *CustomerAPI) List(c *gin.Context) gohttp.Response {
	var queryParam schema.CustomerQueryParam
	if err := c.ShouldBindQuery(&queryParam); err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	validate := validator.New()
	if err := validate.Struct(queryParam); err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	customer, err := u.service.List(c, &queryParam)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res []schema.Customer
	copier.Copy(&res, &customer)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}
