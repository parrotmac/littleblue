package main

import (
	"github.com/sirupsen/logrus"

	"github.com/parrotmac/littleblue/pkg"
)

func main() {
	// Create application from configuration
	// Load configuration
	config := &littleblue.Config{}
	err := config.LoadConfig(".")
	if err != nil {
		logrus.Fatalln(err)
	}
	err = config.Validate()
	if err != nil {
		logrus.Fatalln(err)
	}

	a := littleblue.App{
		Config: config,
	}
	a.Run()
}
