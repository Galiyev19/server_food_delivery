package app

import (
	"fmt"
	"food_delivery/internal/config"
	"log"
)

func Run() {
	// TODO: init config: json

	config := config.New()
	if err := config.InitConfig("config.json", config); err != nil {
		log.Fatal("ERROR: %v", err)
	}

	fmt.Println("INIT CONFIG")

	// TODO: init storage: sqlite3

	// TODO : init router

	// TODO: run server
}
