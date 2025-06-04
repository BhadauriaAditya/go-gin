package database

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func ProvideGormDB() (*gorm.DB, error) {
	err := InitAllDatabases()
	if err != nil {
		return nil, err
	}
	return GetDB("gin") // Assuming 'gin' is the main DB
}

var DatabaseSet = wire.NewSet(ProvideGormDB)