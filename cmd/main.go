package main

import (
	"log"
	"github.com/bogdanCap/go-api/app"
)

func main() {
	application, err := app.New()

	if err != nil {
		log.Fatal(err)
	}
 
	application.Run()
}
