package config

import (
	"log"

	"github.com/caarlos0/env"
)

type PortainerConfig struct {
	Url            string `env:"PORTAINER_URL"`
	Username       string `env:"PORTAINER_USERNAME"`
	Password       string `env:"PORTAINER_PASSWORD"`
	EndPointId     string `env:"ENDPOINT_ID"`
	ManifestLogin  string `env:"MANIFEST_LOGIN"`
	ManifestLogin2 string `env:"MANIFEST_LOGIN2"`
	ByteManifest   []byte
}

var Portainer PortainerConfig

func init() {
	env.Parse(&Portainer)
	log.Println("ManifestLogin: ", Portainer.ManifestLogin)
	Portainer.ByteManifest = []byte(Portainer.ManifestLogin)
	log.Println("ByteManifest: ", Portainer.ByteManifest)
}
