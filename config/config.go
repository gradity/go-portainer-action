package config

import (
	"log"

	"github.com/caarlos0/env"
)

type PortainerConfig struct {
	Url      string `env:"PORTAINER_URL"`
	Username string `env:"PORTAINER_USERNAME"`
	Password string `env:"PORTAINER_PASSWORD"`
}

var Portainer PortainerConfig

func init() {
	env.Parse(&Portainer)
	log.Println(Portainer)
}
