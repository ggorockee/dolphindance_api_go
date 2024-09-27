package config

import (
	"fmt"

	"github.com/ggorockee/toolbox"
	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string, defaultValue string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return toolbox.Getenv(key, defaultValue)
}
