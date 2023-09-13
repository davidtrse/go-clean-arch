package di

import (
	"github.com/davidtrse/go-clean-arch/pkg/connection"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	return connection.DB
}
