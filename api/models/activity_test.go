package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleActivity(t *testing.T) {
	activity := Activity{OriginID: 1}

	sampleActivity, err := activity.GetActivitySample(activity.OriginID)

	expectedActivity := Activity{
		Steps:           2890,
		CaloriesBurned:  201,
		StandHrs:        7,
		ExerciseMinutes: 4,
	}

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedActivity.Steps, sampleActivity.Steps)
	assert.Equal(t, expectedActivity.CaloriesBurned, sampleActivity.CaloriesBurned)
	assert.Equal(t, expectedActivity.StandHrs, sampleActivity.StandHrs)
	assert.Equal(t, expectedActivity.ExerciseMinutes, sampleActivity.ExerciseMinutes)
}
