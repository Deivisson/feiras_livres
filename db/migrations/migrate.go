package migrations

import (
	"gorm.io/gorm"
)

func Load(db *gorm.DB) {
	createFairTable(db)
}
