package main

import (
	"github.com/cansulting/elabox-system-tools/foundation/app"
)

func main() {
	controller, err := app.NewControllerWithDebug(&Activity{}, nil, true)
	if err != nil {
		panic(err)
	}
	AppController = controller
	app.RunApp(controller)
}
