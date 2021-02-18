package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Activity struct {
	ID              uint64  `json:"id" gorm:"primary_key;auto_increment"`
	OriginID        uint64  `json:"originId" gorm:"index:idx_originID;not null"` // Used to map activity data to user
	Epochtime       int64   `json:"epochtime" gorm:"not null"`
	Steps           uint32  `json:"steps" gorm:"not null"`
	StandHrs        uint8   `json:"standHours" gorm:"not null"`
	ExerciseMinutes uint16  `json:"exerciseMinutes" gorm:"not null"`
	CaloriesBurned  float64 `json:"caloriesBurned" gorm:"not null"`
}

type Activities []Activity

func (activity *Activity) GetActivitySample(origin uint64) (*Activity, error) {

	now := time.Now()

	sampleActivity := Activity{
		ID:              1,
		OriginID:        origin,
		Epochtime:       now.Unix(),
		Steps:           2890,
		StandHrs:        7,
		ExerciseMinutes: 4,
		CaloriesBurned:  201,
	}

	if origin != 0 {
		return &sampleActivity, nil
	}

	return &Activity{}, errors.New("Error: data not found")
}

// GetActivity method retrieves activity data given an originID and activityID
func (activity *Activity) GetActivity(db *gorm.DB, originID uint64, activityID uint64) (*Activity, error) {

	err := db.Debug().Model(&Activity{}).Where("id = ? AND origin_id = ?", activityID, originID).Take(&activity).Error
	if err != nil {
		return &Activity{}, err
	}

	return activity, nil
}

func (activity *Activity) SaveActivity(db *gorm.DB) (*Activity, error) {
	var err error

	err = db.Debug().Create(&activity).Error
	if err != nil {
		return &Activity{}, err
	}

	return activity, err
}

func (activities *Activities) GetActivitiesInRange(db *gorm.DB, originID uint64, startUnixTime int64, endUnixTime int64) (*Activities, error) {

	err := db.Debug().Model(&Activity{}).Where("origin_id = ? AND epochtime BETWEEN ? AND ?", originID, startUnixTime, endUnixTime).Find(&activities).Error

	if err != nil {
		return &Activities{}, err
	}

	return activities, nil
}

func (activities *Activities) GetActivities(db *gorm.DB, originID uint64) (*Activities, error) {

	err := db.Debug().Model(&Activity{}).Where("origin_id = ?", originID).Find(&activities).Error
	if err != nil {
		return &Activities{}, err
	}

	return activities, nil
}
