package main

import (
	"log"
	"github.com/bogdanCap/go-api/app"
)

func main() {
	app, err := app.New()

	if err != nil {
		log.Fatal(err)
	}

	defer app.Close()
 
	app.Run()
}
