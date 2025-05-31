package main

import (
	"another-restful-api/internal/env"
	"log"
)

func main() {
	app := &application{
		config: config{
			addr: env.GetString("ADDR", ":8080"),
		},
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
