package client_models

type Activity struct {
	ID              uint64 `json:"id" gorm:"primary_key;auto_increment"`
	OriginID        uint64 `json:"originId" gorm:"index:idx_originID;not null"` // Used to map activity data to user
	Epochtime       int64  `json:"epochtime" gorm:"not null"`
	Steps           uint32 `json:"steps" gorm:"not null"`
	StandHrs        uint8  `json:"standHours" gorm:"not null"`
	ExerciseMinutes uint16 `json:"exerciseMinutes" gorm:"not null"`
	CaloriesBurned  uint16 `json:"caloriesBurned" gorm:"not null"`
}
