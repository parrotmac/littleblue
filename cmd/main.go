package main

import (
	"github.com/parrotmac/littleblue/pkg"
)

func main() {
	// Create application from configuration
	app := pkg.NewDefaultApp()
	app.Run()
}
