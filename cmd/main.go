package main

import (
	"log"

	"github.com/gradity/go-portainer-action/service"
)

func main() {
	jwt := service.GetJWTToken()
	log.Println(*jwt)
	lp := jwt.ListEndpoints()
	log.Println(*lp)
}
