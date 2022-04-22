package migration

import (
	"github.com/quangdangfit/gocommon/logger"
	"go.jumia.org/customers/app/interfaces"
	"go.jumia.org/customers/app/models"
	"go.jumia.org/customers/pkg/utils"
	"go.jumia.org/customers/pkg/validators"
	"go.uber.org/dig"
)

// Migrate migrate to database
func Migrate(container *dig.Container) error {
	return container.Invoke(func(
		db interfaces.IDatabase,
	) error {
		var customer []models.Customer
		logger.Info("Migrating Data")
		err := db.GetInstance().AutoMigrate(&models.Customer{})
		if err != nil {
			logger.Error(err)
		}

		db.GetInstance().Model(models.Customer{}).Find(&customer)

		for _, m := range customer {
			m.Country = string(validators.CheckCountry(m.Phone))
			m.CountryCode = string(validators.CheckCountryCode(m.Phone))

			m.State = utils.BoolPointer(validators.CheckPhoneValidity(m.Phone))
			db.GetInstance().Where("id = ?", m.ID).UpdateColumns(&m)
		}

		return nil
	})
}
