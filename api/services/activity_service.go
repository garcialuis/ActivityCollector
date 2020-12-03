package services

import (
	"github.com/garcialuis/ActivityCollector/api/models"
	"github.com/jinzhu/gorm"
)

func GetActivitySample(originID uint64) (*models.Activity, error) {

	activity := models.Activity{}

	return activity.GetActivitySample(originID)
}

func GetActivity(db *gorm.DB, originID uint64) (*models.Activity, error) {

	activity := models.Activity{}

	return activity.GetActivity(db, originID)
}
