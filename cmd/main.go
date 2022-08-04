package main

import (
	"log"

	"github.com/gradity/go-portainer-action/internal"
)

func main() {
	jwt := internal.GetJWTToken()
	log.Println(*jwt)
	lp := internal.ListEndpoints(*jwt)
	log.Println(*lp)
}
