package main

import (
	"log"

	"github.com/luispfcanales/service-email/api"
)

func main() {
	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
