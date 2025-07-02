package migration

import (
	"gorm.io/gorm"
	"secret-check-integrator/app/dbo/entity"
)

func DoMigration(db *gorm.DB) {
	db.AutoMigrate(&entity.Project{})
}
