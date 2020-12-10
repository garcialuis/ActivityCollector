package main

import (
	"github.com/garcialuis/ActivityCollector/api"
)

func main() {

	go func() {
		api.RunRabbitMQ()
	}()

	api.Run()

}
