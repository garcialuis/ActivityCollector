package client

import (
	"os"

	client_models "github.com/garcialuis/ActivityCollector/client/models"
)

type clientConfig struct {
	hostURL string
}

// ActivityCollector implements ActivityCollectorClient
type ActivityCollector struct {
	config clientConfig
}

// ActivityCollectorClient is the interface that defines our
// functions needed to implement the client
type ActivityCollectorClient interface {
	// configure any variables needed for the client to work
	Configure()
	// GetActivity by originID
	GetActivity(originID uint64) client_models.Activity
}

// Configure method gets client config details from environment
func (activity *ActivityCollector) Configure() {

	url := os.Getenv("ACTIVITY_SERVICE_URL")
	if len(url) == 0 {
		url = "http://localhost:8086/"
	}

	activity.config.hostURL = url
}
