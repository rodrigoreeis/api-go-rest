package main

import (
	"github.com/rodrigoreeis/api-go-rest/database"
	"github.com/rodrigoreeis/api-go-rest/routes"
)

func main() {
	database.Connection()
	routes.HandleRequest()
}
