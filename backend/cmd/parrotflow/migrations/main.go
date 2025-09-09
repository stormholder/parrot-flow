package migrations

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"parrotflow/internal/models"
)

func Init(dbPath string) {
	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	database.AutoMigrate(&models.Scenario{})

	// database.AutoMigrate(&db.WebSite{}, &db.Proxy{}, &db.Scenario{}, &db.ExecutionBlob{}, &db.Execution{})
}
