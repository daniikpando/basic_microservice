package configuration

import (
	"os"
	"encoding/json"
	"log"
)

type Configuration struct {
	Engine   string
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

func GetConfiguration() Configuration {
	var config Configuration

	file, err := os.Open("./config.json")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)

	if err != nil {
		log.Fatal(err)
	}
	return config
}