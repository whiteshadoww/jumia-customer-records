package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/logger"
	"go.jumia.org/customers/app/middleware"
	"go.uber.org/dig"

	"go.jumia.org/customers/app/api"
	"go.jumia.org/customers/pkg/http/wrapper"
)

// RegisterAPI register api routes
func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		customerAPI *api.CustomerAPI,
	) error {

		//--------------------------------API-----------------------------------
		apiPath := r.Group("/api/v1", middleware.CORSMiddleware())
		{
			apiPath.GET("/customers", wrapper.Wrap(customerAPI.List))
		}
		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return err
}
