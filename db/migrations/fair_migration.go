package migrations

import (
	"log"

	"github.com/Deivisson/feiras_livres/domain"
	"gorm.io/gorm"
)

func createFairTable(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&domain.Fair{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}
