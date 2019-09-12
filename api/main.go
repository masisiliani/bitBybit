package main

import (
	"fmt"

	"github.com/masisiliani/bitBybit/api/router"
)

func main() {

	fmt.Println("Starting server...")
	router.InicializeRoutes()
}
