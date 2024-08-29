package main

import (
	"fmt"

	"github.com/DaniloFaraum/go-crud-api/router"

	"github.com/DaniloFaraum/go-crud-api/config"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Initialize()
}
