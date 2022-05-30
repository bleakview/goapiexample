package main

import (
	"os"

	"github.com/bleakview/goapiexample/cmd/goapiexample/app"
)

// @title Books DB application
// @version 1.0.0
// @description <br>
// @description <br>
// @description This API handles books operations.
// @description <br>
// @description ## Operations
// @description <br>
// @description With this API you can:
// @description <br>
// @description * **Create book**
// @description * **Update book**
// @description * **Get all books**
// @description * **Get specific book with {id}**

// @contact.name Mustafa Unal
// @contact.url https://github.com/bleakview/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	app := &app.App{}
	app.Initialize()
	port := "4000"
	portSet, isPortSet := os.LookupEnv("PORT")
	if isPortSet {
		port = portSet
	}

	app.Run(":" + port)
}
