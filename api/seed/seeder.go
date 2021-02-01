package seed

import (
	"log"

	"github.com/garcialuis/ActivityCollector/api/models"
	"github.com/jinzhu/gorm"
)

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Activity{}).Error

	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.Activity{}).Error

	if err != nil {
		log.Fatalf("cannot migrate table(s): %v", err)
	}
}
