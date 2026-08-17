package main

import (
	"log"
	"test/app"
)

func main() {
	application, err := app.New()

	if err != nil {
		log.Fatal(err)
	}
 
	application.Run()
}
