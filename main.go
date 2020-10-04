package main

import (
	"github/saurabh7517/glofox_assignment/routes"
	"net/http"
)

const baseAPIPath string = "/api"

func main() {
	routes.SetupRoutes(baseAPIPath)
	http.ListenAndServe(":8080", nil)
}
