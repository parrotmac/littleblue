package main

import (
	"github.com/sirupsen/logrus"

	"github.com/parrotmac/littleblue/pkg/app"
)

func main() {
	// Create application from configuration
	// Load configuration
	config := &app.Config{}
	err := config.LoadConfig(".")
	if err != nil {
		logrus.Fatalln(err)
	}
	err = config.Validate()
	if err != nil {
		logrus.Fatalln(err)
	}

	a := app.App{
		Config: config,
	}
	a.Run()
}
