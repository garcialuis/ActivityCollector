package api

import (
	"os"

	"github.com/garcialuis/ActivityCollector/api/controllers"
	"github.com/garcialuis/ActivityCollector/api/services"
)

var rabbitMQServer = controllers.Server{}

// Run function initializes connection to db and starts server in specified port
func RunRabbitMQ() {

	rabbitMQServer.Initialize(os.Getenv("DB_POSTGRES_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	services.RunConsumer(rabbitMQServer.DB)
}
