package main

import (
	"fmt"
	"github.com/danyaobertan/api-test/initializers"
	"github.com/danyaobertan/api-test/models"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
	fmt.Println("? Migration complete")
}
