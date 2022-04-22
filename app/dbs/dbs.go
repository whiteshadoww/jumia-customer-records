package dbs

import (
	"github.com/quangdangfit/gocommon/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go.jumia.org/customers/app/interfaces"
	"go.jumia.org/customers/config"
)

type database struct {
	db *gorm.DB
}

// NewDatabase return new IDatabase interface
func NewDatabase() interfaces.IDatabase {
	dbConfig := config.Config.Database

	logger.Info(dbConfig.Path)

	db, err := gorm.Open(sqlite.Open(dbConfig.Path))
	if err != nil {
		logger.Fatal("Cannot connect to database: ", err)
	}

	return &database{
		db: db,
	}
}

// GetInstance get database instance
func (d *database) GetInstance() *gorm.DB {
	return d.db.Debug()
}
