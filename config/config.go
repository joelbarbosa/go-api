package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var config appConfigStruct
var doOnce sync.Once

type appConfigStruct struct {
	AppPort string
}

func init() {
	doOnce.Do(func() {
		err := godotenv.Load()
		log.Println("Loading .env file...")
		if err != nil {
			log.Println("Err: Cannot load .env file")
		}

		config = load()
	})
}

func load() appConfigStruct {
	return appConfigStruct{
		AppPort: os.Getenv("APP_PORT"),
	}
}

func Get() appConfigStruct {
	return config
}
