package main

import "github.com/parrotmac/littleblue/pkg"

func main() {
	app := pkg.NewDefaultApp()
	app.Run("0.0.0.0:9000")
}
