package main

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/app"
	"log"
)

func main() {
	a := app.New()

	err := a.InitDeps()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Starting app...")
	if err = a.Start(); err != nil {
		log.Fatal(err)
	}
}
