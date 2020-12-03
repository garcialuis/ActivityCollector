package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Activity struct {
	ID              uint64 `json:"id"`
	OriginID        uint64 `json:"originId"` // Used to map activity data to user
	Epochtime       int64  `json:"epochtime"`
	Steps           uint32 `json:"steps"`
	StandHrs        uint8  `json:"standHours"`
	ExerciseMinutes uint16 `json:"exerciseMinutes"`
	CaloriesBurned  uint16 `json:"caloriesBurned"`
}

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

// GetActivity method retrieves activity data given an originID
func (activity *Activity) GetActivity(db *gorm.DB, origin uint64) (*Activity, error) {

	err := db.Debug().Model(&Activity{}).Where("originID = ?", origin).Take(&activity).Error
	if err != nil {
		return &Activity{}, err
	}

	return activity, nil
}
