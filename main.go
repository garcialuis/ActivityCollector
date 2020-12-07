package main

import (
	"github.com/garcialuis/ActivityCollector/api"
	"github.com/garcialuis/ActivityCollector/api/services"
)

func main() {

	go func() {
		services.RunConsumer()
	}()

	api.Run()

}
