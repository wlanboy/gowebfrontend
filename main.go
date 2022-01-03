package main

import (
	app "github.com/wlanboy/gowebfrontend/v2/application"
)

func main() {
	a := app.GoService{}
	a.Initialize()

	a.Run()

	a.WaitForShutdown()
}
