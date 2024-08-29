package main

import (
	"github.com/DaniloFaraum/go-crud-api/router"

	"github.com/DaniloFaraum/go-crud-api/config"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.NewLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Initialization failed: %v", err)
		return
	}

	router.Initialize()
}
