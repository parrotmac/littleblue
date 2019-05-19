package main

import (
	"github.com/sirupsen/logrus"

	"github.com/parrotmac/littleblue/pkg"
)

func main() {
	// Load configuration
	err := pkg.LoadConfig(".")
	if err != nil {
		logrus.Fatalln(err)
	}
	err = pkg.Config.Validate()
	if err != nil {
		logrus.Fatalln(err)
	}

	// Create application from configuration
	app := pkg.NewDefaultApp(&pkg.Config)
	app.Run()
}
