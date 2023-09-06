package config

import "github.com/tkanos/gonfig"

type config struct {
	DB_username string
	DB_password string
	DB_host     string
	DB_port     string
	DB_name     string
}

func GetConfig () config {

	configuration := config{}

	gonfig.GetConf("config/config.json", &configuration)
	return configuration
}