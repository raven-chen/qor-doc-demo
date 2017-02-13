package migration

import (
	"github.com/raven-chen/qor_doc_demo/app/models"
	"github.com/raven-chen/qor_doc_demo/db"
)

func init() {
	Run(false)
}

func Run(force bool) {
	if db.DB == nil {
		return
	}
	tables := []interface{}{
		&models.User{}, &models.Product{}, &models.ColorVariation{}, &models.Parameter{}, &models.Order{},
	}

	for _, table := range tables {
		if force {
			if err := db.DB.DropTableIfExists(table).Error; err != nil {
				panic(err)
			}
		}

		if err := db.DB.AutoMigrate(table).Error; err == nil {
		} else {
			db.DB.AddError(err)
		}
	}
}
