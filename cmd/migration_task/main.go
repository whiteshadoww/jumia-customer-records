package main

import (
	"github.com/quangdangfit/gocommon/logger"
	"go.jumia.org/customers/app"
	"go.jumia.org/customers/app/migration"
)

func main() {
	container := app.BuildContainer()
	err := migration.Migrate(container)
	if err != nil {
		logger.Error("Failed to create admin: ", err)
	}
}
